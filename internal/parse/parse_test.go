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
	"testing"

	"github.com/transparency-dev/trillian-tessera/internal/parse"
)

func TestCheckpointUnsafe(t *testing.T) {
	testCases := []struct {
		desc       string
		cp         string
		wantOrigin string
		wantSize   uint64
		wantErr    bool
	}{
		{
			desc:       "happy checkpoint",
			cp:         "original.example.com\n42\nqINS1GRFhWHwdkUeqLEoP4yEMkTBBzxBkGwGQlVlVcs=\n",
			wantOrigin: "original.example.com",
			wantSize:   42,
		},
		{
			desc:    "Negative size",
			cp:      "original.example.com\n-42\nqINS1GRFhWHwdkUeqLEoP4yEMkTBBzxBkGwGQlVlVcs=\n",
			wantErr: true,
		},
		{
			desc:       "Bad hash (passes because hashes are not checked)",
			cp:         "original.example.com\n42\nthisisnotright\n",
			wantOrigin: "original.example.com",
			wantSize:   42,
		},
		{
			desc:       "Empty origin",
			cp:         "\n42\nthisisnotright\n",
			wantOrigin: "",
			wantSize:   42,
		},
		{
			desc:    "No origin",
			cp:      "42\nthisisnotright\n",
			wantErr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			origin, size, err := parse.CheckpointUnsafe([]byte(tC.cp))
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
		})
	}
}

func BenchmarkCheckpointUnsafe(b *testing.B) {
	cpRaw := []byte("go.sum database tree\n31700353\nqINS1GRFhWHwdkUeqLEoP4yEMkTBBzxBkGwGQlVlVcs=\n\n— sum.golang.org Az3grnmrIUEDFqHzAElIQCPNoRFRAAdFo47fooyWKMHb89k11GJh5zHIfNCOBmwn/C3YI8oW9/C8DJ87F61QqspBYwM=")
	for i := 0; i < b.N; i++ {
		_, _, err := parse.CheckpointUnsafe(cpRaw)
		if err != nil {
			b.Error(err)
		}
	}
}
