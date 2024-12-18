# Trillian Tessera

[![Go Report Card](https://goreportcard.com/badge/github.com/transparency-dev/trillian-tessera)](https://goreportcard.com/report/github.com/transparency-dev/trillian-tessera)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/transparency-dev/trillian-tessera/badge)](https://scorecard.dev/viewer/?uri=github.com/transparency-dev/trillian-tessera)
[![Benchmarks](https://img.shields.io/badge/Benchmarks-blue.svg)](https://transparency-dev.github.io/trillian-tessera/dev/bench/)
[![Slack Status](https://img.shields.io/badge/Slack-Chat-blue.svg)](https://transparency-dev.slack.com/)

Trillian Tessera is a Go library for building [tile-based transparency logs (tlogs)](https://c2sp.org/tlog-tiles).
It is the logical successor to the approach [Trillian v1][] takes in building and operating logs.

The implementation and its APIs bake-in
[current best-practices based on the lessons learned](https://transparency.dev/articles/tile-based-logs/)
over the past decade of building and operating transparency logs in production environments and at scale.

Tessera was introduced at the Transparency.Dev summit in October 2024.
Watch [Introducing Trillian Tessera](https://www.youtube.com/watch?v=9j_8FbQ9qSc) for all the details,
but here's a summary of the high level goals:

*   [tlog-tiles API][] and storage
*   Support for both cloud and on-premises infrastructure
    *   [GCP](./storage/gcp/)
    *   [AWS](./storage/aws/)
    *   [MySQL](./storage/mysql/)
    *   [POSIX](./storage/posix/)
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

The main non-goal is to support transparency logs using anything other than the [tlog-tiles API][].
While it is possible to deploy a custom personality in front of Tessera that adapts the tlog-tiles API
into any other API, this strategy will lose a lot of the read scaling that Tessera is designed for.

## Status

Tessera is under active development, with the [alpha](https://github.com/orgs/transparency-dev/projects/2/views/3) milestone coming soon. 
Users of GCP, MySQL, and POSIX are welcome to try the relevant [Getting Started](#getting-started) guide.

## Roadmap

Alpha expected by Q4 2024, and production ready in the first half of 2025.

### What’s happening to Trillian v1?

[Trillian v1][] is still in use in production environments by
multiple organisations in multiple ecosystems, and is likely to remain so for the mid-term. 

New ecosystems, or existing ecosystems looking to evolve, should strongly consider planning a
migration to Tessera and adopting the patterns it encourages.
Note that to achieve the full benefits of Tessera, logs must use the [tlog-tiles API][].

## Concepts

This section introduces concepts and terms that will be used throughout the user guide.

### Sequencing

When data is added to a log, it is first stored in memory for some period (this can be controlled via the [batching options](https://pkg.go.dev/github.com/transparency-dev/trillian-tessera#WithBatching)).
If the process dies in this state, the entry will be lost.

Once a batch of entries is processed by the sequencer, the new data will transition from a volatile state to one where it is durably assigned an index.
If the process dies in this state, the entry will be safe, though it will not be available through the read API of the log until the leaf has been [Integrated](#integration).
Once an index number has been issued to a leaf, no other data will ever be issued the same index number.
All index numbers are contiguous and start from 0.

> [!IMPORTANT]
> Within a batch, there is no guarantee about which order index numbers will be assigned.
> The only way to ensure that sequential calls to `Add` are given sequential indices is by blocking until a sequencing batch is completed.
> This can be achieved by configuring a batch size of 1, though this will make sequencing expensive!

### Integration

Integration is a background process that happens when the Tessera storage object has been created.
This process takes sequenced entries and merges them into the log.
Once this process has been completed, a new entry will:
 - Be available via the read API at the index that was returned from sequencing
 - Have Merkle tree hashes that commit to this data being included in the tree
 - Be committed to by the latest Checkpoint (and any Checkpoints issued after this point)

> [!IMPORTANT]
> There is currently no easy way to determine that integration has completed.
> This isn't an issue if the personality process is continually running.
> For personalities that require periods of downtime, [#341](https://github.com/transparency-dev/trillian-tessera/issues/341) tracks adding an API to allow for clean shutdown.

## Usage

### Getting Started

The best place to start is the [codelab](./cmd/conformance#codelab). 
This will walk you through setting up your first log, writing some entries to it via HTTP, and inspecting the contents.

Take a look at the example personalities in the `/cmd/` directory:
  - [posix](./cmd/conformance/posix/): example of operating a log backed by a local filesystem
    - This example runs an HTTP web server that takes arbitrary data and adds it to a file-based log.
  - [mysql](./cmd/conformance/mysql/): example of operating a log that uses MySQL
    - This example is easiest deployed via `docker compose`, which allows for easy setup and teardown.
  - [gcp](./cmd/conformance/gcp/): example of operating a log running in GCP.
    - This example can be deployed via terraform, see the [deployment instructions](./deployment/live/gcp/conformance#manual-deployment).
  - [aws](./cmd/conformance/aws/): example of operating a log running on AWS.
    - This example can be deployed via terraform, see the [deployment instructions](./deployment/live/aws/codelab#aws-codelab-deployment).
  - [posix-oneshot](./cmd/examples/posix-oneshot/): example of a command line tool to add entries to a log stored on the local filesystem
    - This example is not a long-lived process; running the command integrates entries into the log which lives only as files.

The `main.go` files for each of these example personalities try to strike a balance when demonstrating features of Tessera between simplicity, and demonstrating best practices.
Please raise issues against the repo, or chat to us in [Slack](#contact) if you have ideas for making the examples more accessible!

### Writing Personalities

#### Introduction

Tessera is a library written in Go.
It is designed to efficiently serve logs that allow read access via the [tlog-tiles API][].
The code you write that calls Tessera is referred to as a personality, because it tailors the generic library to your ecosystem.

Before starting to write your own personality, it is strongly recommended that you have familiarized yourself with the provided personalities referenced in [Getting Started](#getting-started).
When writing your Tessera personality, the biggest decision you need to make first is which of the native drivers to use:
 *   [GCP](./storage/gcp/)
 *   [AWS](./storage/aws/)
 *   [MySQL](./storage/mysql/)
 *   [POSIX](./storage/posix/)

The easiest drivers to operate and to scale are the cloud implementations: GCP and AWS.
These are the recommended choice for the majority of users running in production.

If you aren't using a cloud provider, then your options are MySQL and POSIX:
- POSIX is the simplest to get started with as it needs little in the way of extra infrastructure, and
  if you already serve static files as part of your business/project this could be a good fit.
- Alternatively, if you are used to operating user-facing applications backed by a RDBMS, then MySQL could
  be a natural fit.

To get a sense of the rough performance you can expect from the different backends, take a look at
[docs/performance.md](/docs/performance.md).

#### Setup

Once you've picked a storage driver, you can start writing your personality!
You'll need to import the Tessera library:
```shell
# This imports the library at main.
# This should be set to the latest release version to get a stable release.
go get github.com/transparency-dev/trillian-tessera@main
```

#### Constructing the Storage Object

Now you'll need to instantiate the storage object for the native driver you are using:
```go
import (
    "context"

    tessera "github.com/transparency-dev/trillian-tessera"
    "github.com/transparency-dev/trillian-tessera/storage/aws"
    "github.com/transparency-dev/trillian-tessera/storage/gcp"
    "github.com/transparency-dev/trillian-tessera/storage/mysql"
    "github.com/transparency-dev/trillian-tessera/storage/posix"
)

func main() {
    // Choose one!
    driver, err := aws.New(ctx, awsConfig)
    driver, err := gcp.New(ctx, gcpConfig)
    driver, err := mysql.New(ctx, db)
    driver, err := posix.New(ctx, dir, doCreate)

    storage := storage.NewAppender(driver)
}
```

See the documentation for each storage implementation to understand the parameters that each takes.
Each of these `New` calls are variadic functions, which is to say they take any number of trailing arguments.
The optional arguments that can be passed in allow Tessera behaviour to be tuned.
Take a look at the functions in the `trillian-tessera` root package named `With*`, e.g. [`WithBatching`](https://pkg.go.dev/github.com/transparency-dev/trillian-tessera#WithBatching) to see the available options are how they should be used.

The final part of configuring this storage object is to set up the mix-ins that you want to use.
Mix-ins are optional libraries you can use to provide common log behaviours without writing it yourself.
The currently supported mix-ins are:
 * Deduplication 
   * [In-memory](https://pkg.go.dev/github.com/transparency-dev/trillian-tessera#InMemoryDedupe) (cheap, but very limited deduplication behaviour)
   * Persistent (expensive, but can strongly ensure the log contains no duplicates)
     * TODO(mhutchinson): link to these implementations when they are written
 * [Synchronous Integration](https://pkg.go.dev/github.com/transparency-dev/trillian-tessera#IntegrationAwaiter)

See [Mix-ins](#mix-ins) after reading the rest of this section for more details.

#### Writing to the Log

Now you should have a storage object configured for your environment, and the correct mix-ins set up.
Now the fun part - writing to the log!

```go
func main() {
    storage, err := ...
    idx, err := storage.Add(ctx, tessera.NewEntry(data))()
```

Whichever storage option you use, writing to the log follows the same pattern: simply call `Add` with a new entry created with the data to be added as a leaf in the log.
This method returns a _future_ of the form `func() (idx uint64, err error)`.
When called, this future function will block until the data passed into `Add` has been sequenced and an index number is assigned (or until failure, in which case an error is returned).
Once this index has been returned, the new data is sequenced, but not necessarily integrated into the log.

As discussed above in [Integration](#integration), sequenced entries will be _asynchronously_ integrated into the log and be made available via the read API.
Some personalities may need to block until this has been performed, e.g. because they will provide the requester with an inclusion proof, which requires integration.
Such personalities are recommended to use [Synchronous Integration](#synchronous-integration) to perform this blocking.

#### Reading from the Log

Data that has been written to the log needs to be made available for clients and verifiers.
Tessera makes the log readable via the [tlog-tiles API][].
In the case of AWS and GCP, the data to be served is written to object storage and served directly by the cloud provider.
The log operator only needs to ensure that these object storage instances are publicly readable, and set up a URL to point to them.

In the case of MySQL and POSIX, the log operator will need to take more steps to make the data available.
POSIX writes out the files exactly as per the API spec, so the log operator can serve these via an HTTP File Server.

MySQL is the odd implementation in that it requires personality code to handle read traffic.
See the example personalities written for MySQL to see how this Go web server should be configured.

## Mix-ins

### Deduplication

Deduplicating entries means that the log will only store each unique entry once.
Deduplication is recommended for logs that take public submissions, such as CT.
While checking for duplicates is somewhat expensive, it protects the log from a type of DoS attack where users can feed the log back into itself, causing it to grow out of control.
It also protects against clients that may send the same request for logging multiple times (perhaps as a programming error, or by design for reliability).

Logs that do not allow public submissions directly to the log may want to operate without deduplication, instead relying on the personality to never generate duplicates.
This can allow for significantly cheaper operation and faster write throughput.

### Synchronous Integration

Synchronous Integration is provided by [`tessera.IntegrationAwaiter`](https://pkg.go.dev/github.com/transparency-dev/trillian-tessera#IntegrationAwaiter).
This allows personality calls to `Add` to block until the new leaf is integrated into the tree.

## Contributing

See [CONTRIBUTING.md](/CONTRIBUTING.md) for details.

## License

This repo is licensed under the Apache 2.0 license, see [LICENSE](/LICENSE) for details

## Contact

- Slack: https://transparency-dev.slack.com/ ([invitation](https://join.slack.com/t/transparency-dev/shared_invite/zt-27pkqo21d-okUFhur7YZ0rFoJVIOPznQ))
- Mailing list: https://groups.google.com/forum/#!forum/trillian-transparency

## Acknowledgements

Tessera builds upon the hard work, experience, and lessons from many _many_ folks involved in
transparency ecosystems over the years.

[tlog-tiles API]: https://c2sp.org/tlog-tiles
[Static CT API]: https://c2sp.org/static-ct-api
[Trillian v1]: https://github.com/google/trillian
