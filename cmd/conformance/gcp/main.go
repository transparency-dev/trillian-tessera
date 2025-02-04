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
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/transparency-dev/merkle/rfc6962"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api"
	"github.com/transparency-dev/trillian-tessera/storage/gcp"
	"golang.org/x/mod/sumdb/note"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"k8s.io/klog/v2"
)

var (
	bucket            = flag.String("bucket", "", "Bucket to use for storing log")
	listen            = flag.String("listen", ":2024", "Address:port to listen on")
	spanner           = flag.String("spanner", "", "Spanner resource URI ('projects/.../...')")
	signer            = flag.String("signer", "", "Note signer to use to sign checkpoints")
	persistentDedup   = flag.Bool("gcp_dedup", false, "EXPERIMENTAL: Set to true to enable persistent dedupe storage")
	additionalSigners = []string{}
)

func init() {
	flag.Func("additional_signer", "Additional note signer for checkpoints, may be specified multiple times", func(s string) error {
		additionalSigners = append(additionalSigners, s)
		return nil
	})
}

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()

	s, a := signerFromFlags()

	// Create our Tessera storage backend:
	gcpCfg := storageConfigFromFlags()
	driver, err := gcp.New(ctx, gcpCfg,
		tessera.WithCheckpointSigner(s, a...),
		tessera.WithCheckpointInterval(10*time.Second),
		tessera.WithBatching(1024, time.Second),
		tessera.WithPushback(10*4096),
	)
	if err != nil {
		klog.Exitf("Failed to create new GCP storage: %v", err)
	}

	dedups := make([]func(tessera.AddFn) tessera.AddFn, 0, 2)
	dedups = append(dedups, tessera.InMemoryDedupe(256))
	// PersistentDedup is currently experimental, so there's no terraform or documentation yet!
	if *persistentDedup {
		dd, err := gcp.NewDedup(ctx, fmt.Sprintf("%s_dedup", *spanner))
		if err != nil {
			klog.Exitf("Failed to create new GCP dedupe: %v", err)
		}
		dedups = append(dedups, dd.AppendDecorator())

		go func() {
			if err := tessera.Follow(ctx, driver, dd.Follower(BundleHasher)); err != nil {
				klog.Exitf("Follow: %v", err)
			}
		}()
	}

	addFn, _, err := tessera.NewAppender(driver, dedups...)
	if err != nil {
		klog.Exit(err)
	}

	// Expose a HTTP handler for the conformance test writes.
	// This should accept arbitrary bytes POSTed to /add, and return an ascii
	// decimal representation of the index assigned to the entry.
	http.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		idx, err := addFn(r.Context(), tessera.NewEntry(b))()
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

	h2s := &http2.Server{}
	h1s := &http.Server{
		Addr:    *listen,
		Handler: h2c.NewHandler(http.DefaultServeMux, h2s),
	}

	if err := h1s.ListenAndServe(); err != nil {
		klog.Exitf("ListenAndServe: %v", err)
	}
}

// storageConfigFromFlags returns a gcp.Config struct populated with values
// provided via flags.
func storageConfigFromFlags() gcp.Config {
	if *bucket == "" {
		klog.Exit("--bucket must be set")
	}
	if *spanner == "" {
		klog.Exit("--spanner must be set")
	}
	return gcp.Config{
		Bucket:  *bucket,
		Spanner: *spanner,
	}
}

func signerFromFlags() (note.Signer, []note.Signer) {
	s, err := note.NewSigner(*signer)
	if err != nil {
		klog.Exitf("Failed to create new signer: %v", err)
	}

	var a []note.Signer
	for _, as := range additionalSigners {
		s, err := note.NewSigner(as)
		if err != nil {
			klog.Exitf("Failed to create additional signer: %v", err)
		}
		a = append(a, s)
	}

	return s, a
}

// BundleHasher parses a C2SP tlog-tile bundle and returns the leaf hashes of each entry it contains.
// TODO: figure out where this should live/how it should work
func BundleHasher(bundle []byte) ([][]byte, error) {
	eb := &api.EntryBundle{}
	if err := eb.UnmarshalText(bundle); err != nil {
		return nil, fmt.Errorf("unmarshal: %v", err)
	}
	r := make([][]byte, 0, len(eb.Entries))
	for _, e := range eb.Entries {
		h := rfc6962.DefaultHasher.HashLeaf(e)
		r = append(r, h[:])
	}
	return r, nil
}
