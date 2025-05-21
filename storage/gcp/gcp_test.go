// Copyright 2024 The Tessera authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gcp

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spannertest"
	gcs "cloud.google.com/go/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/transparency-dev/tessera"
	"github.com/transparency-dev/tessera/api"
	"github.com/transparency-dev/tessera/api/layout"
	"github.com/transparency-dev/tessera/core"
	storage "github.com/transparency-dev/tessera/storage/internal"
)

func newSpannerDB(t *testing.T) func() {
	t.Helper()
	srv, err := spannertest.NewServer("localhost:0")
	if err != nil {
		t.Fatalf("Failed to set up test spanner: %v", err)
	}
	if err := os.Setenv("SPANNER_EMULATOR_HOST", srv.Addr); err != nil {
		t.Fatalf("Setenv: %v", err)
	}
	return srv.Close

}

func TestSpannerSequencerAssignEntries(t *testing.T) {
	ctx := context.Background()
	close := newSpannerDB(t)
	defer close()

	seq, err := newSpannerCoordinator(ctx, "projects/p/instances/i/databases/d", 1000)
	if err != nil {
		t.Fatalf("newSpannerCoordinator: %v", err)
	}

	want := uint64(0)
	for chunks := range 10 {
		entries := []*core.Entry{}
		for i := range 10 + chunks {
			entries = append(entries, core.NewEntry(fmt.Appendf(nil, "item %d/%d", chunks, i)))
		}
		if err := seq.assignEntries(ctx, entries); err != nil {
			t.Fatalf("assignEntries: %v", err)
		}
		for i, e := range entries {
			if got := *e.Index(); got != want {
				t.Errorf("Chunk %d entry %d got seq %d, want %d", chunks, i, got, want)
			}
			want++
		}
	}
}

func TestSpannerSequencerPushback(t *testing.T) {
	ctx := context.Background()

	for _, test := range []struct {
		name           string
		threshold      uint64
		initialEntries int
		wantPushback   bool
	}{
		{
			name:           "no pushback: num < threshold",
			threshold:      10,
			initialEntries: 5,
		},
		{
			name:           "no pushback: num = threshold",
			threshold:      10,
			initialEntries: 10,
		},
		{
			name:           "pushback: initial > threshold",
			threshold:      10,
			initialEntries: 15,
			wantPushback:   true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			close := newSpannerDB(t)
			defer close()

			seq, err := newSpannerCoordinator(ctx, "projects/p/instances/i/databases/d", test.threshold)
			if err != nil {
				t.Fatalf("newSpannerCoordinator: %v", err)
			}
			// Set up the test scenario with the configured number of initial outstanding entries
			entries := []*core.Entry{}
			for i := range test.initialEntries {
				entries = append(entries, core.NewEntry(fmt.Appendf(nil, "initial item %d", i)))
			}
			if err := seq.assignEntries(ctx, entries); err != nil {
				t.Fatalf("initial assignEntries: %v", err)
			}

			// Now perform the test with a single additional entry to check for pushback
			entries = []*core.Entry{core.NewEntry([]byte("additional"))}
			err = seq.assignEntries(ctx, entries)
			if gotPushback := errors.Is(err, tessera.ErrPushback); gotPushback != test.wantPushback {
				t.Fatalf("assignEntries: got pushback %t (%v), want pushback: %t", gotPushback, err, test.wantPushback)
			} else if !gotPushback && err != nil {
				t.Fatalf("assignEntries: %v", err)
			}
		})
	}
}

func TestSpannerSequencerRoundTrip(t *testing.T) {
	ctx := context.Background()
	close := newSpannerDB(t)
	defer close()

	s, err := newSpannerCoordinator(ctx, "projects/p/instances/i/databases/d", 1000)
	if err != nil {
		t.Fatalf("newSpannerCoordinator: %v", err)
	}

	seq := 0
	wantEntries := []storage.SequencedEntry{}
	for chunks := range 10 {
		entries := []*core.Entry{}
		for range 10 + chunks {
			e := core.NewEntry(fmt.Appendf(nil, "item %d", seq))
			entries = append(entries, e)
			wantEntries = append(wantEntries, storage.SequencedEntry{
				BundleData: e.MarshalBundleData(uint64(seq)),
				LeafHash:   e.LeafHash(),
			})
			seq++
		}
		if err := s.assignEntries(ctx, entries); err != nil {
			t.Fatalf("assignEntries: %v", err)
		}
	}

	seenIdx := uint64(0)
	f := func(_ context.Context, fromSeq uint64, entries []storage.SequencedEntry) ([]byte, error) {
		if fromSeq != seenIdx {
			return nil, fmt.Errorf("f called with fromSeq %d, want %d", fromSeq, seenIdx)
		}
		for i, e := range entries {

			if got, want := e, wantEntries[i]; !reflect.DeepEqual(got, want) {
				return nil, fmt.Errorf("entry %d+%d != %d", fromSeq, i, seenIdx)
			}
			seenIdx++
		}
		return fmt.Appendf(nil, "root<%d>", seenIdx), nil
	}

	more, err := s.consumeEntries(ctx, 7, f, false)
	if err != nil {
		t.Errorf("consumeEntries: %v", err)
	}
	if !more {
		t.Errorf("more: false, expected true")
	}
}

func TestCheckDataCompatibility(t *testing.T) {
	ctx := context.Background()
	close := newSpannerDB(t)
	defer close()

	s, err := newSpannerCoordinator(ctx, "projects/p/instances/i/databases/d", 1000)
	if err != nil {
		t.Fatalf("newSpannerCoordinator: %v", err)
	}

	for _, test := range []struct {
		desc    string
		dbV     int64
		wantErr bool
	}{
		{
			desc: "versions match",
			dbV:  SchemaCompatibilityVersion,
		},
		{
			desc:    "data < library",
			dbV:     SchemaCompatibilityVersion + 1,
			wantErr: true,
		},
		{
			desc:    "data > library",
			dbV:     SchemaCompatibilityVersion - 1,
			wantErr: true,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			if _, err := s.dbPool.Apply(ctx, []*spanner.Mutation{spanner.InsertOrUpdate("Tessera", []string{"id", "compatibilityVersion"}, []any{0, test.dbV})}); err != nil {
				t.Fatalf("Failed for force schema version to %d: %v", test.dbV, err)
			}

			err := s.checkDataCompatibility(ctx)
			if gotErr := err != nil; test.wantErr != gotErr {
				t.Fatalf("checkDataCompatibility: %v, wantErr %t", err, test.wantErr)
			}
		})
	}
}

func makeTile(t *testing.T, size uint64) *api.HashTile {
	t.Helper()
	r := &api.HashTile{Nodes: make([][]byte, size)}
	for i := uint64(0); i < size; i++ {
		h := sha256.Sum256(fmt.Appendf(nil, "%d", i))
		r.Nodes[i] = h[:]
	}
	return r
}

func TestTileRoundtrip(t *testing.T) {
	ctx := context.Background()
	m := newMemObjStore()
	s := &logResourceStore{
		objStore: m,
	}

	for _, test := range []struct {
		name     string
		level    uint64
		index    uint64
		logSize  uint64
		tileSize uint64
	}{
		{
			name:     "ok",
			level:    0,
			index:    3 * layout.TileWidth,
			logSize:  3*layout.TileWidth + 20,
			tileSize: 20,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			wantTile := makeTile(t, test.tileSize)
			tRaw, err := wantTile.MarshalText()
			if err != nil {
				t.Fatalf("Failed to marshal tile: %v", err)
			}
			if err := s.setTile(ctx, test.level, test.index, layout.PartialTileSize(test.level, test.index, test.logSize), tRaw); err != nil {
				t.Fatalf("setTile: %v", err)
			}

			expPath := layout.TilePath(test.level, test.index, layout.PartialTileSize(test.level, test.index, test.logSize))
			_, ok := m.mem[expPath]
			if !ok {
				t.Fatalf("want tile at %v but found none", expPath)
			}

			got, err := s.getTiles(ctx, []storage.TileID{{Level: test.level, Index: test.index}}, test.logSize)
			if err != nil {
				t.Fatalf("getTile: %v", err)
			}
			if !cmp.Equal(got[0], wantTile) {
				t.Fatal("roundtrip returned different data")
			}
		})
	}
}

func makeBundle(t *testing.T, idx uint64, size int) []byte {
	t.Helper()
	r := &bytes.Buffer{}
	if size == 0 {
		size = layout.EntryBundleWidth
	}
	for i := range size {
		e := core.NewEntry(fmt.Appendf(nil, "%d:%d", idx, i))
		if _, err := r.Write(e.MarshalBundleData(uint64(i))); err != nil {
			t.Fatalf("MarshalBundleEntry: %v", err)
		}
	}
	return r.Bytes()
}

func TestBundleRoundtrip(t *testing.T) {
	ctx := context.Background()
	m := newMemObjStore()
	s := &logResourceStore{
		objStore:    m,
		entriesPath: layout.EntriesPath,
	}

	for _, test := range []struct {
		name       string
		index      uint64
		logSize    uint64
		bundleSize int
	}{
		{
			name:       "ok",
			index:      3 * layout.EntryBundleWidth,
			logSize:    3*layout.EntryBundleWidth + 20,
			bundleSize: 20,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			wantBundle := makeBundle(t, test.index, test.bundleSize)
			if err := s.setEntryBundle(ctx, test.index, uint8(test.bundleSize), wantBundle); err != nil {
				t.Fatalf("setEntryBundle: %v", err)
			}

			expPath := layout.EntriesPath(test.index, layout.PartialTileSize(0, test.index, test.logSize))
			_, ok := m.mem[expPath]
			if !ok {
				t.Fatalf("want bundle at %v but found none", expPath)
			}

			got, err := s.getEntryBundle(ctx, test.index, layout.PartialTileSize(0, test.index, test.logSize))
			if err != nil {
				t.Fatalf("getEntryBundle: %v", err)
			}
			if !cmp.Equal(got, wantBundle) {
				t.Fatal("roundtrip returned different data")
			}
		})
	}
}

func TestStreamEntries(t *testing.T) {
	ctx := context.Background()
	m := newMemObjStore()

	logSize1 := 12345
	logSize2 := 100045

	var logSize atomic.Uint64
	logSize.Store(uint64(logSize1))

	s := &LogReader{
		lrs: logResourceStore{
			objStore:    m,
			entriesPath: layout.EntriesPath,
		},
		integratedSize: func(context.Context) (uint64, error) { return logSize.Load(), nil },
	}

	// Populate entry bundles:
	// first to logSize1 (so we're sure we've got the partial bundle)
	for r, idx := logSize1, uint64(0); r > 0; idx++ {
		sz := min(r, layout.EntryBundleWidth)
		b := makeBundle(t, idx, sz)
		if err := s.lrs.setEntryBundle(ctx, idx, uint8(sz), b); err != nil {
			t.Fatalf("setEntryBundle(%d): %v", idx, err)
		}
		r -= sz
	}
	// Then on to logSize2
	for r, idx := logSize2, uint64(0); r > 0; idx++ {
		sz := min(r, layout.EntryBundleWidth)
		b := makeBundle(t, idx, sz)
		if err := s.lrs.setEntryBundle(ctx, idx, uint8(sz), b); err != nil {
			t.Fatalf("setEntryBundle(%d): %v", idx, err)
		}
		r -= sz
	}

	// Finally, try to stream all the bundles back.
	// We'll first try to stream up to logSize1, then when we reach it we'll
	// make the tree appear to grow to logSize2 to test resuming.
	seenEntries := uint64(0)
	next, stop := s.StreamEntries(ctx, 0)

	for {
		gotRI, _, gotErr := next()
		if gotErr != nil {
			if errors.Is(gotErr, tessera.ErrNoMoreEntries) {
				break
			}
			t.Fatalf("gotErr after %d: %v", seenEntries, gotErr)
		}
		if e := gotRI.Index*layout.EntryBundleWidth + uint64(gotRI.First); e != seenEntries {
			t.Fatalf("got idx %d, want %d", e, seenEntries)
		}
		seenEntries += uint64(gotRI.N)
		t.Logf("got RI %d / %d", gotRI.Index, seenEntries)

		switch seenEntries {
		case uint64(logSize1):
			// We've fetched all the entries from the original tree size, now we'll make
			// the tree appear to have grown to the final size.
			// The stream should start returning bundles again until we've consumed them all.
			t.Log("Reached logSize, growing tree")
			logSize.Store(uint64(logSize2))
			time.Sleep(time.Second)
		case uint64(logSize2):
			// We've seen all the entries we created, stop the iterator
			stop()
		}
	}
}

func TestPublishTree(t *testing.T) {
	ctx := context.Background()
	for _, test := range []struct {
		name            string
		publishInterval time.Duration
		attempts        []time.Duration
		wantUpdates     int
	}{
		{
			name:            "works ok",
			publishInterval: 100 * time.Millisecond,
			attempts:        []time.Duration{1 * time.Second},
			wantUpdates:     1,
		}, {
			name:            "too soon, skip update",
			publishInterval: 10 * time.Second,
			attempts:        []time.Duration{100 * time.Millisecond},
			wantUpdates:     0,
		}, {
			name:            "too soon, skip update, but recovers",
			publishInterval: 2 * time.Second,
			attempts:        []time.Duration{100 * time.Millisecond, 2 * time.Second},
			wantUpdates:     1,
		}, {
			name:            "many attempts, eventually one succeeds",
			publishInterval: 1 * time.Second,
			attempts:        []time.Duration{300 * time.Millisecond, 300 * time.Millisecond, 300 * time.Millisecond, 300 * time.Millisecond},
			wantUpdates:     1,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			closeDB := newSpannerDB(t)
			defer closeDB()
			s, err := newSpannerCoordinator(ctx, "projects/p/instances/i/databases/d", 1000)
			if err != nil {
				t.Fatalf("newSpannerCoordinator: %v", err)
			}
			defer s.dbPool.Close()

			m := newMemObjStore()
			storage := &Appender{
				logStore: &logResourceStore{
					objStore:    m,
					entriesPath: layout.EntriesPath,
				},
				sequencer: s,
				newCP: func(_ context.Context, size uint64, hash []byte) ([]byte, error) {
					return fmt.Appendf(nil, "%d/%x,", size, hash), nil
				},
			}
			// Call init so we've got a zero-sized checkpoint to work with.
			if err := storage.init(ctx); err != nil {
				t.Fatalf("storage.init: %v", err)
			}
			if err := s.publishTree(ctx, test.publishInterval, storage.publishCheckpoint); err != nil {
				t.Fatalf("publishTree: %v", err)
			}
			cpOld := []byte("bananas")
			if err := m.setObject(ctx, layout.CheckpointPath, cpOld, nil, "", ""); err != nil {
				t.Fatalf("setObject(bananas): %v", err)
			}
			updatesSeen := 0
			for _, d := range test.attempts {
				time.Sleep(d)
				if err := s.publishTree(ctx, test.publishInterval, storage.publishCheckpoint); err != nil {
					t.Fatalf("publishTree: %v", err)
				}
				cpNew, _, err := m.getObject(ctx, layout.CheckpointPath)
				if err != nil {
					t.Fatalf("getObject: %v", err)
				}
				if !bytes.Equal(cpOld, cpNew) {
					updatesSeen++
					cpOld = cpNew
				}
			}
			if updatesSeen != test.wantUpdates {
				t.Fatalf("Saw %d updates, want %d", updatesSeen, test.wantUpdates)
			}
		})
	}
}

type memObjStore struct {
	sync.RWMutex
	mem map[string][]byte
}

func newMemObjStore() *memObjStore {
	return &memObjStore{
		mem: make(map[string][]byte),
	}
}

func (m *memObjStore) getObject(_ context.Context, obj string) ([]byte, int64, error) {
	m.RLock()
	defer m.RUnlock()

	d, ok := m.mem[obj]
	if !ok {
		return nil, -1, fmt.Errorf("obj %q not found: %w", obj, gcs.ErrObjectNotExist)
	}
	return d, 1, nil
}

// TODO(phboneff): add content type tests
func (m *memObjStore) setObject(_ context.Context, obj string, data []byte, cond *gcs.Conditions, _, _ string) error {
	m.Lock()
	defer m.Unlock()

	d, ok := m.mem[obj]
	if cond != nil {
		if ok && cond.DoesNotExist {
			if !bytes.Equal(d, data) {
				return errors.New("precondition failed and data not identical")
			}
			return nil
		}
	}
	m.mem[obj] = data
	return nil
}
