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
	"fmt"

	"github.com/transparency-dev/merkle/rfc6962"
	"github.com/transparency-dev/tessera/api"
	"github.com/transparency-dev/tessera/api/layout"
	"github.com/transparency-dev/tessera/internal/stream"
)

// NoMoreEntries is a sentinel error returned by StreamEntries when no more entries will be returned by calls to the next function.
var ErrNoMoreEntries = stream.ErrNoMoreEntries

// LogReader provides read-only access to the log.
type LogReader interface {
	// ReadCheckpoint returns the latest checkpoint available.
	// If no checkpoint is available then os.ErrNotExist should be returned.
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

	// NextIndex returns the first as-yet unassigned index.
	//
	// In a quiescent log, this will be the same as the checkpoint size. In a log with entries actively
	// being added, this number will be higher since it will take sequenced but not-yet-integrated/not-yet-published
	// entries into account.
	NextIndex(ctx context.Context) (uint64, error)

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

// Antispam is the public type for antispam implementations. Clients are expected to create one of these
// and then pass it back in to Tessera, but are not expected to call any methods on it, nor implement it.
// For this reason, we use `any` in order to leave the API flexible.
//
// Intended to be used via WithAntispam option.
type Antispam any

// antispam is the internal interface for Antispam. This is only intended for implementation and usage
// within Tessera.
type antispam interface {
	// Decorator must return a function which knows how to decorate an Appender's Add function in order
	// to return an index previously assigned to an entry with the same identity hash, if one exists, or
	// delegate to the next Add function in the chain otherwise.
	Decorator() func(AddFn) AddFn
	// Follower should return a structure which will populate the anti-spam index by tailing the contents
	// of the log, using the provided function to turn entry bundles into identity hashes.
	Follower(func(entryBundle []byte) ([][]byte, error)) stream.Follower
}

// identityHash calculates the antispam identity hash for the provided (single) leaf entry data.
func identityHash(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
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
		h := identityHash(e)
		r = append(r, h[:])
	}
	return r, nil
}

// defaultMerkleLeafHasher parses a C2SP tlog-tile bundle and returns the Merkle leaf hashes of each entry it contains.
func defaultMerkleLeafHasher(bundle []byte) ([][]byte, error) {
	eb := &api.EntryBundle{}
	if err := eb.UnmarshalText(bundle); err != nil {
		return nil, fmt.Errorf("unmarshal: %v", err)
	}
	r := make([][]byte, 0, len(eb.Entries))
	for _, e := range eb.Entries {
		h := rfc6962.DefaultHasher.HashLeaf(e)
		r = append(r, h[:])
	}
	return r, nil
}
