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
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/internal/parse"
	"golang.org/x/sync/errgroup"
)

// StoreFn is the signature of a function which knows how to store a static resource
// at the given path.
type StoreFn func(ctx context.Context, path string, data []byte) error

// Fetcher describes a type which can fetch static resources from a source log, like
// the .*Fetcher implementations in the client package.
type Fetcher interface {
	ReadCheckpoint(ctx context.Context) ([]byte, error)
	ReadTile(ctx context.Context, l, i uint64, p uint8) ([]byte, error)
	ReadEntryBundle(_ context.Context, i uint64, p uint8) ([]byte, error)
}

// Mirror uses the provided Fetcher to pull static tlog-tile resources and checkpoint from
// a source log, and stores them using the provided StoreFn.
//
// The checkpoint will only be stored once all static resources have been successfully mirrored.
// Errors fetching or storing operations will cause the operation to be retried.
//
// Note that this function _only copies the data_; no self-consistency or correctness checking of
// the copied tiles/entries/checkpoint is undertaken.
func Mirror(ctx context.Context, src Fetcher, numWorkers uint, store StoreFn) error {
	sourceCP, err := src.ReadCheckpoint(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch initial source checkpoint: %v", err)
	}
	_, srcSize, _, err := parse.CheckpointUnsafe(sourceCP)
	if err != nil {
		return fmt.Errorf("invalid CP: %v", err)
	}
	log.Printf("Source log size: %d", srcSize)

	totalResources := (srcSize/layout.TileWidth)*2 - 1
	resourcesFetched := atomic.Uint64{}

	work := make(chan job, numWorkers)

	go func() {
		defer close(work)

		stride := srcSize / uint64(numWorkers)
		if r := stride % layout.TileWidth; r != 0 {
			stride += (layout.TileWidth - r)
		}
		log.Printf("Stride: %d", stride)

		for ext, l := srcSize-1, uint64(0); ext > 0; ext, l = uint64(ext>>layout.TileHeight), l+1 {
			for from := uint64(0); from < ext; {
				N := min(stride, ext-from)
				select {
				case <-ctx.Done():
					return
				case work <- job{level: l, from: from, N: N}:
					log.Printf("Level %d, [%d, %d) ", l, from, from+N)
				}
				from = from + N
			}
		}
		log.Println("No more work")
	}()

	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				p := float64(resourcesFetched.Load()*100) / float64(totalResources)
				log.Printf("Progress %d of %d resources (%0.2f)", resourcesFetched.Load(), totalResources, p)
			}
		}
	}()

	g := errgroup.Group{}
	for i := range numWorkers {
		g.Go(func() error {
			for j := range work {
				log.Printf("Worker %d: working on %s", i, j)
				for ri := range layout.Range(j.from, j.N, srcSize>>(j.level*layout.TileHeight)) {
					if err := retry.Do(func() error {
						d, err := src.ReadTile(ctx, j.level, ri.Index, ri.Partial)
						if err != nil {
							log.Println(err.Error())
							return err
						}
						if err := store(ctx, layout.TilePath(j.level, ri.Index, ri.Partial), d); err != nil {
							return err
						}
						resourcesFetched.Add(1)
						return nil
					}); err != nil {
						log.Println(err.Error())
						return err
					}
				}
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to migrate static resources: %v", err)
	}
	return store(ctx, layout.CheckpointPath, sourceCP)
}

type job struct {
	level   uint64
	from, N uint64
}

func (j job) String() string {
	return fmt.Sprintf("Level: %d, Range: [%d, %d)", j.level, j.from, j.from+j.N)
}
