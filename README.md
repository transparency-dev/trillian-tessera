# Trillian Tessera

[![Slack Status](https://img.shields.io/badge/Slack-Chat-blue.svg)](https://transparency-dev.slack.com/)

Trillian Tessera is a Go library for building
[tile-based transparency logs (tlogs)](https://github.com/C2SP/C2SP/blob/main/tlog-tiles.md).
It is the logical successor to the approach Trillian v1 takes in building and operating logs.

The implementation and its APIs bake-in
[current best-practices based on the lessons learned](https://transparency.dev/articles/tile-based-logs/)
over the past decade of building and operating transparency logs in production environments and at scale.

Tessera goals:

*   [Tiles-native API](https://github.com/C2SP/C2SP/blob/main/tlog-tiles.md) and storage
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
    [Static CT API](https://github.com/C2SP/C2SP/blob/main/static-ct-api.md) compliant log.

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

#### Usage

```go

import (
    tessera "github.com/transparency-dev/trillian-tessera"
)

// TODO...

```

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
