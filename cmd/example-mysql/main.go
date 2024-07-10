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
	"crypto/sha256"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/storage/mysql"
	"k8s.io/klog/v2"
)

var (
	mysqlURI          = flag.String("mysql_uri", "user:password@tcp(db:3306)/tessera", "Connection string for a MySQL database")
	dbConnMaxLifetime = flag.Duration("db_conn_max_lifetime", 3*time.Minute, "")
	dbMaxOpenConns    = flag.Int("db_max_open_conns", 64, "")
	dbMaxIdleConns    = flag.Int("db_max_idle_conns", 64, "")
	listen            = flag.String("listen", ":2024", "Address:port to listen on")
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

	storage, err := mysql.New(ctx, db)
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
		level, index, width, err := parseTileLevelIndexWidth(r.PathValue("level"), r.PathValue("index"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if _, werr := w.Write([]byte(fmt.Sprintf("Malformed URL: %s", err.Error()))); werr != nil {
				klog.Errorf("/tile/{level}/{index...}: %v", werr)
			}
			return
		}

		tile, err := storage.ReadTile(r.Context(), level, index)
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			klog.Errorf("/tile/{level}/{index...}: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Only add the immutable Cache-Control header when the number of hash in the returned tile match the requested tile width.
		// This ensures the response will not be cached when returning a partial tile on a full tile request.
		if len(tile)/32 == int(width) {
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		}

		if _, err := w.Write(tile); err != nil {
			klog.Errorf("/tile/{level}/{index...}: %v", err)
			return
		}
	})

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

	if err := http.ListenAndServe(*listen, http.DefaultServeMux); err != nil {
		klog.Exitf("ListenAndServe: %v", err)
	}
}

// parseTileLevelWidthIndex takes level and index in string, validates and returns the level, index and width in uint64.
//
// Examples:
// "/tile/0/x001/x234/067" means level 0 and index 1234067 of a full tile.
// "/tile/0/x001/x234/067.p/8" means level 0, index 1234067 and width 8 of a partial tile.
func parseTileLevelIndexWidth(level, index string) (uint64, uint64, uint64, error) {
	l, err := strconv.ParseUint(level, 10, 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse tile level")
	}

	var i, w uint64
	switch indexPaths := strings.Split(index, "/"); len(indexPaths) {
	// Full tile = 3
	// Partial tile = 4
	case 3, 4:
		indexPath := strings.Join(indexPaths[0:3], "")
		indexPath = strings.ReplaceAll(indexPath, "x", "")
		indexPath = strings.ReplaceAll(indexPath, ".p", "")
		i, err = strconv.ParseUint(indexPath, 10, 64)
		if err != nil {
			return 0, 0, 0, fmt.Errorf("failed to parse tile index")
		}

		w = 256
		if len(indexPaths) == 4 {
			w, err = strconv.ParseUint(indexPaths[3], 10, 64)
			if err != nil {
				return 0, 0, 0, fmt.Errorf("failed to parse tile index")
			}
		}
	default:
		return 0, 0, 0, fmt.Errorf("failed to parse tile index")
	}

	return l, i, w, nil
}
