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

// gcp is a simple personality allowing to run conformance/compliance/performance tests and showing how to use the Tessera GCP storage implmentation.
package main

import (
	"context"
	"crypto/sha256"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	kms "cloud.google.com/go/kms/apiv1"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/storage/gcp"
	"golang.org/x/mod/sumdb/note"
	"k8s.io/klog/v2"
)

var (
	bucket     = flag.String("bucket", "", "Bucket to use for storing log")
	listen     = flag.String("listen", ":2024", "Address:port to listen on")
	project    = flag.String("project", os.Getenv("GOOGLE_CLOUD_PROJECT"), "GCP Project, take from env if unset")
	spanner    = flag.String("spanner", "", "Spanner resource URI ('projects/.../...')")
	kmsKeyName = flag.String("kms_key", "", "GCP KMS key name for signing checkpoints")
	origin     = flag.String("origin", "", "Log origin string")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()

	if *origin == "" {
		klog.Exit("Must supply --origin")
	}

	signer, verifier, kmsClose := signerFromFlags(ctx)
	defer func() {
		if err := kmsClose(); err != nil {
			klog.Errorf("kmsClose(): %v", err)
		}
	}()

	// Create our Tessera storage backend:
	gcpCfg := storageConfigFromFlags()
	storage, err := gcp.New(ctx, gcpCfg,
		tessera.WithCheckpointSignerVerifier(signer, verifier),
		tessera.WithBatching(1024, time.Second),
		tessera.WithPushback(10*4096),
	)
	if err != nil {
		klog.Exitf("Failed to create new GCP storage: %v", err)
	}

	// Expose a HTTP handler for the conformance test writes.
	// This should accept arbitary bytes POSTed to /add, and return an ascii
	// decimal representation of the index assigned to the entry.
	http.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer func() {
			_ = r.Body.Close()
		}()

		idx, err := storage.Add(r.Context(), tessera.NewEntry(b))()
		if err != nil {
			if errors.Is(err, tessera.ErrPushback) {
				w.Header().Add("Retry-After", "1")
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		// Write out the assigned index
		_, _ = w.Write([]byte(fmt.Sprintf("%d", idx)))
	})

	if err := http.ListenAndServe(*listen, http.DefaultServeMux); err != nil {
		klog.Exitf("ListenAndServe: %v", err)
	}
}

// storageConfigFromFlags returns a gcp.Config struct populated with values
// provided via flags.
func storageConfigFromFlags() gcp.Config {
	if *project == "" {
		klog.Exit("--project flag or GOOGLE_CLOUD_PROJECT env must be set.")
	}
	if *bucket == "" {
		klog.Exit("--bucket must be set")
	}
	if *spanner == "" {
		klog.Exit("--spanner must be set")
	}
	return gcp.Config{
		ProjectID: *project,
		Bucket:    *bucket,
		Spanner:   *spanner,
	}
}

// signerFromFlags creates and returns a new KMSSigner from the flags, along with a close func which
// should be called when we're finished with the signer.
func signerFromFlags(ctx context.Context) (note.Signer, note.Verifier, func() error) {
	kmClient, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		klog.Fatalf("Failed to create KeyManagementClient: %v", err)
	}
	signer, err := NewKMSSigner(ctx, kmClient, *kmsKeyName, *origin)
	if err != nil {
		klog.Exitf("Failed to create new signer: %v", err)
	}
	vRaw, err := VerifierKeyString(ctx, kmClient, *kmsKeyName, *origin)
	if err != nil {
		klog.Exitf("Failed to create verifier string: %v", err)
	}
	verifier, err := note.NewVerifier(vRaw)
	if err != nil {
		klog.Exitf("Failed to create verifier from %q: %v", vRaw, err)
	}

	return signer, verifier, kmClient.Close
}
