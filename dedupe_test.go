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

package tessera_test

import (
	"context"
	"crypto/sha256"
	"fmt"
	"sync"
	"testing"

	tessera "github.com/transparency-dev/trillian-tessera"
)

func TestDedupe(t *testing.T) {
	ctx := context.Background()
	testCases := []struct {
		desc     string
		newValue string
		wantIdx  uint64
	}{
		{
			desc:     "first element",
			newValue: "foo",
			wantIdx:  1,
		},
		{
			desc:     "third element",
			newValue: "baz",
			wantIdx:  3,
		},
		{
			desc:     "new element",
			newValue: "omega",
			wantIdx:  4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			idx := uint64(1)
			delegate := func(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
				thisIdx := idx
				idx++
				return func() (uint64, error) {
					return thisIdx, nil
				}
			}
			dedupeAdd := tessera.InMemoryDedupe(delegate, 256)

			// Add foo, bar, baz to prime the cache to make things interesting
			dedupeAdd(ctx, tessera.NewEntry([]byte("foo")))
			dedupeAdd(ctx, tessera.NewEntry([]byte("bar")))
			dedupeAdd(ctx, tessera.NewEntry([]byte("baz")))

			idx, err := dedupeAdd(ctx, tessera.NewEntry([]byte(tC.newValue)))()
			if err != nil {
				t.Fatal(err)
			}
			if idx != tC.wantIdx {
				t.Errorf("got != want (%d != %d)", idx, tC.wantIdx)
			}
		})
	}
}

func BenchmarkDedupe(b *testing.B) {
	ctx := context.Background()
	// Outer loop is for benchmark calibration, inside here is each individual run of the benchmark
	for i := 0; i < b.N; i++ {
		idx := uint64(1)
		delegate := func(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
			thisIdx := idx
			idx++
			return func() (uint64, error) {
				return thisIdx, nil
			}
		}
		dedupeAdd := tessera.InMemoryDedupe(delegate, 256)
		wg := &sync.WaitGroup{}
		// Loop to create a bunch of leaves in parallel to test lock contention
		for leafIndex := range 1024 {
			wg.Add(1)
			go func(index int) {
				_, err := dedupeAdd(ctx, tessera.NewEntry([]byte(fmt.Sprintf("leaf with value %d", index%sha256.Size))))()
				if err != nil {
					b.Error(err)
				}
				wg.Done()
			}(leafIndex)
		}
		wg.Wait()
	}
}
