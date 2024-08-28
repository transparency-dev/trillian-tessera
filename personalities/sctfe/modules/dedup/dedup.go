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

// Package bbolt implements SCTFE storage systems for deduplication.
//
// The interfaces are defined in sctfe/storage.go
package dedup

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/client"
	"golang.org/x/crypto/cryptobyte"
	"golang.org/x/mod/sumdb/note"
	"k8s.io/klog/v2"
)

// KV holds a LeafID and an Idx for deduplication
type KV struct {
	K []byte
	V uint64
}

type DedupStorage interface {
	Add(kvs []KV) error
	Get(key []byte) (uint64, bool, error)
}

type LocalDedupStorage interface {
	Add(kvs []KV) error
	Get(key []byte) (uint64, bool, error)
	LogSize() (uint64, error)
}

type LocalBEDedup struct {
	DedupStorage
	LogSize func() (uint64, error) // returns the largest idx Add has successfully been called with
	fetcher client.Fetcher
}

func NewLocalBestEffortDedup(ctx context.Context, lds LocalDedupStorage, t time.Duration, f client.Fetcher, v note.Verifier, origin string) *LocalBEDedup {
	ret := &LocalBEDedup{DedupStorage: lds, LogSize: lds.LogSize, fetcher: f}
	go func() {
		tck := time.NewTicker(t)
		defer tck.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-tck.C:
				if err := ret.sync(ctx, origin, v); err != nil {
					klog.Warningf("error updating deduplication data: %v", err)
				}
			}
		}
	}()
	return ret
}

func (d *LocalBEDedup) sync(ctx context.Context, origin string, v note.Verifier) error {
	ckpt, _, _, err := client.FetchCheckpoint(ctx, d.fetcher, v, origin)
	if err != nil {
		return fmt.Errorf("FetchCheckpoint: %v", err)
	}
	oldSize, err := d.LogSize()
	if err != nil {
		return fmt.Errorf("OldSize(): %v", err)
	}

	// TODO(phboneff): add parallelism
	// Greatly inspired by https://github.com/FiloSottile/sunlight/blob/main/tile.go and
	// https://github.com/transparency-dev/trillian-tessera/blob/main/client/client.go
	if ckpt.Size > oldSize {
		for i := oldSize / 256; i <= ckpt.Size/256; i++ {
			kvs := []KV{}
			p := layout.EntriesPath(i, ckpt.Size)
			eRaw, err := d.fetcher(ctx, p)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					return fmt.Errorf("leaf bundle at index %d not found: %v", i, err)
				}
				return fmt.Errorf("failed to fetch leaf bundle at index %d: %v", i, err)
			}
			s := cryptobyte.String(eRaw)

			for len(s) > 0 {
				var timestamp uint64
				var entryType uint16
				var extensions, fingerprints cryptobyte.String
				if !s.ReadUint64(&timestamp) || !s.ReadUint16(&entryType) || timestamp > math.MaxInt64 {
					return fmt.Errorf("invalid data tile")
				}
				crt := []byte{}
				switch entryType {
				case 0: // x509_entry
					if !s.ReadUint24LengthPrefixed((*cryptobyte.String)(&crt)) ||
						// TODO(phboneff): remove below?
						!s.ReadUint16LengthPrefixed(&extensions) ||
						!s.ReadUint16LengthPrefixed(&fingerprints) {
						return fmt.Errorf("invalid data tile x509_entry")
					}
				case 1: // precert_entry
					IssuerKeyHash := [32]byte{}
					var defangedCrt, extensions cryptobyte.String
					if !s.CopyBytes(IssuerKeyHash[:]) ||
						!s.ReadUint24LengthPrefixed(&defangedCrt) ||
						!s.ReadUint16LengthPrefixed(&extensions) ||
						!s.ReadUint24LengthPrefixed((*cryptobyte.String)(&crt)) ||
						// TODO(phboneff): remove below?
						!s.ReadUint16LengthPrefixed(&fingerprints) {
						return fmt.Errorf("invalid data tile precert_entry")
					}
				default:
					return fmt.Errorf("invalid data tile: unknown type %d", entryType)
				}
				k := sha256.Sum256(crt)
				kvs = append(kvs, KV{K: k[:], V: i*256 + uint64(len(kvs))})
			}

			if err := d.Add(kvs); err != nil {
				return fmt.Errorf("error storing deduplication data for tile %d: %v", i, err)
			}
		}
	}
	return nil
}
