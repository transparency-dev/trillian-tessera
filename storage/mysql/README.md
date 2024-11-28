# MySQL Storage for Trillian Tessera

This directory contains the implementation of a storage backend for Trillian Tessera using MySQL. This allows Tessera to leverage MySQL as its underlying database for storing checkpoint, entry hashes and data in tiles format.

## Design

See [MySQL storage design documentation](/docs/design/mysql_storage.md).

### Requirements

- A running MySQL server instance. This storage implementation has been tested against MySQL 8.4.

## Usage

### Constructing the Storage Object

```go
import (
    "context"

    tessera "github.com/transparency-dev/trillian-tessera"
    "github.com/transparency-dev/trillian-tessera/storage/mysql"
)

func main() {
    mysqlURI := "user:password@tcp(db:3306)/tessera"
    db, err := sql.Open("mysql", mysqlURI)
    if err != nil {
        klog.Exitf("Failed to connect to DB: %v", err)
    }

    storage, err := mysql.New(ctx, db)
    if err != nil {
        klog.Exitf("Failed to create new MySQL storage: %v", err)
    }
}
```

### Example personality

See [MySQL conformance example](/cmd/conformance/mysql/).

## Future Work

- [Separate sequencing and integration](https://github.com/transparency-dev/trillian-tessera/pull/282)
