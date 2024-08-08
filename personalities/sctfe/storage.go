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

	"github.com/google/certificate-transparency-go/x509"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/ctonly"
	"golang.org/x/sync/errgroup"
)

// Storage provides all the storage primitives necessary to write to a ct-static-api log.
type Storage interface {
	// Add assign an index to the provided Entry, stages the entry for integration, and return it the assigned index.
	Add(context.Context, *ctonly.Entry) (uint64, error)
	// AddIssuerChain stores all certificates in the chain in a content-addressable store under their sha256 hash.
	AddIssuerChain(context.Context, []*x509.Certificate) error
}

type IssuerStorage interface {
	Exists(ctx context.Context, key [32]byte) (bool, error)
	Add(ctx context.Context, key [32]byte, data []byte) error
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
		issuers:   issuerStorage,
	}
	return ctStorage, nil
}

// Add stores CT entries.
func (cts *CTStorage) Add(ctx context.Context, entry *ctonly.Entry) (uint64, error) {
	// TODO(phboneff): add deduplication and chain storage
	return cts.storeData(ctx, entry)
}

// AddIssuerChain stores every certificate in the chain under its sha256.
// If an object is already stored under this hash, continues.
func (cts *CTStorage) AddIssuerChain(ctx context.Context, chain []*x509.Certificate) error {
	errG := errgroup.Group{}
	for _, c := range chain {
		errG.Go(func() error {
			key := sha256.Sum256(c.Raw)
			// We first try and see if this issuer cert has already been stored since reads
			// are cheaper than writes.
			// TODO(phboneff): monitor usage, eventually write directly depending on usage patterns
			ok, err := cts.issuers.Exists(ctx, key)
			if err != nil {
				return fmt.Errorf("error checking if issuer %q exists: %s", hex.EncodeToString(key[:]), err)
			}
			if !ok {
				if err = cts.issuers.Add(ctx, key, c.Raw); err != nil {
					return fmt.Errorf("error adding certificate for issuer %q: %v", hex.EncodeToString(key[:]), err)

				}
			}
			return nil
		})
	}
	if err := errG.Wait(); err != nil {
		return err
	}
	return nil
}
