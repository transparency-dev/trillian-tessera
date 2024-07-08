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

// Package api contains the tiles definitions from the tlog-tiles spec.
package api

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// HashTile represents a tile within the Merkle hash tree.
// Leaf HashTiles will have a corresponding EntryBundle, where each
// entry in the EntryBundle slice hashes to the value at the same
// index in the Nodes slice.
type HashTile struct {
	// Nodes stores the leaf hash nodes in this tile.
	// Note that only non-ephemeral nodes are stored.
	Nodes [][]byte
}

// MarshalText implements encoding/TextMarshaller and writes out an HashTile
// instance as sequences of concatenated hashes as specified by the tlog-tiles spec.
func (t HashTile) MarshalText() ([]byte, error) {
	r := &bytes.Buffer{}
	for _, n := range t.Nodes {
		if _, err := r.Write(n); err != nil {
			return nil, err
		}
	}
	return r.Bytes(), nil
}

// UnmarshalText implements encoding/TextUnmarshaler and reads HashTiles
// which are encoded using the tlog-tiles spec.
func (t *HashTile) UnmarshalText(raw []byte) error {
	if len(raw)%32 != 0 {
		return fmt.Errorf("%d is not a multiple of 32", len(raw))
	}
	nodes := make([][]byte, 0, len(raw)/32)
	for index := 0; index < len(raw); index += 32 {
		data := raw[index : index+32]
		nodes = append(nodes, data)
	}
	t.Nodes = nodes
	return nil
}

// EntryBundle represents a sequence of entries in the log.
// These entries correspond to a leaf tile in the hash tree.
type EntryBundle struct {
	// Entries stores the leaf entries of the log, in order.
	// Note that only non-ephemeral nodes are stored.
	Entries [][]byte
}

// MarshalText implements encoding/TextMarshaller and writes out an EntryBundle
// instance as sequences of big-endian uint16 length-prefixed log entries,
// as specified by the tlog-tiles spec.
func (t EntryBundle) MarshalText() ([]byte, error) {
	r := &bytes.Buffer{}
	sizeBs := make([]byte, 2)
	for _, n := range t.Entries {
		binary.BigEndian.PutUint16(sizeBs, uint16(len(n)))
		if _, err := r.Write(sizeBs); err != nil {
			return nil, err
		}
		if _, err := r.Write(n); err != nil {
			return nil, err
		}
	}
	return r.Bytes(), nil
}

// UnmarshalText implements encoding/TextUnmarshaler and reads EntryBundles
// which are encoded using the tlog-tiles spec.
func (t *EntryBundle) UnmarshalText(raw []byte) error {
	nodes := make([][]byte, 0)
	for index := 0; index < len(raw); {
		dataIndex := index + 2
		if dataIndex > len(raw) {
			return fmt.Errorf("dangling bytes at byte index %d in data of %d bytes", index, len(raw))
		}
		size := int(binary.BigEndian.Uint16(raw[index:dataIndex]))
		dataEnd := dataIndex + size
		if dataEnd > len(raw) {
			return fmt.Errorf("require %d bytes from byte index %d, but size is %d", size, dataIndex, len(raw))
		}
		data := raw[dataIndex:dataEnd]
		nodes = append(nodes, data)
		index = dataIndex + size
	}
	t.Entries = nodes
	return nil
}