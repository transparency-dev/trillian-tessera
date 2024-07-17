// Copyright 2021 Google LLC. All Rights Reserved.
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

package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/transparency-dev/formats/log"
	"github.com/transparency-dev/merkle/compact"
	"github.com/transparency-dev/merkle/rfc6962"
	"github.com/transparency-dev/serverless-log/api"
	"golang.org/x/mod/sumdb/note"
)

var (
	testOrigin      = "example.com/testdata"
	testLogVerifier = mustMakeVerifier("astra+cad5a3d2+AZJqeuyE/GnknsCNh1eCtDtwdAwKBddOlS8M2eI1Jt4b")
	// Built using serverless/testdata/build_log.sh
	testRawCheckpoints, testCheckpoints = mustLoadTestCheckpoints()
)

func mustMakeVerifier(vs string) note.Verifier {
	v, err := note.NewVerifier(vs)
	if err != nil {
		panic(fmt.Errorf("NewVerifier(%q): %v", vs, err))
	}
	return v
}

func mustLoadTestCheckpoints() ([][]byte, []log.Checkpoint) {
	raws, cps := make([][]byte, 0), make([]log.Checkpoint, 0)
	for i := 0; ; i++ {
		cpName := fmt.Sprintf("checkpoint.%d", i)
		r, err := testLogFetcher(context.Background(), cpName)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				// Probably just no more checkpoints left
				break
			}
			panic(err)
		}
		cp, _, _, err := log.ParseCheckpoint(r, testOrigin, testLogVerifier)
		if err != nil {
			panic(fmt.Errorf("ParseCheckpoint(%s): %v", cpName, err))
		}
		raws, cps = append(raws, r), append(cps, *cp)
	}
	if len(raws) == 0 {
		panic("no checkpoints loaded")
	}
	return raws, cps
}

// testLogFetcher is a fetcher which reads from the checked-in golden test log
// data stored in ../testdata/log
func testLogFetcher(_ context.Context, p string) ([]byte, error) {
	path := filepath.Join("../testdata/log", p)
	return os.ReadFile(path)
}

// fetchCheckpointShim allows fetcher requests for checkpoints to be intercepted.
type fetchCheckpointShim struct {
	// Checkpoints holds raw checkpoints to be returned when the fetcher is asked to retrieve a checkpoint path.
	// The zero-th entry will be returned until Advance is called.
	Checkpoints [][]byte
}

// Fetcher intercepts requests for the checkpoint file, returning the zero-th
// entry in the Checkpoints field. All other requests are passed through
// to the delegate fetcher.
func (f *fetchCheckpointShim) Fetcher(deleg Fetcher) Fetcher {
	return func(ctx context.Context, path string) ([]byte, error) {
		if strings.HasSuffix(path, "checkpoint") {
			if len(f.Checkpoints) == 0 {
				return nil, os.ErrNotExist
			}
			r := f.Checkpoints[0]
			return r, nil
		}
		return deleg(ctx, path)
	}
}

// Advance causes subsequent intercepted checkpoint requests to return
// the next entry in the Checkpoints slice.
func (f *fetchCheckpointShim) Advance() {
	f.Checkpoints = f.Checkpoints[1:]
}

func TestCheckLogStateTracker(t *testing.T) {
	ctx := context.Background()
	h := rfc6962.DefaultHasher

	for _, test := range []struct {
		desc       string
		cpRaws     [][]byte
		wantCpRaws [][]byte
	}{
		{
			desc: "Consistent",
			cpRaws: [][]byte{
				testRawCheckpoints[0],
				testRawCheckpoints[2],
				testRawCheckpoints[3],
				testRawCheckpoints[5],
				testRawCheckpoints[6],
				testRawCheckpoints[10],
			},
			wantCpRaws: [][]byte{
				testRawCheckpoints[0],
				testRawCheckpoints[2],
				testRawCheckpoints[3],
				testRawCheckpoints[5],
				testRawCheckpoints[6],
				testRawCheckpoints[10],
			},
		}, {
			desc: "Identical CP",
			cpRaws: [][]byte{
				testRawCheckpoints[0],
				testRawCheckpoints[0],
				testRawCheckpoints[0],
				testRawCheckpoints[0],
			},
			wantCpRaws: [][]byte{
				testRawCheckpoints[0],
				testRawCheckpoints[0],
				testRawCheckpoints[0],
				testRawCheckpoints[0],
			},
		}, {
			desc: "Identical CP pairs",
			cpRaws: [][]byte{
				testRawCheckpoints[0],
				testRawCheckpoints[0],
				testRawCheckpoints[5],
				testRawCheckpoints[5],
			},
			wantCpRaws: [][]byte{
				testRawCheckpoints[0],
				testRawCheckpoints[0],
				testRawCheckpoints[5],
				testRawCheckpoints[5],
			},
		}, {
			desc: "Out of order",
			cpRaws: [][]byte{
				testRawCheckpoints[5],
				testRawCheckpoints[2],
				testRawCheckpoints[0],
				testRawCheckpoints[3],
			},
			wantCpRaws: [][]byte{
				testRawCheckpoints[5],
				testRawCheckpoints[5],
				testRawCheckpoints[5],
				testRawCheckpoints[5],
			},
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			shim := fetchCheckpointShim{Checkpoints: test.cpRaws}
			f := shim.Fetcher(testLogFetcher)
			lst, err := NewLogStateTracker(ctx, f, h, testRawCheckpoints[0], testLogVerifier, testOrigin, UnilateralConsensus(f))
			if err != nil {
				t.Fatalf("NewLogStateTracker: %v", err)
			}

			for i := range test.cpRaws {
				_, _, newCP, err := lst.Update(ctx)
				if err != nil {
					t.Errorf("Update %d: %v", i, err)
				}
				if got, want := newCP, test.wantCpRaws[i]; !bytes.Equal(got, want) {
					t.Errorf("Update moved to:\n%s\nwant:\n%s", string(got), string(want))
				}

				shim.Advance()
			}
		})
	}
}

func TestCheckConsistency(t *testing.T) {
	ctx := context.Background()

	h := rfc6962.DefaultHasher

	for _, test := range []struct {
		desc    string
		cp      []log.Checkpoint
		wantErr bool
	}{
		{
			desc: "2 CP",
			cp: []log.Checkpoint{
				testCheckpoints[2],
				testCheckpoints[5],
			},
		}, {
			desc: "5 CP",
			cp: []log.Checkpoint{
				testCheckpoints[0],
				testCheckpoints[2],
				testCheckpoints[3],
				testCheckpoints[5],
				testCheckpoints[6],
			},
		}, {
			desc: "big CPs",
			cp: []log.Checkpoint{
				testCheckpoints[3],
				testCheckpoints[7],
				testCheckpoints[8],
			},
		}, {
			desc: "Identical CP",
			cp: []log.Checkpoint{
				testCheckpoints[0],
				testCheckpoints[0],
				testCheckpoints[0],
				testCheckpoints[0],
			},
		}, {
			desc: "Identical CP pairs",
			cp: []log.Checkpoint{
				testCheckpoints[0],
				testCheckpoints[0],
				testCheckpoints[5],
				testCheckpoints[5],
			},
		}, {
			desc: "Out of order",
			cp: []log.Checkpoint{
				testCheckpoints[5],
				testCheckpoints[2],
				testCheckpoints[0],
				testCheckpoints[3],
			},
		}, {
			desc:    "no checkpoints",
			cp:      []log.Checkpoint{},
			wantErr: true,
		}, {
			desc: "one checkpoint",
			cp: []log.Checkpoint{
				testCheckpoints[3],
			},
			wantErr: true,
		}, {
			desc: "two inconsistent CPs",
			cp: []log.Checkpoint{
				{
					Size: 2,
					Hash: []byte("This is a banana"),
				},
				testCheckpoints[4],
			},
			wantErr: true,
		}, {
			desc: "Inconsistent",
			cp: []log.Checkpoint{
				testCheckpoints[5],
				testCheckpoints[2],
				{
					Size: 4,
					Hash: []byte("This is a banana"),
				},
				testCheckpoints[3],
			},
			wantErr: true,
		}, {
			desc: "Inconsistent - clashing CPs",
			cp: []log.Checkpoint{
				{
					Size: 2,
					Hash: []byte("This is a banana"),
				},
				{
					Size: 2,
					Hash: []byte("This is NOT a banana"),
				},
			},
			wantErr: true,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			err := CheckConsistency(ctx, h, testLogFetcher, test.cp)
			if gotErr := err != nil; gotErr != test.wantErr {
				t.Fatalf("wantErr: %t, got %v", test.wantErr, err)
			}
		})
	}
}

func TestNodeCacheHandlesInvalidRequest(t *testing.T) {
	ctx := context.Background()
	wantBytes := []byte("one")
	f := func(_ context.Context, _, _ uint64) (*api.Tile, error) {
		return &api.Tile{
			Nodes: [][]byte{wantBytes},
		}, nil
	}

	// Large tree, but we're emulating skew since f, above, will return a tile which only knows about 1
	// leaf.
	nc := newNodeCache(f, 10)

	if got, err := nc.GetNode(ctx, compact.NewNodeID(0, 0)); err != nil {
		t.Errorf("got %v, want no error", err)
	} else if !bytes.Equal(got, wantBytes) {
		t.Errorf("got %v, want %v", got, wantBytes)
	}

	if _, err := nc.GetNode(ctx, compact.NewNodeID(0, 1)); err == nil {
		t.Error("got no error, want error because ID is out of range")
	}
}

func TestHandleZeroRoot(t *testing.T) {
	zeroCP := testCheckpoints[0]
	if zeroCP.Size != 0 {
		t.Fatal("BadData: checkpoint has non-zero size")
	}
	if len(zeroCP.Hash) == 0 {
		t.Fatal("BadTestData: checkpoint.0 has empty root hash")
	}
	if _, err := NewProofBuilder(context.Background(), zeroCP, rfc6962.DefaultHasher.HashChildren, testLogFetcher); err != nil {
		t.Fatalf("NewProofBuilder: %v", err)
	}
}
