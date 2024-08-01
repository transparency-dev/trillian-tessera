package mysql_test

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"testing"

	"k8s.io/klog/v2"
)

var mysqlURI = flag.String("mysql_uri", "root:root@tcp(db:3306)/test_tessera", "Connection string for a MySQL database")

func TestMain(m *testing.M) {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()

	db, err := sql.Open("mysql", *mysqlURI)
	if err != nil {
		klog.Errorf("MySQL not available, skipping all MySQL storage tests")
		return
	}
	if err := db.PingContext(ctx); err != nil {
		klog.Errorf("MySQL not available, skipping all MySQL storage tests")
		return
	}
	klog.Info("Successfully connected to MySQL test database")

	os.Exit(m.Run())
}
