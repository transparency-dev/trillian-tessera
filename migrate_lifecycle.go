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

package tessera

import (
	"context"
	"fmt"

	"github.com/transparency-dev/trillian-tessera/api/layout"
)

// MigrationTarget describes the contract of the Migration lifecycle.
//
// This lifecycle mode is used to migrate C2SP tlog-tiles and static-ct
// compliant logs into Tessera.
type MigrationTarget interface {
	// SetEntryBundle stores the provided serialised entry bundle at the location implied by the provided
	// entry bundle index and partial size.
	//
	// Bundles may be set in any order (not just consecutively), and the implementation should integrate
	// them into the local tree in the most efficient way possible.
	//
	// Writes should be idempotent; repeated calls to set the same bundle with the same data should not
	// return an error.
	SetEntryBundle(ctx context.Context, idx uint64, partial uint8, bundle []byte) error
	// AwaitIntegration should block until the local integrated tree has grown to the provided size,
	// and should return the locally calculated root hash derived from the integration of the contents of
	// entry bundles set using SetEntryBundle above.
	AwaitIntegration(ctx context.Context, size uint64) ([]byte, error)
	// IntegratedSize returns the current size of the locally integrated log.
	IntegratedSize(ctx context.Context) (uint64, error)
}

// UnbundlerFunc is a function which knows how to turn a serialised entry bundle into a slice of
// []byte representing each of the entries within the bundle.
type UnbundlerFunc func(entryBundle []byte) ([][]byte, error)

// NewMigrationTarget returns a MigrationTarget, which allows a personality to "import" a C2SP
// tlog-tiles or static-ct compliant log into a Tessera instance.
//
// TODO(al): bundleHasher should be implicit from WithCTLayout being present or not.
// TODO(al): AppendOptions should be somehow replaced - perhaps MigrationOptions, or some other way of limiting options to those which make sense for this lifecycle mode.
func NewMigrationTarget(ctx context.Context, d Driver, bundleHasher UnbundlerFunc, opts *MigrationOptions) (MigrationTarget, error) {
	type migrateLifecycle interface {
		MigrationTarget(context.Context, UnbundlerFunc, *MigrationOptions) (MigrationTarget, LogReader, error)
	}
	lc, ok := d.(migrateLifecycle)
	if !ok {
		return nil, fmt.Errorf("driver %T does not implement MigrationTarget lifecycle", d)
	}
	m, r, err := lc.MigrationTarget(ctx, bundleHasher, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to init MigrationTarget lifecycle: %v", err)
	}
	for _, f := range opts.followers {
		go f(ctx, r)
	}
	return m, nil
}

func NewMigrationOptions() *MigrationOptions {
	return &MigrationOptions{
		entriesPath: layout.EntriesPath,
	}
}

// MigrationOptions holds migration lifecycle settings for all storage implementations.
type MigrationOptions struct {
	// entriesPath knows how to format entry bundle paths.
	entriesPath func(n uint64, p uint8) string
	followers   []func(context.Context, LogReader)
}

func (o MigrationOptions) EntriesPath() func(uint64, uint8) string {
	return o.entriesPath
}

func (o MigrationOptions) Followers() []func(context.Context, LogReader) {
	return o.followers
}

func (o *MigrationOptions) WithAntispam(as Antispam) *MigrationOptions {
	if as != nil {
		o.followers = append(o.followers, func(ctx context.Context, lr LogReader) {
			as.Populate(ctx, lr, defaultIDHasher)
		})
	}
	return o
}
