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
)

// Queue knows how to queue up a number of entries before calling a FlushFunc with
// a slice of all queued entries, in the same order as they were added, after either
// a defined period of time has passed, or a defined number of entries were added.
type Queue struct {
	buf   *buffer.Buffer
	flush FlushFunc

	inFlightMu sync.Mutex
	inFlight   map[string][]chan Index
}

// FlushFunc is the signature of a function which will receive the slice of queued entries.
// It should return the index assigned to the first entry in the provided slice.
type FlushFunc func(ctx context.Context, entries [][]byte) (index uint64, err error)

// NewQueue creates a new queue with the specified maximum age and size.
//
// The provided FlushFunc will be called with a slice containing the contents of the queue, in
// the same order as they were added, when either the oldest entry in the queue has been there
// for maxAge, or the size of the queue reaches maxSize.
func NewQueue(maxAge time.Duration, maxSize uint, f FlushFunc) *Queue {
	q := &Queue{
		flush:    f,
		inFlight: make(map[string][]chan Index, maxSize),
	}
	q.buf = buffer.New(
		buffer.WithSize(maxSize),
		buffer.WithFlushInterval(maxAge),
		buffer.WithFlusher(buffer.FlusherFunc(q.doFlush)),
	)
	return q
}

// Index represents the index assigned to an entry by the FlushFunc, or an error.
type Index struct {
	N   uint64
	Err error
}

// squashDupes keeps track of all in-flight requests, enabling dupe squashing for entries currently in the queue.
// Returns true if the provided entry is a dupe and should NOT be added to the queue.
func (q *Queue) squashDupes(e entry) bool {
	q.inFlightMu.Lock()
	defer q.inFlightMu.Unlock()

	k := string(e.data)
	l, isKnown := q.inFlight[k]
	q.inFlight[k] = append(l, e.index)
	return isKnown
}

func (q *Queue) Add(ctx context.Context, e []byte) <-chan Index {
	entry := entry{
		data:  e,
		index: make(chan Index, 1),
	}
	if q.squashDupes(entry) {
		// This entry is already in the queue, so no need to add it again.
		return entry.index
	}
	if err := q.buf.Push(entry); err != nil {
		entry.index <- Index{Err: err}
		close(entry.index)
	}
	return entry.index
}

func (q *Queue) doFlush(items []interface{}) {
	entries := make([]entry, len(items))
	entriesData := make([][]byte, len(items))
	for i, t := range items {
		entries[i] = t.(entry)
		entriesData[i] = entries[i].data
	}

	go func() {
		s, err := q.flush(context.TODO(), entriesData)

		// Send assigned indices to all the waiting Add() requests, including dupes.
		q.inFlightMu.Lock()
		defer q.inFlightMu.Unlock()

		for i, e := range entries {
			for _, dd := range q.inFlight[string(e.data)] {
				dd <- Index{N: s + uint64(i), Err: err}
				close(dd)
			}
		}
	}()

}

type entry struct {
	data  []byte
	index chan Index
}
