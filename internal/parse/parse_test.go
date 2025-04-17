// Copyright 2024 The Tessera authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package parse_test

import (
	"bytes"
	"encoding/base64"
	"testing"

	"github.com/transparency-dev/trillian-tessera/internal/parse"
)

func TestCheckpointUnsafe(t *testing.T) {
	testCases := []struct {
		desc       string
		cp         string
		wantOrigin string
		wantSize   uint64
		wantHash   []byte
		wantErr    bool
	}{
		{
			desc:       "happy checkpoint",
			cp:         "original.example.com\n42\nqINS1GRFhWHwdkUeqLEoP4yEMkTBBzxBkGwGQlVlVcs=\n",
			wantOrigin: "original.example.com",
			wantSize:   42,
			wantHash:   mustDecodeB64(t, "qINS1GRFhWHwdkUeqLEoP4yEMkTBBzxBkGwGQlVlVcs="),
		},
		{
			desc:    "Negative size",
			cp:      "original.example.com\n-42\nqINS1GRFhWHwdkUeqLEoP4yEMkTBBzxBkGwGQlVlVcs=\n",
			wantErr: true,
		},
		{
			desc:    "Bad hash",
			cp:      "original.example.com\n42\nthisisnotright\n",
			wantErr: true,
		},
		{
			desc:       "Empty origin",
			cp:         "\n42\nqINS1GRFhWHwdkUeqLEoP4yEMkTBBzxBkGwGQlVlVcs=\n",
			wantOrigin: "",
			wantSize:   42,
			wantHash:   mustDecodeB64(t, "qINS1GRFhWHwdkUeqLEoP4yEMkTBBzxBkGwGQlVlVcs="),
		},
		{
			desc:    "No origin",
			cp:      "42\nqINS1GRFhWHwdkUeqLEoP4yEMkTBBzxBkGwGQlVlVcs=\n",
			wantErr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			origin, size, hash, err := parse.CheckpointUnsafe([]byte(tC.cp))
			if gotErr := err != nil; gotErr != tC.wantErr {
				t.Fatalf("gotErr != wantErr (%t != %t): %v", gotErr, tC.wantErr, err)
			}
			if tC.wantErr {
				return
			}
			if tC.wantOrigin != origin {
				t.Errorf("origin: got != want (%v != %v)", origin, tC.wantOrigin)
			}
			if tC.wantSize != size {
				t.Errorf("size : got != want (%v != %v)", size, tC.wantSize)
			}
			if !bytes.Equal(tC.wantHash, hash) {
				t.Errorf("hash : got != want (%v != %v)", hash, tC.wantHash)
			}
		})
	}
}

func mustDecodeB64(t *testing.T, encoded string) []byte {
	t.Helper()
	res, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func BenchmarkCheckpointUnsafe(b *testing.B) {
	cpRaw := []byte("go.sum database tree\n31700353\nqINS1GRFhWHwdkUeqLEoP4yEMkTBBzxBkGwGQlVlVcs=\n\n— sum.golang.org Az3grnmrIUEDFqHzAElIQCPNoRFRAAdFo47fooyWKMHb89k11GJh5zHIfNCOBmwn/C3YI8oW9/C8DJ87F61QqspBYwM=")
	for b.Loop() {
		_, _, _, err := parse.CheckpointUnsafe(cpRaw)
		if err != nil {
			b.Error(err)
		}
	}
}
