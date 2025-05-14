// Copyright 2025 The Tessera authors. All Rights Reserved.
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

package fsck

import (
	"bytes"
	"context"
	"fmt"

	"github.com/transparency-dev/merkle/compact"
	"github.com/transparency-dev/merkle/rfc6962"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/internal/parse"
	"github.com/transparency-dev/trillian-tessera/internal/stream"
	"golang.org/x/sync/errgroup"
	"k8s.io/klog/v2"
)

type Fetcher interface {
	ReadCheckpoint(ctx context.Context) ([]byte, error)
	ReadTile(ctx context.Context, l, i uint64, p uint8) ([]byte, error)
	ReadEntryBundle(ctx context.Context, i uint64, p uint8) ([]byte, error)
}

func Check(ctx context.Context, f Fetcher, bundleHasher func([]byte) ([][]byte, error)) error {
	cp, err := f.ReadCheckpoint(ctx)
	if err != nil {
		klog.Exitf("fetch initial source checkpoint: %v", err)
	}
	klog.V(1).Infof("Fsck: checkpoint:\n%s", cp)
	// TODO(al): Should really open this with the pubK to verify the checkpoint is well formed.
	_, logSize, logRoot, err := parse.CheckpointUnsafe(cp)
	if err != nil {
		klog.Exitf("Failed to parse checkpoint: %v", err)
	}
	klog.Infof("Fsck: checking log of size %d", logSize)

	N := uint(1)

	fTree := fsckTree{
		fetcher:           f,
		bundleHasher:      bundleHasher,
		tree:              (&compact.RangeFactory{Hash: rfc6962.DefaultHasher.HashChildren}).NewEmptyRange(0),
		sourceSize:        logSize,
		pendingTiles:      make(map[compact.NodeID]*api.HashTile),
		expectedResources: make(chan resource, N),
	}

	getSize := func(_ context.Context) (uint64, error) { return logSize, nil }
	next, cancel := stream.StreamAdaptor(ctx, N, getSize, f.ReadEntryBundle, 0)
	defer cancel()

	eg := errgroup.Group{}
	eg.Go(fTree.compareResources(ctx))

	for {
		ri, b, err := next()
		if err != nil {
			klog.Warningf("next: %v", err)
			break
		}
		if err := fTree.AppendBundle(ri, b); err != nil {
			klog.Warningf("AppendBundle(%v): %v", ri, err)
			break
		}
		if fTree.tree.End() == logSize {
			break
		}
	}
	fTree.flushPartialTiles()
	close(fTree.expectedResources)

	if err := eg.Wait(); err != nil {
		klog.Exitf("Failed: %v", err)
	}

	gotRoot, err := fTree.tree.GetRootHash(nil)
	switch {
	case err != nil:
		klog.Exitf("Failed to calculate root: %v", err)
	case !bytes.Equal(gotRoot, logRoot):
		klog.Exitf("Calculated root %x, but checkpoint claims %x", gotRoot, logRoot)
	default:
		klog.Infof("Successfully fsck'd log with size %d and root %x", logSize, gotRoot)
	}

	return nil
}

type resource struct {
	level, index uint64
	partial      uint8
	content      []byte
}

type fsckTree struct {
	fetcher      Fetcher
	bundleHasher func([]byte) ([][]byte, error)
	tree         *compact.Range
	sourceSize   uint64

	pendingTiles map[compact.NodeID]*api.HashTile

	expectedResources chan resource
}

func (f *fsckTree) AppendBundle(ri layout.RangeInfo, data []byte) error {
	hs, err := f.bundleHasher(data)
	if err != nil {
		return err
	}
	for i := ri.First; i < ri.First+ri.N; i++ {
		if err := f.tree.Append(hs[i], f.visit); err != nil {
			return err
		}
	}
	return nil
}

func (f *fsckTree) visit(id compact.NodeID, h []byte) {
	// We're only storing the lowest level of hash in the tiles, so early-out in other cases.
	if id.Level%layout.TileHeight != 0 {
		return
	}
	tLevel, tIdx, hIdx := id.Level/layout.TileHeight, id.Index/layout.EntryBundleWidth, id.Index%layout.EntryBundleWidth
	k := compact.NodeID{Level: tLevel, Index: tIdx}
	t, ok := f.pendingTiles[k]
	if !ok {
		t = &api.HashTile{}
		f.pendingTiles[k] = t
	}
	if hIdx != uint64(len(t.Nodes)) {
		klog.Exitf("LOGIC ERROR: got tile (l: %d, idx: %d) node index %d, for tile with %d nodes", tLevel, tIdx, hIdx, len(t.Nodes))
	}
	t.Nodes = append(t.Nodes, h)
	// TODO: make this better
	if len(t.Nodes) == layout.EntryBundleWidth {
		c, err := t.MarshalText()
		if err != nil {
			klog.Exitf("Failed to marshal tile: %v", err)
		}
		f.expectedResources <- resource{
			level:   uint64(tLevel),
			index:   tIdx,
			partial: uint8(len(t.Nodes)),
			content: c,
		}
		delete(f.pendingTiles, k)
	}
}

func (f *fsckTree) flushPartialTiles() {
	for k, t := range f.pendingTiles {
		c, err := t.MarshalText()
		if err != nil {
			klog.Exitf("Failed to marshal tile: %v", err)
		}
		f.expectedResources <- resource{
			level:   uint64(k.Level),
			index:   k.Index,
			partial: uint8(len(t.Nodes)),
			content: c,
		}
		delete(f.pendingTiles, k)
	}
}

func (f *fsckTree) compareResources(ctx context.Context) func() error {
	return func() error {
		for r := range f.expectedResources {
			data, err := f.fetcher.ReadTile(ctx, r.level, r.index, r.partial)
			if err != nil {
				return err
			}
			p := layout.TilePath(r.level, r.index, r.partial)
			if !bytes.Equal(data, r.content) {
				return fmt.Errorf("%s: log has %x expected %x", p, data, r.content)
			}
			klog.V(2).Infof("%s ok", p)
		}
		return nil
	}
}
