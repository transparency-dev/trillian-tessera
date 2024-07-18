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

package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	ct "github.com/google/certificate-transparency-go"
	"github.com/google/certificate-transparency-go/trillian/ctfe"
	"github.com/google/certificate-transparency-go/x509"
	"github.com/transparency-dev/trillian-tessera/ctonly"
)

func parseAddChain(r *http.Request) (e *ctonly.Entry, code int, err error) {
	e, code, err = parseChainOrPreChain(r.Context(), r.Body)
	if err != nil {
		return nil, code, err
	}
	if e.IsPrecert {
		return nil, http.StatusBadRequest, fmt.Errorf("pre-certificate submitted to add-chain")
	}
	return
}

func parsePreChain(r *http.Request) (e *ctonly.Entry, code int, err error) {
	e, code, err = parseChainOrPreChain(r.Context(), r.Body)
	if err != nil {
		return nil, code, err
	}
	if !e.IsPrecert {
		return nil, http.StatusBadRequest, fmt.Errorf("certificate submitted to add-pre-chain")
	}
	return
}

func parseChainOrPreChain(ctx context.Context, reqBody io.ReadCloser) (e *ctonly.Entry, code int, err error) {
	body, err := io.ReadAll(reqBody)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to read body: %w", err)
	}
	var req struct {
		Chain [][]byte
	}
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("failed to parse request: %w", err)
	}
	if len(req.Chain) == 0 {
		return nil, http.StatusBadRequest, fmt.Errorf("empty chain")
	}
	notBefore := time.Date(2000, 0, 0, 0, 0, 0, 0, time.UTC)
	notAfter := time.Date(3000, 0, 0, 0, 0, 0, 0, time.UTC)

	chain, err := ctfe.ValidateChain(req.Chain, ctfe.NewCertValidationOpts(rootsPool, time.Time{}, false, false, &notBefore, &notAfter, false, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}))
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("invalid chain: %w", err)
	}
	e = &ctonly.Entry{Certificate: chain[0].Raw}
	issuers := chain[1:]
	if isPrecert, err := ctfe.IsPrecertificate(chain[0]); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("invalid precertificate: %w", err)
	} else if isPrecert {
		if len(issuers) == 0 {
			return nil, http.StatusBadRequest, fmt.Errorf("missing precertificate issuer")
		}

		var preIssuer *x509.Certificate
		if ct.IsPreIssuer(issuers[0]) {
			preIssuer = issuers[0]
			issuers = issuers[1:]
			if len(issuers) == 0 {
				return nil, http.StatusBadRequest, fmt.Errorf("missing precertificate signing certificate issuer")
			}
		}

		defangedTBS, err := x509.BuildPrecertTBS(chain[0].RawTBSCertificate, preIssuer)
		if err != nil {
			return nil, http.StatusInternalServerError, fmt.Errorf("failed to build TBSCertificate: %w", err)
		}

		e.IsPrecert = true
		e.Certificate = defangedTBS
		e.Precertificate = chain[0].Raw
		if preIssuer != nil {
			e.PrecertSigningCert = preIssuer.Raw
		}
		kh := sha256.Sum256(issuers[0].RawSubjectPublicKeyInfo)
		e.IssuerKeyHash = kh[:]
	}

	return e, 0, nil
}

func buildResponse(ctx context.Context, seq uint64, timestamp uint64, merkleLeaf []byte) (resp []byte, code int, err error) {

	ext, err := ctonly.Extensions{LeafIndex: seq}.Marshal()
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to encode extensions: %w", err)
	}
	// The digitally-signed data of an SCT is technically not a MerkleTreeLeaf,
	// but it's a completely identical structure, except for the second field,
	// which is a SignatureType of value 0 and length 1 instead of a
	// MerkleLeafType of value 0 and length 1.
	sctSignature := []byte("signature")
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to sign SCT: %w", err)
	}

	rsp, err := json.Marshal(&ct.AddChainResponse{
		SCTVersion: ct.V1,
		Timestamp:  uint64(timestamp),
		ID:         []byte("LogID"),
		Extensions: base64.StdEncoding.EncodeToString(ext),
		Signature:  sctSignature,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to encode response: %w", err)
	}

	return rsp, http.StatusOK, nil
}
