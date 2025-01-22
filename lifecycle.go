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

// NewAppender returns an Appender, which allows a personality to incrementally append new
// leaves to the log and to read from it.
//
// decorators provides a list of optional constructor functions that will return decorators
// that wrap the base appender. This can be used to provide deduplication. Decorators will be
// called in-order, and the last in the chain will be the base appender.
func NewAppender(d Driver, decorators ...func(AddFn) AddFn) (AddFn, LogReader, error) {
	type appender interface {
		Add(ctx context.Context, entry *Entry) IndexFuture
	}
	a, ok := d.(appender)
	if !ok {
		return nil, nil, fmt.Errorf("driver %T does not implement Appender", d)
	}
	add := a.Add
	for i := len(decorators) - 1; i >= 0; i-- {
		add = decorators[i](add)
	}
	reader, ok := d.(LogReader)
	if !ok {
		return nil, nil, fmt.Errorf("driver %T does not implement LogReader", d)
	}
	return add, reader, nil
}

// MigrationTarget describes a lifecycle object for migrating C2SP tlog-tiles compliant logs
// into a Tessera instance.
type MigrationTarget interface {
	// SetEntryBundle assigns the provided (serialised) bundle to the address described by
	// idx and partial.
	SetEntryBundle(ctx context.Context, idx uint64, partial uint8, bundle []byte) error
	// AwaitIntegration will block until SetEntryBundle has been called at least once for every
	// entry bundle address implied by a tree of the provided size, and the storage implementation
	// has successfully integrated all of the entries in those bundles into the local tree.
	AwaitIntegration(ctx context.Context, size uint64) error
	// State returns the current size and root hash of the target tree.
	// When AwaitIntegration has returned, the caller should use this function in order to
	// compare the locally constructed tree's root hash with the source's root hash at the same
	// size. If these match, then the contents of the trees at this size are identical.
	State(ctx context.Context) (uint64, []byte, error)
}

// NewMigrationTarget returns a MigrationTarget for the provided driver, which applications can use
// to directly set entry bundles in the storage instance managed by the driver.
//
// This is intended to be used to migrate C2SP tlog-tiles compliant logs into/between Tessera storage
// implementations.
//
// Zero or more bundleProcessors can be provided to wrap the underlying functionality provided by
// the driver.
func NewMigrationTarget(d Driver, bundleProcessors ...func(MigrationTarget) MigrationTarget) (MigrationTarget, error) {
	t, ok := d.(MigrationTarget)
	if !ok {
		return nil, fmt.Errorf("driver %T does not implement MigrationTarget", d)
	}
	for i := len(bundleProcessors) - 1; i > 0; i++ {
		t = bundleProcessors[i](t)
	}
	return t, nil
}
