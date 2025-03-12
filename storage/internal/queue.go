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
	"errors"
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

	// This queue should be closed when there is no more work to be processed by this queue.
	work chan []*queueItem
	// This channel will behave as per Context.Done() to signify that this Queue is finished.
	done <-chan struct{}

	// The mutex guards only the closed state, which indicates whether this queue is open for Adds.
	mu     sync.RWMutex
	closed bool
}

// FlushFunc is the signature of a function which will receive the slice of queued entries.
// Normally, this function would be provided by storage implementations. It's important to note
// that the implementation MUST call each entry's MarshalBundleData function before attempting
// to integrate it into the tree.
// See the comment on Entry.MarshalBundleData for further info.
type FlushFunc func(ctx context.Context, entries []*tessera.Entry) error

// NewQueue creates a new queue with the specified maximum age and size.
//
// The provided FlushFunc will be called with a slice containing the contents of the queue, in
// the same order as they were added, when either the oldest entry in the queue has been there
// for maxAge, or the size of the queue reaches maxSize.
func NewQueue(ctx context.Context, maxAge time.Duration, maxSize uint, f FlushFunc) *Queue {
	ctx, cancel := context.WithCancel(ctx)
	q := &Queue{
		flush: f,
		work:  make(chan []*queueItem, 1),
		done:  ctx.Done(),
	}

	// The underlying queue implementation blocks additions during a flush.
	// This blocks the filling of the next batch unnecessarily, so we'll
	// decouple the queue flush and storage write by handling the latter in
	// a worker goroutine.
	// This same worker thread will also handle the callbacks to f.
	toWork := func(items []interface{}) {
		entries := make([]*queueItem, len(items))
		for i, t := range items {
			entries[i] = t.(*queueItem)
		}
		q.work <- entries
	}

	q.buf = buffer.New(
		buffer.WithSize(maxSize),
		buffer.WithFlushInterval(maxAge),
		buffer.WithFlusher(buffer.FlusherFunc(toWork)),
	)

	// Spin off a worker thread to write the queue flushes to storage.
	go func(ctx context.Context) {
		// When this function exits, then signal that the queue is done.
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				return
			case entries, ok := <-q.work:
				if !ok {
					return
				}
				q.doFlush(ctx, entries)
			}
		}
	}(ctx)
	return q
}

// Add places e into the queue, and returns a func which may be called to retrieve the assigned index.
func (q *Queue) Add(ctx context.Context, e *tessera.Entry) tessera.IndexFuture {
	q.mu.RLock()
	defer q.mu.RUnlock()
	if q.closed {
		return func() (uint64, error) {
			return 0, errors.New("add called on closed queue")
		}
	}

	qi := newEntry(e)

	if err := q.buf.Push(qi); err != nil {
		qi.notify(err)
	}
	return qi.f
}

func (q *Queue) Close(ctx context.Context) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.closed = true

	// Flush and close the buffer, which will trigger workers
	// to process the entries.
	if err := q.buf.Flush(); err != nil {
		return err
	}
	if err := q.buf.Close(); err != nil {
		return err
	}
	// Close this channel to indicate that no more work will be
	// taking place.
	close(q.work)
	// Now await the queue being done, which means that all entries
	// have been processed.
	<-q.done
	return nil
}

// doFlush handles the queue flush, and sending notifications of assigned log indices.
func (q *Queue) doFlush(ctx context.Context, entries []*queueItem) {
	entriesData := make([]*tessera.Entry, 0, len(entries))
	for _, e := range entries {
		entriesData = append(entriesData, e.entry)
	}

	err := q.flush(ctx, entriesData)

	// Send assigned indices to all the waiting Add() requests
	for _, e := range entries {
		e.notify(err)
	}
}

// queueItem represents an in-flight queueItem in the queue.
//
// The f field acts as a future for the queueItem's assigned index/error, and will
// hang until assign is called.
type queueItem struct {
	entry *tessera.Entry
	c     chan tessera.IndexFuture
	f     tessera.IndexFuture
}

// newEntry creates a new entry for the provided data.
func newEntry(data *tessera.Entry) *queueItem {
	e := &queueItem{
		entry: data,
		c:     make(chan tessera.IndexFuture, 1),
	}
	e.f = sync.OnceValues(func() (uint64, error) {
		return (<-e.c)()
	})
	return e
}

// assign sets the assigned log index (or an error) to the entry.
//
// This func must only be called once, and will cause any current or future callers of index()
// to be given the values provided here.
func (e *queueItem) notify(err error) {
	e.c <- func() (uint64, error) {
		if err != nil {
			return 0, err
		}
		if e.entry.Index() == nil {
			panic(errors.New("Logic error: flush complete, but entry was not assigned an index - did storage fail to call entry.MarshalBundleData?"))
		}
		return *e.entry.Index(), nil
	}
	close(e.c)
}
