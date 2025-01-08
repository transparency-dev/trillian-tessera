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
	"encoding/json"
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
	"github.com/transparency-dev/trillian-tessera/storage/internal"
	"k8s.io/klog/v2"
)

const (
	dirPerm  = 0o755
	filePerm = 0o644
	stateDir = ".state"

	minCheckpointInterval = time.Second
)

// Storage implements storage functions for a POSIX filesystem.
// It leverages the POSIX atomic operations.
type Storage struct {
	mu    sync.Mutex
	path  string
	queue *storage.Queue

	curSize uint64
	newCP   options.NewCPFunc

	cpUpdated chan struct{}

	entriesPath options.EntriesPathFunc
}

// NewTreeFunc is the signature of a function which receives information about newly integrated trees.
type NewTreeFunc func(size uint64, root []byte) error

// New creates a new POSIX storage.
// - path is a directory in which the log should be stored
// - create must only be set when first creating the log, and will create the directory structure and an empty checkpoint
func New(ctx context.Context, path string, create bool, opts ...func(*options.StorageOptions)) (tessera.Driver, error) {
	opt := storage.ResolveStorageOptions(opts...)
	if opt.CheckpointInterval < minCheckpointInterval {
		return nil, fmt.Errorf("requested CheckpointInterval (%v) is less than minimum permitted %v", opt.CheckpointInterval, minCheckpointInterval)
	}

	r := &Storage{
		path:        path,
		newCP:       opt.NewCP,
		entriesPath: opt.EntriesPath,
		cpUpdated:   make(chan struct{}),
	}
	if err := r.initialise(create); err != nil {
		return nil, err
	}
	r.queue = storage.NewQueue(ctx, opt.BatchMaxAge, opt.BatchMaxSize, r.sequenceBatch)

	go func(ctx context.Context, i time.Duration) {
		t := time.NewTicker(i)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-r.cpUpdated:
			case <-t.C:
			}
			if err := r.publishCheckpoint(i); err != nil {
				klog.Warningf("publishCheckpoint: %v", err)
			}
		}
	}(ctx, opt.CheckpointInterval)

	return r, nil
}

// lockFile creates/opens a lock file at the specified path, and flocks it.
// Once locked, the caller perform whatever operations are necessary, before
// calling the returned function to unlock it.
//
// Note that a) this is advisory, and b) should use an non-API specified file
// (e.g. <something>.lock>) to avoid inherent brittleness of the `fcntrl` API
// (*any* `Close` operation on this file (even if it's a different FD) from
// this PID, or overwriting of the file by *any* process breaks the lock.)
func lockFile(p string) (func() error, error) {
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
			return f.Close, err
		}
	}
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
func (s *Storage) ReadEntryBundle(_ context.Context, index uint64, p uint8) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.path, s.entriesPath(index, p)))
}

func (s *Storage) ReadTile(_ context.Context, level, index uint64, p uint8) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.path, layout.TilePath(level, index, p)))
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
	// - The POSIX `lockForTreeUpdate()` ensures that distinct tasks are serialised.
	s.mu.Lock()
	unlock, err := lockFile(filepath.Join(s.path, stateDir, "treeState.lock"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := unlock(); err != nil {
			panic(err)
		}
		s.mu.Unlock()
	}()

	size, _, err := s.readTreeState()
	if err != nil {
		return err
	}
	s.curSize = size
	klog.V(1).Infof("Sequencing from %d", s.curSize)

	if len(entries) == 0 {
		return nil
	}
	currTile := &bytes.Buffer{}
	seq := s.curSize
	bundleIndex, entriesInBundle := seq/layout.EntryBundleWidth, seq%layout.EntryBundleWidth
	if entriesInBundle > 0 {
		// If the latest bundle is partial, we need to read the data it contains in for our newer, larger, bundle.
		part, err := s.ReadEntryBundle(ctx, bundleIndex, uint8(s.curSize%layout.EntryBundleWidth))
		if err != nil {
			return err
		}
		if _, err := currTile.Write(part); err != nil {
			return fmt.Errorf("failed to write partial bundle into buffer: %v", err)
		}
	}
	writeBundle := func(bundleIndex uint64, partialSize uint8) error {
		bf := filepath.Join(s.path, s.entriesPath(bundleIndex, partialSize))
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
		if entriesInBundle == layout.EntryBundleWidth {
			//  This bundle is full, so we need to write it out...
			// ... and prepare the next entry bundle for any remaining entries in the batch
			if err := writeBundle(bundleIndex, 0); err != nil {
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
		// This check should be redundant since this is [currently] checked above, but an overflow around the uint8 below could
		// potentially be bad news if that check was broken/defeated as we'd be writing invalid bundle data, so do a belt-and-braces
		// check and bail if need be.
		if entriesInBundle > layout.EntryBundleWidth {
			return fmt.Errorf("logic error: entriesInBundle(%d) > max bundle size %d", entriesInBundle, layout.EntryBundleWidth)
		}
		if err := writeBundle(bundleIndex, uint8(entriesInBundle)); err != nil {
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

// doIntegrate handles integrating new entries into the log, and updating the tree state.
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

	klog.Infof("New tree state: %d, %x", newSize, newRoot)
	if err := s.writeTreeState(newSize, newRoot); err != nil {
		return fmt.Errorf("failed to write new tree state: %v", err)
	}

	return nil
}

func (s *Storage) readTiles(ctx context.Context, tileIDs []storage.TileID, treeSize uint64) ([]*api.HashTile, error) {
	r := make([]*api.HashTile, 0, len(tileIDs))
	for _, id := range tileIDs {
		t, err := s.readTile(ctx, id.Level, id.Index, layout.PartialTileSize(id.Level, id.Index, treeSize))
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
func (s *Storage) readTile(ctx context.Context, level, index uint64, p uint8) (*api.HashTile, error) {
	t, err := s.ReadTile(ctx, level, index, p)
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
	if tileSize == 0 || tileSize > layout.TileWidth {
		return fmt.Errorf("tileSize %d must be > 0 and <= %d", tileSize, layout.TileWidth)
	}
	t, err := tile.MarshalText()
	if err != nil {
		return fmt.Errorf("failed to marshal tile: %w", err)
	}

	tPath := filepath.Join(s.path, layout.TilePath(level, index, layout.PartialTileSize(level, index, logSize)))
	tDir := filepath.Dir(tPath)
	if err := os.MkdirAll(tDir, dirPerm); err != nil {
		return fmt.Errorf("failed to create directory %q: %w", tDir, err)
	}

	if err := createExclusive(tPath, t); err != nil {
		return err
	}

	if tileSize == layout.TileWidth {
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
		if err := s.writeTreeState(0, rfc6962.DefaultHasher.EmptyRoot()); err != nil {
			return fmt.Errorf("failed to write tree-state checkpoint: %v", err)
		}
		if err := s.publishCheckpoint(0); err != nil {
			return fmt.Errorf("failed to publish checkpoint: %v", err)
		}
	}
	curSize, _, err := s.readTreeState()
	if err != nil {
		return fmt.Errorf("failed to load checkpoint for log: %v", err)
	}
	s.curSize = curSize

	return nil
}

type treeState struct {
	Size uint64 `json:"size"`
	Root []byte `json:"root"`
}

// writeTreeState stores the current tree size and root hash on disk.
func (s *Storage) writeTreeState(size uint64, root []byte) error {
	raw, err := json.Marshal(treeState{Size: size, Root: root})
	if err != nil {
		return fmt.Errorf("Marshal: %v", err)
	}

	if err := createExclusive(filepath.Join(s.path, stateDir, "treeState"), raw); err != nil {
		return fmt.Errorf("failed to create private tree state file: %w", err)
	}
	// Notify that we know for sure there's a new checkpoint, but don't block if there's already
	// an outstanding notification in the channel.
	select {
	case s.cpUpdated <- struct{}{}:
	default:
	}
	return nil
}

// readTreeState reads and returns the currently stored tree state.
func (s *Storage) readTreeState() (uint64, []byte, error) {
	p := filepath.Join(s.path, stateDir, "treeState")
	raw, err := os.ReadFile(p)
	if err != nil {
		return 0, nil, fmt.Errorf("ReadFile(%q): %v", p, err)
	}
	ts := &treeState{}
	if err := json.Unmarshal(raw, ts); err != nil {
		return 0, nil, fmt.Errorf("Unmarshal: %v", err)
	}
	return ts.Size, ts.Root, nil
}

// publishCheckpoint checks whether the currently published checkpoint (if any) is more than
// minStaleness old, and, if so, creates and published a fresh checkpoint from the current
// stored tree state.
func (s *Storage) publishCheckpoint(minStaleness time.Duration) error {
	// Lock the destination "published" checkpoint location:
	lockPath := filepath.Join(s.path, stateDir, "publish.lock")
	unlock, err := lockFile(lockPath)
	if err != nil {
		return fmt.Errorf("lockFile(%s): %v", lockPath, err)
	}
	defer func() {
		if err := unlock(); err != nil {
			klog.Warningf("unlock(%s): %v", lockPath, err)
		}
	}()

	info, err := os.Stat(filepath.Join(s.path, layout.CheckpointPath))
	if errors.Is(err, os.ErrNotExist) {
		klog.V(1).Infof("No checkpoint exists, publishing")
	} else if err != nil {
		return fmt.Errorf("stat(%s): %v", layout.CheckpointPath, err)
	} else {
		if d := time.Since(info.ModTime()); d < minStaleness {
			klog.V(1).Infof("publishCheckpoint: skipping publish because previous checkpoint published %v ago, less than %v", d, minStaleness)
			return nil
		}
	}
	size, root, err := s.readTreeState()
	if err != nil {
		return fmt.Errorf("readTreeState: %v", err)
	}
	cpRaw, err := s.newCP(size, root)
	if err != nil {
		return fmt.Errorf("newCP: %v", err)
	}

	if err := createExclusive(filepath.Join(s.path, layout.CheckpointPath), cpRaw); err != nil {
		return fmt.Errorf("createExclusive(%s): %v", layout.CheckpointPath, err)
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
