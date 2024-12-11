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

package migrate

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/client"
	"golang.org/x/sync/errgroup"
	"k8s.io/klog/v2"
)

type migrate struct {
	storage    MigrationStorage
	getCP      client.CheckpointFetcherFunc
	getTile    client.TileFetcherFunc
	getEntries client.EntryBundleFetcherFunc

	todo chan span

	tilesToMigrate   uint64
	bundlesToMigrate uint64
	tilesMigrated    atomic.Uint64
	bundlesMigrated  atomic.Uint64
}

// span represents the number of tiles at a given tile-level.
type span struct {
	level int
	start uint64
	N     uint64
}

type MigrationStorage interface {
	SetTile(ctx context.Context, level, index uint64, partial uint8, tile []byte) error
	SetEntryBundle(ctx context.Context, index uint64, partial uint8, bundle []byte) error
	SetState(ctx context.Context, treeSize uint64, rootHash []byte) error
}

func Migrate(ctx context.Context, stateDB string, getCP client.CheckpointFetcherFunc, getTile client.TileFetcherFunc, getEntries client.EntryBundleFetcherFunc, storage MigrationStorage) error {
	// TODO store state & resume
	m := &migrate{
		storage:    storage,
		getCP:      getCP,
		getTile:    getTile,
		getEntries: getEntries,
		todo:       make(chan span, 100),
	}

	// init
	cp, err := getCP(ctx)
	if err != nil {
		return fmt.Errorf("fetch initial source checkpoint: %v", err)
	}
	bits := strings.Split(string(cp), "\n")
	size, err := strconv.ParseUint(bits[1], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid CP size %q: %v", bits[1], err)
	}
	rootHash, err := base64.StdEncoding.DecodeString(bits[2])
	if err != nil {
		return fmt.Errorf("invalid checkpoint roothash %q: %v", bits[2], err)
	}

	// figure out what needs copying
	go m.populateSpans(size)

	// Print stats
	go func() {
		for {
			time.Sleep(time.Second)
			tn := m.tilesMigrated.Load()
			tnp := float64(tn*100) / float64(m.tilesToMigrate)
			bn := m.bundlesMigrated.Load()
			bnp := float64(tn*100) / float64(m.bundlesToMigrate)
			klog.Infof("tiles: %d (%.2f%%)  bundles: %d (%.2f%%)", tn, tnp, bn, bnp)
		}
	}()

	// Do the copying
	eg := errgroup.Group{}
	for i := 0; i < 1000; i++ {
		eg.Go(func() error {
			return m.migrateRange(ctx)

		})
	}
	if err := eg.Wait(); err != nil {
		return fmt.Errorf("migrate failed to copy resources: %v", err)
	}
	return storage.SetState(ctx, size, rootHash)
}

func calcExpectedCounts(treeSize uint64) (uint64, uint64) {
	tiles := uint64(0)
	bundles := uint64(0)
	levelSize := treeSize
	for level := 0; levelSize > 0; level++ {
		numFull, partial := levelSize/layout.TileWidth, levelSize%layout.TileWidth
		n := numFull
		if partial > 0 {
			n++
		}
		tiles += n
		if level == 0 {
			bundles = n
		}
		levelSize >>= layout.TileHeight
	}
	return tiles, bundles
}

func (m *migrate) populateSpans(treeSize uint64) {
	m.tilesToMigrate, m.bundlesToMigrate = calcExpectedCounts(treeSize)
	klog.Infof("Spans for treeSize %d", treeSize)
	klog.Infof("total resources to fetch %d tiles + %d bundles = %d", m.tilesToMigrate, m.bundlesToMigrate, m.tilesToMigrate+m.bundlesToMigrate)

	levelSize := treeSize
	for level := 0; levelSize > 0; level++ {
		numFull, partial := levelSize/layout.TileWidth, levelSize%layout.TileWidth
		for j := uint64(0); j < numFull; j++ {
			m.todo <- span{level: level, start: j, N: layout.TileWidth}
			if level == 0 {
				m.todo <- span{level: -1, start: j, N: layout.TileWidth}
			}
		}
		if partial > 0 {
			m.todo <- span{level: level, start: numFull, N: partial}
			if level == 0 {
				m.todo <- span{level: -1, start: numFull, N: partial}
			}

		}
		levelSize >>= layout.TileHeight
	}
	close(m.todo)
}

func (m *migrate) migrateRange(ctx context.Context) error {
	for s := range m.todo {
		if s.N == layout.TileWidth {
			s.N = 0
		}
		if s.level == -1 {
			d, err := m.getEntries(ctx, s.start, uint8(s.N))
			if err != nil {
				return fmt.Errorf("failed to fetch entrybundle %d (p=%d): %v", s.start, s.N, err)
			}
			if err := m.storage.SetEntryBundle(ctx, s.start, uint8(s.N), d); err != nil {
				return fmt.Errorf("failed to store entrybundle %d (p=%d): %v", s.start, s.N, err)
			}
			m.bundlesMigrated.Add(1)
		} else {
			d, err := m.getTile(ctx, uint64(s.level), s.start, uint8(s.N))
			if err != nil {
				return fmt.Errorf("failed to fetch tile level %d index %d (p=%d): %v", s.level, s.start, s.N, err)
			}
			if err := m.storage.SetTile(ctx, uint64(s.level), s.start, uint8(s.N), d); err != nil {
				return fmt.Errorf("failed to store tile level %d index %d (p=%d): %v", s.level, s.start, s.N, err)
			}
			m.tilesMigrated.Add(1)
		}
	}
	return nil
}
