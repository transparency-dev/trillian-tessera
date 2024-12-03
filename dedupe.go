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
	"time"

	"github.com/globocom/go-buffer"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"k8s.io/klog/v2"
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

// DedupeStorage describes a struct which can store and retrieve hash to index mappings.
type DedupeStorage interface {
	// Set must persist the provided id->index mappings.
	// It must return an error if it fails to store one or more of the mappings, but it does not
	// have to store all mappings atomically.
	Set(context.Context, []DedupeMapping) error
	// Index must return the index of a previously stored ID, or nil if the ID is unknown.
	Index(context.Context, []byte) (*uint64, error)
}

// DedupeMapping represents an ID -> index mapping.
type DedupeMapping struct {
	ID  []byte
	Idx uint64
}

type persistentDedup struct {
	ctx      context.Context
	storage  DedupeStorage
	delegate func(ctx context.Context, e *Entry) IndexFuture

	mu          sync.Mutex
	numLookups  uint64
	numWrites   uint64
	numDBDedups uint64
	numPushErrs uint64

	buf *buffer.Buffer
}

// PersistentDedup returns a wrapped Add method which will return the previously seen index for the given entry, if such an entry exists
func PersistentDedupe(ctx context.Context, s DedupeStorage, delegate func(ctx context.Context, e *Entry) IndexFuture) func(context.Context, *Entry) IndexFuture {
	r := &persistentDedup{
		ctx:      ctx,
		storage:  s,
		delegate: delegate,
	}

	// TODO(al): Make these configurable
	r.buf = buffer.New(
		buffer.WithSize(64),
		buffer.WithFlushInterval(200*time.Millisecond),
		buffer.WithFlusher(buffer.FlusherFunc(r.flush)),
		buffer.WithPushTimeout(15*time.Second),
	)
	go func(ctx context.Context) {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				r.mu.Lock()
				klog.V(1).Infof("DEDUP: # Writes %d, # Lookups %d, # DB hits %v, # buffer Push discards %d", r.numWrites, r.numLookups, r.numDBDedups, r.numPushErrs)
				r.mu.Unlock()
			}
		}
	}(ctx)
	return r.add
}

// add adds the entry to the underlying delegate only if e isn't already known. In either case,
// an IndexFuture will be returned that the client can use to get the sequence number of this entry.
func (d *persistentDedup) add(ctx context.Context, e *Entry) IndexFuture {
	idx, err := d.Index(ctx, e.Identity())
	if err != nil {
		return func() (uint64, error) { return 0, err }
	}
	if idx != nil {
		return func() (uint64, error) { return *idx, nil }
	}

	i, err := d.delegate(ctx, e)()
	if err != nil {
		return func() (uint64, error) { return 0, err }
	}

	err = d.Set(ctx, e.Identity(), i)
	return func() (uint64, error) {
		return i, err
	}
}

func (d *persistentDedup) inc(p *uint64) {
	d.mu.Lock()
	defer d.mu.Unlock()
	(*p)++
}

func (d *persistentDedup) Index(ctx context.Context, h []byte) (*uint64, error) {
	d.inc(&d.numLookups)
	r, err := d.storage.Index(ctx, h)
	if r != nil {
		d.inc(&d.numDBDedups)
	}
	return r, err
}

func (d *persistentDedup) Set(_ context.Context, h []byte, idx uint64) error {
	err := d.buf.Push(DedupeMapping{ID: h, Idx: idx})
	if err != nil {
		d.inc(&d.numPushErrs)
		// This means there's pressure flushing dedup writes out, so discard this write.
		if err != buffer.ErrTimeout {
			return err
		}
	}
	return nil
}

func (d *persistentDedup) flush(items []interface{}) {
	entries := make([]DedupeMapping, len(items))
	for i := range items {
		entries[i] = items[i].(DedupeMapping)
	}

	ctx, c := context.WithTimeout(d.ctx, 15*time.Second)
	defer c()

	if err := d.storage.Set(ctx, entries); err != nil {
		klog.Infof("Failed to flush dedup entries: %v", err)
		return
	}
	d.mu.Lock()
	d.numWrites += uint64(len(entries))
	d.mu.Unlock()
}
