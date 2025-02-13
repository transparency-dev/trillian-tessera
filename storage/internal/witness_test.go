// Copyright 2025 Google LLC. All Rights Reserved.
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

package storage

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"golang.org/x/mod/sumdb/note"
)

const (
	log_vkey  = "LilLog+b65b2501+Af40T+NgLuQzeKqU1mbUL4pcmQwCVDK67QSmdJ3Q8LTl"
	log_skey  = "PRIVATE+KEY+LilLog+b65b2501+AbKaNq0e8nx6WOuOH0eYAgPeKPtk8KM3fZBhwr5qzo+p"
	wit1_vkey = "Wit1+55ee4561+AVhZSmQj9+SoL+p/nN0Hh76xXmF7QcHfytUrI1XfSClk"
	wit1_skey = "PRIVATE+KEY+Wit1+55ee4561+AeadRiG7XM4XiieCHzD8lxysXMwcViy5nYsoXURWGrlE"
	wit2_vkey = "Wit2+85ecc407+AWVbwFJte9wMQIPSnEnj4KibeO6vSIOEDUTDp3o63c2x"
	wit2_skey = "PRIVATE+KEY+Wit2+85ecc407+AfPTvxw5eUcqSgivo2vaiC7JPOMUZ/9baHPSDrWqgdGm"
	cp        = "LilLog\n" +
		"34840403\n" +
		"Ux/vc6m0VqNe7o2MbLNrCSwFzFvGBCGNClW2x3up/YI=\n"
)

func TestWitness_UpdateRequest(t *testing.T) {
	cpSize := uint64(34840403)
	s, err := note.NewSigner(log_skey)
	if err != nil {
		t.Fatal(err)
	}
	logSignedCheckpoint, err := note.Sign(&note.Note{Text: cp}, s)
	if err != nil {
		t.Fatal(err)
	}
	sig1 := sigForSigner(t, cp, wit1_skey)

	testCases := []struct {
		desc     string
		proof    [][]byte
		witSize  uint64
		wantErr  bool
		wantBody string
	}{
		{
			desc:     "size 0 no proof needed",
			witSize:  0,
			wantBody: fmt.Sprintf("old 0\n\n%s", logSignedCheckpoint),
		},
		{
			desc:     "non zero size requires proof",
			witSize:  22,
			proof:    [][]byte{[]byte("hello"), []byte("world")},
			wantBody: fmt.Sprintf("old 22\naGVsbG8=\nd29ybGQ=\n\n%s", logSignedCheckpoint),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			var gotBody string
			w := witness{
				url:  "https://example.com/thislittlelogofmine/add-checkpoint",
				size: tC.witSize,
				post: func(ctx context.Context, url string, body string) (postResponse, error) {
					gotBody = body
					return postResponse{
						statusCode: 200,
						body:       append(sig1, '\n'),
					}, nil
				},
				fetchProof: func(ctx context.Context, from uint64, to uint64) [][]byte {
					return tC.proof
				},
			}

			_, err := w.update(ctx, logSignedCheckpoint, cpSize)
			if got, want := err != nil, tC.wantErr; got != want {
				t.Fatalf("got != want (%t != %t): %v", got, want, err)
			}
			if tC.wantErr {
				return
			}

			if gotBody != tC.wantBody {
				t.Errorf("body does not match expected: %q", gotBody)
			}
		})
	}
}

func TestWitness_UpdateResponse(t *testing.T) {
	sig1 := sigForSigner(t, cp, wit1_skey)
	sig2 := sigForSigner(t, cp, wit2_skey)

	testCases := []struct {
		desc    string
		pr      postResponse
		pre     error
		wantErr bool
	}{
		{
			desc: "all good",
			pr: postResponse{
				statusCode: 200,
				body:       sig1,
			},
		}, {
			desc: "all good, two sigs",
			pr: postResponse{
				statusCode: 200,
				body:       append(append(sig1, byte('\n')), sig2...),
			},
		}, {
			desc: "404 is an error",
			pr: postResponse{
				statusCode: 404,
			},
			wantErr: true,
		}, {
			desc: "403 is an error",
			pr: postResponse{
				statusCode: 403,
			},
			wantErr: true,
		}, {
			desc: "422 is an error",
			pr: postResponse{
				statusCode: 422,
			},
			wantErr: true,
		}, {
			desc: "409 with no headers is error",
			pr: postResponse{
				statusCode: 409,
			},
			wantErr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			w := witness{
				url:  "https://example.com/thislittlelogofmine/add-checkpoint",
				size: 0,
				post: func(ctx context.Context, url string, body string) (postResponse, error) {
					return tC.pr, tC.pre
				},
				fetchProof: func(ctx context.Context, from uint64, to uint64) [][]byte {
					return [][]byte{}
				},
			}

			resp, err := w.update(ctx, []byte{}, 0)
			if got, want := err != nil, tC.wantErr; got != want {
				t.Fatalf("got != want (%t != %t): %v", got, want, err)
			}
			if err == nil {
				if !bytes.Equal(resp, tC.pr.body) {
					t.Errorf("expected result %q but got %q", tC.pr.body, resp)
				}
			}
		})
	}
}

func sigForSigner(t *testing.T, cp, skey string) []byte {
	t.Helper()
	s, err := note.NewSigner(skey)
	if err != nil {
		t.Fatal(err)
	}
	witSignedCheckpoint, err := note.Sign(&note.Note{Text: cp}, s)
	if err != nil {
		t.Fatal(err)
	}
	return bytes.Trim(witSignedCheckpoint[len(cp):], "\n")
}
