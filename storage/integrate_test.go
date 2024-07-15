// Copyright 2024 Google LLC. All Rights Reserved.
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

package storage

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"sync"
	"testing"

	"github.com/transparency-dev/merkle/compact"
	"github.com/transparency-dev/trillian-tessera/api"
)

func TestNewRangeFetchesTiles(t *testing.T) {
	ctx := context.Background()
	m := newMemTileStore[api.HashTile]()
	tb := NewTreeBuilder(m.getTile)

	treeSize := uint64(0x102030)
	wantIDs := []compact.NodeID{
		{Level: 0, Index: 0x1020},
		{Level: 1, Index: 0x10},
		{Level: 2, Index: 0x0},
	}

	for _, id := range wantIDs {
		if err := m.setTile(ctx, id, treeSize, zeroTile(256)); err != nil {
			t.Fatalf("setTile: %v", err)
		}
	}

	_, err := tb.newRange(ctx, treeSize)
	if err != nil {
		t.Fatalf("newRange(%d): %v", treeSize, err)
	}
}

func TestTileVisit(t *testing.T) {
	ctx := context.Background()
	m := newMemTileStore[fullTile]()
	treeSize := uint64(0x102030)

	for _, test := range []struct {
		name      string
		visits    map[compact.NodeID][]byte
		wantTiles map[compact.NodeID]*api.HashTile
	}{
		{
			name: "ok - single tile",
			visits: map[compact.NodeID][]byte{
				compact.NewNodeID(0, 0): {0},
				compact.NewNodeID(0, 1): {1},
				compact.NewNodeID(1, 1): {2},
			},
			wantTiles: map[compact.NodeID]*api.HashTile{
				compact.NewNodeID(0, 0): {Nodes: [][]byte{{0}, {1}}},
			},
		},
		{
			name: "ok - multiple tiles",
			visits: map[compact.NodeID][]byte{
				compact.NewNodeID(0, 0):     {0},
				compact.NewNodeID(0, 1*256): {1},
				compact.NewNodeID(8, 2*256): {2},
			},
			wantTiles: map[compact.NodeID]*api.HashTile{
				compact.NewNodeID(0, 0): {Nodes: [][]byte{{0}}},
				compact.NewNodeID(0, 1): {Nodes: [][]byte{{1}}},
				compact.NewNodeID(1, 2): {Nodes: [][]byte{{2}}},
			},
		},
	} {
		twc := newTileWriteCache(treeSize, m.getTile)
		v := twc.Visitor(ctx)
		for id, k := range test.visits {
			v(id, k)
		}
		if err := twc.Err(); err != nil {
			t.Fatalf("Got err: %v", err)
		}
		gotTiles := twc.Tiles()
		for id, wantTile := range test.wantTiles {
			gotTile, ok := gotTiles[id]
			if !ok {
				t.Errorf("Missing tile %v", id)
				continue
			}
			if !reflect.DeepEqual(gotTile, wantTile) {
				t.Errorf("Got tile with unexpected data at %v:\ngot:\n%x\nwant:\n%x", id, gotTile, wantTile)
			}
			delete(gotTiles, id)
			delete(test.wantTiles, id)
		}
		if l := len(gotTiles); l > 0 {
			t.Errorf("got unexpected tiles: %v", gotTiles)
		}
		if l := len(test.wantTiles); l > 0 {
			t.Errorf("did not get expected tiles: %v", test.wantTiles)
		}
	}
}

// zerotile creates a new api.HashTile of the provided size, whose leaves are all a single zero byte.
func zeroTile(size uint64) *api.HashTile {
	r := &api.HashTile{
		Nodes: make([][]byte, size),
	}
	for i := range r.Nodes {
		r.Nodes[i] = []byte{0}
	}
	return r
}

type memTileStoreKey struct {
	id       compact.NodeID
	treeSize uint64
}

func (k memTileStoreKey) String() string {
	return fmt.Sprintf("[L: 0x%x, I: 0x%x, S: 0x%x]", k.id.Level, k.id.Index, k.treeSize)
}

type memTileStore[T any] struct {
	sync.RWMutex
	mem map[memTileStoreKey]*T
}

func newMemTileStore[T any]() *memTileStore[T] {
	return &memTileStore[T]{
		mem: make(map[memTileStoreKey]*T),
	}
}

func (m *memTileStore[T]) getTile(_ context.Context, id compact.NodeID, treeSize uint64) (*T, error) {
	m.RLock()
	defer m.RUnlock()

	k := memTileStoreKey{id: id, treeSize: treeSize}
	d, ok := m.mem[k]
	if !ok {
		return nil, fmt.Errorf("tile %q: %w", k, os.ErrNotExist)
	}
	return d, nil
}

func (m *memTileStore[T]) setTile(_ context.Context, id compact.NodeID, treeSize uint64, t *T) error {
	m.Lock()
	defer m.Unlock()

	k := memTileStoreKey{id: id, treeSize: treeSize}
	_, ok := m.mem[k]
	if ok {
		return fmt.Errorf("%q is already present", k)
	}

	m.mem[k] = &(*t)
	return nil
}
