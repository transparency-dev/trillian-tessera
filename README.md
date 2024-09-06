# Trillian Tessera

[![Slack Status](https://img.shields.io/badge/Slack-Chat-blue.svg)](https://transparency-dev.slack.com/)

Trillian Tessera is a Go library for building
[tile-based transparency logs (tlogs)](https://github.com/C2SP/C2SP/blob/main/tlog-tiles.md).
It is the logical successor to the approach Trillian v1 takes in building and operating logs.

The implementation and its APIs bake-in
[current best-practices based on the lessons learned](https://transparency.dev/articles/tile-based-logs/)
over the past decade of building and operating transparency logs in production environments and at scale.

Tessera goals:

*   [tlog-tiles API][] and storage
*   Support for both cloud and on-premises infrastructure
    *   GCP and AWS support will be provided initially
    *   Cloud agnostic MySQL and POSIX filesystem support
*   Make it easy to build and deploy new transparency logs on supported infrastructure
    *   Library instead of microservice architecture
    *   No additional services to manage
    *   Lower TCO for operators compared with Trillian v1
*   Fast sequencing and integration of entries
*   Optional functionality which can be enabled for those ecosystems/logs which need it (only pay the cost for what you need):
    *   "Best-effort" de-duplication of entries
    *   Synchronous integration
*   Broadly similar write-throughput and write-availability, and potentially _far_ higher read-throughput
    and read-availability compared to Trillian v1 (dependent on underlying infrastructure)
*   Enable building of arbitrary log personalities, including support for the peculiarities of a
    [Static CT API][] compliant log.

### Status

Tessera is currently under active development, and is not yet ready for general use. However, early
feedback is welcome.

### Roadmap

Alpha expected by Q4 2024, and production ready in the first half of 2025.

#### What’s happening to Trillian v1?

[Trillian v1](https://github.com/google/trillian) is still in use in production environments by
multiple organisations in multiple ecosystems, and is likely to remain so for the mid-term. 

New ecosystems, or existing ecosystems looking to evolve, should strongly consider planning a
migration to Tessera and adopting the patterns it encourages. 

### Getting started

Take a look at the example personalities in the `/cmd/` directory:
  - [posix](./cmd/conformance/posix/): example of operating a log backed by a local filesystem
    - This example runs an HTTP web server that takes arbitrary data and adds it to a file-based log.
  - [mysql](./cmd/conformance/mysql/): example of operating a log that uses MySQL
    - This example is easiest deployed via `docker compose`, which allows for easy setup and teardown.
  - [gcp](./cmd/conformance/gcp/): example of operating a log running in GCP
    - This example can be deployed via terraform (see the [deployment](./deployment/) directory).
  - [posix-oneshot](./cmd/examples/posix-oneshot/): exmaple of a command line tool to add entries to a log stored on the local filesystem
    - This example is not a long-lived process; running the command integrates entries into the log which lives only as files.

The `main.go` files for each of these example personalities try to strike a balance when demonstrating features of Tessera between simplicity, and demonstrating best practices.
Please raise issues against the repo, or chat to us in [Slack](#contact) if you have ideas for making the examples more accessible!

### Contributing

See [CONTRIBUTING.md](/CONTRIBUTING.md) for details.

### License

This repo is licensed under the Apache 2.0 license, see [LICENSE](/LICENSE) for details

### Contact

- Slack: https://transparency-dev.slack.com/ ([invitation](https://join.slack.com/t/transparency-dev/shared_invite/zt-27pkqo21d-okUFhur7YZ0rFoJVIOPznQ))
- Mailing list: https://groups.google.com/forum/#!forum/trillian-transparency

### Acknowledgements

Tessera builds upon the hard work, experience, and lessons from many _many_ folks involved in
transparency ecosystems over the years.

[tlog-tiles API]: https://c2sp.org/tlog-tiles
[Static CT API]: https://c2sp.org/static-ct-api
