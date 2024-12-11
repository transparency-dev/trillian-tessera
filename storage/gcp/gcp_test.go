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
	"testing"
	"time"

	"cloud.google.com/go/spanner/spannertest"
	"cloud.google.com/go/spanner/spansql"
	gcs "cloud.google.com/go/storage"
	"github.com/google/go-cmp/cmp"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	storage "github.com/transparency-dev/trillian-tessera/storage/internal"
)

func newSpannerDB(t *testing.T) func() {
	t.Helper()
	srv, err := spannertest.NewServer("localhost:0")
	if err != nil {
		t.Fatalf("Failed to set up test spanner: %v", err)
	}
	os.Setenv("SPANNER_EMULATOR_HOST", srv.Addr)
	dml, err := spansql.ParseDDL("", `
			CREATE TABLE SeqCoord (id INT64 NOT NULL, next INT64 NOT NULL,) PRIMARY KEY (id); 
			CREATE TABLE Seq (id INT64 NOT NULL, seq INT64 NOT NULL, v BYTES(MAX),) PRIMARY KEY (id, seq); 
			CREATE TABLE IntCoord (id INT64 NOT NULL, seq INT64 NOT NULL, rootHash BYTES(32) NOT NULL,) PRIMARY KEY (id); 
	`)
	if err != nil {
		t.Fatalf("Invalid DDL: %v", err)
	}
	if err := srv.UpdateDDL(dml); err != nil {
		t.Fatalf("Failed to create schema in test spanner: %v", err)
	}

	return srv.Close

}

func TestSpannerSequencerAssignEntries(t *testing.T) {
	ctx := context.Background()
	close := newSpannerDB(t)
	defer close()

	seq, err := newSpannerSequencer(ctx, "projects/p/instances/i/databases/d", 1000)
	if err != nil {
		t.Fatalf("newSpannerSequencer: %v", err)
	}

	want := uint64(0)
	for chunks := 0; chunks < 10; chunks++ {
		entries := []*tessera.Entry{}
		for i := 0; i < 10+chunks; i++ {
			entries = append(entries, tessera.NewEntry([]byte(fmt.Sprintf("item %d/%d", chunks, i))))
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

			seq, err := newSpannerSequencer(ctx, "projects/p/instances/i/databases/d", test.threshold)
			if err != nil {
				t.Fatalf("newSpannerSequencer: %v", err)
			}
			// Set up the test scenario with the configured number of initial outstanding entries
			entries := []*tessera.Entry{}
			for i := 0; i < test.initialEntries; i++ {
				entries = append(entries, tessera.NewEntry([]byte(fmt.Sprintf("initial item %d", i))))
			}
			if err := seq.assignEntries(ctx, entries); err != nil {
				t.Fatalf("initial assignEntries: %v", err)
			}

			// Now perform the test with a single additional entry to check for pushback
			entries = []*tessera.Entry{tessera.NewEntry([]byte("additional"))}
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

	s, err := newSpannerSequencer(ctx, "projects/p/instances/i/databases/d", 1000)
	if err != nil {
		t.Fatalf("newSpannerSequencer: %v", err)
	}

	seq := 0
	wantEntries := []storage.SequencedEntry{}
	for chunks := 0; chunks < 10; chunks++ {
		entries := []*tessera.Entry{}
		for i := 0; i < 10+chunks; i++ {
			e := tessera.NewEntry([]byte(fmt.Sprintf("item %d", seq)))
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
		return []byte(fmt.Sprintf("root<%d>", seenIdx)), nil
	}

	more, err := s.consumeEntries(ctx, 7, f, false)
	if err != nil {
		t.Errorf("consumeEntries: %v", err)
	}
	if !more {
		t.Errorf("more: false, expected true")
	}
}

func makeTile(t *testing.T, size uint64) *api.HashTile {
	t.Helper()
	r := &api.HashTile{Nodes: make([][]byte, size)}
	for i := uint64(0); i < size; i++ {
		h := sha256.Sum256([]byte(fmt.Sprintf("%d", i)))
		r.Nodes[i] = h[:]
	}
	return r
}

func TestTileRoundtrip(t *testing.T) {
	ctx := context.Background()
	m := newMemObjStore()
	s := &Storage{
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

func makeBundle(t *testing.T, size uint64) []byte {
	t.Helper()
	r := &bytes.Buffer{}
	for i := uint64(0); i < size; i++ {
		e := tessera.NewEntry([]byte(fmt.Sprintf("%d", i)))
		if _, err := r.Write(e.MarshalBundleData(i)); err != nil {
			t.Fatalf("MarshalBundleEntry: %v", err)
		}
	}
	return r.Bytes()
}

func TestBundleRoundtrip(t *testing.T) {
	ctx := context.Background()
	m := newMemObjStore()
	s := &Storage{
		objStore:    m,
		entriesPath: layout.EntriesPath,
	}

	for _, test := range []struct {
		name       string
		index      uint64
		logSize    uint64
		bundleSize uint64
	}{
		{
			name:       "ok",
			index:      3 * layout.EntryBundleWidth,
			logSize:    3*layout.EntryBundleWidth + 20,
			bundleSize: 20,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			wantBundle := makeBundle(t, test.bundleSize)
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

func TestPublishCheckpoint(t *testing.T) {
	ctx := context.Background()

	close := newSpannerDB(t)
	defer close()

	s, err := newSpannerSequencer(ctx, "projects/p/instances/i/databases/d", 1000)
	if err != nil {
		t.Fatalf("newSpannerSequencer: %v", err)
	}

	for _, test := range []struct {
		name            string
		cpModifiedAt    time.Time
		publishInterval time.Duration
		wantUpdate      bool
	}{
		{
			name:            "works ok",
			cpModifiedAt:    time.Now().Add(-15 * time.Second),
			publishInterval: 10 * time.Second,
			wantUpdate:      true,
		}, {
			name:            "too soon, skip update",
			cpModifiedAt:    time.Now().Add(-5 * time.Second),
			publishInterval: 10 * time.Second,
			wantUpdate:      false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			m := newMemObjStore()
			storage := &Storage{
				objStore:    m,
				sequencer:   s,
				entriesPath: layout.EntriesPath,
				newCP:       func(size uint64, hash []byte) ([]byte, error) { return []byte(fmt.Sprintf("%d/%x,", size, hash)), nil },
			}
			// Call init so we've got a zero-sized checkpoint to work with.
			if err := storage.init(ctx); err != nil {
				t.Fatalf("storage.init: %v", err)
			}
			cpOld := []byte("bananas")
			if err := m.setObject(ctx, layout.CheckpointPath, cpOld, nil, "", ""); err != nil {
				t.Fatalf("setObject(bananas): %v", err)
			}
			m.lMod = test.cpModifiedAt
			if err := storage.publishCheckpoint(ctx, test.publishInterval); err != nil {
				t.Fatalf("publishCheckpoint: %v", err)
			}
			cpNew, _, err := m.getObject(ctx, layout.CheckpointPath)
			cpUpdated := !bytes.Equal(cpOld, cpNew)
			if err != nil {
				if !errors.Is(err, gcs.ErrObjectNotExist) {
					t.Fatalf("getObject: %v", err)
				}
				cpUpdated = false
			}
			if test.wantUpdate != cpUpdated {
				t.Fatalf("got cpUpdated=%t, want %t", cpUpdated, test.wantUpdate)
			}
		})
	}

}

type memObjStore struct {
	sync.RWMutex
	mem  map[string][]byte
	lMod time.Time
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

func (m *memObjStore) lastModified(_ context.Context, obj string) (time.Time, error) {
	return m.lMod, nil
}
