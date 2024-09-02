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

package sctfe

import (
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/google/certificate-transparency-go/tls"
	"github.com/transparency-dev/formats/log"
	"github.com/transparency-dev/trillian-tessera/personalities/sctfe/modules/dedup"
	"golang.org/x/crypto/cryptobyte"
	"golang.org/x/mod/sumdb/note"

	ct "github.com/google/certificate-transparency-go"
)

func buildV1SCT(signer crypto.Signer, leaf *ct.MerkleTreeLeaf) (*ct.SignedCertificateTimestamp, error) {
	// Serialize SCT signature input to get the bytes that need to be signed
	sctInput := ct.SignedCertificateTimestamp{
		SCTVersion: ct.V1,
		Timestamp:  leaf.TimestampedEntry.Timestamp,
		Extensions: leaf.TimestampedEntry.Extensions,
	}
	data, err := ct.SerializeSCTSignatureInput(sctInput, ct.LogEntry{Leaf: *leaf})
	if err != nil {
		return nil, fmt.Errorf("failed to serialize SCT data: %v", err)
	}

	h := sha256.Sum256(data)
	signature, err := signer.Sign(rand.Reader, h[:], crypto.SHA256)
	if err != nil {
		return nil, fmt.Errorf("failed to sign SCT data: %v", err)
	}

	digitallySigned := ct.DigitallySigned{
		Algorithm: tls.SignatureAndHashAlgorithm{
			Hash:      tls.SHA256,
			Signature: tls.SignatureAlgorithmFromPubKey(signer.Public()),
		},
		Signature: signature,
	}

	logID, err := GetCTLogID(signer.Public())
	if err != nil {
		return nil, fmt.Errorf("failed to get logID for signing: %v", err)
	}

	return &ct.SignedCertificateTimestamp{
		SCTVersion: ct.V1,
		LogID:      ct.LogID{KeyID: logID},
		Timestamp:  sctInput.Timestamp,
		Extensions: sctInput.Extensions,
		Signature:  digitallySigned,
	}, nil
}

type RFC6962NoteSignature struct {
	timestamp uint64
	signature ct.DigitallySigned
}

// buildCp builds a https://c2sp.org/static-ct-api checkpoint.
// TODO(phboneff): add tests
func buildCp(signer crypto.Signer, size uint64, timeMilli uint64, hash []byte) ([]byte, error) {
	sth := ct.SignedTreeHead{
		Version:   ct.V1,
		TreeSize:  size,
		Timestamp: timeMilli,
	}
	copy(sth.SHA256RootHash[:], hash)

	sthBytes, err := ct.SerializeSTHSignatureInput(sth)
	if err != nil {
		return nil, fmt.Errorf("ct.SerializeSTHSignatureInput(): %v", err)
	}

	h := sha256.Sum256(sthBytes)
	signature, err := signer.Sign(rand.Reader, h[:], crypto.SHA256)
	if err != nil {
		return nil, err
	}

	rfc6962Note := RFC6962NoteSignature{
		timestamp: sth.Timestamp,
		signature: ct.DigitallySigned{
			Algorithm: tls.SignatureAndHashAlgorithm{
				Hash:      tls.SHA256,
				Signature: tls.SignatureAlgorithmFromPubKey(signer.Public()),
			},
			Signature: signature,
		},
	}

	sig, err := tls.Marshal(rfc6962Note)
	if err != nil {
		return nil, fmt.Errorf("couldn't encode RFC6962NoteSignature: %w", err)
	}

	return sig, nil
}

// CpSigner implements note.Signer. It can generate https://c2sp.org/static-ct-api checkpoints.
type CpSigner struct {
	sthSigner  crypto.Signer
	origin     string
	keyHash    uint32
	timeSource TimeSource
}

// Sign takes an unsigned checkpoint, and signs it with a https://c2sp.org/static-ct-api signature.
// Returns an error if the message doesn't parse as a checkpoint, or if the
// checkpoint origin doesn't match with the Signer's origin.
// TODO(phboneff): add tests
func (cts *CpSigner) Sign(msg []byte) ([]byte, error) {
	ckpt := &log.Checkpoint{}
	rest, err := ckpt.Unmarshal(msg)

	if len(rest) != 0 {
		return nil, fmt.Errorf("checkpoint contains trailing data: %s", string(rest))
	} else if err != nil {
		return nil, fmt.Errorf("ckpt.Unmarshal: %v", err)
	} else if ckpt.Origin != cts.origin {
		return nil, fmt.Errorf("checkpoint's origin %s doesn't match signer's origin %s", ckpt.Origin, cts.origin)
	}

	// TODO(phboneff): make sure that it's ok to generate the timestamp here
	t := uint64(cts.timeSource.Now().UnixMilli())
	sig, err := buildCp(cts.sthSigner, ckpt.Size, t, ckpt.Hash[:])
	if err != nil {
		return nil, fmt.Errorf("coudn't sign CT checkpoint: %v", err)
	}
	return sig, nil
}

func (cts *CpSigner) Name() string {
	return cts.origin
}

func (cts *CpSigner) KeyHash() uint32 {
	return cts.keyHash
}

// NewCpSigner returns a new note signer that can sign https://c2sp.org/static-ct-api checkpoints.
// TODO(phboneff): add tests
func NewCpSigner(signer crypto.Signer, origin string, logID [32]byte, timeSource TimeSource) note.Signer {
	h := sha256.New()
	h.Write([]byte(origin))
	h.Write([]byte{0x0A}) // newline
	h.Write([]byte{0x05}) // signature type
	h.Write(logID[:])
	sum := h.Sum(nil)

	ctSigner := &CpSigner{
		sthSigner:  signer,
		origin:     origin,
		keyHash:    binary.BigEndian.Uint32(sum),
		timeSource: timeSource,
	}
	return ctSigner
}

// DedupFromBundle converts a bundle into an array of {leafID, idx}.
//
// The index of a leaf is computed from its position in the log, instead of parsing SCTs.
// Greatly inspired by https://github.com/FiloSottile/sunlight/blob/main/tile.go
func DedupFromBundle(bundle []byte, bundleIdx uint64) ([]dedup.KV, error) {
	kvs := []dedup.KV{}
	s := cryptobyte.String(bundle)

	for len(s) > 0 {
		var timestamp uint64
		var entryType uint16
		var extensions, fingerprints cryptobyte.String
		if !s.ReadUint64(&timestamp) || !s.ReadUint16(&entryType) || timestamp > math.MaxInt64 {
			return nil, fmt.Errorf("invalid data tile")
		}
		crt := []byte{}
		switch entryType {
		case 0: // x509_entry
			if !s.ReadUint24LengthPrefixed((*cryptobyte.String)(&crt)) ||
				!s.ReadUint16LengthPrefixed(&extensions) ||
				!s.ReadUint16LengthPrefixed(&fingerprints) {
				return nil, fmt.Errorf("invalid data tile x509_entry")
			}
		case 1: // precert_entry
			IssuerKeyHash := [32]byte{}
			var defangedCrt, extensions cryptobyte.String
			if !s.CopyBytes(IssuerKeyHash[:]) ||
				!s.ReadUint24LengthPrefixed(&defangedCrt) ||
				!s.ReadUint16LengthPrefixed(&extensions) ||
				!s.ReadUint24LengthPrefixed((*cryptobyte.String)(&crt)) ||
				!s.ReadUint16LengthPrefixed(&fingerprints) {
				return nil, fmt.Errorf("invalid data tile precert_entry")
			}
		default:
			return nil, fmt.Errorf("invalid data tile: unknown type %d", entryType)
		}
		k := sha256.Sum256(crt)
		kvs = append(kvs, dedup.KV{K: k[:], V: bundleIdx*256 + uint64(len(kvs))})
	}
	return kvs, nil
}
