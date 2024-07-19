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

// ct-example-gcp is a simple personality showing how to use the Tessera GCP storage
// implmentation with CT Tiles API support
package main

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/certificate-transparency-go/x509util"
	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/storage/gcp"
	"golang.org/x/crypto/hkdf"
	"golang.org/x/mod/sumdb/note"
	"k8s.io/klog/v2"
)

var (
	bucket  = flag.String("bucket", "", "Bucket to use for storing log")
	listen  = flag.String("listen", ":2024", "Address:port to listen on")
	project = flag.String("project", os.Getenv("GOOGLE_CLOUD_PROJECT"), "GCP Project, take from env if unset")
	spanner = flag.String("spanner", "", "Spanner resource URI ('projects/.../...')")
	signer  = flag.String("signer", "", "Path to file containing log private key")
	roots   = flag.String("roots", "/etc/ssl/certs/*", "Glob of files to use for root certs (PEM format). Individual files may contain multiple PEM entries")
)

var (
	rootsPool = getRoots()
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()

	gcpCfg := gcp.Config{
		ProjectID: *project,
		Bucket:    *bucket,
		Spanner:   *spanner,
	}
	signer, logK := signerFromFlags()
	storage, err := gcp.New(ctx, gcpCfg, tessera.WithCheckpointSigner(signer))
	add := tessera.NewCertificateTransparencySequencedWriter(storage)
	if err != nil {
		klog.Exitf("Failed to create new GCP storage: %v", err)
	}

	l := log{
		k:   logK,
		add: add,
	}

	http.HandleFunc("POST /ct/v1/add-chain", func(w http.ResponseWriter, r *http.Request) {
		rsp, code, err := l.parseAddChain(r)
		if err != nil {
			klog.V(3).Infof("parseAddChain: %v", err)
			http.Error(w, err.Error(), code)
			return
		}
		_, _ = w.Write(rsp)
	})
	http.HandleFunc("POST /ct/v1/add-pre-chain", func(w http.ResponseWriter, r *http.Request) {
		rsp, code, err := l.parsePreChain(r)
		if err != nil {
			klog.V(3).Infof("parsePreChain: %v", err)
			http.Error(w, err.Error(), code)
			return
		}
		_, _ = w.Write(rsp)
	})

	// TODO: remove this proxy
	serveGCS := func(w http.ResponseWriter, r *http.Request) {
		resource := strings.TrimLeft(r.URL.Path, "/")
		b, err := storage.Get(r.Context(), resource)
		if err != nil {
			klog.V(1).Infof("Get: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Get: %v", err)))
			return
		}
		_, _ = w.Write(b)
	}
	http.HandleFunc("GET /checkpoint", serveGCS)
	http.HandleFunc("GET /tile/", serveGCS)

	if err := http.ListenAndServe(*listen, http.DefaultServeMux); err != nil {
		klog.Exitf("ListenAndServe: %v", err)
	}
}

func signerFromFlags() (note.Signer, *ecdsa.PrivateKey) {
	raw, err := os.ReadFile(*signer)
	if err != nil {
		klog.Exitf("Failed to read secret key file %q: %v", *signer, err)
	}
	signer, err := note.NewSigner(string(raw))
	if err != nil {
		klog.Exitf("Failed to create new signer: %v", err)
	}

	// Look away now. 🙈
	logK, err := ecdsa.GenerateKey(elliptic.P256(), hkdf.New(crypto.SHA256.New, raw, []byte("log key"), nil))
	if err != nil {
		panic(fmt.Errorf("failed to generate log key: %v", err))
	}
	return signer, logK
}

func getRoots() *x509util.PEMCertPool {
	files, err := filepath.Glob(*roots)
	if err != nil {
		klog.Errorf("glob(%q): %v", *roots, err)
	}

	roots := []byte{}

	for _, file := range files {
		pem, err := os.ReadFile(file)
		if err != nil {
			klog.Fatalf("Read(%q): %v", file, err)
		}
		roots = append(roots, pem...)
	}
	pool := x509util.NewPEMCertPool()
	pool.AppendCertsFromPEM(roots)

	return pool
}
