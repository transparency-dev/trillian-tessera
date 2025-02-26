// Copyright 2025 The Tessera authors. All Rights Reserved.
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

// Package witness contains the implementation for sending out a checkpoint to witnesses
// and retrieving sufficient signatures to satisfy a policy.
package witness

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"

	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/internal/parse"
)

// ProofFetchFn is the signature of a function that can return the consistency proof
// for append-only between the two given tree sizes.
type ProofFetchFn func(ctx context.Context, from, to uint64) [][]byte

// NewWitnessGateway returns a WitnessGateway that will send out new checkpoints to witnesses
// in the group, and will ensure that the policy is satisfied before returning. All outbound
// requests will be done using the given client.
func NewWitnessGateway(group tessera.WitnessGroup, client *http.Client, fetchProof ProofFetchFn) WitnessGateway {
	urls := group.URLs()
	slices.Sort(urls)
	urls = slices.Compact(urls)
	witnesses := make([]witness, 0, len(urls))
	postFnImpl := func(ctx context.Context, url string, body string) (pr postResponse, err error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(body))
		if err != nil {
			return postResponse{}, err
		}
		httpResp, err := client.Do(req)
		if err != nil {
			return postResponse{}, err
		}
		rb, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return postResponse{}, fmt.Errorf("failed to read response body: %v", err)
		}
		_ = httpResp.Body.Close()
		return postResponse{
			statusCode: httpResp.StatusCode,
			body:       rb,
			headers:    httpResp.Header,
		}, nil
	}
	for _, u := range urls {
		witnesses = append(witnesses, witness{
			url:        u,
			size:       0,
			post:       postFnImpl,
			fetchProof: fetchProof,
		})
	}
	return WitnessGateway{
		group:     group,
		witnesses: witnesses,
	}
}

// WitnessGateway allows a log implementation to send out a checkpoint to witnesses.
type WitnessGateway struct {
	group     tessera.WitnessGroup
	witnesses []witness
}

// Witness sends out a new checkpoint (which must be signed by the log), to all witnesses
// and returns the checkpoint as soon as the policy the WitnessGateway was constructed with
// is Satisfied.
func (wg WitnessGateway) Witness(ctx context.Context, cp []byte) ([]byte, error) {
	// TODO(mhutchinson): also set a deadline?
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var waitGroup sync.WaitGroup
	_, size, err := parse.CheckpointUnsafe(cp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse checkpoint from log: %v", err)
	}

	type sigOrErr struct {
		sig []byte
		err error
	}
	results := make(chan sigOrErr)

	// Kick off a goroutine for each witness and send result to results chan
	for _, w := range wg.witnesses {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			sig, err := w.update(ctx, cp, size)
			results <- sigOrErr{
				sig: sig,
				err: err,
			}
		}()
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	// Consume the results coming back from each witness
	var sigBlock bytes.Buffer
	sigBlock.Write(cp)
	for r := range results {
		if r.err != nil {
			err = errors.Join(err, r.err)
			continue
		}
		// Add new signature to the new note we're building
		sigBlock.Write(r.sig)

		// See whether the group is satisfied now
		if newCp := sigBlock.Bytes(); wg.group.Satisfied(newCp) {
			return newCp, nil
		}
	}

	// We can only get here if all witnesses have returned and we're still not satisfied.
	return sigBlock.Bytes(), err
}

// postResponse contains the parts of the witness response needed for logic flow in order to
// allow testing.
type postResponse struct {
	statusCode int
	body       []byte
	headers    map[string][]string
}

// postFn wraps http.Post in order to allow for easier testing.
type postFn func(ctx context.Context, url, body string) (pr postResponse, err error)

// witness is the log's model of a witness's view of this log.
// It has a URL which is the address to which updates to this log's state can be posted to the witness,
// using the https://github.com/C2SP/C2SP/blob/main/tlog-witness.md spec.
// It also has the size of the checkpoint that the log thinks that the witness last signed.
// This is important for sending update proofs.
// This is defaulted to zero on startup and calibrated after the first request, which is expected by the spec:
// `If a client doesn't have information on the latest cosigned checkpoint, it MAY initially make a request with a old size of zero to obtain it`
type witness struct {
	url        string
	size       uint64
	post       postFn
	fetchProof ProofFetchFn
}

func (w witness) update(ctx context.Context, cp []byte, size uint64) ([]byte, error) {
	var proof [][]byte
	if w.size > 0 {
		proof = w.fetchProof(ctx, w.size, size)
	}

	// The request body MUST be a sequence of
	// - a previous size line,
	// - zero or more consistency proof lines,
	// - and an empty line,
	// - followed by a [checkpoint][].
	body := fmt.Sprintf("old %d\n", w.size)
	for _, p := range proof {
		body += base64.StdEncoding.EncodeToString(p) + "\n"
	}
	body += "\n"
	body += string(cp)

	resp, err := w.post(ctx, w.url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to post to witness at %q: %v", w.url, err)
	}

	switch resp.statusCode {
	case http.StatusOK:
		if len(resp.body) == 0 {
			return nil, errors.New("expected response body from witness, but got empty body")
		}
		return resp.body, nil
	case http.StatusConflict:
		// Two cases here: the first is a situation we can recover from, the second isn't.

		// The witness MUST check that the old size matches the size of the latest checkpoint it cosigned
		// for the checkpoint's origin (or zero if it never cosigned a checkpoint for that origin).
		// If it doesn't match, the witness MUST respond with a "409 Conflict" HTTP status code.
		// The response body MUST consist of the tree size of the latest cosigned checkpoint in decimal,
		// followed by a newline (U+000A). The response MUST have a Content-Type of text/x.tlog.size
		ct := resp.headers["Content-Type"]
		if len(ct) == 1 && ct[0] == "text/x.tlog.size" {
			bodyStr := string(resp.body)
			newWitSize, err := strconv.ParseInt(bodyStr, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("witness at %q replied with x.tlog.size but body %q could not be parsed as decimal", w.url, bodyStr)
			}
			if newWitSize < int64(w.size) {
				// This case should not happen unless the witness is misbehaving
				return nil, fmt.Errorf("witness at %q replied with x.tlog.size %d, smaller than known size %d", w.url, newWitSize, w.size)
			}
			w.size = uint64(newWitSize)
			// Witnesses could cause this recursion to go on for longer than expected if the value they kept returning
			// this case with slightly larger values. Consider putting a max recursion cap if context timeout isn't enough.
			return w.update(ctx, cp, size)
		}

		// If the old size matches the checkpoint size, the witness MUST check that the root hashes are also identical.
		// If they don't match, the witness MUST respond with a "409 Conflict" HTTP status code.
		return nil, fmt.Errorf("witness at %q says old root hash did not match previous for size %d: %d", w.url, w.size, resp.statusCode)
	case http.StatusNotFound:
		// If the checkpoint origin is unknown, the witness MUST respond with a "404 Not Found" HTTP status code.
		return nil, fmt.Errorf("witness at %q says checkpoint origin is unknown: %d", w.url, resp.statusCode)
	case http.StatusForbidden:
		// If none of the signatures verify against a trusted public key, the witness MUST respond with a "403 Forbidden" HTTP status code.
		return nil, fmt.Errorf("witness at %q says no signatures verify against trusted public key: %d", w.url, resp.statusCode)
	case http.StatusBadRequest:
		// The old size MUST be equal to or lower than the checkpoint size.
		// Otherwise, the witness MUST respond with a "400 Bad Request" HTTP status code.
		return nil, fmt.Errorf("witness at %q says old checkpoint size of %d is too large: %d", w.url, w.size, resp.statusCode)
	case http.StatusUnprocessableEntity:
		//  If the Merkle Consistency Proof doesn't verify, the witness MUST respond with a "422 Unprocessable Entity" HTTP status code.
		return nil, fmt.Errorf("witness at %q says that the consistency proof is bad: %d", w.url, resp.statusCode)
	default:
		return nil, fmt.Errorf("got bad status code: %v", resp.statusCode)
	}
}
