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
	"sync/atomic"
	"testing"
	"time"

	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/storage"
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
			assignMu := sync.Mutex{}
			assignedItems := make([]tessera.Entry, test.numItems)
			assignedIndex := uint64(0)

			// flushFunc mimics sequencing storage - it takes entries, assigns them to
			// positions in assignedItems and returns the first assigned positition for each batch.
			flushFunc := func(_ context.Context, entries []tessera.Entry) (uint64, error) {
				assignMu.Lock()
				defer assignMu.Unlock()

				s := assignedIndex
				for _, e := range entries {
					assignedItems[assignedIndex] = e
					assignedIndex++
				}
				return s, nil
			}

			// Create the Queue
			q := storage.NewQueue(test.maxWait, uint(test.maxEntries), flushFunc)

			// Now submit a bunch of entries
			adds := make([]storage.IndexFunc, test.numItems)
			wantEntries := make([]tessera.Entry, test.numItems)
			for i := uint64(0); i < test.numItems; i++ {
				wantEntries[i] = tessera.NewEntry([]byte(fmt.Sprintf("item %d", i)))
				adds[i] = q.Add(context.TODO(), wantEntries[i])
			}

			for i, r := range adds {
				N, err := r()
				if err != nil {
					t.Errorf("Add: %v", err)
					return
				}
				if got, want := assignedItems[N], wantEntries[i]; !reflect.DeepEqual(got, want) {
					t.Errorf("Got item@%d %q, want %q", N, got, want)
				}
			}
		})
	}
}

func TestDedup(t *testing.T) {
	idx := atomic.Uint64{}

	q := storage.NewQueue(time.Second, 10 /*maxSize*/, func(ctx context.Context, entries []tessera.Entry) (index uint64, err error) {
		r := idx.Load()
		idx.Add(1)
		return r, nil
	})

	numEntries := 10
	adds := []storage.IndexFunc{}
	for i := 0; i < numEntries; i++ {
		adds = append(adds, q.Add(context.TODO(), tessera.NewEntry([]byte("Have I seen this before?"))))
	}

	firstN, err := adds[0]()
	if err != nil {
		t.Fatalf("Add: %v", err)
	}
	for i := 1; i < len(adds); i++ {
		N, err := adds[i]()
		if err != nil {
			t.Errorf("[%d] got %v", i, err)
		}
		if N != firstN {
			t.Errorf("[%d] got seq %d, want %d", i, N, firstN)
		}
	}
}
