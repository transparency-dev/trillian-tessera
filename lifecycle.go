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

	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/storage/antispam"
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

// LogFollower provides read-only access to the log with an API tailored to bulk in-order
// reads of all integrated entry bundles.
//
// This API allows the storage implementation to tailor its approach to
type LogFollower interface {
	// IntegratedSize returns the size of the currently integrated tree.
	// Note that this _may_ be larger than the currently _published_ checkpoint.
	IntegratedSize(ctx context.Context) (uint64, error)

	// StreamEntryBundles returns functions which act like a pull iterator for subsequent entry bundles starting at the given index.
	//
	// Implementations must:
	//  - truncate the requested range if any or all of it is beyond the extent of the currently integrated tree.
	//  - cease iterating if next() produces an error, or cancel is called. next should continue to return an error if called again after either of these cases.
	StreamEntryRange(ctx context.Context, fromIdx, N, treeSize uint64) (next func() (layout.RangeInfo, []byte, error), cancel func())
}

// Appender allows personalities access to the lifecycle methods associated with logs
// in sequencing mode. This only has a single method, but other methods are likely to be added
// such as a Shutdown method for #341.
type Appender struct {
	Add AddFn
	// TODO(#341): add this method and implement it in all drivers
	// Shutdown func(ctx context.Context)
}

// NewAppender returns an Appender, which allows a personality to incrementally append new
// leaves to the log and to read from it.
//
// decorators provides a list of optional constructor functions that will return decorators
// that wrap the base appender. This can be used to provide deduplication. Decorators will be
// called in-order, and the last in the chain will be the base appender.
func NewAppender(ctx context.Context, d Driver, opts ...func(*AppendOptions)) (*Appender, LogReader, error) {
	resolved := resolveAppendOptions(opts...)
	type appendLifecycle interface {
		Appender(context.Context, *AppendOptions) (*Appender, LogReader, error)
	}
	lc, ok := d.(appendLifecycle)
	if !ok {
		return nil, nil, fmt.Errorf("driver %T does not implement Appender lifecycle", d)
	}
	a, r, err := lc.Appender(ctx, resolved)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to init appender lifecycle: %v", err)
	}
	for i := len(resolved.AddDecorators) - 1; i >= 0; i-- {
		a.Add = resolved.AddDecorators[i](a.Add)
	}
	return a, r, nil
}

type Antispam interface {
	AddDecorator() func(AddFn) AddFn
	Populate(context.Context, antispam.LogFollower, func(entryBundle []byte) ([][]byte, error))
}

func WithAntispam(inMemEntries uint, as Antispam) func(*AppendOptions) {
	return func(o *AppendOptions) {
		o.AddDecorators = append(o.AddDecorators, InMemoryDedupe(inMemEntries))
		if as != nil {
			o.AddDecorators = append(o.AddDecorators, as.AddDecorator())
			o.Followers = append(o.Followers, func(ctx context.Context, lf LogFollower) error {
				return as.Populate(ctx, lf, idHasher)
			})
		}
	}
}

// AppendOptions holds settings for all storage implementations.
type AppendOptions struct {
	// NewCP knows how to format and sign checkpoints.
	NewCP func(size uint64, hash []byte) ([]byte, error)

	BatchMaxAge  time.Duration
	BatchMaxSize uint

	PushbackMaxOutstanding uint

	// EntriesPath knows how to format entry bundle paths.
	EntriesPath func(n uint64, p uint8) string

	CheckpointInterval time.Duration

	AddDecorators []func(AddFn) AddFn
	Followers     []func(context.Context, LogFollower) error
}

// resolveAppendOptions turns a variadic array of storage options into an AppendOptions instance.
func resolveAppendOptions(opts ...func(*AppendOptions)) *AppendOptions {
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
