// Package shizzle is all the stuff that gets linked to everywhere and doesn't link out.
// It's a way of avoiding package cycles, essentially.
package shizzle

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/transparency-dev/merkle/rfc6962"
	"github.com/transparency-dev/trillian-tessera/api"
)

// Add adds a new entry to be sequenced.
// This method quickly returns an IndexFuture, which will return the index assigned
// to the new leaf. Until this index is obtained from the future, the leaf is not durably
// added to the log, and terminating the process may lead to this leaf being lost.
// Once the future resolves and returns an index, the leaf is durably sequenced and will
// be preserved even in the process terminates.
//
// Once a leaf is sequenced, it will be integrated into the tree soon (generally single digit
// seconds). Until it is integrated and published, clients of the log will not be able to
// verifiably access this value. Personalities that require blocking until the leaf is integrated
// can use the PublicationAwaiter to wrap the call to this method.
type AddFn func(ctx context.Context, entry *Entry) IndexFuture

// IndexFuture is the signature of a function which can return an assigned index or error.
//
// Implementations of this func are likely to be "futures", or a promise to return this data at
// some point in the future, and as such will block when called if the data isn't yet available.
type IndexFuture func() (Index, error)

// Index represents a durably assigned index for some entry.
type Index struct {
	// Index is the location in the log to which a particular entry has been assigned.
	Index uint64
	// IsDup is true if Index represents a previously assigned index for an identical entry.
	IsDup bool
}

// Entry represents an entry in a log.
type Entry struct {
	// We keep the all data in exported fields inside an unexported interal struct.
	// This allows us to use gob to serialise the entry data (relying on the backwards-compatibility
	// it provides), while also keeping these fields private which allows us to deter bad practice
	// by forcing use of the API to set these values to safe values.
	internal struct {
		Data     []byte
		Identity []byte
		LeafHash []byte
		Index    *uint64
	}

	// marshalForBundle knows how to convert this entry's Data into a marshalled bundle entry.
	marshalForBundle func(index uint64) []byte
}

// Data returns the raw entry bytes which will form the entry in the log.
func (e Entry) Data() []byte { return e.internal.Data }

// Identity returns an identity which may be used to de-duplicate entries and they are being added to the log.
func (e Entry) Identity() []byte { return e.internal.Identity }

// LeafHash is the Merkle leaf hash which will be used for this entry in the log.
// Note that in almost all cases, this should be the RFC6962 definition of a leaf hash.
func (e Entry) LeafHash() []byte { return e.internal.LeafHash }

// Index returns the index assigned to the entry in the log, or nil if no index has been assigned.
func (e Entry) Index() *uint64 { return e.internal.Index }

// MarshalBundleData returns this entry's data in a format ready to be appended to an EntryBundle.
//
// Note that MarshalBundleData _may_ be called multiple times, potentially with different values for index
// (e.g. if there's a failure in the storage when trying to persist the assignment), so index should not
// be considered final until the storage Add method has returned successfully with the durably assigned index.
func (e *Entry) MarshalBundleData(index uint64) []byte {
	e.internal.Index = &index
	return e.marshalForBundle(index)
}

// NewEntry creates a new Entry object with leaf data.
func NewEntry(data []byte) *Entry {
	e := &Entry{}
	e.internal.Data = data
	h := identityHash(e.internal.Data)
	e.internal.Identity = h[:]
	e.internal.LeafHash = rfc6962.DefaultHasher.HashLeaf(e.internal.Data)
	// By default we will marshal ourselves into a bundle using the mechanism described
	// by https://c2sp.org/tlog-tiles:
	e.marshalForBundle = func(_ uint64) []byte {
		r := make([]byte, 0, 2+len(e.internal.Data))
		r = binary.BigEndian.AppendUint16(r, uint16(len(e.internal.Data)))
		r = append(r, e.internal.Data...)
		return r
	}
	return e
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
