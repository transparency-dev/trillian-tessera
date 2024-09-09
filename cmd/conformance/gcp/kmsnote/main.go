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

// kmsnote is a tool for creating a note verifier from an Ed25519 public key
// retrieved from KMS.
package main

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"log"

	kms "cloud.google.com/go/kms/apiv1"
	"golang.org/x/mod/sumdb/note"

	"cloud.google.com/go/kms/apiv1/kmspb"

	"k8s.io/klog/v2"
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

var (
	keyID = flag.String("key_id", "", "cryptoKeyVersion ID string ('projects/.../locations/.../keyRings/.../cryptoKeys/.../cryptoKeyVersions/...')")
	name  = flag.String("name", "", "Name for generated note Verifier")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	ctx := context.Background()
	kmClient, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		log.Fatalf("failed to create KeyManagementClient: %v", err)
	}
	defer kmClient.Close()

	v, err := verifierKeyString(ctx, kmClient, *keyID, *name)
	if err != nil {
		klog.Exitf("Failed to generate verifier string: %v", err)
	}

	fmt.Println(v)
}

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

// VerifierKeyString returns a string which can be used to create a note
// verifier based on a GCP KMS
// [Ed25519](https://pkg.go.dev/golang.org/x/mod/sumdb/note#hdr-Generating_Keys)
// key.
func verifierKeyString(ctx context.Context, c *kms.KeyManagementClient, kmsKeyName, noteKeyName string) (string, error) {
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
