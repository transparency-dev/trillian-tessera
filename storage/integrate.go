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
	"errors"
	"fmt"
	"os"
	"reflect"
	"sync"

	"github.com/transparency-dev/merkle/compact"
	"github.com/transparency-dev/merkle/rfc6962"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"golang.org/x/sync/errgroup"
	"k8s.io/klog/v2"
)

// fullTile represents a "fully populated" tile, i.e. it has all non-ephemeral internaly nodes
// implied by the leaves.
type fullTile struct {
	sync.RWMutex

	inner  map[compact.NodeID][]byte
	leaves [][]byte
}

// newFullTile creates and populates a fullTile struct based on the passed in HashTile data.
func newFullTile(h *api.HashTile) *fullTile {
	ft := &fullTile{
		inner:  make(map[compact.NodeID][]byte),
		leaves: make([][]byte, 0, 256),
	}

	if h != nil {
		r := (&compact.RangeFactory{Hash: rfc6962.DefaultHasher.HashChildren}).NewEmptyRange(0)
		for _, h := range h.Nodes {
			r.Append(h, ft.set)
		}
	}
	return ft
}

// set allows setting of individual leaf/inner nodes.
// It's intended to be used as a visitor for compact.Range.
func (f *fullTile) set(id compact.NodeID, hash []byte) {
	f.Lock()
	defer f.Unlock()

	if id.Level == 0 {
		if l, idx := uint64(len(f.leaves)), id.Index; idx >= l {
			f.leaves = append(f.leaves, make([][]byte, idx-l+1)...)
		}
		f.leaves[id.Index] = hash
	} else {
		f.inner[id] = hash
	}
}

// get allows access to individual leaf/inner nodes.
func (f *fullTile) get(id compact.NodeID) []byte {
	f.RLock()
	defer f.RUnlock()

	if id.Level == 0 {
		if l := uint64(len(f.leaves)); id.Index >= l {
			return nil
		}
		return f.leaves[id.Index]
	}
	return f.inner[id]
}

func (f *fullTile) equals(other *fullTile) bool {
	return reflect.DeepEqual(f, other)
}

type GetTileFunc func(ctx context.Context, tileLevel uint64, tileIndex uint64, treeSize uint64) (*api.HashTile, error)
type getFullTileFunc func(ctx context.Context, tileLevel uint64, tileIndex uint64, treeSize uint64) (*fullTile, error)

type TreeBuilder struct {
	sync.Mutex
	getTile getFullTileFunc
	rf      *compact.RangeFactory
}

func NewTreeBuilder(getTile GetTileFunc) *TreeBuilder {
	readCache := tileReadCache{entries: make(map[string]*fullTile)}
	getFullTile := func(ctx context.Context, tileLevel uint64, tileIndex uint64, treeSize uint64) (*fullTile, error) {
		r, ok := readCache.get(tileLevel, tileIndex, treeSize)
		if ok {
			return r, nil
		}
		t, err := getTile(ctx, tileLevel, tileIndex, treeSize)
		if err != nil {
			return nil, err
		}
		ft := newFullTile(t)
		readCache.set(tileLevel, tileIndex, treeSize, ft)
		return ft, nil
	}
	return &TreeBuilder{
		getTile: getFullTile,
		rf:      &compact.RangeFactory{Hash: rfc6962.DefaultHasher.HashChildren},
	}
}

func (t *TreeBuilder) newRange(ctx context.Context, treeSize uint64) (*compact.Range, error) {
	rangeNodes := compact.RangeNodes(0, treeSize, nil)
	errG := errgroup.Group{}
	hashes := make([][]byte, len(rangeNodes))
	for i, id := range rangeNodes {
		i := i
		id := id
		errG.Go(func() error {
			tLevel, tIndex, nLevel, nIndex := layout.NodeCoordsToTileAddress(uint64(id.Level), id.Index)
			ft, err := t.getTile(ctx, uint64(tLevel), tIndex, treeSize)
			if err != nil {
				return err
			}
			h := ft.get(compact.NodeID{Level: nLevel, Index: nIndex})
			if h == nil {
				return fmt.Errorf("missing node: [%d/%d@%d]", id.Level, id.Index, treeSize)
			}
			hashes[i] = h
			return nil
		})
	}
	if err := errG.Wait(); err != nil {
		return nil, err
	}
	return t.rf.NewRange(0, treeSize, hashes)
}

func (t *TreeBuilder) Integrate(ctx context.Context, fromSize uint64, entries [][]byte) (newSize uint64, rootHash []byte, tiles map[compact.NodeID]*api.HashTile, err error) {
	t.Lock()
	defer t.Unlock()

	baseRange, err := t.newRange(ctx, fromSize)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("failed to create range covering existing log: %w", err)
	}

	// Initialise a compact range representation, and verify the stored state.
	r, err := baseRange.GetRootHash(nil)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("invalid log state, unable to recalculate root: %w", err)
	}

	klog.V(1).Infof("Loaded state with roothash %x", r)
	// Create a new compact range which represents the update to the tree
	newRange := t.rf.NewEmptyRange(fromSize)
	if len(entries) == 0 {
		klog.V(1).Infof("Nothing to do.")
		// Nothing to do, nothing done.
		return fromSize, r, nil, nil
	}
	tc := newTileWriteCache(fromSize, t.getTile)
	visitor := tc.Visitor(ctx)
	for _, e := range entries {
		lh := rfc6962.DefaultHasher.HashLeaf(e)
		// Update range and set nodes
		if err := newRange.Append(lh, visitor); err != nil {
			return 0, nil, nil, fmt.Errorf("newRange.Append(): %v", err)
		}

	}

	// Merge the update range into the old tree
	if err := baseRange.AppendRange(newRange, visitor); err != nil {
		return 0, nil, nil, fmt.Errorf("failed to merge new range onto existing log: %w", err)
	}

	// Calculate the new root hash - don't pass in the tileCache visitor here since
	// this will construct any ephemeral nodes and we do not want to store those.
	newRoot, err := baseRange.GetRootHash(nil)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("failed to calculate new root hash: %w", err)
	}

	if err := tc.Err(); err != nil {
		return 0, nil, nil, err
	}

	// All calculation is now complete, all that remains is to store the new
	// tiles and updated log state.
	klog.V(1).Infof("New log state: size 0x%x hash: %x", baseRange.End(), newRoot)

	return baseRange.End(), newRoot, tc.Tiles(), nil

}

// tileReadCache is a structure which provides a very simple thread-safe map of tiles.
type tileReadCache struct {
	sync.RWMutex

	hits    int
	entries map[string]*fullTile
}

func newTileReadCache() tileReadCache {
	return tileReadCache{
		entries: make(map[string]*fullTile),
	}
}

// get returns a previously set tile and true, or, if no such tile is in the cache, returns nil and false.
func (r *tileReadCache) get(tileLevel, tileIndex, treeSize uint64) (*fullTile, bool) {
	r.RLock()
	defer r.RUnlock()
	k := layout.TilePath(tileLevel, tileIndex, treeSize)
	e, ok := r.entries[k]
	if ok {
		r.hits++
	}
	return e, ok
}

// set associates the given tile coords with a tile.
func (r *tileReadCache) set(tileLevel, tileIndex, treeSize uint64, t *fullTile) {
	r.Lock()
	defer r.Unlock()
	k := layout.TilePath(tileLevel, tileIndex, treeSize)
	if e, ok := r.entries[k]; ok && !e.equals(t) {
		if klog.V(2).Enabled() {
			klog.Infof("OVERWRITE TILE %v:\nExisting:%x\nNew\n%x", k, e, t)
		}
		panic(fmt.Errorf("Attempting to overwrite %v with different content", k))
	}
	r.entries[k] = t
}

// tileWriteCache is a simple cache for storing the newly created tiles produced by
// the integration of new leaves into the tree.
//
// Calls to Visit will cause the map of tiles to become filled with the set of
// `dirty` tiles which need to be flushed back to storage to preserve the updated
// tree state.
//
// Note that by itself, this cache does not update any persisted state.
type tileWriteCache struct {
	sync.Mutex
	m   map[compact.NodeID]*fullTile
	err []error

	treeSize uint64
	getTile  getFullTileFunc
}

func newTileWriteCache(treeSize uint64, getTile getFullTileFunc) *tileWriteCache {
	return &tileWriteCache{
		m:        make(map[compact.NodeID]*fullTile),
		treeSize: treeSize,
		getTile:  getTile,
	}
}

func (tc *tileWriteCache) Err() error {
	return errors.Join(tc.err...)
}

// Visit should be called once for each newly set non-ephemeral node in the
// tree.
//
// If the tile containing id has not been seen before, this method will fetch
// it from disk (or create a new empty in-memory tile if it doesn't exist), and
// update it by setting the node corresponding to id to the value hash.
func (tc *tileWriteCache) Visitor(ctx context.Context) compact.VisitFn {
	return func(id compact.NodeID, hash []byte) {
		tc.Lock()
		defer tc.Unlock()
		//klog.V(3).Infof("VISIT %v", id)
		tileLevel, tileIndex, nodeLevel, nodeIndex := layout.NodeCoordsToTileAddress(uint64(id.Level), uint64(id.Index))
		tileKey := compact.NodeID{Level: uint(tileLevel), Index: tileIndex}
		tile := tc.m[tileKey]
		if tile == nil {
			var err error
			tile, err = tc.getTile(ctx, tileLevel, tileIndex, tc.treeSize)
			if err != nil {
				if !errors.Is(err, os.ErrNotExist) {
					tc.err = append(tc.err, err)
					return
				}
				// This is a brand new tile.
				tile = newFullTile(nil)
			}
			tc.m[tileKey] = tile
		}
		// Update the tile with the new node hash.
		idx := compact.NodeID{Level: nodeLevel, Index: nodeIndex}
		tile.set(idx, hash)
	}
}

// Tiles returns all visited tiles.
func (tc *tileWriteCache) Tiles() map[compact.NodeID]*api.HashTile {
	newTiles := make(map[compact.NodeID]*api.HashTile)
	for k, t := range tc.m {
		newTiles[k] = &api.HashTile{Nodes: t.leaves}
	}
	return newTiles
}
