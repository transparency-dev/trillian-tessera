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
	// Add assigns an index to the provided Entry, stages the entry for integration, and return it the assigned index.
	Add(context.Context, *ctonly.Entry) (uint64, error)
	// AddIssuerChain stores every the chain certificate in a content-addressable store under their sha256 hash.
	AddIssuerChain(context.Context, []*x509.Certificate) error
}

type KV struct {
	K []byte
	V []byte
}

type IssuerStorage interface {
	Exists(ctx context.Context, key []byte) (bool, error)
	AddIssuers(ctx context.Context, kv []KV) error
}

// CTStorage implements Storage.
type CTStorage struct {
	storeData func(context.Context, *ctonly.Entry) (uint64, error)
	issuers   IssuerStorage
}

// NewCTStorage instantiates a CTStorage object.
func NewCTSTorage(logStorage tessera.Storage, issuerStorage IssuerStorage) (*CTStorage, error) {
	ctStorage := &CTStorage{
		storeData: tessera.NewCertificateTransparencySequencedWriter(logStorage),
		issuers:   NewCachedIssuerStorage(issuerStorage),
	}
	return ctStorage, nil
}

// Add stores CT entries.
func (cts *CTStorage) Add(ctx context.Context, entry *ctonly.Entry) (uint64, error) {
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
	if err := cts.issuers.AddIssuers(ctx, kvs); err != nil {
		return fmt.Errorf("error storing intermediates: %v", err)
	}
	return nil
}

// cachedIssuerStorage wraps an IssuerStorage, and keeps a copy the sha256 of certs it contains.
//
// This is intended to make querying faster. It does not keep a copy of the certs, only sha256.
// Only up to N keys will be stored locally.
// TODO(phboneff): add monitoring for the number of keys
type cachedIssuerStorage struct {
	sync.Mutex
	m map[string]struct{}
	N int // maximum number of entries allowed in m
	s IssuerStorage
}

// Exists checks if the key exists in the local cache, if not checks in the underlying storage.
// If it finds it there, caches the key locally.
func (c *cachedIssuerStorage) Exists(ctx context.Context, key []byte) (bool, error) {
	_, ok := c.m[string(key)]
	if ok {
		klog.V(2).Infof("Exists: found %q in local key cache", key)
		return true, nil
	}
	ok, err := c.s.Exists(ctx, key)
	if err != nil {
		return false, fmt.Errorf("error checking if issuer %q exists in the underlying IssuerStorage: %s", key, err)
	}
	if ok {
		c.Lock()
		c.m[string(key)] = struct{}{}
		c.Unlock()
	}
	return ok, nil
}

// AddIssuers first adds the issuers to the underlying storage, then caches their sha256 locally.
//
// Only up to c.N issuer sha256 will be cached.
func (c *cachedIssuerStorage) AddIssuers(ctx context.Context, kv []KV) error {
	req := []KV{}
	for _, kv := range kv {
		b, err := c.Exists(ctx, kv.K)
		if err != nil {
			return fmt.Errorf("error checking if issuer %q has been sotred previously: %v", string(kv.K), err)
		}
		if !b {
			req = append(req, kv)
		}
	}
	if err := c.s.AddIssuers(ctx, req); err != nil {
		return fmt.Errorf("AddIssuers: error storing issuer data for in the underlying IssuerStorage: %v", err)
	}
	for _, kv := range req {
		if len(c.m) >= c.N {
			klog.V(2).Infof("Add: local issuer cache full, will stop caching issuers.")
			return nil
		}
		c.Lock()
		c.m[string(kv.K)] = struct{}{}
		c.Unlock()
	}
	return nil
}

func NewCachedIssuerStorage(s IssuerStorage) *cachedIssuerStorage {
	return &cachedIssuerStorage{s: s, N: maxCachedIssuerKeys, m: make(map[string]struct{})}
}
