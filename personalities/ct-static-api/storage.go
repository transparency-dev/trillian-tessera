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

package ctfe

import (
	"context"

	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/ctonly"
)

// Storage provides all the storage primitives necessary to write to a ct-static-api log.
type Storage interface {
	// 	Add assign an index to the provided Entry, stages the entry for integration, and return it the assigned index.
	Add(context.Context, *ctonly.Entry) (uint64, error)
}

// CTStorage implements Storage.
type CTStorage struct {
	storeData func(context.Context, *ctonly.Entry) (uint64, error)
	// TODO(phboneff): add storeExtraData
	// TODO(phboneff): add dedupe
}

// NewCTStorage instantiates a CTStorage object.
func NewCTSTorage(logStorage tessera.Storage) (*CTStorage, error) {
	ctStorage := &CTStorage{
		storeData: tessera.NewCertificateTransparencySequencedWriter(logStorage),
	}
	return ctStorage, nil
}

// Add stores CT entries.
func (cts CTStorage) Add(ctx context.Context, entry *ctonly.Entry) (uint64, error) {
	// TODO(phboneff): add deduplication and chain storage
	return cts.storeData(ctx, entry)
}
