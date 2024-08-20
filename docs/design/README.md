# Design docs

This directory contains design documentation for Tessera.

It's probably wise to start with the [philosophy](philosophy.md]) doc first in order to provide
the context around the approach and design trade-offs made herein.

## Storage

Tessera supports multiple backend storage implementations, each of these has an associated
"one-pager" design doc:

* [GCP](storage_gcp.md)
* [MySQL](storage_mysql.md)
* [POSIX filesystem](storage_posix.md)

