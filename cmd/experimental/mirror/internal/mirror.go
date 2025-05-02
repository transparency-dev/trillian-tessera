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

// Package mirror provides support for the infrastructure-specific mirror tools.
package mirror

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"log"
	"os"
	"sync/atomic"

	"github.com/avast/retry-go/v4"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/internal/parse"
	"golang.org/x/sync/errgroup"
)

// Target describes a type which can store log static resources.
type Target interface {
	ReadCheckpoint(ctx context.Context) ([]byte, error)
	WriteCheckpoint(ctx context.Context, data []byte) error
	WriteTile(ctx context.Context, l, i uint64, p uint8, data []byte) error
	WriteEntryBundle(ctx context.Context, i uint64, p uint8, data []byte) error
}

// Source describes a type which can fetch static resources from a source log, like
// the .*Fetcher implementations in the client package.
type Source interface {
	ReadCheckpoint(ctx context.Context) ([]byte, error)
	ReadTile(ctx context.Context, l, i uint64, p uint8) ([]byte, error)
	ReadEntryBundle(_ context.Context, i uint64, p uint8) ([]byte, error)
}

// Mirror is a struct which knows how to use the Src and Store functions to copy a tlog-tiles compliant
// log from one location to another.
//
// The checkpoint will only be stored once all static resources have been successfully copied.
// Errors fetching or storing operations will cause the operation to be retried a few times before eventually giving up.
//
// Note that this function _only copies the data_; no self-consistency or correctness checking of
// the copied tiles/entries/checkpoint is undertaken.
type Mirror struct {
	NumWorkers uint
	Source     Source
	Target     Target

	totalResources   uint64
	resourcesFetched atomic.Uint64
}

// Run performs the copy operation.
//
// This is a long-lived operation, returning only once ctx becomes Done, the copy is completed,
// or an error occurs during the operation.
func (m *Mirror) Run(ctx context.Context) error {
	sourceCP, srcSize, err := fetchAndParseCP(ctx, m.Source.ReadCheckpoint)
	if err != nil {
		return fmt.Errorf("failed to fetch source source checkpoint size: %v", err)
	}
	_, targetSize, err := fetchAndParseCP(ctx, m.Target.ReadCheckpoint)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to read checkpoint in target: %v", err)
	}

	delta := srcSize - targetSize
	stride := delta / uint64(m.NumWorkers)
	if r := stride % layout.TileWidth; r != 0 {
		stride += (layout.TileWidth - r)
	}

	log.Printf("Source log size: %d, target log size: %d, ∆ %d, stride: %d", srcSize, targetSize, delta, stride)

	if delta == 0 {
		return nil
	}

	m.totalResources = calcNumResources(srcSize, targetSize, stride)
	m.resourcesFetched = atomic.Uint64{}

	work := make(chan job, m.NumWorkers)

	go func() {
		defer close(work)

		for j := range jobs(srcSize, targetSize, stride) {
			select {
			case <-ctx.Done():
				return
			case work <- j:
				log.Printf("Job: %s", j)
			}
		}
		log.Println("No more work")
	}()

	g := errgroup.Group{}
	for i := range m.NumWorkers {
		g.Go(func() error {
			for j := range work {
				log.Printf("Worker %d: working on %s", i, j)
				for ri := range layout.Range(j.from, j.N, srcSize>>(j.level*layout.TileHeight)) {
					if err := retry.Do(m.copyTile(ctx, j.level, ri.Index, ri.Partial)); err != nil {
						log.Println(err.Error())
						return err
					}

					if j.level == 0 {
						if err := retry.Do(m.copyBundle(ctx, ri.Index, ri.Partial)); err != nil {
							log.Println(err.Error())
							return err
						}
					}
				}
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to migrate static resources: %v", err)
	}
	return m.Target.WriteCheckpoint(ctx, sourceCP)
}

// Progress returns the total number of resources present in the source log, and the number of resources
// successfully copied to the destination so far.
func (m *Mirror) Progress() (uint64, uint64) {
	return m.totalResources, m.resourcesFetched.Load()
}

// copyTile reads a tile from the source log and stores it into the same location in the destination log.
func (m *Mirror) copyTile(ctx context.Context, l, i uint64, p uint8) func() error {
	return func() error {
		d, err := m.Source.ReadTile(ctx, l, i, p)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		if err := m.Target.WriteTile(ctx, l, i, p, d); err != nil {
			return err
		}
		m.resourcesFetched.Add(1)
		return nil
	}
}

// copyBundle reads an entry bundle from the source log and stores it into the same location in the destination log.
func (m *Mirror) copyBundle(ctx context.Context, i uint64, p uint8) func() error {
	return func() error {
		d, err := m.Source.ReadEntryBundle(ctx, i, p)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		if err := m.Target.WriteEntryBundle(ctx, i, p, d); err != nil {
			return err
		}
		m.resourcesFetched.Add(1)
		return nil
	}
}

type job struct {
	level, from, N uint64
}

func (j job) String() string {
	return fmt.Sprintf("Level: %d, Range: [%d, %d)", j.level, j.from, j.from+j.N)
}

func jobs(srcSize, targetSize, stride uint64) iter.Seq[job] {
	return func(yield func(job) bool) {
		for start, ext, l := targetSize, srcSize, uint64(0); ext > 0; start, ext, l = start>>layout.TileHeight, uint64(ext>>layout.TileHeight), l+1 {
			for from := start; from < ext; {
				N := stride
				// If we're starting from a partial tile, then just fetch the remainder first so we're tile-aligned from then on.
				if r := from % layout.TileWidth; r != 0 {
					N = layout.TileWidth - r
				}
				N = min(N, ext-from)
				if !yield(job{level: l, from: from, N: N}) {
					return
				}
				from = from + N
			}
		}

	}
}

// calcNumResources calculates the number of new static resources which need to be mirrored, given the
// size of the source and target.
func calcNumResources(srcSize, targetSize, stride uint64) uint64 {
	leafBundles := uint64(0)
	tiles := uint64(0)

	for j := range jobs(srcSize, targetSize, stride) {
		nTiles := (j.N + layout.TileWidth - 1) / layout.TileWidth
		tiles += nTiles
		if j.level == 0 {
			leafBundles += nTiles
		}
	}

	return leafBundles + tiles
}

func fetchAndParseCP(ctx context.Context, f func(context.Context) ([]byte, error)) ([]byte, uint64, error) {
	cp, err := f(ctx)
	if err != nil {
		return nil, 0, err
	}
	_, size, _, err := parse.CheckpointUnsafe(cp)
	return cp, size, err
}
