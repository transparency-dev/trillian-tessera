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

// Package driver contains the contract that driver implementations need to implement.
// A driver is an implementation of a log at a very high "black box" level. Keeping
// operations at a high functional level means that implementations can take care of their
// internals in the most idiomatic and performant way. The opposite approach to this is
// to force all drivers to have granular operations for sequencing, integration, creating
// a new checkpoint, etc. This is the approach that Trillian v1 took and led to a lot of
// complexity for maintainers and deployers.
package driver

import (
	"context"

	tessera "github.com/transparency-dev/trillian-tessera"
)

// Readers contains functions that allow read-only access to the log.
type Readers struct {
	// ReadCheckpoint returns the latest checkpoint available.
	// If no checkpoint is available then it must return os.ErrNotExist.
	ReadCheckpoint func(ctx context.Context) ([]byte, error)

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
	//
	// If no tile is available then it must return os.ErrNotExist.
	ReadTile func(ctx context.Context, level, index uint64, p uint8) ([]byte, error)

	// ReadEntryBundle returns the raw marshalled leaf bundle at the given coordinates, if
	// it exists.
	// The expected usage and corresponding behaviours are similar to ReadTile.
	//
	// If no entry bundle is available then it must return os.ErrNotExist.
	ReadEntryBundle func(ctx context.Context, index uint64, p uint8) ([]byte, error)
}

// Appenders contains functions that allow new entries to be appended to a log.
type Appenders struct {
	// Add should duably assign an index to the provided Entry, returning a future to access that value.
	//
	// Implementations MUST call MarshalBundleData method on the entry before persisting/integrating it.
	Add tessera.AddFn
}
