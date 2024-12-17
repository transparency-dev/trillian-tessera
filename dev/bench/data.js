window.BENCHMARK_DATA = {
  "lastUpdate": 1734451086283,
  "repoUrl": "https://github.com/transparency-dev/trillian-tessera",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "mhutchinson@gmail.com",
            "name": "Martin Hutchinson",
            "username": "mhutchinson"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "cfc5fe29bd7faaa2f04a3cd7de2acb466ad104ec",
          "message": "Set up benchmark action (#405)\n\nTowards #338.",
          "timestamp": "2024-12-09T16:55:02Z",
          "tree_id": "040711c895795a1d879d614076ca48a4c7acfb29",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/cfc5fe29bd7faaa2f04a3cd7de2acb466ad104ec"
        },
        "date": 1733764025361,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4115825,
            "unit": "ns/op\t  701840 B/op\t   19684 allocs/op",
            "extra": "291 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4115825,
            "unit": "ns/op",
            "extra": "291 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 701840,
            "unit": "B/op",
            "extra": "291 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19684,
            "unit": "allocs/op",
            "extra": "291 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7759,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "264615 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7759,
            "unit": "ns/op",
            "extra": "264615 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "264615 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "264615 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.9,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10444465 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.9,
            "unit": "ns/op",
            "extra": "10444465 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10444465 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10444465 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 349612,
            "unit": "ns/op\t  298554 B/op\t    3439 allocs/op",
            "extra": "3966 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 349612,
            "unit": "ns/op",
            "extra": "3966 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298554,
            "unit": "B/op",
            "extra": "3966 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3439,
            "unit": "allocs/op",
            "extra": "3966 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "rogerng@google.com",
            "name": "Roger Ng",
            "username": "roger2hk"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "8c25bbffbeb1c9ad2dd9fc231a2bbe8b9e8e1052",
          "message": "Remove AWS `dsn` from error log to avoid password leak (#402)\n\n* Remove AWS `dsn` from error log to avoid password leak\n\n* Remove AWS `dsn` from error log to avoid password leak",
          "timestamp": "2024-12-09T17:08:55Z",
          "tree_id": "8adf1a376206d61554870aad6f247f11d9533e9f",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/8c25bbffbeb1c9ad2dd9fc231a2bbe8b9e8e1052"
        },
        "date": 1733764175580,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1162974,
            "unit": "ns/op\t  689551 B/op\t   19561 allocs/op",
            "extra": "1027 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1162974,
            "unit": "ns/op",
            "extra": "1027 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689551,
            "unit": "B/op",
            "extra": "1027 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19561,
            "unit": "allocs/op",
            "extra": "1027 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1970,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "610368 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1970,
            "unit": "ns/op",
            "extra": "610368 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "610368 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "610368 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 125.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10346365 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 125.7,
            "unit": "ns/op",
            "extra": "10346365 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10346365 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10346365 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 370653,
            "unit": "ns/op\t  297952 B/op\t    3435 allocs/op",
            "extra": "4084 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 370653,
            "unit": "ns/op",
            "extra": "4084 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 297952,
            "unit": "B/op",
            "extra": "4084 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3435,
            "unit": "allocs/op",
            "extra": "4084 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "rogerng@google.com",
            "name": "Roger Ng",
            "username": "roger2hk"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "20c25b17f760cf543d59e7fb20ef2a01e8eb21ef",
          "message": "Fix `go-routine` typo in AWS (#407)",
          "timestamp": "2024-12-09T18:04:23Z",
          "tree_id": "2ce3ef3d659feb220e9114161b69237832e1dbb0",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/20c25b17f760cf543d59e7fb20ef2a01e8eb21ef"
        },
        "date": 1733767501668,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1143573,
            "unit": "ns/op\t  689286 B/op\t   19558 allocs/op",
            "extra": "1050 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1143573,
            "unit": "ns/op",
            "extra": "1050 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689286,
            "unit": "B/op",
            "extra": "1050 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19558,
            "unit": "allocs/op",
            "extra": "1050 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1954,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "566126 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1954,
            "unit": "ns/op",
            "extra": "566126 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "566126 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "566126 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 115.6,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10428170 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 115.6,
            "unit": "ns/op",
            "extra": "10428170 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10428170 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10428170 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 366892,
            "unit": "ns/op\t  298006 B/op\t    3435 allocs/op",
            "extra": "4098 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 366892,
            "unit": "ns/op",
            "extra": "4098 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298006,
            "unit": "B/op",
            "extra": "4098 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3435,
            "unit": "allocs/op",
            "extra": "4098 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "rogerng@google.com",
            "name": "Roger Ng",
            "username": "roger2hk"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "225c6e010280e115723055e8149bb55882f1691e",
          "message": "Update log message and comment in POSIX (#408)",
          "timestamp": "2024-12-09T18:12:26Z",
          "tree_id": "d0d30d45ddec2880934ca76afa0bab6102c98720",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/225c6e010280e115723055e8149bb55882f1691e"
        },
        "date": 1733767978684,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1159144,
            "unit": "ns/op\t  689120 B/op\t   19556 allocs/op",
            "extra": "1078 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1159144,
            "unit": "ns/op",
            "extra": "1078 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689120,
            "unit": "B/op",
            "extra": "1078 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19556,
            "unit": "allocs/op",
            "extra": "1078 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1983,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "558681 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1983,
            "unit": "ns/op",
            "extra": "558681 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "558681 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "558681 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10406815 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116,
            "unit": "ns/op",
            "extra": "10406815 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10406815 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10406815 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 360351,
            "unit": "ns/op\t  298101 B/op\t    3436 allocs/op",
            "extra": "4143 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 360351,
            "unit": "ns/op",
            "extra": "4143 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298101,
            "unit": "B/op",
            "extra": "4143 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3436,
            "unit": "allocs/op",
            "extra": "4143 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "al@google.com",
            "name": "Al Cutter",
            "username": "AlCutter"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6039e35f361f3fe23d5e93e63e711b413d082084",
          "message": "Fix POSIX full-bundle handling (#406)",
          "timestamp": "2024-12-09T18:29:01Z",
          "tree_id": "2d4befd3e4193eccb535ccccaea37430f05232cc",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/6039e35f361f3fe23d5e93e63e711b413d082084"
        },
        "date": 1733768979115,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1310245,
            "unit": "ns/op\t  689646 B/op\t   19561 allocs/op",
            "extra": "952 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1310245,
            "unit": "ns/op",
            "extra": "952 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689646,
            "unit": "B/op",
            "extra": "952 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19561,
            "unit": "allocs/op",
            "extra": "952 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 2162,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "566762 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 2162,
            "unit": "ns/op",
            "extra": "566762 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "566762 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "566762 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 151.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "8347254 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 151.7,
            "unit": "ns/op",
            "extra": "8347254 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "8347254 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8347254 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 385598,
            "unit": "ns/op\t  297810 B/op\t    3433 allocs/op",
            "extra": "3544 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 385598,
            "unit": "ns/op",
            "extra": "3544 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 297810,
            "unit": "B/op",
            "extra": "3544 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3433,
            "unit": "allocs/op",
            "extra": "3544 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "e9b7fd2bcfad8e5881b6cd66eed8abaafbfd25c6",
          "message": "Bump the all-go-deps group with 7 updates (#411)\n\nBumps the all-go-deps group with 7 updates:\n\n| Package | From | To |\n| --- | --- | --- |\n| [cloud.google.com/go/storage](https://github.com/googleapis/google-cloud-go) | `1.47.0` | `1.48.0` |\n| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.70.0` | `1.71.0` |\n| [google.golang.org/api](https://github.com/googleapis/google-api-go-client) | `0.209.0` | `0.210.0` |\n| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.68.0` | `1.68.1` |\n| [golang.org/x/crypto](https://github.com/golang/crypto) | `0.29.0` | `0.30.0` |\n| [golang.org/x/net](https://github.com/golang/net) | `0.31.0` | `0.32.0` |\n| [golang.org/x/sync](https://github.com/golang/sync) | `0.9.0` | `0.10.0` |\n\n\nUpdates `cloud.google.com/go/storage` from 1.47.0 to 1.48.0\n- [Release notes](https://github.com/googleapis/google-cloud-go/releases)\n- [Changelog](https://github.com/googleapis/google-cloud-go/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-cloud-go/compare/spanner/v1.47.0...spanner/v1.48.0)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.70.0 to 1.71.0\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.70.0...service/s3/v1.71.0)\n\nUpdates `google.golang.org/api` from 0.209.0 to 0.210.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.209.0...v0.210.0)\n\nUpdates `google.golang.org/grpc` from 1.68.0 to 1.68.1\n- [Release notes](https://github.com/grpc/grpc-go/releases)\n- [Commits](https://github.com/grpc/grpc-go/compare/v1.68.0...v1.68.1)\n\nUpdates `golang.org/x/crypto` from 0.29.0 to 0.30.0\n- [Commits](https://github.com/golang/crypto/compare/v0.29.0...v0.30.0)\n\nUpdates `golang.org/x/net` from 0.31.0 to 0.32.0\n- [Commits](https://github.com/golang/net/compare/v0.31.0...v0.32.0)\n\nUpdates `golang.org/x/sync` from 0.9.0 to 0.10.0\n- [Commits](https://github.com/golang/sync/compare/v0.9.0...v0.10.0)\n\n---\nupdated-dependencies:\n- dependency-name: cloud.google.com/go/storage\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/grpc\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/crypto\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/net\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/sync\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-12-09T21:18:01Z",
          "tree_id": "e374dde65348a1ee0173039f6d28959b4915bcca",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/e9b7fd2bcfad8e5881b6cd66eed8abaafbfd25c6"
        },
        "date": 1733779155323,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4313741,
            "unit": "ns/op\t  705106 B/op\t   19717 allocs/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4313741,
            "unit": "ns/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 705106,
            "unit": "B/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19717,
            "unit": "allocs/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6501,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "171648 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6501,
            "unit": "ns/op",
            "extra": "171648 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "171648 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "171648 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 118.6,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10508440 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 118.6,
            "unit": "ns/op",
            "extra": "10508440 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10508440 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10508440 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 348864,
            "unit": "ns/op\t  298001 B/op\t    3435 allocs/op",
            "extra": "4060 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 348864,
            "unit": "ns/op",
            "extra": "4060 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298001,
            "unit": "B/op",
            "extra": "4060 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3435,
            "unit": "allocs/op",
            "extra": "4060 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "e9b7fd2bcfad8e5881b6cd66eed8abaafbfd25c6",
          "message": "Bump the all-go-deps group with 7 updates (#411)\n\nBumps the all-go-deps group with 7 updates:\n\n| Package | From | To |\n| --- | --- | --- |\n| [cloud.google.com/go/storage](https://github.com/googleapis/google-cloud-go) | `1.47.0` | `1.48.0` |\n| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.70.0` | `1.71.0` |\n| [google.golang.org/api](https://github.com/googleapis/google-api-go-client) | `0.209.0` | `0.210.0` |\n| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.68.0` | `1.68.1` |\n| [golang.org/x/crypto](https://github.com/golang/crypto) | `0.29.0` | `0.30.0` |\n| [golang.org/x/net](https://github.com/golang/net) | `0.31.0` | `0.32.0` |\n| [golang.org/x/sync](https://github.com/golang/sync) | `0.9.0` | `0.10.0` |\n\n\nUpdates `cloud.google.com/go/storage` from 1.47.0 to 1.48.0\n- [Release notes](https://github.com/googleapis/google-cloud-go/releases)\n- [Changelog](https://github.com/googleapis/google-cloud-go/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-cloud-go/compare/spanner/v1.47.0...spanner/v1.48.0)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.70.0 to 1.71.0\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.70.0...service/s3/v1.71.0)\n\nUpdates `google.golang.org/api` from 0.209.0 to 0.210.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.209.0...v0.210.0)\n\nUpdates `google.golang.org/grpc` from 1.68.0 to 1.68.1\n- [Release notes](https://github.com/grpc/grpc-go/releases)\n- [Commits](https://github.com/grpc/grpc-go/compare/v1.68.0...v1.68.1)\n\nUpdates `golang.org/x/crypto` from 0.29.0 to 0.30.0\n- [Commits](https://github.com/golang/crypto/compare/v0.29.0...v0.30.0)\n\nUpdates `golang.org/x/net` from 0.31.0 to 0.32.0\n- [Commits](https://github.com/golang/net/compare/v0.31.0...v0.32.0)\n\nUpdates `golang.org/x/sync` from 0.9.0 to 0.10.0\n- [Commits](https://github.com/golang/sync/compare/v0.9.0...v0.10.0)\n\n---\nupdated-dependencies:\n- dependency-name: cloud.google.com/go/storage\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/grpc\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/crypto\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/net\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/sync\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-12-09T21:18:01Z",
          "tree_id": "e374dde65348a1ee0173039f6d28959b4915bcca",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/e9b7fd2bcfad8e5881b6cd66eed8abaafbfd25c6"
        },
        "date": 1733779430316,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4367262,
            "unit": "ns/op\t  701797 B/op\t   19686 allocs/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4367262,
            "unit": "ns/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 701797,
            "unit": "B/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19686,
            "unit": "allocs/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7338,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "229556 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7338,
            "unit": "ns/op",
            "extra": "229556 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "229556 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "229556 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "9882890 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.7,
            "unit": "ns/op",
            "extra": "9882890 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "9882890 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9882890 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 358261,
            "unit": "ns/op\t  298920 B/op\t    3442 allocs/op",
            "extra": "3930 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 358261,
            "unit": "ns/op",
            "extra": "3930 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298920,
            "unit": "B/op",
            "extra": "3930 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3442,
            "unit": "allocs/op",
            "extra": "3930 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "a1bd4df67d5550ec04c24c670613b277713834c3",
          "message": "Bump the all-gha-deps group with 2 updates (#410)\n\nBumps the all-gha-deps group with 2 updates: [actions/setup-go](https://github.com/actions/setup-go) and [github/codeql-action](https://github.com/github/codeql-action).\n\n\nUpdates `actions/setup-go` from 4 to 5\n- [Release notes](https://github.com/actions/setup-go/releases)\n- [Commits](https://github.com/actions/setup-go/compare/v4...v5)\n\nUpdates `github/codeql-action` from 3.27.5 to 3.27.6\n- [Release notes](https://github.com/github/codeql-action/releases)\n- [Changelog](https://github.com/github/codeql-action/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/github/codeql-action/compare/f09c1c0a94de965c15400f5634aa42fac8fb8f88...aa578102511db1f4524ed59b8cc2bae4f6e88195)\n\n---\nupdated-dependencies:\n- dependency-name: actions/setup-go\n  dependency-type: direct:production\n  update-type: version-update:semver-major\n  dependency-group: all-gha-deps\n- dependency-name: github/codeql-action\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-12-09T21:23:22Z",
          "tree_id": "10b1905923db9d9b7738053cb7dbb81e52faa91a",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/a1bd4df67d5550ec04c24c670613b277713834c3"
        },
        "date": 1733779445728,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4870474,
            "unit": "ns/op\t  707757 B/op\t   19746 allocs/op",
            "extra": "280 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4870474,
            "unit": "ns/op",
            "extra": "280 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 707757,
            "unit": "B/op",
            "extra": "280 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19746,
            "unit": "allocs/op",
            "extra": "280 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 5699,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "234171 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 5699,
            "unit": "ns/op",
            "extra": "234171 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "234171 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "234171 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.3,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10469600 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.3,
            "unit": "ns/op",
            "extra": "10469600 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10469600 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10469600 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 347108,
            "unit": "ns/op\t  297977 B/op\t    3435 allocs/op",
            "extra": "4110 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 347108,
            "unit": "ns/op",
            "extra": "4110 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 297977,
            "unit": "B/op",
            "extra": "4110 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3435,
            "unit": "allocs/op",
            "extra": "4110 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "rogerng@google.com",
            "name": "Roger Ng",
            "username": "roger2hk"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "246624ff212553244c80872a1ccd7b82140d591d",
          "message": "Replace `256` with `layout.TileWidth`/`layout.EntryBundleWidth` (#409)\n\n* Replace `256` with `layout.TileWidth`/`layout.EntryBundleWidth`\n\n* Revert changes related to batching and dedup",
          "timestamp": "2024-12-10T12:02:26Z",
          "tree_id": "c9596d4165b3af2dcb4592ebb2ccaa1e0fc493ed",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/246624ff212553244c80872a1ccd7b82140d591d"
        },
        "date": 1733832183178,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4948944,
            "unit": "ns/op\t  705926 B/op\t   19727 allocs/op",
            "extra": "285 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4948944,
            "unit": "ns/op",
            "extra": "285 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 705926,
            "unit": "B/op",
            "extra": "285 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19727,
            "unit": "allocs/op",
            "extra": "285 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 5477,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "247970 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 5477,
            "unit": "ns/op",
            "extra": "247970 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "247970 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "247970 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 114.5,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10466448 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 114.5,
            "unit": "ns/op",
            "extra": "10466448 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10466448 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10466448 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 350053,
            "unit": "ns/op\t  298190 B/op\t    3437 allocs/op",
            "extra": "4159 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 350053,
            "unit": "ns/op",
            "extra": "4159 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298190,
            "unit": "B/op",
            "extra": "4159 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3437,
            "unit": "allocs/op",
            "extra": "4159 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "mhutchinson@gmail.com",
            "name": "Martin Hutchinson",
            "username": "mhutchinson"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "03827bf4c8bcb75f8a57889ede48b1e62959949d",
          "message": "[GitHub Workflows] Give names to steps (#412)\n\nIt's a lot easier to read these steps when they all have names. This should not affect operation.\r\n\r\nUse hash pinning for actions",
          "timestamp": "2024-12-10T13:18:47Z",
          "tree_id": "ff202910d3309a7dbe5ef5ee9ae1b535891dfcbf",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/03827bf4c8bcb75f8a57889ede48b1e62959949d"
        },
        "date": 1733836764507,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 5189277,
            "unit": "ns/op\t  706738 B/op\t   19737 allocs/op",
            "extra": "290 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 5189277,
            "unit": "ns/op",
            "extra": "290 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 706738,
            "unit": "B/op",
            "extra": "290 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19737,
            "unit": "allocs/op",
            "extra": "290 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6038,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "231562 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6038,
            "unit": "ns/op",
            "extra": "231562 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "231562 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "231562 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 115.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10372189 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 115.7,
            "unit": "ns/op",
            "extra": "10372189 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10372189 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10372189 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 365958,
            "unit": "ns/op\t  298177 B/op\t    3437 allocs/op",
            "extra": "3883 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 365958,
            "unit": "ns/op",
            "extra": "3883 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298177,
            "unit": "B/op",
            "extra": "3883 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3437,
            "unit": "allocs/op",
            "extra": "3883 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "mhutchinson@gmail.com",
            "name": "Martin Hutchinson",
            "username": "mhutchinson"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6faae3723174df4ec9b2d71177c9ba0076f1b426",
          "message": "[Benchmarks] Comment on PRs with benchmark scores (#413)\n\nThis should always leave a comment with a summary of benchmark results. Furthermore, if the alert threshold is triggered then it will leave another comment making it clear that this has happened.\r\n\r\nI fully expect that we'll want to disable at least one of these, and possibly tune the alert threshold. But by starting with too much, it will be fast to see and back this off to a useful level.",
          "timestamp": "2024-12-10T16:55:58Z",
          "tree_id": "d1c6e4251056c3971042c4863e8386c6820def45",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/6faae3723174df4ec9b2d71177c9ba0076f1b426"
        },
        "date": 1733849797275,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4607969,
            "unit": "ns/op\t  702297 B/op\t   19691 allocs/op",
            "extra": "282 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4607969,
            "unit": "ns/op",
            "extra": "282 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702297,
            "unit": "B/op",
            "extra": "282 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19691,
            "unit": "allocs/op",
            "extra": "282 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 5658,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "192346 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 5658,
            "unit": "ns/op",
            "extra": "192346 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "192346 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "192346 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 115.3,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10407111 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 115.3,
            "unit": "ns/op",
            "extra": "10407111 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10407111 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10407111 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 350775,
            "unit": "ns/op\t  297709 B/op\t    3432 allocs/op",
            "extra": "3795 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 350775,
            "unit": "ns/op",
            "extra": "3795 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 297709,
            "unit": "B/op",
            "extra": "3795 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3432,
            "unit": "allocs/op",
            "extra": "3795 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "lapo@lapo.it",
            "name": "Lapo Luchini",
            "username": "lapo-luchini"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "e7adf39b88cf311749ccb4be0e2fb2c01e561b1f",
          "message": "Add cache headers to POSIX version too (#416)",
          "timestamp": "2024-12-12T15:52:08Z",
          "tree_id": "49539a88ae56b6ba6b2fd61d00b0fd292d367871",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/e7adf39b88cf311749ccb4be0e2fb2c01e561b1f"
        },
        "date": 1734018767475,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 5071984,
            "unit": "ns/op\t  701736 B/op\t   19684 allocs/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 5071984,
            "unit": "ns/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 701736,
            "unit": "B/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19684,
            "unit": "allocs/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6474,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "187246 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6474,
            "unit": "ns/op",
            "extra": "187246 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "187246 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "187246 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.3,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10356183 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.3,
            "unit": "ns/op",
            "extra": "10356183 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10356183 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10356183 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 378060,
            "unit": "ns/op\t  297932 B/op\t    3435 allocs/op",
            "extra": "4069 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 378060,
            "unit": "ns/op",
            "extra": "4069 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 297932,
            "unit": "B/op",
            "extra": "4069 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3435,
            "unit": "allocs/op",
            "extra": "4069 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "bobcallaway@users.noreply.github.com",
            "name": "Bob Callaway",
            "username": "bobcallaway"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "f20f700c21f5a30524f4fa80cea1951205c6bd35",
          "message": "pin actions and reduce unnecessary perms & creds (#418)\n\nSigned-off-by: Bob Callaway <bcallaway@google.com>",
          "timestamp": "2024-12-12T15:59:27Z",
          "tree_id": "46b8cabd8a075d2f41c5dbb6f86e7725ba2daace",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/f20f700c21f5a30524f4fa80cea1951205c6bd35"
        },
        "date": 1734019208856,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4679378,
            "unit": "ns/op\t  702197 B/op\t   19690 allocs/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4679378,
            "unit": "ns/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702197,
            "unit": "B/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19690,
            "unit": "allocs/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7081,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "234507 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7081,
            "unit": "ns/op",
            "extra": "234507 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "234507 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "234507 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.2,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10349031 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.2,
            "unit": "ns/op",
            "extra": "10349031 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10349031 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10349031 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 351701,
            "unit": "ns/op\t  297690 B/op\t    3432 allocs/op",
            "extra": "3789 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 351701,
            "unit": "ns/op",
            "extra": "3789 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 297690,
            "unit": "B/op",
            "extra": "3789 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3432,
            "unit": "allocs/op",
            "extra": "3789 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "e79f2eac95fa6744d2f9acf4f335e5a9070a05e7",
          "message": "Bump golang.org/x/crypto from 0.30.0 to 0.31.0 in the go_modules group (#415)\n\nBumps the go_modules group with 1 update: [golang.org/x/crypto](https://github.com/golang/crypto).\n\n\nUpdates `golang.org/x/crypto` from 0.30.0 to 0.31.0\n- [Commits](https://github.com/golang/crypto/compare/v0.30.0...v0.31.0)\n\n---\nupdated-dependencies:\n- dependency-name: golang.org/x/crypto\n  dependency-type: direct:production\n  dependency-group: go_modules\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-12-16T09:03:02Z",
          "tree_id": "769552fdee6f21c984378539e0f9d47778d69743",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/e79f2eac95fa6744d2f9acf4f335e5a9070a05e7"
        },
        "date": 1734339858137,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4300502,
            "unit": "ns/op\t  704506 B/op\t   19714 allocs/op",
            "extra": "277 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4300502,
            "unit": "ns/op",
            "extra": "277 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 704506,
            "unit": "B/op",
            "extra": "277 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19714,
            "unit": "allocs/op",
            "extra": "277 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7270,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "186208 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7270,
            "unit": "ns/op",
            "extra": "186208 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "186208 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "186208 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 120.5,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10261484 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 120.5,
            "unit": "ns/op",
            "extra": "10261484 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10261484 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10261484 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 359211,
            "unit": "ns/op\t  298441 B/op\t    3439 allocs/op",
            "extra": "3901 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 359211,
            "unit": "ns/op",
            "extra": "3901 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298441,
            "unit": "B/op",
            "extra": "3901 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3439,
            "unit": "allocs/op",
            "extra": "3901 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "al@google.com",
            "name": "Al Cutter",
            "username": "AlCutter"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "f8889f5495a55801ccc417bd76c0bce96deb05db",
          "message": "Remove public from cache-control (#417)",
          "timestamp": "2024-12-16T13:52:01Z",
          "tree_id": "cb26a1aa4ac4b88848ffa96e809b82fa43a44f12",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/f8889f5495a55801ccc417bd76c0bce96deb05db"
        },
        "date": 1734357159068,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4450768,
            "unit": "ns/op\t  702878 B/op\t   19696 allocs/op",
            "extra": "302 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4450768,
            "unit": "ns/op",
            "extra": "302 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702878,
            "unit": "B/op",
            "extra": "302 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19696,
            "unit": "allocs/op",
            "extra": "302 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6529,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "238344 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6529,
            "unit": "ns/op",
            "extra": "238344 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "238344 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "238344 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.6,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10464921 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.6,
            "unit": "ns/op",
            "extra": "10464921 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10464921 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10464921 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 341975,
            "unit": "ns/op\t  298317 B/op\t    3438 allocs/op",
            "extra": "4182 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 341975,
            "unit": "ns/op",
            "extra": "4182 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298317,
            "unit": "B/op",
            "extra": "4182 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3438,
            "unit": "allocs/op",
            "extra": "4182 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "mhutchinson@gmail.com",
            "name": "Martin Hutchinson",
            "username": "mhutchinson"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "166040096b065062fbb962c8792a5ca9d68e61c5",
          "message": "Added link to benchmarks (#419)\n\nMake sure there's always an easy to find link. #338.",
          "timestamp": "2024-12-16T14:38:15Z",
          "tree_id": "9837d4d98aef2a61c8162bd6a26414f54d79b039",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/166040096b065062fbb962c8792a5ca9d68e61c5"
        },
        "date": 1734359935246,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 5087257,
            "unit": "ns/op\t  704957 B/op\t   19719 allocs/op",
            "extra": "270 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 5087257,
            "unit": "ns/op",
            "extra": "270 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 704957,
            "unit": "B/op",
            "extra": "270 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19719,
            "unit": "allocs/op",
            "extra": "270 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7427,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "329950 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7427,
            "unit": "ns/op",
            "extra": "329950 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "329950 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "329950 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 115.2,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10347676 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 115.2,
            "unit": "ns/op",
            "extra": "10347676 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10347676 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10347676 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 354165,
            "unit": "ns/op\t  298232 B/op\t    3438 allocs/op",
            "extra": "4178 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 354165,
            "unit": "ns/op",
            "extra": "4178 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298232,
            "unit": "B/op",
            "extra": "4178 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3438,
            "unit": "allocs/op",
            "extra": "4178 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "al@google.com",
            "name": "Al Cutter",
            "username": "AlCutter"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "c4d2331e71c6a8bf4b9cbf600bf758a2c6b61e15",
          "message": "Fix CT serialisation (#421)",
          "timestamp": "2024-12-17T12:45:43Z",
          "tree_id": "85ee88eca76c40e147a3cf17fd8b5e44a53b3322",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/c4d2331e71c6a8bf4b9cbf600bf758a2c6b61e15"
        },
        "date": 1734439580009,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4607846,
            "unit": "ns/op\t  704686 B/op\t   19714 allocs/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4607846,
            "unit": "ns/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 704686,
            "unit": "B/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19714,
            "unit": "allocs/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6559,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "207412 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6559,
            "unit": "ns/op",
            "extra": "207412 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "207412 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "207412 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 115.1,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10350336 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 115.1,
            "unit": "ns/op",
            "extra": "10350336 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10350336 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10350336 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 351018,
            "unit": "ns/op\t  298005 B/op\t    3435 allocs/op",
            "extra": "4059 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 351018,
            "unit": "ns/op",
            "extra": "4059 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298005,
            "unit": "B/op",
            "extra": "4059 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3435,
            "unit": "allocs/op",
            "extra": "4059 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "mhutchinson@gmail.com",
            "name": "Martin Hutchinson",
            "username": "mhutchinson"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "21df5fa9abd8b4527559365340a03dd58775766c",
          "message": "[Hammer] Throttle has mutex guarding changes (#424)\n\nThis should avoid the weird number that can be reported when changing the limits at runtime. This also cleans up the readability of the main supply loop, avoiding the need to have a label to break to.",
          "timestamp": "2024-12-17T15:57:27Z",
          "tree_id": "385df12a9742a721f9d2a1278176232c83b47ebf",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/21df5fa9abd8b4527559365340a03dd58775766c"
        },
        "date": 1734451085882,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4722910,
            "unit": "ns/op\t  703929 B/op\t   19705 allocs/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4722910,
            "unit": "ns/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 703929,
            "unit": "B/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19705,
            "unit": "allocs/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 5578,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "207372 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 5578,
            "unit": "ns/op",
            "extra": "207372 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "207372 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "207372 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.6,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10383918 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.6,
            "unit": "ns/op",
            "extra": "10383918 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10383918 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10383918 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 346659,
            "unit": "ns/op\t  298512 B/op\t    3439 allocs/op",
            "extra": "3904 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 346659,
            "unit": "ns/op",
            "extra": "3904 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298512,
            "unit": "B/op",
            "extra": "3904 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3439,
            "unit": "allocs/op",
            "extra": "3904 times\n4 procs"
          }
        ]
      }
    ]
  }
}