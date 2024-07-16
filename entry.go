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
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/transparency-dev/merkle/rfc6962"
)

// Entry represents an entry in a log.
type Entry struct {
	data     []byte
	identity []byte
	leafHash []byte
}

func (e Entry) Data() []byte     { return e.data }
func (e Entry) Identity() []byte { return e.identity }
func (e Entry) LeafHash() []byte { return e.leafHash }

// NewEntry creates a new Entry object with leaf data.
func NewEntry(data []byte, opts ...EntryOpt) Entry {
	e := Entry{
		data: data,
	}
	for _, opt := range opts {
		opt(&e)
	}
	if e.identity == nil {
		h := sha256.Sum256(e.data)
		e.identity = h[:]
	}
	if e.leafHash == nil {
		e.leafHash = rfc6962.DefaultHasher.HashLeaf(e.data)

	}
	return e
}

func (e *Entry) MarshalBinary() ([]byte, error) {
	buf := make([]byte, 0, len(e.data)+len(e.identity)+len(e.leafHash)+12)
	buf = binary.AppendVarint(buf, int64(len(e.data)))
	buf = append(buf, e.data...)
	buf = binary.AppendVarint(buf, int64(len(e.identity)))
	buf = append(buf, e.identity...)
	buf = binary.AppendVarint(buf, int64(len(e.leafHash)))
	buf = append(buf, e.leafHash...)
	return buf, nil
}

func (e *Entry) UnmarshalBinary(buf []byte) error {
	l, n := binary.Varint(buf)
	buf = buf[n:]
	e.data, buf = buf[:l], buf[l:]
	l, n = binary.Varint(buf)
	buf = buf[n:]
	e.identity, buf = buf[:l], buf[l:]
	l, n = binary.Varint(buf)
	buf = buf[n:]
	e.leafHash, buf = buf[:l], buf[l:]
	if l := len(buf); l != 0 {
		return fmt.Errorf("%d trailing bytes", l)
	}
	return nil
}

// EntryOpt is the signature of options for creating new Entry instances.
type EntryOpt func(e *Entry)

// NewEntryWithIdentity creates a new Entry with leaf data and a semantic identity.
func WithIdentity(identity []byte) EntryOpt {
	return func(e *Entry) {
		e.identity = identity
	}
}
