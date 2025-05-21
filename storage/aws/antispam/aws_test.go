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

package aws_test

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/transparency-dev/tessera"
	"github.com/transparency-dev/tessera/api"
	"github.com/transparency-dev/tessera/core"
	aws "github.com/transparency-dev/tessera/storage/aws/antispam"
	"github.com/transparency-dev/tessera/testonly"
	"k8s.io/klog/v2"
)

var (
	mySQLURI            = flag.String("mysql_uri", "root:root@tcp(localhost:3306)/test_tessera", "Connection string for a MySQL database")
	isMySQLTestOptional = flag.Bool("is_mysql_test_optional", true, "Boolean value to control whether the MySQL test is optional")
)

// TestMain inits flags and runs tests.
func TestMain(m *testing.M) {
	klog.InitFlags(nil)
	// m.Run() will parse flags
	os.Exit(m.Run())
}

func TestAntispam(t *testing.T) {
	ctx := t.Context()
	if canSkipMySQLTest(t, ctx) {
		klog.Warningf("MySQL not available, skipping %s", t.Name())
		t.Skip("MySQL not available, skipping test")
	}
	mustDropTables(t, ctx)
	antispam, err := aws.NewAntispam(ctx, *mySQLURI, aws.AntispamOpts{})
	if err != nil {
		t.Fatal(err)
	}
	fl, shutdown := testonly.NewTestLog(t, tessera.NewAppendOptions().WithCheckpointInterval(time.Second))
	defer func() {
		if err := shutdown(t.Context()); err != nil {
			t.Logf("shutdown: %v", err)
		}
	}()
	addFn := antispam.Decorator()(fl.Appender.Add)
	follower := antispam.Follower(testBundleHasher)
	go follower.Follow(ctx, fl.LogReader)

	pos, err := follower.EntriesProcessed(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if pos != 0 {
		t.Error("expected initial position to be 0")
	}

	var idx1 core.Index
	idxf1 := addFn(ctx, core.NewEntry([]byte("one")))
	idxf2 := addFn(ctx, core.NewEntry([]byte("two")))
	if idx1, err = idxf1(); err != nil {
		t.Fatal(err)
	}
	if _, err := idxf2(); err != nil {
		t.Fatal(err)
	}
	for {
		if idx, err := follower.EntriesProcessed(ctx); err != nil {
			t.Fatal(err)
		} else if idx == 2 {
			break
		}
	}
	dupIdx, err := addFn(ctx, core.NewEntry([]byte("one")))()
	if err != nil {
		t.Error(err)
	}
	if !dupIdx.IsDup {
		t.Error("expected dupe but it wasn't marked as such")
	}
	if dupIdx.Index != idx1.Index {
		t.Errorf("expected idx %d but got %d", idx1.Index, dupIdx.Index)
	}
}

// canSkipMySQLTest checks if the test MySQL db is available and, if not, if the test can be skipped.
//
// Use this method before every MySQL test, and if it returns true, skip the test.
//
// If is_mysql_test_optional is set to true and MySQL database cannot be opened or pinged,
// the test will fail immediately. Otherwise, the test will be skipped if the test is optional
// and the database is not available.
func canSkipMySQLTest(t *testing.T, ctx context.Context) bool {
	t.Helper()

	db, err := sql.Open("mysql", *mySQLURI)
	if err != nil {
		if *isMySQLTestOptional {
			return true
		}
		t.Fatalf("failed to open MySQL test db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatalf("failed to close MySQL database: %v", err)
		}
	}()
	if err := db.PingContext(ctx); err != nil {
		if *isMySQLTestOptional {
			return true
		}
		t.Fatalf("failed to ping MySQL test db: %v", err)
	}
	return false
}

// mustDropTables drops the `Seq`, `SeqCoord` and `IntCoord` tables.
// Call this function before every MySQL test.
func mustDropTables(t *testing.T, ctx context.Context) {
	t.Helper()

	db, err := sql.Open("mysql", *mySQLURI)
	if err != nil {
		t.Fatalf("failed to connect to db: %v", *mySQLURI)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatalf("failed to close db: %v", err)
		}
	}()

	if _, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS `AntispamMeta`, `AntispamIDSeq`, `AntispamFollowCoord`"); err != nil {
		t.Fatalf("failed to drop all tables: %v", err)
	}
}

func testIDHash(d []byte) []byte {
	r := sha256.Sum256(d)
	return r[:]
}

func testBundleHasher(b []byte) ([][]byte, error) {
	bun := &api.EntryBundle{}
	err := bun.UnmarshalText(b)
	if err != nil {
		return nil, err
	}
	r := make([][]byte, len(bun.Entries))
	for i, e := range bun.Entries {
		r[i] = testIDHash(e)
	}
	return r, err
}
