package mysql_test

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"testing"

	"k8s.io/klog/v2"
)

var (
	mysqlURI            = flag.String("mysql_uri", "root:root@tcp(localhost:3306)/test_tessera", "Connection string for a MySQL database")
	isMySQLTestOptional = flag.Bool("is_mysql_test_optional", true, "Boolean value to control whether the MySQL test is optional")
)

func TestMain(m *testing.M) {
	klog.InitFlags(nil)
	flag.Parse()
	ctx := context.Background()

	db, err := sql.Open("mysql", *mysqlURI)
	if err != nil {
		if *isMySQLTestOptional {
			klog.Warning("MySQL not available, skipping all MySQL storage tests")
			return
		}
		klog.Fatalf("Failed to open MySQL test db: %v", err)
	}
	if err := db.PingContext(ctx); err != nil {
		if *isMySQLTestOptional {
			klog.Warning("MySQL not available, skipping all MySQL storage tests")
			return
		}
		klog.Fatalf("Failed to ping MySQL test db: %v", err)
	}
	klog.Info("Successfully connected to MySQL test database")

	os.Exit(m.Run())
}
