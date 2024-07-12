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
	"sync"
	"testing"

	"cloud.google.com/go/spanner/spannertest"
	"cloud.google.com/go/spanner/spansql"
	gcs "cloud.google.com/go/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
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
			CREATE TABLE IntCoord (id INT64 NOT NULL, seq INT64 NOT NULL,) PRIMARY KEY (id); 
	`)
	if err != nil {
		t.Fatalf("Invalid DDL: %v", err)
	}
	if err := srv.UpdateDDL(dml); err != nil {
		t.Fatalf("Failed to create schema in test spanner: %v", err)
	}

	return srv.Close

}

func TestSpannerSequencer(t *testing.T) {
	ctx := context.Background()
	close := newSpannerDB(t)
	defer close()

	seq, err := newSpannerSequencer(ctx, "projects/p/instances/i/databases/d")
	if err != nil {
		t.Fatalf("newSpannerSequencer: %v", err)
	}

	want := uint64(0)
	for chunks := 0; chunks < 10; chunks++ {
		entries := [][]byte{}
		for i := 0; i < 10+chunks; i++ {
			entries = append(entries, []byte(fmt.Sprintf("item %d/%d", chunks, i)))
		}
		got, err := seq.assignEntries(ctx, entries)
		if err != nil {
			t.Fatalf("assignEntries: %v", err)
		}
		if got != want {
			t.Errorf("Chunk %d got seq %d, want %d", chunks, got, want)
		}
		want += uint64(len(entries))
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
			index:    3 * 256,
			logSize:  3*256 + 20,
			tileSize: 20,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			wantTile := makeTile(t, test.tileSize)
			if err := s.setTile(ctx, test.level, test.index, test.logSize, wantTile); err != nil {
				t.Fatalf("setTile: %v", err)
			}

			expPath := layout.TilePath(test.level, test.index, test.logSize)
			_, ok := m.mem[expPath]
			if !ok {
				t.Fatalf("want tile at %v but found none", expPath)
			}

			got, err := s.getTile(ctx, test.level, test.index, test.logSize)
			if err != nil {
				t.Fatalf("getTile: %v", err)
			}
			if !cmp.Equal(got, wantTile) {
				t.Fatal("roundtrip returned different data")
			}
		})
	}
}

func makeBundle(t *testing.T, size uint64) *api.EntryBundle {
	t.Helper()
	r := &api.EntryBundle{Entries: make([][]byte, size)}
	for i := uint64(0); i < size; i++ {
		r.Entries[i] = []byte(fmt.Sprintf("%d", i))
	}
	return r
}

func TestBundleRoundtrip(t *testing.T) {
	ctx := context.Background()
	m := newMemObjStore()
	s := &Storage{
		objStore: m,
	}

	for _, test := range []struct {
		name       string
		index      uint64
		logSize    uint64
		bundleSize uint64
	}{
		{
			name:       "ok",
			index:      3 * 256,
			logSize:    3*256 + 20,
			bundleSize: 20,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			wantBundle := makeBundle(t, test.bundleSize)
			if err := s.setEntryBundle(ctx, test.index, test.logSize, wantBundle); err != nil {
				t.Fatalf("setEntryBundle: %v", err)
			}

			expPath := layout.EntriesPath(test.index, test.logSize)
			_, ok := m.mem[expPath]
			if !ok {
				t.Fatalf("want bundle at %v but found none", expPath)
			}

			got, err := s.getEntryBundle(ctx, test.index, test.logSize)
			if err != nil {
				t.Fatalf("getEntryBundle: %v", err)
			}
			if !cmp.Equal(got, wantBundle) {
				t.Fatal("roundtrip returned different data")
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

func (m *memObjStore) setObject(_ context.Context, obj string, data []byte, cond *gcs.Conditions) error {
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
