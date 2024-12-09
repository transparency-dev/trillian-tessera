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

// Package client provides client support for interacting with logs that
// uses the [tlog-tiles API].
//
// [tlog-tiles API]: https://c2sp.org/tlog-tiles
package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/transparency-dev/formats/log"
	"github.com/transparency-dev/merkle/compact"
	"github.com/transparency-dev/merkle/proof"
	"github.com/transparency-dev/merkle/rfc6962"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"golang.org/x/mod/sumdb/note"
)

var (
	hasher = rfc6962.DefaultHasher
)

// CheckpointFetcherFunc is the signature of a function which can retrieve the latest
// checkpoint from a log's data storage.
//
// Note that the implementation of this MUST return (either directly or wrapped)
// an os.ErrIsNotExist when the file referenced by path does not exist, e.g. a HTTP
// based implementation MUST return this error when it receives a 404 StatusCode.
type CheckpointFetcherFunc func(ctx context.Context) ([]byte, error)

// TileFetcherFunc is the signature of a function which can fetch the raw data
// for a given tile.
//
// Note that the implementation of this MUST return (either directly or wrapped)
// an os.ErrIsNotExist when the file referenced by path does not exist, e.g. a HTTP
// based implementation MUST return this error when it receives a 404 StatusCode.
type TileFetcherFunc func(ctx context.Context, level, index uint64, p uint8) ([]byte, error)

// EntryBundleFetcherFunc is the signature of a function which can fetch the raw data
// for a given entry bundle.
//
// Note that the implementation of this MUST return (either directly or wrapped)
// an os.ErrIsNotExist when the file referenced by path does not exist, e.g. a HTTP
// based implementation MUST return this error when it receives a 404 StatusCode.
type EntryBundleFetcherFunc func(ctx context.Context, bundleIndex uint64, p uint8) ([]byte, error)

// ConsensusCheckpointFunc is a function which returns the largest checkpoint known which is
// signed by logSigV and satisfies some consensus algorithm.
//
// This is intended to provide a hook for adding a consensus view of a log, e.g. via witnessing.
type ConsensusCheckpointFunc func(ctx context.Context, logSigV note.Verifier, origin string) (*log.Checkpoint, []byte, *note.Note, error)

// UnilateralConsensus blindly trusts the source log, returning the checkpoint it provided.
func UnilateralConsensus(f CheckpointFetcherFunc) ConsensusCheckpointFunc {
	return func(ctx context.Context, logSigV note.Verifier, origin string) (*log.Checkpoint, []byte, *note.Note, error) {
		return FetchCheckpoint(ctx, f, logSigV, origin)
	}
}

// FetchCheckpoint retrieves and opens a checkpoint from the log.
// Returns both the parsed structure and the raw serialised checkpoint.
func FetchCheckpoint(ctx context.Context, f CheckpointFetcherFunc, v note.Verifier, origin string) (*log.Checkpoint, []byte, *note.Note, error) {
	cpRaw, err := f(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	cp, _, n, err := log.ParseCheckpoint(cpRaw, origin, v)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to parse Checkpoint: %v", err)
	}
	return cp, cpRaw, n, nil
}

// ProofBuilder knows how to build inclusion and consistency proofs from tiles.
// Since the tiles commit only to immutable nodes, the job of building proofs is slightly
// more complex as proofs can touch "ephemeral" nodes, so these need to be synthesized.
type ProofBuilder struct {
	cp        log.Checkpoint
	nodeCache nodeCache
}

// NewProofBuilder creates a new ProofBuilder object for a given tree size.
// The returned ProofBuilder can be re-used for proofs related to a given tree size, but
// it is not thread-safe and should not be accessed concurrently.
func NewProofBuilder(ctx context.Context, cp log.Checkpoint, f TileFetcherFunc) (*ProofBuilder, error) {
	pb := &ProofBuilder{
		cp:        cp,
		nodeCache: newNodeCache(f, cp.Size),
	}
	// Can't re-create the root of a zero size checkpoint other than by convention,
	// so return early here in that case.
	if cp.Size == 0 {
		return pb, nil
	}

	hashes, err := FetchRangeNodes(ctx, cp.Size, f)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch range nodes: %w", err)
	}
	// Create a compact range which represents the state of the log.
	r, err := (&compact.RangeFactory{Hash: hasher.HashChildren}).NewRange(0, cp.Size, hashes)
	if err != nil {
		return nil, err
	}

	// Recreate the root hash so that:
	// a) we validate the self-integrity of the log state, and
	// b) we calculate (and cache) and ephemeral nodes present in the tree,
	//    this is important since they could be required by proofs.
	sr, err := r.GetRootHash(pb.nodeCache.SetEphemeralNode)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(cp.Hash, sr) {
		return nil, fmt.Errorf("invalid checkpoint hash %x, expected %x", cp.Hash, sr)
	}
	return pb, nil
}

// InclusionProof constructs an inclusion proof for the leaf at index in a tree of
// the given size.
// This function uses the passed-in function to retrieve tiles containing any log tree
// nodes necessary to build the proof.
func (pb *ProofBuilder) InclusionProof(ctx context.Context, index uint64) ([][]byte, error) {
	nodes, err := proof.Inclusion(index, pb.cp.Size)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate inclusion proof node list: %w", err)
	}
	return pb.fetchNodes(ctx, nodes)
}

// ConsistencyProof constructs a consistency proof between the two passed in tree sizes.
// This function uses the passed-in function to retrieve tiles containing any log tree
// nodes necessary to build the proof.
func (pb *ProofBuilder) ConsistencyProof(ctx context.Context, smaller, larger uint64) ([][]byte, error) {
	nodes, err := proof.Consistency(smaller, larger)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate consistency proof node list: %w", err)
	}
	return pb.fetchNodes(ctx, nodes)
}

// fetchNodes retrieves the specified proof nodes via pb's nodeCache.
func (pb *ProofBuilder) fetchNodes(ctx context.Context, nodes proof.Nodes) ([][]byte, error) {
	hashes := make([][]byte, 0, len(nodes.IDs))
	// TODO(al) parallelise this.
	for _, id := range nodes.IDs {
		h, err := pb.nodeCache.GetNode(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("failed to get node (%v): %w", id, err)
		}
		hashes = append(hashes, h)
	}
	var err error
	if hashes, err = nodes.Rehash(hashes, hasher.HashChildren); err != nil {
		return nil, fmt.Errorf("failed to rehash proof: %w", err)
	}
	return hashes, nil
}

// FetchRangeNodes returns the set of nodes representing the compact range covering
// a log of size s.
func FetchRangeNodes(ctx context.Context, s uint64, f TileFetcherFunc) ([][]byte, error) {
	nc := newNodeCache(f, s)
	nIDs := make([]compact.NodeID, 0, compact.RangeSize(0, s))
	nIDs = compact.RangeNodes(0, s, nIDs)
	hashes := make([][]byte, 0, len(nIDs))
	for _, n := range nIDs {
		h, err := nc.GetNode(ctx, n)
		if err != nil {
			return nil, err
		}
		hashes = append(hashes, h)
	}
	return hashes, nil
}

// FetchLeafHashes fetches N consecutive leaf hashes starting with the leaf at index first.
func FetchLeafHashes(ctx context.Context, f TileFetcherFunc, first, N, logSize uint64) ([][]byte, error) {
	nc := newNodeCache(f, logSize)
	hashes := make([][]byte, 0, N)
	for i, end := first, first+N; i < end; i++ {
		nID := compact.NodeID{Level: 0, Index: i}
		h, err := nc.GetNode(ctx, nID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch node %v: %v", nID, err)
		}
		hashes = append(hashes, h)
	}
	return hashes, nil
}

// nodeCache hides the tiles abstraction away, and improves
// performance by caching tiles it's seen.
// Not threadsafe, and intended to be only used throughout the course
// of a single request.
type nodeCache struct {
	logSize   uint64
	ephemeral map[compact.NodeID][]byte
	tiles     map[tileKey]api.HashTile
	getTile   TileFetcherFunc
}

// tileKey is used as a key in nodeCache's tile map.
type tileKey struct {
	tileLevel uint64
	tileIndex uint64
}

// newNodeCache creates a new nodeCache instance for a given log size.
func newNodeCache(f TileFetcherFunc, logSize uint64) nodeCache {
	return nodeCache{
		logSize:   logSize,
		ephemeral: make(map[compact.NodeID][]byte),
		tiles:     make(map[tileKey]api.HashTile),
		getTile:   f,
	}
}

// SetEphemeralNode stored a derived "ephemeral" tree node.
func (n *nodeCache) SetEphemeralNode(id compact.NodeID, h []byte) {
	n.ephemeral[id] = h
}

// GetNode returns the internal log tree node hash for the specified node ID.
// A previously set ephemeral node will be returned if id matches, otherwise
// the tile containing the requested node will be fetched and cached, and the
// node hash returned.
func (n *nodeCache) GetNode(ctx context.Context, id compact.NodeID) ([]byte, error) {
	// First check for ephemeral nodes:
	if e := n.ephemeral[id]; len(e) != 0 {
		return e, nil
	}
	// Otherwise look in fetched tiles:
	tileLevel, tileIndex, nodeLevel, nodeIndex := layout.NodeCoordsToTileAddress(uint64(id.Level), uint64(id.Index))
	tKey := tileKey{tileLevel, tileIndex}
	t, ok := n.tiles[tKey]
	if !ok {
		tileRaw, err := n.getTile(ctx, tileLevel, tileIndex, layout.PartialTileSize(tileLevel, tileIndex, n.logSize))
		if err != nil {
			return nil, fmt.Errorf("failed to fetch tile: %w", err)
		}
		var tile api.HashTile
		if err := tile.UnmarshalText(tileRaw); err != nil {
			return nil, fmt.Errorf("failed to parse tile: %w", err)
		}
		t = tile
		n.tiles[tKey] = tile
	}
	// We've got the tile, now we need to look up (or calculate) the node inside of it
	numLeaves := 1 << nodeLevel
	firstLeaf := int(nodeIndex) * numLeaves
	lastLeaf := firstLeaf + numLeaves
	if lastLeaf > len(t.Nodes) {
		return nil, fmt.Errorf("require leaf nodes [%d, %d) but only got %d leaves", firstLeaf, lastLeaf, len(t.Nodes))
	}
	rf := compact.RangeFactory{Hash: hasher.HashChildren}
	r := rf.NewEmptyRange(0)
	for _, l := range t.Nodes[firstLeaf:lastLeaf] {
		if err := r.Append(l, nil); err != nil {
			return nil, fmt.Errorf("Append: %v", err)
		}
	}
	return r.GetRootHash(nil)
}

// GetEntryBundle fetches the entry bundle at the given _tile index_.
func GetEntryBundle(ctx context.Context, f EntryBundleFetcherFunc, i, logSize uint64) (api.EntryBundle, error) {
	bundle := api.EntryBundle{}
	sRaw, err := f(ctx, i, layout.PartialTileSize(0, i, logSize))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return bundle, fmt.Errorf("leaf bundle at index %d not found: %v", i, err)
		}
		return bundle, fmt.Errorf("failed to fetch leaf bundle at index %d: %v", i, err)
	}
	if err := bundle.UnmarshalText(sRaw); err != nil {
		return bundle, fmt.Errorf("failed to parse EntryBundle at index %d: %v", i, err)
	}
	return bundle, nil
}

// LogStateTracker represents a client-side view of a target log's state.
// This tracker handles verification that updates to the tracked log state are
// consistent with previously seen states.
type LogStateTracker struct {
	CPFetcher   CheckpointFetcherFunc
	TileFetcher TileFetcherFunc
	// Origin is the expected first line of checkpoints from the log.
	Origin              string
	ConsensusCheckpoint ConsensusCheckpointFunc

	// LatestConsistentRaw holds the raw bytes of the latest proven-consistent
	// LogState seen by this tracker.
	LatestConsistentRaw []byte
	// LatestConsistent is the deserialised form of LatestConsistentRaw
	LatestConsistent log.Checkpoint
	// The note with signatures and other metadata about the checkpoint
	CheckpointNote *note.Note
	// ProofBuilder for building proofs at LatestConsistent checkpoint.
	ProofBuilder *ProofBuilder

	CpSigVerifier note.Verifier
}

// NewLogStateTracker creates a newly initialised tracker.
// If a serialised LogState representation is provided then this is used as the
// initial tracked state, otherwise a log state is fetched from the target log.
func NewLogStateTracker(ctx context.Context, cpF CheckpointFetcherFunc, tF TileFetcherFunc, checkpointRaw []byte, nV note.Verifier, origin string, cc ConsensusCheckpointFunc) (LogStateTracker, error) {
	ret := LogStateTracker{
		ConsensusCheckpoint: cc,
		CPFetcher:           cpF,
		TileFetcher:         tF,
		LatestConsistent:    log.Checkpoint{},
		CheckpointNote:      nil,
		CpSigVerifier:       nV,
		Origin:              origin,
	}
	if len(checkpointRaw) > 0 {
		ret.LatestConsistentRaw = checkpointRaw
		cp, _, _, err := log.ParseCheckpoint(checkpointRaw, origin, nV)
		if err != nil {
			return ret, err
		}
		ret.LatestConsistent = *cp
		ret.ProofBuilder, err = NewProofBuilder(ctx, ret.LatestConsistent, ret.TileFetcher)
		if err != nil {
			return ret, fmt.Errorf("NewProofBuilder: %v", err)
		}
		return ret, nil
	}
	_, _, _, err := ret.Update(ctx)
	return ret, err
}

// ErrInconsistency should be returned when there has been an error proving consistency
// between log states.
// The raw log state representations are included as-returned by the target log, this
// ensures that evidence of inconsistent log updates are available to the caller of
// the method(s) returning this error.
type ErrInconsistency struct {
	SmallerRaw []byte
	LargerRaw  []byte
	Proof      [][]byte

	Wrapped error
}

func (e ErrInconsistency) Unwrap() error {
	return e.Wrapped
}

func (e ErrInconsistency) Error() string {
	return fmt.Sprintf("log consistency check failed: %s", e.Wrapped)
}

// Update attempts to update the local view of the target log's state.
// If a more recent logstate is found, this method will attempt to prove
// that it is consistent with the local state before updating the tracker's
// view.
// Returns the old checkpoint, consistency proof, and newer checkpoint used to update.
// If the LatestConsistent checkpoint is 0 sized, no consistency proof will be returned
// since it would be meaningless to do so.
func (lst *LogStateTracker) Update(ctx context.Context) ([]byte, [][]byte, []byte, error) {
	c, cRaw, cn, err := lst.ConsensusCheckpoint(ctx, lst.CpSigVerifier, lst.Origin)
	if err != nil {
		return nil, nil, nil, err
	}
	builder, err := NewProofBuilder(ctx, *c, lst.TileFetcher)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create proof builder: %w", err)
	}
	var p [][]byte
	if lst.LatestConsistent.Size > 0 {
		if c.Size <= lst.LatestConsistent.Size {
			return lst.LatestConsistentRaw, p, lst.LatestConsistentRaw, nil
		}
		p, err = builder.ConsistencyProof(ctx, lst.LatestConsistent.Size, c.Size)
		if err != nil {
			return nil, nil, nil, err
		}
		if err := proof.VerifyConsistency(hasher, lst.LatestConsistent.Size, c.Size, p, lst.LatestConsistent.Hash, c.Hash); err != nil {
			return nil, nil, nil, ErrInconsistency{
				SmallerRaw: lst.LatestConsistentRaw,
				LargerRaw:  cRaw,
				Proof:      p,
				Wrapped:    err,
			}
		}
		// Update is consistent,

	}
	oldRaw := lst.LatestConsistentRaw
	lst.LatestConsistentRaw, lst.LatestConsistent, lst.CheckpointNote = cRaw, *c, cn
	lst.ProofBuilder = builder
	return oldRaw, p, lst.LatestConsistentRaw, nil
}

// CheckConsistency is a wapper function which simplifies verifying consistency between two or more checkpoints.
func CheckConsistency(ctx context.Context, f TileFetcherFunc, cp []log.Checkpoint) error {
	if l := len(cp); l < 2 {
		return fmt.Errorf("passed %d checkpoints, need at least 2", l)
	}
	sort.Slice(cp, func(i, j int) bool {
		return cp[i].Size < cp[j].Size
	})
	pb, err := NewProofBuilder(ctx, cp[len(cp)-1], f)
	if err != nil {
		return fmt.Errorf("failed to create proofbuilder: %v", err)
	}

	// Go through list of checkpoints pairwise, checking consistency.
	a, b := cp[0], cp[1]
	for i := 0; i < len(cp)-1; i, a, b = i+1, cp[i], cp[i+1] {
		if a.Size == b.Size {
			if bytes.Equal(a.Hash, b.Hash) {
				continue
			}
			return fmt.Errorf("two checkpoints with same size (%d) but different hashes (%x vs %x)", a.Size, a.Hash, b.Hash)
		}
		if a.Size > 0 {
			cp, err := pb.ConsistencyProof(ctx, a.Size, b.Size)
			if err != nil {
				return fmt.Errorf("failed to fetch consistency between sizes %d, %d: %v", a.Size, b.Size, err)
			}
			if err := proof.VerifyConsistency(hasher, a.Size, b.Size, cp, a.Hash, b.Hash); err != nil {
				return fmt.Errorf("invalid consistency proof between sizes %d, %d: %v", a.Size, b.Size, err)
			}
		}
	}
	return nil
}
