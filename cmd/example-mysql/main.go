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

// example-mysql is a simple personality showing how to use the Tessera MySQL storage implmentation.
package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"github.com/transparency-dev/trillian-tessera/storage/mysql"
	"golang.org/x/mod/sumdb/note"
	"k8s.io/klog/v2"
)

var (
	mysqlURI          = flag.String("mysql_uri", "user:password@tcp(db:3306)/tessera", "Connection string for a MySQL database")
	dbConnMaxLifetime = flag.Duration("db_conn_max_lifetime", 3*time.Minute, "")
	dbMaxOpenConns    = flag.Int("db_max_open_conns", 64, "")
	dbMaxIdleConns    = flag.Int("db_max_idle_conns", 64, "")
	initSchemaPath    = flag.String("init_schema_path", "", "Location of the schema file if database initialization is needed")
	listen            = flag.String("listen", ":2024", "Address:port to listen on")
	privateKeyPath    = flag.String("private_key_path", "", "Location of private key file")
	publicKeyPath     = flag.String("public_key_path", "", "Location of public key file")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()

	db, err := sql.Open("mysql", *mysqlURI)
	if err != nil {
		klog.Exitf("Failed to connect to DB: %v", err)
	}
	db.SetConnMaxLifetime(*dbConnMaxLifetime)
	db.SetMaxOpenConns(*dbMaxOpenConns)
	db.SetMaxIdleConns(*dbMaxIdleConns)

	initDatabaseSchema(ctx)

	rawPrivateKey, err := os.ReadFile(*privateKeyPath)
	if err != nil {
		klog.Exitf("Failed to read private key file %q: %v", *privateKeyPath, err)
	}
	noteSigner, err := note.NewSigner(string(rawPrivateKey))
	if err != nil {
		klog.Exitf("Failed to create new signer: %v", err)
	}
	rawPublicKey, err := os.ReadFile(*publicKeyPath)
	if err != nil {
		klog.Exitf("Failed to read public key file %q: %v", *publicKeyPath, err)
	}
	noteVerifier, err := note.NewVerifier(string(rawPublicKey))
	if err != nil {
		klog.Exitf("Failed to create new verifier: %v", err)
	}

	storage, err := mysql.New(ctx, db, tessera.WithCheckpointSignerVerifier(noteSigner, noteVerifier))
	if err != nil {
		klog.Exitf("Failed to create new MySQL storage: %v", err)
	}

	http.HandleFunc("GET /checkpoint", func(w http.ResponseWriter, r *http.Request) {
		checkpoint, err := storage.ReadCheckpoint(r.Context())
		if err != nil {
			klog.Errorf("/checkpoint: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(checkpoint); err != nil {
			klog.Errorf("/checkpoint: %v", err)
			return
		}
	})

	http.HandleFunc("GET /tile/{level}/{index...}", func(w http.ResponseWriter, r *http.Request) {
		level, index, width, err := layout.ParseTileLevelIndexWidth(r.PathValue("level"), r.PathValue("index"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if _, werr := w.Write([]byte(fmt.Sprintf("Malformed URL: %s", err.Error()))); werr != nil {
				klog.Errorf("/tile/{level}/{index...}: %v", werr)
			}
			return
		}

		tile, err := storage.ReadTile(r.Context(), level, index, width)
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			klog.Errorf("/tile/{level}/{index...}: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")

		if _, err := w.Write(tile); err != nil {
			klog.Errorf("/tile/{level}/{index...}: %v", err)
			return
		}
	})

	http.HandleFunc("GET /tile/entries/{index...}", func(w http.ResponseWriter, r *http.Request) {
		index, _, err := layout.ParseTileIndexWidth(r.PathValue("index"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if _, werr := w.Write([]byte(fmt.Sprintf("Malformed URL: %s", err.Error()))); werr != nil {
				klog.Errorf("/tile/entries/{index...}: %v", werr)
			}
			return
		}

		entryBundle, err := storage.ReadEntryBundle(r.Context(), index/256)
		if err != nil {
			// TODO: Move this error back into storage implementation.
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			klog.Errorf("/tile/entries/{index...}: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// TODO: Add immutable Cache-Control header.

		if _, err := w.Write(entryBundle); err != nil {
			klog.Errorf("/tile/entries/{index...}: %v", err)
			return
		}
	})

	http.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer func() {
			if err := r.Body.Close(); err != nil {
				klog.Warningf("/add: %v", err)
			}
		}()
		idx, err := storage.Add(r.Context(), tessera.NewEntry(b))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		if _, err = w.Write([]byte(fmt.Sprintf("%d", idx))); err != nil {
			klog.Errorf("/add: %v", err)
			return
		}
	})

	if err := http.ListenAndServe(*listen, http.DefaultServeMux); err != nil {
		klog.Exitf("ListenAndServe: %v", err)
	}
}

func initDatabaseSchema(ctx context.Context) {
	if *initSchemaPath != "" {
		klog.Infof("Initializing database schema")

		db, err := sql.Open("mysql", *mysqlURI+"?multiStatements=true")
		if err != nil {
			klog.Exitf("Failed to connect to DB: %v", err)
		}
		defer func() {
			if err := db.Close(); err != nil {
				klog.Warningf("Failed to close db: %v", err)
			}
		}()

		rawSchema, err := os.ReadFile(*initSchemaPath)
		if err != nil {
			klog.Exitf("Failed to read init schema file %q: %v", *initSchemaPath, err)
		}
		if _, err := db.ExecContext(ctx, string(rawSchema)); err != nil {
			klog.Exitf("Failed to execute init database schema: %v", err)
		}

		klog.Infof("Database schema initialized")
	}
}
