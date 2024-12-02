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
	"sync"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

// InMemoryDedupe wraps an Add function to prevent duplicate entries being written to the underlying
// storage by keeping an in-memory cache of recently seen entries.
// Where an existing entry has already been `Add`ed, the previous `IndexFuture` will be returned.
// When no entry is found in the cache, the delegate method will be called to store the entry, and
// the result will be registered in the cache.
//
// Internally this uses a cache with a max size configured by the size parameter.
// If the entry being `Add`ed is not found in the cache, then it calls the delegate.
//
// This object can be used in isolation, or in conjunction with a persistent dedupe implementation.
// When using this with a persistent dedupe, the persistent layer should be the delegate of this
// InMemoryDedupe. This allows recent duplicates to be deduplicated in memory, reducing the need to
// make calls to a persistent storage.
func InMemoryDedupe(delegate func(ctx context.Context, e *Entry) IndexFuture, size uint) func(context.Context, *Entry) IndexFuture {
	dedupe := &inMemoryDedupe{
		delegate: delegate,
		cache:    expirable.NewLRU[string, IndexFuture](int(size), nil, 0),
	}
	return dedupe.add
}

type inMemoryDedupe struct {
	delegate func(ctx context.Context, e *Entry) IndexFuture
	mu       sync.Mutex // cache is thread safe, but this mutex allows us to do conditional writes
	cache    *expirable.LRU[string, IndexFuture]
}

// Add adds the entry to the underlying delegate only if e hasn't been recently seen. In either case,
// an IndexFuture will be returned that the client can use to get the sequence number of this entry.
func (d *inMemoryDedupe) add(ctx context.Context, e *Entry) IndexFuture {
	id := string(e.Identity())
	d.mu.Lock()
	defer d.mu.Unlock()

	f, ok := d.cache.Get(id)
	if !ok {
		f = d.delegate(ctx, e)
		d.cache.Add(id, f)
	}
	return f
}
