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
	"time"

	"github.com/globocom/go-buffer"
)

// Queue knows how to queue up a number of entries before calling a FlushFunc with
// a slice of all queued entries, in the same order as they were added, after either
// a defined period of time has passed, or a defined number of entries were added.
type Queue struct {
	buf   *buffer.Buffer
	flush FlushFunc
}

// FlushFunc is the signature of a function which will receive the slice of queued entries.
// It should return the index assigned to the first entry in the provided slice.
type FlushFunc func(ctx context.Context, entries [][]byte) (index uint64, err error)

func NewQueue(maxWait time.Duration, maxSize uint, f FlushFunc) *Queue {
	q := &Queue{
		flush: f,
	}
	q.buf = buffer.New(
		buffer.WithSize(maxSize),
		buffer.WithFlushInterval(maxWait),
		buffer.WithFlusher(buffer.FlusherFunc(q.doFlush)),
	)
	return q
}

type Index struct {
	N   uint64
	Err error
}

func (q *Queue) Add(ctx context.Context, e []byte) <-chan Index {
	entry := entry{
		data:  e,
		index: make(chan Index, 1),
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
		for i, e := range entries {
			e.index <- Index{N: s + uint64(i), Err: err}
			close(e.index)
		}
	}()

}

type entry struct {
	data  []byte
	index chan Index
}
