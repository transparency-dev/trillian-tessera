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

package tessera

import (
	"context"
	"crypto/sha256"
	"fmt"
	"sync"
	"testing"

	"github.com/transparency-dev/trillian-tessera/shizzle"
)

func TestDedupe(t *testing.T) {
	ctx := context.Background()
	testCases := []struct {
		desc     string
		newValue string
		wantIdx  uint64
		wantDup  bool
	}{
		{
			desc:     "first element",
			newValue: "foo",
			wantIdx:  1,
			wantDup:  true,
		},
		{
			desc:     "third element",
			newValue: "baz",
			wantIdx:  3,
			wantDup:  true,
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
			delegate := func(ctx context.Context, e *shizzle.Entry) shizzle.IndexFuture {
				thisIdx := idx
				idx++
				return func() (shizzle.Index, error) {
					return shizzle.Index{Index: thisIdx}, nil
				}
			}
			dedupeAdd := newInMemoryDedupe(256)(delegate)

			// Add foo, bar, baz to prime the cache to make things interesting
			for _, s := range []string{"foo", "bar", "baz"} {
				if _, err := dedupeAdd(ctx, shizzle.NewEntry([]byte(s)))(); err != nil {
					t.Fatalf("dedupeAdd(%q): %v", s, err)
				}
			}

			i, err := dedupeAdd(ctx, shizzle.NewEntry([]byte(tC.newValue)))()
			if err != nil {
				t.Fatalf("dedupeAdd(%q): %v", tC.newValue, err)
			}
			if i.Index != tC.wantIdx {
				t.Errorf("got Index != want Index (%d != %d)", i.Index, tC.wantIdx)
			}
			if i.IsDup != tC.wantDup {
				t.Errorf("got IsDup != want IsDup(%t != %t)", i.IsDup, tC.wantDup)

			}
		})
	}
}

func BenchmarkDedupe(b *testing.B) {
	ctx := context.Background()
	// Outer loop is for benchmark calibration, inside here is each individual run of the benchmark
	for b.Loop() {
		idx := uint64(1)
		delegate := func(ctx context.Context, e *shizzle.Entry) shizzle.IndexFuture {
			thisIdx := idx
			idx++
			return func() (shizzle.Index, error) {
				return shizzle.Index{Index: thisIdx}, nil
			}
		}
		dedupeAdd := newInMemoryDedupe(256)(delegate)
		wg := &sync.WaitGroup{}
		// Loop to create a bunch of leaves in parallel to test lock contention
		for leafIndex := range 1024 {
			wg.Add(1)
			go func(index int) {
				_, err := dedupeAdd(ctx, shizzle.NewEntry(fmt.Appendf(nil, "leaf with value %d", index%sha256.Size)))()
				if err != nil {
					b.Error(err)
				}
				wg.Done()
			}(leafIndex)
		}
		wg.Wait()
	}
}
