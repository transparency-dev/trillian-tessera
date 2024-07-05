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

package gcp

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
	"testing"

	gcs "cloud.google.com/go/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/transparency-dev/trillian-tessera/api/layout"
)

func TestTileSuffix(t *testing.T) {
	for _, test := range []struct {
		name string
		size uint64
		want string
	}{
		{
			name: "no suffix",
			size: 256 * 23,
			want: "",
		}, {
			name: "no suffix on zero",
			size: 0,
			want: "",
		}, {
			name: "has suffix",
			size: 256*23 + 3,
			want: ".p/3",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got, want := tileSuffix(test.size), test.want; got != want {
				t.Fatalf("got %s want %s", got, want)
			}
		})
	}
}

func makeTile(t *testing.T, size uint64) [][]byte {
	t.Helper()
	r := make([][]byte, size)
	for i := uint64(0); i < size; i++ {
		h := sha256.Sum256([]byte(fmt.Sprintf("%d", i)))
		r[i] = h[:]
	}
	return r
}

func TestTileRoundtrip(t *testing.T) {
	ctx := context.Background()
	m := newMemObjStore()
	s := &Storage{
		objStore: m,
	}

	for _, test := range []struct {
		name     string
		level    uint64
		index    uint64
		logSize  uint64
		tileSize uint64
		wantErr  bool
	}{
		{
			name:     "ok",
			level:    0,
			index:    3 * 256,
			logSize:  3*256 + 20,
			tileSize: 20,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			wantTile := makeTile(t, test.tileSize)
			err := s.setTile(ctx, test.level, test.index, wantTile)
			if gotErr := err != nil; gotErr != test.wantErr {
				t.Fatalf("setTile: %v want err %t", err, test.wantErr)
			}

			expPath := layout.TilePath(test.level, test.index) + tileSuffix(test.tileSize)
			_, ok := m.mem[expPath]
			if !ok {
				t.Fatalf("want tile at %v but found none", expPath)
			}

			got, err := s.getTile(ctx, test.level, test.index, test.logSize)
			if err != nil {
				t.Fatalf("getTile: %v", err)
			}
			if !cmp.Equal(got, wantTile) {
				t.Fatal("roundtrip returned different data")
			}
		})
	}
}

type memObjStore struct {
	sync.RWMutex
	mem map[string][]byte
}

func newMemObjStore() *memObjStore {
	return &memObjStore{
		mem: make(map[string][]byte),
	}
}

func (m *memObjStore) getObject(_ context.Context, obj string) ([]byte, int64, error) {
	m.RLock()
	defer m.RUnlock()

	d, ok := m.mem[obj]
	if !ok {
		return nil, -1, fmt.Errorf("obj %q not found: %w", obj, gcs.ErrObjectNotExist)
	}
	return d, 1, nil
}

func (m *memObjStore) setObject(_ context.Context, obj string, data []byte, cond *gcs.Conditions) error {
	m.Lock()
	defer m.Unlock()

	d, ok := m.mem[obj]
	if cond != nil {
		if ok && cond.DoesNotExist {
			if !bytes.Equal(d, data) {
				return errors.New("precondition failed and data not identical")
			}
			return nil
		}
	}
	m.mem[obj] = data
	return nil
}
