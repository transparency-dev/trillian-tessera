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

package tessera

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/transparency-dev/tessera/ctonly"
)

func TestCTEntriesPath(t *testing.T) {
	for _, test := range []struct {
		N        uint64
		p        uint8
		wantPath string
	}{
		{
			N:        0,
			wantPath: "tile/data/000",
		},
		{
			N:        0,
			p:        8,
			wantPath: "tile/data/000.p/8",
		}, {
			N:        255,
			wantPath: "tile/data/255",
		}, {
			N:        255,
			p:        253,
			wantPath: "tile/data/255.p/253",
		}, {
			N:        256,
			wantPath: "tile/data/256",
		}, {
			N:        123456789000,
			wantPath: "tile/data/x123/x456/x789/000",
		},
	} {
		desc := fmt.Sprintf("N %d", test.N)
		t.Run(desc, func(t *testing.T) {
			gotPath := ctEntriesPath(test.N, test.p)
			if gotPath != test.wantPath {
				t.Errorf("got file %q want %q", gotPath, test.wantPath)
			}
		})
	}
}

var (
	testCert              = []byte("I am a Certificate")
	testPrecert           = []byte("I am a Precertificate")
	testPrecertTBS        = []byte("I am a Precertificate TBS")
	testIssuerKeyHash     = sha256.Sum256([]byte("I'm an IssuerKey"))
	testFingerprintsChain = [][32]byte{
		sha256.Sum256([]byte("one")),
		sha256.Sum256([]byte("two")),
	}
)

func TestCTIdentityHasher(t *testing.T) {
	for _, test := range []struct {
		name    string
		entries []ctonly.Entry
	}{
		{
			name: "Single Certificate",
			entries: []ctonly.Entry{
				{
					Timestamp:         1234,
					IsPrecert:         false,
					Certificate:       testCert,
					FingerprintsChain: testFingerprintsChain,
				},
			},
		},
		{
			name: "Single Preertificate",
			entries: []ctonly.Entry{
				{
					Timestamp:         1234,
					IsPrecert:         true,
					Certificate:       testPrecertTBS,
					Precertificate:    testPrecert,
					IssuerKeyHash:     testIssuerKeyHash[:],
					FingerprintsChain: testFingerprintsChain,
				},
			},
		},
		{
			name: "Mixed bag",
			entries: []ctonly.Entry{
				{
					Timestamp:         1234,
					IsPrecert:         true,
					Certificate:       testPrecertTBS,
					Precertificate:    testPrecert,
					IssuerKeyHash:     testIssuerKeyHash[:],
					FingerprintsChain: testFingerprintsChain,
				}, {
					Timestamp:         1234,
					IsPrecert:         false,
					Certificate:       testCert,
					FingerprintsChain: testFingerprintsChain,
				},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			bundle := []byte{}
			wantIDs := [][]byte{}
			for _, e := range test.entries {
				bundle = append(bundle, e.LeafData(123)...)
				wantIDs = append(wantIDs, e.Identity())
			}

			gotIDs, gotErr := ctBundleIDHasher(bundle)
			if gotErr != nil {
				t.Fatalf("ctBundleIDHasher: %v", gotErr)
			}
			if lg, lw := len(gotIDs), len(wantIDs); lg != lw {
				t.Fatalf("got %d hashes, want %d", lg, lw)
			}
			for i := range gotIDs {
				if !bytes.Equal(gotIDs[i], wantIDs[i]) {
					t.Fatalf("%d: got ID hash %x, want %x", i, gotIDs[i], wantIDs[i])
				}
			}

		})
	}
}

func TestCTMerkleLeafHasher(t *testing.T) {
	for _, test := range []struct {
		name    string
		entries []ctonly.Entry
	}{
		{
			name: "Single Certificate",
			entries: []ctonly.Entry{
				{
					Timestamp:         1234,
					IsPrecert:         false,
					Certificate:       testCert,
					FingerprintsChain: testFingerprintsChain,
				},
			},
		},
		{
			name: "Single Preertificate",
			entries: []ctonly.Entry{
				{
					Timestamp:         1234,
					IsPrecert:         true,
					Certificate:       testPrecertTBS,
					Precertificate:    testPrecert,
					IssuerKeyHash:     testIssuerKeyHash[:],
					FingerprintsChain: testFingerprintsChain,
				},
			},
		},
		{
			name: "Mixed bag",
			entries: []ctonly.Entry{
				{
					Timestamp:         1234,
					IsPrecert:         true,
					Certificate:       testPrecertTBS,
					Precertificate:    testPrecert,
					IssuerKeyHash:     testIssuerKeyHash[:],
					FingerprintsChain: testFingerprintsChain,
				}, {
					Timestamp:         1234,
					IsPrecert:         false,
					Certificate:       testCert,
					FingerprintsChain: testFingerprintsChain,
				},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			bundle := []byte{}
			wantIDs := [][]byte{}
			for _, e := range test.entries {
				bundle = append(bundle, e.LeafData(123)...)
				wantIDs = append(wantIDs, e.MerkleLeafHash(123))
			}

			gotIDs, gotErr := ctMerkleLeafHasher(bundle)
			if gotErr != nil {
				t.Fatalf("ctMerkleLeafHasher: %v", gotErr)
			}
			if lg, lw := len(gotIDs), len(wantIDs); lg != lw {
				t.Fatalf("got %d hashes, want %d", lg, lw)
			}
			for i := range gotIDs {
				if !bytes.Equal(gotIDs[i], wantIDs[i]) {
					t.Fatalf("%d: got ID hash %x, want %x", i, gotIDs[i], wantIDs[i])
				}
			}

		})
	}
}
