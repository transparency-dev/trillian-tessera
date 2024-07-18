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

// package api_test contains tests for the api package.
package api_test

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api"
)

func TestHashTile_MarshalTileRoundtrip(t *testing.T) {
	for _, test := range []struct {
		size int
	}{
		{
			size: 1,
		}, {
			size: 255,
		}, {
			size: 11,
		}, {
			size: 42,
		},
	} {
		t.Run(fmt.Sprintf("tile size %d", test.size), func(t *testing.T) {
			tile := api.HashTile{Nodes: make([][]byte, 0, test.size)}
			for i := 0; i < test.size; i++ {
				// Fill in the leaf index
				tile.Nodes = append(tile.Nodes, make([]byte, 32))
				if _, err := rand.Read(tile.Nodes[i]); err != nil {
					t.Error(err)
				}
			}

			raw, err := tile.MarshalText()
			if err != nil {
				t.Fatalf("MarshalText() = %v", err)
			}

			tile2 := api.HashTile{}
			if err := tile2.UnmarshalText(raw); err != nil {
				t.Fatalf("UnmarshalText() = %v", err)
			}

			if diff := cmp.Diff(tile, tile2); len(diff) != 0 {
				t.Fatalf("Got tile with diff: %s", diff)
			}
		})
	}
}

func TestLeafBundle_MarshalTileRoundtrip(t *testing.T) {
	for _, test := range []struct {
		size int
	}{
		{
			size: 1,
		}, {
			size: 255,
		}, {
			size: 11,
		}, {
			size: 42,
		},
	} {
		t.Run(fmt.Sprintf("tile size %d", test.size), func(t *testing.T) {
			bundleRaw := &bytes.Buffer{}
			want := make([][]byte, test.size)
			for i := 0; i < test.size; i++ {
				// Fill in the leaf index
				want[i] = make([]byte, i*100)
				if _, err := rand.Read(want[i]); err != nil {
					t.Error(err)
				}
				if err := tessera.NewEntry(want[i]).WriteBundleEntry(bundleRaw); err != nil {
					t.Fatalf("Write: %v", err)
				}
			}

			tile2 := api.EntryBundle{}
			if err := tile2.UnmarshalText(bundleRaw.Bytes()); err != nil {
				t.Fatalf("UnmarshalText() = %v", err)
			}

			for i := 0; i < test.size; i++ {
				if got, want := tile2.Entries[i], want[i]; !bytes.Equal(got, want) {
					t.Errorf("%d: want %x, got %x", i, got, want)
				}
			}
		})
	}
}

func TestLeafBundle_UnmarshalText(t *testing.T) {
	for _, test := range []struct {
		desc    string
		input   []byte
		wantErr bool
	}{
		{
			desc:    "no data",
			input:   []byte{},
			wantErr: false,
		},
		{
			desc:    "insufficient data",
			input:   []byte{0x0, 0x02, 'a'},
			wantErr: true,
		},
		{
			desc:    "empty nodes",
			input:   []byte{0x0, 0x0, 0x0, 0x1, 'a', 0x0, 0x0},
			wantErr: false,
		},
		{
			desc:    "insufficient integer bytes",
			input:   []byte{0x1},
			wantErr: true,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			tile := api.EntryBundle{}
			err := tile.UnmarshalText(test.input)
			if gotErr := err != nil; gotErr != test.wantErr {
				t.Errorf("wantErr: %t, got %v", test.wantErr, err)
			}
		})
	}
}
