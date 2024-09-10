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
//
// This tool will fetch the specified public key from Cloud KMS, and use it to
// create an Ed25519 note.Verifier string using the provided name, which will
// be printed to stdout.
//
// Example usage:
//
//	$ go run github.com/transparency-dev/trillian-tessera/cmd/conformance/gcp/kmsnote \
//		  --key_id=projects/trillian-tessera/locations/us-central1/keyRings/ci-conformance/cryptoKeys/log-signer/cryptoKeyVersions/1 \
//		  --name="ci-conformance"
//
//	ci-conformance+1e5feae8+ARe2PncWChrXMrQtDOqJrKkn2XXOGwsJ+amwLnaWyhEK
package main

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"os"

	kms "cloud.google.com/go/kms/apiv1"
	"golang.org/x/mod/sumdb/note"

	"cloud.google.com/go/kms/apiv1/kmspb"

	"k8s.io/klog/v2"
)

var (
	keyID  = flag.String("key_id", "", "cryptoKeyVersion ID string ('projects/.../locations/.../keyRings/.../cryptoKeys/.../cryptoKeyVersions/...'), see https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions")
	name   = flag.String("name", "", "Name for generated note Verifier")
	output = flag.String("output", "", "Optional filename to write the generated note to, leave unset to write to stdout")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	ctx := context.Background()
	kmClient, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		klog.Exitf("Failed to create KeyManagementClient: %v", err)
	}
	defer kmClient.Close()

	v, err := verifierKeyString(ctx, kmClient, *keyID, *name)
	if err != nil {
		klog.Exitf("Failed to generate verifier string: %v", err)
	}

	if *output == "" {
		fmt.Println(v)
	} else {
		if err := os.WriteFile(*output, []byte(v), 0o644); err != nil {
			klog.Exitf("Failed to write output to %q: %v", *output, err)
		}
	}
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
