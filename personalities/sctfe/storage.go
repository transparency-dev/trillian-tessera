// Copyright 2016 Google LLC. All Rights Reserved.
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

package sctfe

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/google/certificate-transparency-go/x509"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/ctonly"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/modules/dedup"
	"k8s.io/klog/v2"
)

const (
	// Each key is 64 bytes long, so this will take up to 64MB.
	// A CT log references ~15k unique issuer certifiates in 2024, so this gives plenty of space
	// if we ever run into this limit, we should re-think how it works.
	maxCachedIssuerKeys = 1 << 20
)

// Storage provides all the storage primitives necessary to write to a ct-static-api log.
type Storage interface {
	// Add assigns an index to the provided Entry, stages the entry for integration, and returns a future for the assigned index.
	Add(context.Context, *ctonly.Entry) tessera.IndexFuture
	// AddIssuerChain stores every the chain certificate in a content-addressable store under their sha256 hash.
	AddIssuerChain(context.Context, []*x509.Certificate) error
	// AddCertIndex stores the index of certificate in a log under its hash.
	AddCertIndex(context.Context, *x509.Certificate, uint64) error
	// GetCertIndex gets the index of certificate in a log from its hash.
	GetCertIndex(context.Context, *x509.Certificate) (uint64, bool, error)
}

type KV struct {
	K []byte
	V []byte
}

// IssuerStorage issuer certificates under their hex encoded sha256.
type IssuerStorage interface {
	AddIssuersIfNotExist(ctx context.Context, kv []KV) error
}

// CTStorage implements Storage.
type CTStorage struct {
	storeData    func(context.Context, *ctonly.Entry) tessera.IndexFuture
	storeIssuers func(context.Context, []KV) error
	dedupStorage dedup.BEDedupStorage
}

// NewCTStorage instantiates a CTStorage object.
func NewCTSTorage(logStorage tessera.Storage, issuerStorage IssuerStorage, dedupStorage dedup.BEDedupStorage) (*CTStorage, error) {
	ctStorage := &CTStorage{
		storeData:    tessera.NewCertificateTransparencySequencedWriter(logStorage),
		storeIssuers: cachedStoreIssuers(issuerStorage),
		dedupStorage: dedupStorage,
	}
	return ctStorage, nil
}

// Add stores CT entries.
func (cts *CTStorage) Add(ctx context.Context, entry *ctonly.Entry) tessera.IndexFuture {
	// TODO(phboneff): add deduplication and chain storage
	return cts.storeData(ctx, entry)
}

// AddIssuerChain stores every chain certificate under its sha256.
//
// If an object is already stored under this hash, continues.
func (cts *CTStorage) AddIssuerChain(ctx context.Context, chain []*x509.Certificate) error {
	kvs := []KV{}
	for _, c := range chain {
		id := sha256.Sum256(c.Raw)
		key := []byte(hex.EncodeToString(id[:]))
		kvs = append(kvs, KV{K: key, V: c.Raw})
	}
	if err := cts.storeIssuers(ctx, kvs); err != nil {
		return fmt.Errorf("error storing intermediates: %v", err)
	}
	return nil
}

// cachedStoreIssuers returns a caching wrapper for an IssuerStorage
//
// This is intended to make querying faster. It does not keep a copy of the certs, only sha256.
// Only up to maxCachedIssuerKeys keys will be stored locally.
func cachedStoreIssuers(s IssuerStorage) func(context.Context, []KV) error {
	var mu sync.RWMutex
	m := make(map[string]struct{})
	return func(ctx context.Context, kv []KV) error {
		req := []KV{}
		for _, kv := range kv {
			mu.RLock()
			_, ok := m[string(kv.K)]
			mu.RUnlock()
			if ok {
				klog.V(2).Infof("cachedStoreIssuers wrapper: found %q in local key cache", kv.K)
				continue
			}
			req = append(req, kv)
		}
		if err := s.AddIssuersIfNotExist(ctx, req); err != nil {
			return fmt.Errorf("AddIssuersIfNotExist()s: error storing issuer data in the underlying IssuerStorage: %v", err)
		}
		for _, kv := range req {
			if len(m) >= maxCachedIssuerKeys {
				klog.V(2).Infof("cachedStoreIssuers wrapper: local issuer cache full, will stop caching issuers.")
				return nil
			}
			mu.Lock()
			m[string(kv.K)] = struct{}{}
			mu.Unlock()
		}
		return nil
	}
}

// AddCertIndex stores <cert_hash, index> in the deduplication storage.
func (cts CTStorage) AddCertIndex(ctx context.Context, c *x509.Certificate, idx uint64) error {
	key := sha256.Sum256(c.Raw)
	if err := cts.dedupStorage.Add(ctx, []dedup.LeafIdx{{LeafID: key[:], Idx: idx}}); err != nil {
		return fmt.Errorf("error storing index %d of %q: %v", idx, hex.EncodeToString(key[:]), err)
	}
	return nil
}

// GetCertIndex fetches the index of a given certificate from the deduplication storage.
func (cts CTStorage) GetCertIndex(ctx context.Context, c *x509.Certificate) (uint64, bool, error) {
	key := sha256.Sum256(c.Raw)
	idx, ok, err := cts.dedupStorage.Get(ctx, key[:])
	if err != nil {
		return 0, false, fmt.Errorf("error fetching index of %q: %v", hex.EncodeToString(key[:]), err)
	}
	return idx, ok, nil
}
