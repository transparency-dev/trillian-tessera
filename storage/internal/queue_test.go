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

package storage_test

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/transparency-dev/tessera/core"
	storage "github.com/transparency-dev/tessera/storage/internal"
)

func TestQueue(t *testing.T) {
	for _, test := range []struct {
		name       string
		numItems   uint64
		maxEntries int
		maxWait    time.Duration
	}{
		{
			name:       "small",
			numItems:   100,
			maxEntries: 200,
			maxWait:    time.Second,
		}, {
			name:       "more items than queue space",
			numItems:   100,
			maxEntries: 20,
			maxWait:    time.Second,
		}, {
			name:       "much flushing",
			numItems:   100,
			maxEntries: 100,
			maxWait:    time.Microsecond,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			assignMu := sync.Mutex{}
			assignedItems := make([]*core.Entry, test.numItems)
			assignedIndex := uint64(0)

			// flushFunc mimics sequencing storage - it takes entries, assigns them to
			// positions in assignedItems.
			flushFunc := func(_ context.Context, entries []*core.Entry) error {
				assignMu.Lock()
				defer assignMu.Unlock()

				for _, e := range entries {
					_ = e.MarshalBundleData(assignedIndex)
					assignedItems[assignedIndex] = e
					assignedIndex++
				}
				return nil
			}

			// Create the Queue
			q := storage.NewQueue(ctx, test.maxWait, uint(test.maxEntries), flushFunc)

			// Now submit a bunch of entries
			adds := make([]core.IndexFuture, test.numItems)
			wantEntries := make([]*core.Entry, test.numItems)
			for i := uint64(0); i < test.numItems; i++ {
				d := fmt.Appendf(nil, "item %d", i)
				wantEntries[i] = core.NewEntry(d)
				adds[i] = q.Add(ctx, wantEntries[i])
			}

			for i, r := range adds {
				N, err := r()
				if err != nil {
					t.Errorf("Add: %v", err)
					return
				}
				if got, want := assignedItems[N.Index].Data(), wantEntries[i].Data(); !reflect.DeepEqual(got, want) {
					t.Errorf("Got item@%d %v, want %v", N.Index, got, want)
				}
			}
		})
	}
}
