// Copyright 2024 Google LLC. All Rights Reserved.
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

// Package storage provides implementations and shared components for tessera storage backends.
package storage

import (
	"context"
	"sync"
	"time"

	"github.com/globocom/go-buffer"
	tessera "github.com/transparency-dev/trillian-tessera"
)

// Queue knows how to queue up a number of entries in order, taking care of deduplication as they're added.
//
// When the buffered queue grows past a defined size, or the age of the oldest entry in the
// queue reaches a defined threshold, the queue will call a provided FlushFunc with
// a slice containing all queued entries in the same order as they were added.
//
// If multiple identical entries are added to the queue between flushes, the queue will deduplicate them by
// passing only the first through to the FlushFunc, and returning the index assigned to that entry to all
// duplicate add calls.
// Note that this deduplication only applies to "in-flight" entries currently in the queue; entries added
// after a flush will not be deduped against those added before the flush.
type Queue struct {
	buf   *buffer.Buffer
	flush FlushFunc

	inFlightMu sync.Mutex
	inFlight   map[string]*entry
}

// IndexFunc is a function which returns an assigned log index, or an error.
type IndexFunc func() (idx uint64, err error)

// FlushFunc is the signature of a function which will receive the slice of queued entries.
// It should return the index assigned to the first entry in the provided slice.
type FlushFunc func(ctx context.Context, entries []tessera.Entry) (index uint64, err error)

// NewQueue creates a new queue with the specified maximum age and size.
//
// The provided FlushFunc will be called with a slice containing the contents of the queue, in
// the same order as they were added, when either the oldest entry in the queue has been there
// for maxAge, or the size of the queue reaches maxSize.
func NewQueue(ctx context.Context, maxAge time.Duration, maxSize uint, f FlushFunc) *Queue {
	q := &Queue{
		flush:    f,
		inFlight: make(map[string]*entry, maxSize),
	}

	// The underlying queue implementation blocks additions during a flush.
	// This blocks the filling of the next batch unnecessarily, so we'll
	// decouple the queue flush and storage write by handling the later in
	// a worker goroutine.
	// This same worker thread will also handle the callbacks to f.
	work := make(chan []*entry, 1)
	toWork := func(items []interface{}) {
		entries := make([]*entry, len(items))
		for i, t := range items {
			entries[i] = t.(*entry)
		}
		work <- entries

	}

	q.buf = buffer.New(
		buffer.WithSize(maxSize),
		buffer.WithFlushInterval(maxAge),
		buffer.WithFlusher(buffer.FlusherFunc(toWork)),
	)

	// This will allow
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case entries := <-work:
				q.doFlush(ctx, entries)
			}
		}
	}(ctx)
	return q
}

// squashDupes keeps track of all in-flight requests, enabling dupe squashing for entries currently in the queue.
// Returns an entry struct, and a bool which is true if the provided entry is a dupe and should NOT be added to the queue.
func (q *Queue) squashDupes(e tessera.Entry) (*entry, bool) {
	q.inFlightMu.Lock()
	defer q.inFlightMu.Unlock()

	k := string(e.Identity())
	entry, isKnown := q.inFlight[k]
	if !isKnown {
		entry = newEntry(e)
		q.inFlight[k] = entry
	}
	return entry, isKnown
}

// Add places e into the queue, and returns a func which may be called to retrieve the assigned index.
func (q *Queue) Add(ctx context.Context, e tessera.Entry) IndexFunc {
	entry, isDupe := q.squashDupes(e)
	if isDupe {
		// This entry is already in the queue, so no need to add it again.
		return entry.index
	}
	if err := q.buf.Push(entry); err != nil {
		entry.assign(0, err)
	}
	return entry.index
}

// doFlush handles the queue flush, and sending notifications of assigned log indices.
//
// To prevent blocking the queue longer than necessary, the notifications happen in a
// separate goroutine.
func (q *Queue) doFlush(ctx context.Context, entries []*entry) {
	entriesData := make([]tessera.Entry, 0, len(entries))
	for _, e := range entries {
		entriesData = append(entriesData, e.data)
	}

	s, err := q.flush(ctx, entriesData)

	// Send assigned indices to all the waiting Add() requests, including dupes.
	q.inFlightMu.Lock()
	defer q.inFlightMu.Unlock()

	for i, e := range entries {
		e.assign(s+uint64(i), err)
		k := string(e.data.Identity())
		delete(q.inFlight, k)
	}
}

// entry represents an in-flight entry in the queue.
//
// The index field acts as a Future for the entry's assigned index/error, and will
// hang until assign is called.
type entry struct {
	data  tessera.Entry
	c     chan IndexFunc
	index IndexFunc
}

// newEntry creates a new entry for the provided data.
func newEntry(data tessera.Entry) *entry {
	e := &entry{
		data: data,
		c:    make(chan IndexFunc, 1),
	}
	e.index = sync.OnceValues(func() (uint64, error) {
		return (<-e.c)()
	})
	return e
}

// assign sets the assigned log index (or an error) to the entry.
//
// This func must only be called once, and will cause any current or future callers of index()
// to be given the values provided here.
func (e *entry) assign(idx uint64, err error) {
	e.c <- func() (uint64, error) {
		return idx, err
	}
	close(e.c)
}
