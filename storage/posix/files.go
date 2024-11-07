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

	"github.com/transparency-dev/merkle/rfc6962"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	storage "github.com/transparency-dev/trillian-tessera/storage/internal"
	"k8s.io/klog/v2"
)

const (
	dirPerm  = 0o755
	filePerm = 0o644
)

// Storage implements storage functions for a POSIX filesystem.
// It leverages the POSIX atomic operations.
type Storage struct {
	sync.Mutex
	path  string
	queue *storage.Queue

	cpFile *os.File

	curSize uint64
	newCP   tessera.NewCPFunc
	parseCP tessera.ParseCPFunc

	entriesPath tessera.EntriesPathFunc
}

// NewTreeFunc is the signature of a function which receives information about newly integrated trees.
type NewTreeFunc func(size uint64, root []byte) error

// New creates a new POSIX storage.
// - path is a directory in which the log should be stored
// - create must only be set when first creating the log, and will create the directory structure and an empty checkpoint
func New(ctx context.Context, path string, create bool, opts ...func(*tessera.StorageOptions)) (*Storage, error) {
	opt := tessera.ResolveStorageOptions(opts...)

	r := &Storage{
		path:        path,
		newCP:       opt.NewCP,
		parseCP:     opt.ParseCP,
		entriesPath: opt.EntriesPath,
	}
	if err := r.initialise(create); err != nil {
		return nil, err
	}
	r.queue = storage.NewQueue(ctx, opt.BatchMaxAge, opt.BatchMaxSize, r.sequenceBatch)

	return r, nil
}

func (s *Storage) curTree() (uint64, []byte, error) {
	cpRaw, err := readCheckpoint(s.path)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to read log checkpoint: %q", err)
	}
	cp, err := s.parseCP(cpRaw)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to parse Checkpoint: %q", err)
	}
	return cp.Size, cp.Hash, nil
}

// lockCP places a POSIX advisory lock for the checkpoint.
// Note that a) this is advisory, and b) we use an adjacent file to the checkpoint
// (`checkpoint.lock`) to avoid inherent brittleness of the `fcntrl` API (*any* `Close`
// operation on this file (even if it's a different FD) from this PID, or overwriting
// of the file by *any* process breaks the lock.)
func (s *Storage) lockCP() error {
	var err error
	if s.cpFile != nil {
		panic("not unlocked")
	}
	s.cpFile, err = os.OpenFile(filepath.Join(s.path, layout.CheckpointPath+".lock"), syscall.O_CREAT|syscall.O_RDWR|syscall.O_CLOEXEC, filePerm)
	if err != nil {
		return err
	}

	flockT := syscall.Flock_t{
		Type:   syscall.F_WRLCK,
		Whence: io.SeekStart,
		Start:  0,
		Len:    0,
	}
	for {
		if err := syscall.FcntlFlock(s.cpFile.Fd(), syscall.F_SETLKW, &flockT); err != syscall.EINTR {
			return err
		}
	}
}

// unlockCP unlocks the `checkpoint.lock` file.
func (s *Storage) unlockCP() error {
	if s.cpFile == nil {
		panic(errors.New("not locked"))
	}
	if err := s.cpFile.Close(); err != nil {
		return err
	}
	s.cpFile = nil
	return nil
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

	size, _, err := s.curTree()
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
	tb := storage.NewTreeBuilder(func(ctx context.Context, tileIDs []storage.TileID, treeSize uint64) ([]*api.HashTile, error) {
		n, err := s.readTiles(ctx, tileIDs, treeSize)
		if err != nil {
			return nil, fmt.Errorf("getTiles: %w", err)
		}
		return n, nil
	})

	newSize, newRoot, tiles, err := tb.Integrate(ctx, fromSeq, entries)
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
		if err := writeCheckpoint(s.path, cpRaw); err != nil {
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
		if err := os.MkdirAll(s.path, dirPerm); err != nil {
			return fmt.Errorf("failed to create log directory: %q", err)
		}
		n, err := s.newCP(0, rfc6962.DefaultHasher.EmptyRoot())
		if err != nil {
			return fmt.Errorf("failed to sign empty checkpoint: %v", err)
		}
		if err := writeCheckpoint(s.path, n); err != nil {
			return fmt.Errorf("failed to write empty checkpoint: %v", err)
		}
	}
	curSize, _, err := s.curTree()
	if err != nil {
		return fmt.Errorf("failed to load checkpoint for log: %v", err)
	}
	s.curSize = curSize

	return nil
}

// writeCheckpoint stores a raw log checkpoint on disk.
func writeCheckpoint(path string, newCPRaw []byte) error {
	if err := createExclusive(filepath.Join(path, layout.CheckpointPath), newCPRaw); err != nil {
		return fmt.Errorf("failed to create checkpoint file: %w", err)
	}
	return nil
}

// readcheckpoint returns the latest stored checkpoint.
func readCheckpoint(path string) ([]byte, error) {
	return os.ReadFile(filepath.Join(path, layout.CheckpointPath))
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
