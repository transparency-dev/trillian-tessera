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
	"crypto/ed25519"
	"crypto/sha256"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"errors"

	kms "cloud.google.com/go/kms/apiv1"
	"golang.org/x/mod/sumdb/note"

	"cloud.google.com/go/kms/apiv1/kmspb"
)

const (
	// KeyVersionNameFormat is the GCP resource identifier for a key version.
	// google.cloud.kms.v1.CryptoKeyVersion.name
	// https://cloud.google.com/php/docs/reference/cloud-kms/latest/V1.CryptoKeyVersion
	KeyVersionNameFormat = "projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%d"
	// From
	// https://cs.opensource.google/go/x/mod/+/refs/tags/v0.12.0:sumdb/note/note.go;l=232;drc=baa5c2d058db25484c20d76985ba394e73176132
	algEd25519 = 1
)

func publicKeyFromPEM(pemKey []byte) ([]byte, error) {
	block, _ := pem.Decode(pemKey)
	if block == nil {
		return nil, errors.New("failed to decode pemKey")
	}

	k, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := k.(ed25519.PublicKey)
	if !ok {
		return nil, errors.New("failed to assert ed25519.PublicKey type")
	}

	return publicKey, nil
}

// keyHash calculates the ed25519 key hash from the key name and public key.
func keyHash(keyName string, publicKey []byte) (uint32, error) {
	h := sha256.New()
	h.Write([]byte(keyName))
	h.Write([]byte("\n"))
	prefixedPublicKey := append([]byte{algEd25519}, publicKey...)
	h.Write(prefixedPublicKey)
	sum := h.Sum(nil)

	return binary.BigEndian.Uint32(sum), nil
}

// Signer is an implementation of a
// [note signer](https://pkg.go.dev/golang.org/x/mod/sumdb/note#Signer) which
// interfaces with GCP KMS.
type Signer struct {
	// ctx must be stored because Signer is used as an implementation of the
	// note.Signer interface, which does not allow for a context in the Sign
	// method. However, the KMS AsymmetricSign API requires a context.
	ctx        context.Context
	client     *kms.KeyManagementClient
	keyHash    uint32
	keyName    string
	kmsKeyName string
}

// NewKMSSigner creates a signer which uses an Ed25519 key in GCP KMS.
// See https://cloud.google.com/kms/docs/algorithms#elliptic_curve_signing_algorithms
//
// kmsKeyName is the GCP KMS name of the key to be used.
// noteKeyName is the value used as the signer name in the note signature.
func NewKMSSigner(ctx context.Context, c *kms.KeyManagementClient, kmsKeyName, noteKeyName string) (*Signer, error) {
	s := &Signer{}

	s.client = c
	s.ctx = ctx
	s.keyName = noteKeyName
	s.kmsKeyName = kmsKeyName

	// Set keyHash.
	req := &kmspb.GetPublicKeyRequest{
		Name: kmsKeyName,
	}
	resp, err := c.GetPublicKey(ctx, req)
	if err != nil {
		return nil, err
	}

	publicKey, err := publicKeyFromPEM([]byte(resp.Pem))
	if err != nil {
		return nil, err
	}

	kh, err := keyHash(s.keyName, publicKey)
	if err != nil {
		return nil, err
	}
	s.keyHash = kh

	return s, nil
}

// Name identifies the key that this Signer uses.
func (s *Signer) Name() string {
	return s.keyName
}

// KeyHash returns the computed key hash of the signer's public key and name.
// It is used as a hint in identifying the correct key to verify with.
func (s *Signer) KeyHash() uint32 {
	return s.keyHash
}

// Sign returns a signature for the given message.
func (s *Signer) Sign(msg []byte) ([]byte, error) {
	req := &kmspb.AsymmetricSignRequest{
		Name: s.kmsKeyName,
		Data: msg,
	}
	resp, err := s.client.AsymmetricSign(s.ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetSignature(), nil
}

// VerifierKeyString returns a string which can be used to create a note
// verifier based on a GCP KMS
// [Ed25519](https://pkg.go.dev/golang.org/x/mod/sumdb/note#hdr-Generating_Keys)
// key.
func VerifierKeyString(ctx context.Context, c *kms.KeyManagementClient, kmsKeyName, noteKeyName string) (string, error) {
	req := &kmspb.GetPublicKeyRequest{
		Name: kmsKeyName,
	}
	resp, err := c.GetPublicKey(ctx, req)
	if err != nil {
		return "", err
	}

	publicKey, err := publicKeyFromPEM([]byte(resp.Pem))
	if err != nil {
		return "", err
	}

	return note.NewEd25519VerifierKey(noteKeyName, publicKey)
}
