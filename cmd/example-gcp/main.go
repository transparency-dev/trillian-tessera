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

// example-gcp is a simple personality showing how to use the Tessera GCP storage
// implmentation.
package main

import (
	"crypto/sha256"
	"flag"
	"io"
	"net/http"
	"os"

	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/storage/gcp"
	"k8s.io/klog/v2"
)

var (
	project = flag.String("project", os.Getenv("GOOGLE_CLOUD_PROJECT"), "GCP Project, take from env if unset")
	listen  = flag.String("listen", ":2024", "Address:port to listen on")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	_, err := gcp.New()
	if err != nil {
		klog.Exitf("Failed to create new GCP storage: %v", err)
	}

	http.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		id := sha256.Sum256(b)
		_ = tessera.NewEntry(b, tessera.WithIdentity(id[:]))

		// TODO: Add entry to log and return assigned index.
	})
}
