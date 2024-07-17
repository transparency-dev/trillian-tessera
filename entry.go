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

// Package tessera provides an implementation of a tile-based logging framework.
package tessera

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/transparency-dev/merkle/rfc6962"
)

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

	// indexFunc will be called when an index is assigned to the entry.
	indexFunc func(index uint64)

	// writeBundleEntry knows how to writeBundleEntry this bundle in the correct format for appending to a serialised bundle.
	writeBundleEntry func(io.Writer) error
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

// AssignIndex will be called by the log when an index is assigned to the entry.
func (e *Entry) AssignIndex(index uint64) {
	e.internal.Index = &index
	if e.indexFunc != nil {
		e.indexFunc(index)
	}
}

// NewEntry creates a new Entry object with leaf data.
func NewEntry(data []byte, opts ...EntryOpt) *Entry {
	e := &Entry{}
	e.internal.Data = data
	for _, opt := range opts {
		opt(e)
	}
	if e.internal.Identity == nil {
		h := sha256.Sum256(e.internal.Data)
		e.internal.Identity = h[:]
	}
	if e.internal.LeafHash == nil {
		e.internal.LeafHash = rfc6962.DefaultHasher.HashLeaf(e.internal.Data)
	}
	if e.writeBundleEntry == nil {
		// By default we will marshal ourselves into a bundle using the mechanism described
		// by https://c2sp.org/tlog-tiles:
		e.writeBundleEntry = func(w io.Writer) error {
			if err := binary.Write(w, binary.BigEndian, uint16(len(e.internal.Data))); err != nil {
				return err
			}
			if n, err := w.Write(e.internal.Data); err != nil {
				return err
			} else if l := len(e.internal.Data); n != l {
				return fmt.Errorf("short write %d bytes, want %d", n, l)
			}
			return nil
		}
	}
	return e
}

// WriteBundleEntry serialises the data this Entry contains into a suitable format for
// adding to an EntryBundle.
func (e Entry) WriteBundleEntry(w io.Writer) error {
	return e.writeBundleEntry(w)
}

func (e Entry) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}
	enc := gob.NewEncoder(b)
	if err := enc.Encode(e.internal); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (e *Entry) UnmarshalBinary(buf []byte) error {
	dec := gob.NewDecoder(bytes.NewReader(buf))
	return dec.Decode(&e.internal)
}

// EntryOpt is the signature of options for creating new Entry instances.
type EntryOpt func(e *Entry)

// WithIdentity is an option to create Entries with an explicit identity.
//
// The provided identity may be used to deduplicate entries as they're being
// added to the log, if such behaviour is supported and enabled on the
// storage implementation.
func WithIdentity(identity []byte) EntryOpt {
	return func(e *Entry) {
		e.internal.Identity = identity
	}
}
