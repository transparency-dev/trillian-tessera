// Copyright 2025 The Tessera authors. All Rights Reserved.
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

// mysql-migrate is a command-line tool for migrating data from a tlog-tiles
// compliant log, into a Tessera log instance.
package main

import (
	"context"
	"database/sql"
	"encoding/base64"
	"flag"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	tessera "github.com/transparency-dev/trillian-tessera"
	"github.com/transparency-dev/trillian-tessera/client"
	"github.com/transparency-dev/trillian-tessera/cmd/experimental/migrate/internal"
	"github.com/transparency-dev/trillian-tessera/storage/mysql"
	"k8s.io/klog/v2"
)

var (
	mysqlURI          = flag.String("mysql_uri", "user:password@tcp(db:3306)/tessera", "Connection string for a MySQL database")
	dbConnMaxLifetime = flag.Duration("db_conn_max_lifetime", 3*time.Minute, "")
	dbMaxOpenConns    = flag.Int("db_max_open_conns", 64, "")
	dbMaxIdleConns    = flag.Int("db_max_idle_conns", 64, "")
	initSchemaPath    = flag.String("init_schema_path", "", "Location of the schema file if database initialization is needed")

	sourceURL  = flag.String("source_url", "", "Base URL for the source log.")
	numWorkers = flag.Int("num_workers", 30, "Number of migration worker goroutines.")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()

	srcURL, err := url.Parse(*sourceURL)
	if err != nil {
		klog.Exitf("Invalid --source_url %q: %v", *sourceURL, err)
	}
	src, err := client.NewHTTPFetcher(srcURL, nil)
	if err != nil {
		klog.Exitf("Failed to create HTTP fetcher: %v", err)
	}
	sourceCP, err := src.ReadCheckpoint(ctx)
	if err != nil {
		klog.Exitf("Failed to read source checkpoint: %v", err)
	}
	bits := strings.Split(string(sourceCP), "\n")
	sourceSize, err := strconv.ParseUint(bits[1], 10, 64)
	if err != nil {
		klog.Exitf("Invalid CP size %q: %v", bits[1], err)
	}
	sourceRoot, err := base64.StdEncoding.DecodeString(bits[2])
	if err != nil {
		klog.Exitf("Invalid checkpoint roothash %q: %v", bits[2], err)
	}

	db := createDatabaseOrDie(ctx)

	// Initialise the Tessera MySQL storage
	driver, err := mysql.New(ctx, db)
	if err != nil {
		klog.Exitf("Failed to create new MySQL storage: %v", err)
	}

	opts := tessera.NewMigrationOptions()

	m, err := tessera.NewMigrationTarget(ctx, driver, opts)
	if err != nil {
		klog.Exitf("Failed to create MigrationTarget: %v", err)
	}

	if err := internal.Migrate(context.Background(), *numWorkers, sourceSize, sourceRoot, src.ReadEntryBundle, m); err != nil {
		klog.Exitf("Migrate failed: %v", err)
	}

	// TODO(#341): wait for any followers or other internal processes to complete
	<-make(chan bool)
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

func createDatabaseOrDie(ctx context.Context) *sql.DB {
	db, err := sql.Open("mysql", *mysqlURI)
	if err != nil {
		klog.Exitf("Failed to connect to DB: %v", err)
	}
	db.SetConnMaxLifetime(*dbConnMaxLifetime)
	db.SetMaxOpenConns(*dbMaxOpenConns)
	db.SetMaxIdleConns(*dbMaxIdleConns)

	initDatabaseSchema(ctx)
	return db
}
