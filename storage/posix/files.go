// Copyright 2024 The Tessera authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package posix

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/transparency-dev/merkle/rfc6962"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/internal/options"
	"github.com/transparency-dev/trillian-tessera/internal/parse"
	storage "github.com/transparency-dev/trillian-tessera/storage/internal"
	"k8s.io/klog/v2"
)

const (
	dirPerm  = 0o755
	filePerm = 0o644
	stateDir = ".state"
)

// Storage implements storage functions for a POSIX filesystem.
// It leverages the POSIX atomic operations.
type Storage struct {
	sync.Mutex
	path  string
	queue *storage.Queue

	cpFile *os.File

	curSize uint64
	newCP   options.NewCPFunc

	entriesPath options.EntriesPathFunc
}

// NewTreeFunc is the signature of a function which receives information about newly integrated trees.
type NewTreeFunc func(size uint64, root []byte) error

// New creates a new POSIX storage.
// - path is a directory in which the log should be stored
// - create must only be set when first creating the log, and will create the directory structure and an empty checkpoint
func New(ctx context.Context, path string, create bool, opts ...func(*options.StorageOptions)) (*Storage, error) {
	opt := storage.ResolveStorageOptions(opts...)

	r := &Storage{
		path:        path,
		newCP:       opt.NewCP,
		entriesPath: opt.EntriesPath,
	}
	if err := r.initialise(create); err != nil {
		return nil, err
	}
	r.queue = storage.NewQueue(ctx, opt.BatchMaxAge, opt.BatchMaxSize, r.sequenceBatch)

	go func(ctx context.Context, i time.Duration) {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(i):
				if err := r.publishCheckpoint(i); err != nil {
					klog.Warningf("publishCheckpoint: %v", err)
				}
			}
		}
	}(ctx, opt.CheckpointInterval)

	return r, nil
}

func (s *Storage) curTree() (uint64, error) {
	rawCp, err := s.readCheckpoint()
	if err != nil {
		return 0, fmt.Errorf("failed to read log checkpoint: %q", err)
	}
	_, size, err := parse.CheckpointUnsafe(rawCp)
	return size, err
}

// lockCP places a POSIX advisory lock for the checkpoint.
// Note that a) this is advisory, and b) we use an adjacent file to the checkpoint
// (`checkpoint.lock`) to avoid inherent brittleness of the `fcntrl` API (*any* `Close`
// operation on this file (even if it's a different FD) from this PID, or overwriting
// of the file by *any* process breaks the lock.)
func (s *Storage) lockCP() error {
	f, err := lockFile(filepath.Join(s.path, stateDir, layout.CheckpointPath+".lock"))
	if err != nil {
		return err
	}
	s.cpFile = f
	return nil
}

// unlockCP unlocks the `checkpoint.lock` file.
func (s *Storage) unlockCP() error {
	if s.cpFile == nil {
		panic(errors.New("not locked"))
	}
	if err := unlockFile(s.cpFile); err != nil {
		return err
	}
	s.cpFile = nil
	return nil
}

// lockFile creates/opens the file at the specified path, and flocks it.
// Once locked, the caller can use the returned file handle to perform whatever
// operations are necessary, before calling unlockFile with the handle.
func lockFile(p string) (*os.File, error) {
	f, err := os.OpenFile(p, syscall.O_CREAT|syscall.O_RDWR|syscall.O_CLOEXEC, filePerm)
	if err != nil {
		return nil, err
	}

	flockT := syscall.Flock_t{
		Type:   syscall.F_WRLCK,
		Whence: io.SeekStart,
		Start:  0,
		Len:    0,
	}
	// Keep trying until we manage to get an answer without being interrupted.
	for {
		if err := syscall.FcntlFlock(f.Fd(), syscall.F_SETLKW, &flockT); err != syscall.EINTR {
			return f, err
		}
	}
}

// unlockFile simply closes the locked file.
// This is enough to release the flock taken in lockFile previously.
func unlockFile(f *os.File) error {
	return f.Close()
}

// Add takes an entry and queues it for inclusion in the log.
// Upon placing the entry in an in-memory queue to be sequenced, it returns a future that will
// evaluate to either the sequence number assigned to this entry, or an error.
// This future is made available when the entry is queued. Any further calls to Add after
// this returns will guarantee that the later entry appears later in the log than any
// earlier entries. Concurrent calls to Add are supported, but the order they are queued and
// thus included in the log is non-deterministic.
//
// If the future resolves to a non-error state then it means that the entry is both
// sequenced and integrated into the log. This means that a checkpoint will be available
// that commits to this entry.
//
// It is recommended that the caller keeps the process running until all futures returned
// by this method have successfully evaluated. Terminating earlier than this will likely
// mean that some of the entries added are not committed to by a checkpoint, and thus are
// not considered to be in the log.
func (s *Storage) Add(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
	return s.queue.Add(ctx, e)
}

func (s *Storage) ReadCheckpoint(_ context.Context) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.path, layout.CheckpointPath))
}

// ReadEntryBundle retrieves the Nth entries bundle for a log of the given size.
func (s *Storage) ReadEntryBundle(_ context.Context, index, logSize uint64) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.path, s.entriesPath(index, logSize)))
}

func (s *Storage) ReadTile(_ context.Context, level, index, logSize uint64) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.path, layout.TilePath(level, index, logSize)))
}

// sequenceBatch writes the entries from the provided batch into the entry bundle files of the log.
//
// This func starts filling entries bundles at the next available slot in the log, ensuring that the
// sequenced entries are contiguous from the zeroth entry (i.e left-hand dense).
// We try to minimise the number of partially complete entry bundles by writing entries in chunks rather
// than one-by-one.
func (s *Storage) sequenceBatch(ctx context.Context, entries []*tessera.Entry) error {
	// Double locking:
	// - The mutex `Lock()` ensures that multiple concurrent calls to this function within a task are serialised.
	// - The POSIX `LockCP()` ensures that distinct tasks are serialised.
	s.Lock()
	if err := s.lockCP(); err != nil {
		panic(err)
	}
	defer func() {
		if err := s.unlockCP(); err != nil {
			panic(err)
		}
		s.Unlock()
	}()

	size, err := s.curTree()
	if err != nil {
		return err
	}
	s.curSize = size
	klog.V(1).Infof("Sequencing from %d", s.curSize)

	if len(entries) == 0 {
		return nil
	}
	currTile := &bytes.Buffer{}
	newSize := s.curSize + uint64(len(entries))
	seq := s.curSize
	bundleIndex, entriesInBundle := seq/uint64(256), seq%uint64(256)
	if entriesInBundle > 0 {
		// If the latest bundle is partial, we need to read the data it contains in for our newer, larger, bundle.
		part, err := s.ReadEntryBundle(ctx, bundleIndex, s.curSize)
		if err != nil {
			return err
		}
		if _, err := currTile.Write(part); err != nil {
			return fmt.Errorf("failed to write partial bundle into buffer: %v", err)
		}
	}
	writeBundle := func(bundleIndex uint64) error {
		bf := filepath.Join(s.path, s.entriesPath(bundleIndex, newSize))
		if err := os.MkdirAll(filepath.Dir(bf), dirPerm); err != nil {
			return fmt.Errorf("failed to make entries directory structure: %w", err)
		}
		if err := createExclusive(bf, currTile.Bytes()); err != nil {
			if !errors.Is(err, os.ErrExist) {
				return err
			}
		}
		return nil
	}

	seqEntries := make([]storage.SequencedEntry, 0, len(entries))
	// Add new entries to the bundle
	for i, e := range entries {
		bundleData := e.MarshalBundleData(seq + uint64(i))
		if _, err := currTile.Write(bundleData); err != nil {
			return fmt.Errorf("failed to write entry %d to currTile: %v", i, err)
		}
		seqEntries = append(seqEntries, storage.SequencedEntry{
			BundleData: bundleData,
			LeafHash:   e.LeafHash(),
		})

		entriesInBundle++
		if entriesInBundle == uint64(256) {
			//  This bundle is full, so we need to write it out...
			// ... and prepare the next entry bundle for any remaining entries in the batch
			if err := writeBundle(bundleIndex); err != nil {
				return err
			}
			bundleIndex++
			entriesInBundle = 0
			currTile = &bytes.Buffer{}
		}
	}
	// If we have a partial bundle remaining once we've added all the entries from the batch,
	// this needs writing out too.
	if entriesInBundle > 0 {
		if err := writeBundle(bundleIndex); err != nil {
			return err
		}
	}

	// For simplicity, in-line the integration of these new entries into the Merkle structure too.
	if err := s.doIntegrate(ctx, seq, seqEntries); err != nil {
		klog.Errorf("Integrate failed: %v", err)
		return err
	}
	return nil
}

// doIntegrate handles integrating new entries into the log, and updating the checkpoint.
func (s *Storage) doIntegrate(ctx context.Context, fromSeq uint64, entries []storage.SequencedEntry) error {
	getTiles := func(ctx context.Context, tileIDs []storage.TileID, treeSize uint64) ([]*api.HashTile, error) {
		n, err := s.readTiles(ctx, tileIDs, treeSize)
		if err != nil {
			return nil, fmt.Errorf("getTiles: %w", err)
		}
		return n, nil
	}

	newSize, newRoot, tiles, err := storage.Integrate(ctx, getTiles, fromSeq, entries)
	if err != nil {
		klog.Errorf("Integrate: %v", err)
		return fmt.Errorf("Integrate: %v", err)
	}
	for k, v := range tiles {
		if err := s.storeTile(ctx, uint64(k.Level), k.Index, newSize, v); err != nil {
			return fmt.Errorf("failed to set tile(%v): %v", k, err)
		}
	}

	klog.Infof("New CP: %d, %x", newSize, newRoot)
	if s.newCP != nil {
		cpRaw, err := s.newCP(newSize, newRoot)
		if err != nil {
			return fmt.Errorf("newCP: %v", err)
		}
		if err := s.writeCheckpoint(cpRaw); err != nil {
			return fmt.Errorf("failed to write new checkpoint: %v", err)
		}
	}

	return nil
}

func (s *Storage) readTiles(ctx context.Context, tileIDs []storage.TileID, treeSize uint64) ([]*api.HashTile, error) {
	r := make([]*api.HashTile, 0, len(tileIDs))
	for _, id := range tileIDs {
		t, err := s.readTile(ctx, id.Level, id.Index, treeSize)
		if err != nil {
			return nil, err
		}
		r = append(r, t)
	}
	return r, nil
}

// readTile returns the parsed tile at the given tile-level and tile-index.
// If no complete tile exists at that location, it will attempt to find a
// partial tile for the given tree size at that location.
func (s *Storage) readTile(ctx context.Context, level, index, logSize uint64) (*api.HashTile, error) {
	t, err := s.ReadTile(ctx, level, index, logSize)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// We'll signal to higher levels that it wasn't found by retuning a nil for this tile.
			return nil, nil
		}
		return nil, err
	}

	var tile api.HashTile
	if err := tile.UnmarshalText(t); err != nil {
		return nil, fmt.Errorf("failed to parse tile: %w", err)
	}
	return &tile, nil
}

// storeTile writes a tile out to disk.
// Fully populated tiles are stored at the path corresponding to the level &
// index parameters, partially populated (i.e. right-hand edge) tiles are
// stored with a .xx suffix where xx is the number of "tile leaves" in hex.
func (s *Storage) storeTile(_ context.Context, level, index, logSize uint64, tile *api.HashTile) error {
	tileSize := uint64(len(tile.Nodes))
	klog.V(2).Infof("StoreTile: level %d index %x ts: %x", level, index, tileSize)
	if tileSize == 0 || tileSize > 256 {
		return fmt.Errorf("tileSize %d must be > 0 and <= 256", tileSize)
	}
	t, err := tile.MarshalText()
	if err != nil {
		return fmt.Errorf("failed to marshal tile: %w", err)
	}

	tPath := filepath.Join(s.path, layout.TilePath(level, index, logSize))
	tDir := filepath.Dir(tPath)
	if err := os.MkdirAll(tDir, dirPerm); err != nil {
		return fmt.Errorf("failed to create directory %q: %w", tDir, err)
	}

	if err := createExclusive(tPath, t); err != nil {
		return err
	}

	if tileSize == 256 {
		partials, err := filepath.Glob(fmt.Sprintf("%s.p/*", tPath))
		if err != nil {
			return fmt.Errorf("failed to list partial tiles for clean up; %w", err)
		}
		// Clean up old partial tiles by symlinking them to the new full tile.
		for _, p := range partials {
			klog.V(2).Infof("relink partial %s to %s", p, tPath)
			// We have to do a little dance here to get POSIX atomicity:
			// 1. Create a new temporary symlink to the full tile
			// 2. Rename the temporary symlink over the top of the old partial tile
			tmp := fmt.Sprintf("%s.link", tPath)
			_ = os.Remove(tmp)
			if err := os.Symlink(tPath, tmp); err != nil {
				return fmt.Errorf("failed to create temp link to full tile: %w", err)
			}
			if err := os.Rename(tmp, p); err != nil {
				return fmt.Errorf("failed to rename temp link over partial tile: %w", err)
			}
		}
	}

	return nil
}

// initialise ensures that the storage location is valid by loading the checkpoint from this location.
// If `create` is set to true, then this will first ensure that the directory path is created, and
// an empty checkpoint is created in this directory.
func (s *Storage) initialise(create bool) error {
	if create {
		// Create the directory structure and write out an empty checkpoint
		klog.Infof("Initializing directory for POSIX log at %q (this should only happen ONCE per log!)", s.path)
		if err := os.MkdirAll(filepath.Join(s.path, stateDir), dirPerm); err != nil {
			return fmt.Errorf("failed to create log directory: %q", err)
		}
		n, err := s.newCP(0, rfc6962.DefaultHasher.EmptyRoot())
		if err != nil {
			return fmt.Errorf("failed to sign empty checkpoint: %v", err)
		}
		if err := s.writeCheckpoint(n); err != nil {
			return fmt.Errorf("failed to write empty checkpoint: %v", err)
		}
	}
	curSize, err := s.curTree()
	if err != nil {
		return fmt.Errorf("failed to load checkpoint for log: %v", err)
	}
	s.curSize = curSize

	return nil
}

// writeCheckpoint stores a raw log checkpoint on disk.
func (s *Storage) writeCheckpoint(newCPRaw []byte) error {
	if err := createExclusive(filepath.Join(s.path, stateDir, layout.CheckpointPath), newCPRaw); err != nil {
		return fmt.Errorf("failed to create private checkpoint file: %w", err)
	}
	return nil
}

// readcheckpoint returns the latest stored private checkpoint.
func (s *Storage) readCheckpoint() ([]byte, error) {
	return os.ReadFile(filepath.Join(s.path, stateDir, layout.CheckpointPath))
}

// publishCheckpoint makes the most recently checkpoint from the stateDir available at the
// (public) tlog-tiles specified Checkpoint location.
func (s *Storage) publishCheckpoint(minStaleness time.Duration) error {
	// Lock the destination "published" checkpoint location:
	cpF, err := lockFile(filepath.Join(s.path, layout.CheckpointPath))
	if err != nil {
		return fmt.Errorf("lockFile(%s): %v", layout.CheckpointPath, err)
	}
	defer func() {
		if err := cpF.Close(); err != nil {
			klog.Warningf("close(%s): %v", layout.CheckpointPath, err)
		}
	}()

	info, err := cpF.Stat()
	if err != nil {
		return fmt.Errorf("stat(%s): %v", layout.CheckpointPath, err)
	}
	if d := time.Since(info.ModTime()); d < minStaleness {
		klog.V(1).Infof("publishCheckpoint: skipping publish because previous checkpoint published %v ago, less than %v", d, minStaleness)
		return nil
	}
	cpRaw, err := os.ReadFile(filepath.Join(s.path, stateDir, layout.CheckpointPath))
	if err != nil {
		return fmt.Errorf("read state checkpoint: %v", err)
	}
	// This write can fail, e.g. if the storage is full, but this is not fatal:
	// - the source of truth is the checkpoint in the state directory
	// - we'll try again after another publish interval.
	if _, err := cpF.Write(cpRaw); err != nil {
		return fmt.Errorf("write(%s): %v", layout.CheckpointPath, err)
	}
	klog.Infof("Published latest checkpoint")
	return nil
}

// createExclusive creates a file at the given path and name before writing the data in d to it.
// It will error if the file already exists, or it's unable to fully write the
// data & close the file.
func createExclusive(f string, d []byte) error {
	tmpName := f + ".temp"
	if err := os.WriteFile(tmpName, d, filePerm); err != nil {
		return fmt.Errorf("unable to write data to temporary file: %w", err)
	}
	if err := os.Rename(tmpName, f); err != nil {
		return err
	}
	return nil
}
