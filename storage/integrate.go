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

	"github.com/transparency-dev/merkle/compact"
	"github.com/transparency-dev/merkle/rfc6962"
	"github.com/transparency-dev/serverless-log/client"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"golang.org/x/sync/errgroup"
	"k8s.io/klog"
)

// fullTile represents a "fully populated" tile, i.e. it has all non-ephemeral internaly nodes
// implied by the leaves.
type fullTile struct {
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
func (f fullTile) set(id compact.NodeID, hash []byte) {
	if id.Level == 0 {
		if l := uint64(len(f.leaves)); id.Index > l {
			f.leaves = append(f.leaves, make([][]byte, l-id.Index+1)...)
		}
		f.leaves[id.Index] = hash
	} else {
		f.inner[id] = hash
	}
}

type GetTileFunc func(ctx context.Context, tileLevel uint64, tileIndex uint64, treeSize uint64) (*api.HashTile, error)

type TreeBuilder struct {
	tileWriteCache *tileWriteCache
	getTile        func(ctx context.Context, tileLevel uint64, tileIndex uint64) (*api.HashTile, error)
}

func NewTreeBuilder(getTile func(ctx context.Context, tileLevel uint64, tileIndex uint64, treeSize uint64) (*api.HashTile, error)) *TreeBuilder {
	getFullTile := func(ctx context.Context, tileLevel uint64, tileIndex uint64) (*fullTile, error) {
		t, err := getTile(ctx, tileLevel, tileIndex)
		if err != nil {
			return nil, fmt.Errorf("getTile: %v", err)
		}
		return newFullTile(t), nil
	}
	return &TreeBuilder{
		tileWriteCache: newTileWriteCache(getFullTile),
		getTile:        getTile,
	}
}

func (t *TreeBuilder) Integrate(ctx context.Context, fromSize uint64, entries [][]byte) (newSize uint64, rootHash []byte, err error) {
	rc := readCache{entries: make(map[string]fullTile)}
	defer func() {
		klog.Infof("read cache hits: %d", rc.hits)
	}()
	getTile := func(l, i uint64) (*api.HashTile, error) {
		r, ok := rc.get(l, i)
		if ok {
			return r, nil
		}
		t, err := t.getTile(ctx, l, i, fromSize)
		if err != nil {
			return nil, err
		}
		rc.set(l, i, t)
		return t, nil
	}
	prewarmCache(compact.RangeNodes(0, fromSize, nil), getTile)

	hashes, err := client.FetchRangeNodes(ctx, fromSize, func(_ context.Context, l, i uint64) (*api.Tile, error) {
		return getTile(l, i)
	})
	if err != nil {
		return 0, nil, fmt.Errorf("failed to fetch compact range nodes: %w", err)
	}

	rf := compact.RangeFactory{Hash: h.HashChildren}
	baseRange, err := rf.NewRange(0, fromSize, hashes)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create range covering existing log: %w", err)
	}

	// Initialise a compact range representation, and verify the stored state.
	r, err := baseRange.GetRootHash(nil)
	if err != nil {
		return 0, nil, fmt.Errorf("invalid log state, unable to recalculate root: %w", err)
	}

	klog.V(1).Infof("Loaded state with roothash %x", r)

	// Create a new compact range which represents the update to the tree
	newRange := rf.NewEmptyRange(fromSize)
	tc := tileCache{m: make(map[tileKey]*api.Tile), getTile: getTile}
	if len(batch) == 0 {
		klog.V(1).Infof("Nothing to do.")
		// Nothing to do, nothing done.
		return fromSize, r, nil
	}
	for _, e := range batch {
		lh := h.HashLeaf(e)
		// Update range and set nodes
		if err := newRange.Append(lh, tc.Visit); err != nil {
			return 0, nil, fmt.Errorf("newRange.Append(): %v", err)
		}

	}

	// Merge the update range into the old tree
	if err := baseRange.AppendRange(newRange, tc.Visit); err != nil {
		return 0, nil, fmt.Errorf("failed to merge new range onto existing log: %w", err)
	}

	// Calculate the new root hash - don't pass in the tileCache visitor here since
	// this will construct any ephemeral nodes and we do not want to store those.
	newRoot, err := baseRange.GetRootHash(nil)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to calculate new root hash: %w", err)
	}

	// All calculation is now complete, all that remains is to store the new
	// tiles and updated log state.
	klog.V(1).Infof("New log state: size 0x%x hash: %x", baseRange.End(), newRoot)

	eg := errgroup.Group{}
	for k, t := range tc.m {
		k := k
		t := t
		eg.Go(func() error {
			if err := st.StoreTile(ctx, k.level, k.index, t); err != nil {
				return fmt.Errorf("failed to store tile at level %d index %d: %w", k.level, k.index, err)
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return 0, nil, err
	}

	return baseRange.End(), newRoot, nil

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
	m   map[compact.NodeID]*fullTile
	err []error

	getTile func(ctx context.Context, level, index uint64) (*fullTile, error)
}

func newTileWriteCache(getTile func(ctx context.Context, tileLevel uint64, tileIndex uint64) (*fullTile, error)) *tileWriteCache {
	return &tileWriteCache{
		m:       make(map[compact.NodeID]*fullTile),
		getTile: getTile,
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
func (tc tileWriteCache) Visit(ctx context.Context, id compact.NodeID, hash []byte) {
	tileLevel, tileIndex, nodeLevel, nodeIndex := layout.NodeCoordsToTileAddress(uint64(id.Level), uint64(id.Index))
	tileKey := compact.NodeID{Level: uint(tileLevel), Index: tileIndex}
	tile := tc.m[tileKey]
	if tile == nil {
		var err error
		tile, err = tc.getTile(ctx, tileLevel, tileIndex)
		if err != nil {
			if !os.IsNotExist(err) {
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
