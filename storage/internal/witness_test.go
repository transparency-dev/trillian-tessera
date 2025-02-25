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
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	tessera "github.com/transparency-dev/trillian-tessera"
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

func TestWitnessGateway_Update(t *testing.T) {
	logVerifier := mustCreateVerifier(t, log_vkey)
	logSigner := mustCreateSigner(t, log_skey)

	logSignedCheckpoint, err := note.Sign(&note.Note{Text: cp}, logSigner)
	if err != nil {
		t.Fatal(err)
	}

	// Set up a fake server hosting the witnesses.
	// The witnesses just sign the checkpoint with whatever key is requested, they don't check the body at all.
	// An improvement on this would be to make the fake witnesses more realistic, but it's a non-trivial
	// amount of code to add to this already long test!
	var wit1 tessera.Witness
	var wit2 tessera.Witness
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w1u, err := url.Parse(wit1.Url)
		if err != nil {
			t.Fatal(err)
		}
		w2u, err := url.Parse(wit2.Url)
		if err != nil {
			t.Fatal(err)
		}

		switch r.URL.String() {
		case w1u.Path:
			_, _ = w.Write(append(sigForSigner(t, cp, wit1_skey), '\n'))
		case w2u.Path:
			_, _ = w.Write(append(sigForSigner(t, cp, wit2_skey), '\n'))
		default:
			t.Fatalf("Unknown case: %s", r.URL.String())
		}
	}))
	baseUrl, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	wit1, err = tessera.NewWitness(wit1_vkey, baseUrl)
	if err != nil {
		t.Fatal(err)
	}
	wit2, err = tessera.NewWitness(wit2_vkey, baseUrl)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		desc     string
		group    tessera.WitnessGroup
		wantSigs int
	}{
		{
			desc:     "no witnesses",
			group:    tessera.WitnessGroup{},
			wantSigs: 0,
		},
		{
			desc:     "one optional witness",
			group:    tessera.NewWitnessGroup(0, wit1),
			wantSigs: 0,
		},
		{
			desc:     "two optional witnesses",
			group:    tessera.NewWitnessGroup(0, wit1, wit2),
			wantSigs: 0,
		},
		{
			desc:     "one required witness",
			group:    tessera.NewWitnessGroup(1, wit1),
			wantSigs: 1,
		},
		{
			desc:     "one required witness out of 2",
			group:    tessera.NewWitnessGroup(1, wit1, wit2),
			wantSigs: 1,
		},
		{
			desc:     "two required witnesses",
			group:    tessera.NewWitnessGroup(2, wit1, wit2),
			wantSigs: 2,
		},
		{
			desc:     "one required witness twice",
			group:    tessera.NewWitnessGroup(2, wit1, wit1),
			wantSigs: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()

			fetchProof := func(ctx context.Context, from, to uint64) [][]byte {
				return nil
			}
			g := NewWitnessGateway(tC.group, ts.Client(), fetchProof)

			witnessedCP, err := g.Witness(ctx, logSignedCheckpoint)
			if err != nil {
				t.Fatal(err)
			}
			n, err := note.Open(witnessedCP, note.VerifierList(logVerifier, wit1.Key, wit2.Key))
			if err != nil {
				t.Fatalf("failed to open note %q: %v", witnessedCP, err)
			}
			if len(n.Sigs)-1 < tC.wantSigs {
				t.Errorf("wanted %d sigs but got %d", tC.wantSigs, len(n.Sigs)-1)
			}
		})
	}
}

func TestWitness_UpdateRequest(t *testing.T) {
	cpSize := uint64(34840403)
	s := mustCreateSigner(t, log_skey)

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

func mustCreateVerifier(t *testing.T, vkey string) note.Verifier {
	t.Helper()
	verifier, err := note.NewVerifier(vkey)
	if err != nil {
		t.Fatal(err)
	}
	return verifier
}

func mustCreateSigner(t *testing.T, skey string) note.Signer {
	t.Helper()
	signer, err := note.NewSigner(skey)
	if err != nil {
		t.Fatal(err)
	}
	return signer
}
