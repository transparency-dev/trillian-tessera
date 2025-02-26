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
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
)

// NoMoreEntries is a sentinel error returned by StreamEntries when no more entries will be returned by calls to the next function.
var ErrNoMoreEntries = errors.New("no more entries")

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

	// IntegratedSize returns the current size of the integrated tree.
	//
	// This tree will have in place all the static resources the returned size implies, but
	// there may not yet be a checkpoint for this size signed, witnessed, or published.
	//
	// It's ONLY safe to use this value for processes internal to the operation of the log (e.g.
	// populating antispam data structures); it MUST NOT not be used as a substitute for
	// reading the checkpoint when only data which has been publicly committed to by the
	// log should be used. If in doubt, use ReadCheckpoint instead.
	IntegratedSize(ctx context.Context) (uint64, error)

	// StreamEntries() returns functions `next` and `stop` which act like a pull iterator for
	// consecutive entry bundles, starting with the entry bundle which contains the requested entry
	// index.
	//
	// Each call to `next` will return raw entry bundle bytes along with a RangeInfo struct which
	// contains information on which entries within that bundle are to be considered valid.
	//
	// next will hang if it has reached the extent of the current tree, and return once either
	// the tree has grown and more entries are available, or cancel was called.
	//
	// next will cease iterating if either:
	//   - it produces an error (e.g. via the underlying calls to the log storage)
	//   - the returned cancel function is called
	// and will continue to return an error if called again after either of these cases.
	StreamEntries(ctx context.Context, fromEntryIdx uint64) (next func() (layout.RangeInfo, []byte, error), cancel func())
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
	for _, f := range resolved.Followers {
		go f(ctx, r)
	}
	return a, r, nil
}

// Antispam describes the contract that an antispam implementation must meet in order to be used via the
// WithAntispam option below.
type Antispam interface {
	// Decorator must return a function which knows how to decorate an Appender's Add function in order
	// to return an index previously assigned to an entry with the same identity hash, if one exists, or
	// delegate to the next Add function in the chain otherwise.
	Decorator() func(AddFn) AddFn
	// Populate should be a long-running function which uses the provided log reader to build the
	// antispam index used by the dectorator above.
	//
	// Typically, implementations of this function will tail the contents of the log using the provided
	// log reader to stream entry bundles from the log, and, for each entry bundle, use the
	// provided bundle function to convert the bundle into a slice of identity hashes which
	// corresponds to the entries the bundle contains. These hashes should then be used to populate
	// some form of an identity hash -> index mapping.
	//
	// This function will be called automatically by Tessera, and is expected to block until the context
	// is done.
	Populate(context.Context, LogReader, func(entryBundle []byte) ([][]byte, error))
}

func WithAntispam(inMemEntries uint, as Antispam) func(*AppendOptions) {
	return func(o *AppendOptions) {
		o.AddDecorators = append(o.AddDecorators, InMemoryDedupe(inMemEntries))
		if as != nil {
			o.AddDecorators = append(o.AddDecorators, as.Decorator())
			o.Followers = append(o.Followers, func(ctx context.Context, lr LogReader) {
				as.Populate(ctx, lr, defaultIDHasher)
			})
		}
	}
}

// defaultIDHasher returns a list of identity hashes corresponding to entries in the provided bundle.
// Currently, these are simply SHA256 hashes of the raw byte of each entry.
func defaultIDHasher(bundle []byte) ([][]byte, error) {
	eb := &api.EntryBundle{}
	if err := eb.UnmarshalText(bundle); err != nil {
		return nil, fmt.Errorf("unmarshal: %v", err)
	}
	r := make([][]byte, 0, len(eb.Entries))
	for _, e := range eb.Entries {
		h := sha256.Sum256(e)
		r = append(r, h[:])
	}
	return r, nil
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
	Witnesses          WitnessGroup

	AddDecorators []func(AddFn) AddFn
	Followers     []func(context.Context, LogReader)
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

// MigrationTarget describes the contract of the Migration lifecycle.
//
// This lifecycle mode is used to migrate C2SP tlog-tiles and static-ct
// compliant logs into Tessera.
type MigrationTarget interface {
	// SetEntryBundle stores the provided serialised entry bundle at the location implied by the provided
	// entry bundle index and partial size.
	//
	// Bundles may be set in any order (not just consecutively), and the implementation should integrate
	// them into the local tree in the most efficient way possible.
	//
	// Writes should be idempotent; repeated calls to set the same bundle with the same data should not
	// return an error.
	SetEntryBundle(ctx context.Context, idx uint64, partial uint8, bundle []byte) error
	// AwaitIntegration should block until the local integrated tree has grown to the provided size,
	// and should return the locally calculated root hash derived from the integration of the contents of
	// entry bundles set using SetEntryBundle above.
	AwaitIntegration(ctx context.Context, size uint64) ([]byte, error)
	// IntegratedSize returns the current size of the locally integrated log.
	IntegratedSize(ctx context.Context) (uint64, error)
}

// UnbundlerFunc is a function which knows how to turn a serialised entry bundle into a slice of
// []byte representing each of the entries within the bundle.
type UnbundlerFunc func(entryBundle []byte) ([][]byte, error)

// NewMigrationTarget returns a MigrationTarget, which allows a personality to "import" a C2SP
// tlog-tiles or static-ct compliant log into a Tessera instance.
//
// TODO(al): bundleHasher should be implicit from WithCTLayout being present or not.
// TODO(al): AppendOptions should be somehow replaced - perhaps MigrationOptions, or some other way of limiting options to those which make sense for this lifecycle mode.
func NewMigrationTarget(ctx context.Context, d Driver, bundleHasher UnbundlerFunc, opts ...func(*AppendOptions)) (MigrationTarget, error) {
	resolved := resolveAppendOptions(opts...)
	type migrateLifecycle interface {
		MigrationTarget(context.Context, UnbundlerFunc, *AppendOptions) (MigrationTarget, LogReader, error)
	}
	lc, ok := d.(migrateLifecycle)
	if !ok {
		return nil, fmt.Errorf("driver %T does not implement MigrationTarget lifecycle", d)
	}
	m, r, err := lc.MigrationTarget(ctx, bundleHasher, resolved)
	if err != nil {
		return nil, fmt.Errorf("failed to init MigrationTarget lifecycle: %v", err)
	}
	for _, f := range resolved.Followers {
		go f(ctx, r)
	}
	return m, nil
}
