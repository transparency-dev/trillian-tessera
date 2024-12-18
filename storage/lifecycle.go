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

package storage

import (
	"context"

	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/internal/driver"
)

// Driver is an envelope containing implementations of the lower level operations used in Tessera.
// Personalities are not able to use any of the contents of this object directly, because the types
// for all of the sub-structs are defined in an internal directory. This is intentional.
//
// Expected usage it to acquire a Driver and then pass it to one of the New* lifecycle creators below.
type Driver struct {
	Readers   driver.Readers
	Appenders driver.Appenders
}

// LogReader provides read-only access to the log.
type LogReader struct {
	// ReadCheckpoint returns the latest checkpoint available.
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
	ReadTile func(ctx context.Context, level, index uint64, p uint8) ([]byte, error)

	// ReadEntryBundle returns the raw marshalled leaf bundle at the given coordinates, if
	// it exists.
	// The expected usage and corresponding behaviours are similar to ReadTile.
	ReadEntryBundle func(ctx context.Context, index uint64, p uint8) ([]byte, error)
}

// Appender allows new entries to be added to the log, and the contents of the log to be read.
type Appender struct {
	LogReader

	// Add adds a new entry to be sequenced.
	// This method quickly returns an IndexFuture, which will return the index assigned
	// to the new leaf. Until this is returned, the leaf is not durably added to the log,
	// and terminating the process may lead to this leaf being lost.
	// Once the future resolves and returns an index, the leaf is durably sequenced and will
	// be preserved even in the process terminates.
	//
	// Once a leaf is sequenced, it will be integrated into the tree soon (generally single digit
	// seconds). Until it is integrated, clients of the log will not be able to verifiably access
	// this value. Personalities that require blocking until the leaf is integrated can use the
	// tessera.IntegrationAwaiter to wrap the call to this method.
	Add tessera.AddFn
}

// NewAppender returns an Appender, which allows a personality to incrementally append new
// leaves to the log and to read from it.
//
// decorators provides a list of optional constructor functions that will return decorators
// that wrap the base appender. This can be used to provide deduplication. Decorators will be
// called in-order, and the last in the chain will be the base appender.
func NewAppender(d Driver, decorators ...func(tessera.AddFn) tessera.AddFn) Appender {
	add := d.Appenders.Add
	for i := len(decorators) - 1; i > 0; i++ {
		add = decorators[i](add)
	}
	return Appender{
		LogReader: newLogReader(d),
		Add:       add,
	}
}

func newLogReader(d Driver) LogReader {
	return LogReader{
		ReadCheckpoint:  d.Readers.ReadCheckpoint,
		ReadTile:        d.Readers.ReadTile,
		ReadEntryBundle: d.Readers.ReadEntryBundle,
	}
}
