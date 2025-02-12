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
	"fmt"
	"time"

	f_log "github.com/transparency-dev/formats/log"
	"github.com/transparency-dev/trillian-tessera/api/layout"
)

// LogReader provides read-only access to the log.
type LogReader interface {
	// ReadCheckpoint returns the latest checkpoint available.
	ReadCheckpoint(ctx context.Context) ([]byte, error)

	// ReadTile returns the raw marshalled tile at the given coordinates, if it exists.
	// The expected usage for this method is to derive the parameters from a tree size
	// that has been committed to by a checkpoint returned by this log. Whenever such a
	// tree size is used, this method will behave as per the https://c2sp.org/tlog-tiles
	// spec for the /tile/ path.
	//
	// If callers pass in parameters that are not implied by a published tree size, then
	// implementations _may_ act differently from one another, but all will act in ways
	// that are allowed by the spec. For example, if the only published tree size has been
	// for size 2, then asking for a partial tile of 1 may lead to some implementations
	// returning not found, some may return a tile with 1 leaf, and some may return a tile
	// with more leaves.
	ReadTile(ctx context.Context, level, index uint64, p uint8) ([]byte, error)

	// ReadEntryBundle returns the raw marshalled leaf bundle at the given coordinates, if
	// it exists.
	// The expected usage and corresponding behaviours are similar to ReadTile.
	ReadEntryBundle(ctx context.Context, index uint64, p uint8) ([]byte, error)
}

// Appender allows personalities access to the lifecycle methods associated with logs
// in sequencing mode. This only has a single method, but other methods are likely to be added
// such as a Shutdown method for #341.
type Appender struct {
	Add AddFn
	// TODO(#341): add this method and implement it in all drivers
	// Shutdown func(ctx context.Context)
}

type AppenderOptionFn func(*AppendOptions)

// NewAppender returns an Appender, which allows a personality to incrementally append new
// leaves to the log and to read from it.
//
// decorators provides a list of optional constructor functions that will return decorators
// that wrap the base appender. This can be used to provide deduplication. Decorators will be
// called in-order, and the last in the chain will be the base appender.
// TODO(mhutchinson): switch the decorators over to a WithOpt for future flexibility.
func NewAppender(d Driver, opts ...AppenderOptionFn) (*Appender, LogReader, error) {
	resolved := resolveAppendOptions(opts...)
	type appender interface {
		Add(ctx context.Context, entry *Entry) IndexFuture
		// Shutdown(ctx context.Context)
	}
	a, ok := d.(appender)
	if !ok {
		return nil, nil, fmt.Errorf("driver %T does not implement Appender", d)
	}
	add := a.Add
	for i := len(resolved.AddDecorators) - 1; i >= 0; i-- {
		add = resolved.AddDecorators[i](add)
	}
	reader, ok := d.(LogReader)
	if !ok {
		return nil, nil, fmt.Errorf("driver %T does not implement LogReader", d)
	}
	return &Appender{
		Add: add,
		// Shutdown: a.Shutdown,
	}, reader, nil
}

func WithAppendDeduplication(decorators ...func(AddFn) AddFn) func(*AppendOptions) {
	return func(o *AppendOptions) {
		o.AddDecorators = decorators
	}
}

// NewCPFunc is the signature of a function which knows how to format and sign checkpoints.
type NewCPFunc func(size uint64, hash []byte) ([]byte, error)

// ParseCPFunc is the signature of a function which knows how to verify and parse checkpoints.
type ParseCPFunc func(raw []byte) (*f_log.Checkpoint, error)

// EntriesPathFunc is the signature of a function which knows how to format entry bundle paths.
type EntriesPathFunc func(n uint64, p uint8) string

// AppendOptions holds optional settings for all storage implementations.
type AppendOptions struct {
	NewCP NewCPFunc

	BatchMaxAge  time.Duration
	BatchMaxSize uint

	PushbackMaxOutstanding uint

	EntriesPath EntriesPathFunc

	CheckpointInterval time.Duration

	AddDecorators []func(AddFn) AddFn
}

// resolveAppendOptions turns a variadic array of storage options into an AppendOptions instance.
func resolveAppendOptions(opts ...AppenderOptionFn) *AppendOptions {
	defaults := &AppendOptions{
		BatchMaxSize:           DefaultBatchMaxSize,
		BatchMaxAge:            DefaultBatchMaxAge,
		EntriesPath:            layout.EntriesPath,
		CheckpointInterval:     DefaultCheckpointInterval,
		AddDecorators:          make([]func(AddFn) AddFn, 0),
		PushbackMaxOutstanding: DefaultPushbackMaxOutstanding,
	}
	for _, opt := range opts {
		opt(defaults)
	}
	return defaults
}
