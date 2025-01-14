window.BENCHMARK_DATA = {
  "lastUpdate": 1736852371781,
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
          "id": "e22b6b4fe5f570e931af18f7a12d5f49a3da35a4",
          "message": "Bump github/codeql-action in the all-gha-deps group (#422)\n\nBumps the all-gha-deps group with 1 update: [github/codeql-action](https://github.com/github/codeql-action).\n\n\nUpdates `github/codeql-action` from 3.27.6 to 3.27.9\n- [Release notes](https://github.com/github/codeql-action/releases)\n- [Changelog](https://github.com/github/codeql-action/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/github/codeql-action/compare/aa578102511db1f4524ed59b8cc2bae4f6e88195...df409f7d9260372bd5f19e5b04e83cb3c43714ae)\n\n---\nupdated-dependencies:\n- dependency-name: github/codeql-action\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2024-12-17T18:27:38Z",
          "tree_id": "8f61790919961f71f80c6abea0b04b68ca7b12af",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/e22b6b4fe5f570e931af18f7a12d5f49a3da35a4"
        },
        "date": 1734460115432,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4772763,
            "unit": "ns/op\t  703525 B/op\t   19706 allocs/op",
            "extra": "228 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4772763,
            "unit": "ns/op",
            "extra": "228 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 703525,
            "unit": "B/op",
            "extra": "228 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19706,
            "unit": "allocs/op",
            "extra": "228 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6505,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "218013 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6505,
            "unit": "ns/op",
            "extra": "218013 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "218013 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "218013 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 115.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10235055 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 115.7,
            "unit": "ns/op",
            "extra": "10235055 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10235055 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10235055 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 362582,
            "unit": "ns/op\t  298075 B/op\t    3436 allocs/op",
            "extra": "4126 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 362582,
            "unit": "ns/op",
            "extra": "4126 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298075,
            "unit": "B/op",
            "extra": "4126 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3436,
            "unit": "allocs/op",
            "extra": "4126 times\n4 procs"
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
          "id": "976c092e906df2e76d0a630c5f0a024844ad31d7",
          "message": "Rollback changes and add some comments (#426)",
          "timestamp": "2024-12-18T14:38:57Z",
          "tree_id": "3e7350ee72e5bf2cd69c7c75b4ea0b240bc8c717",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/976c092e906df2e76d0a630c5f0a024844ad31d7"
        },
        "date": 1734532774578,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4909137,
            "unit": "ns/op\t  708602 B/op\t   19752 allocs/op",
            "extra": "228 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4909137,
            "unit": "ns/op",
            "extra": "228 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 708602,
            "unit": "B/op",
            "extra": "228 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19752,
            "unit": "allocs/op",
            "extra": "228 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6777,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "228582 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6777,
            "unit": "ns/op",
            "extra": "228582 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "228582 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "228582 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.1,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10350966 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.1,
            "unit": "ns/op",
            "extra": "10350966 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10350966 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10350966 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 349123,
            "unit": "ns/op\t  298197 B/op\t    3437 allocs/op",
            "extra": "4154 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 349123,
            "unit": "ns/op",
            "extra": "4154 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298197,
            "unit": "B/op",
            "extra": "4154 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3437,
            "unit": "allocs/op",
            "extra": "4154 times\n4 procs"
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
          "id": "2cf2e08b77648337b980592f0957e104ed55f627",
          "message": "Open AWS storage to being used with non-AWS MySQL and S3 services (#428)",
          "timestamp": "2024-12-19T18:27:19Z",
          "tree_id": "fe51c53990ef020a70359f056f752b7cc508a82b",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/2cf2e08b77648337b980592f0957e104ed55f627"
        },
        "date": 1734632920699,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4915906,
            "unit": "ns/op\t  704734 B/op\t   19718 allocs/op",
            "extra": "285 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4915906,
            "unit": "ns/op",
            "extra": "285 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 704734,
            "unit": "B/op",
            "extra": "285 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19718,
            "unit": "allocs/op",
            "extra": "285 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 2085,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "585652 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 2085,
            "unit": "ns/op",
            "extra": "585652 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "585652 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "585652 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.3,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10371332 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.3,
            "unit": "ns/op",
            "extra": "10371332 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10371332 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10371332 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 361154,
            "unit": "ns/op\t  298004 B/op\t    3435 allocs/op",
            "extra": "3856 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 361154,
            "unit": "ns/op",
            "extra": "3856 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298004,
            "unit": "B/op",
            "extra": "3856 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3435,
            "unit": "allocs/op",
            "extra": "3856 times\n4 procs"
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
          "id": "a012afa30db74b357cf4d493eb07b867b49c3132",
          "message": "Fix s3.NewFromConfig panic (#429)",
          "timestamp": "2024-12-20T11:48:59Z",
          "tree_id": "9614018b921b11129075d4d9d1f355b3f9be9a48",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/a012afa30db74b357cf4d493eb07b867b49c3132"
        },
        "date": 1734695380315,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4846769,
            "unit": "ns/op\t  702380 B/op\t   19692 allocs/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4846769,
            "unit": "ns/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702380,
            "unit": "B/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19692,
            "unit": "allocs/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8104,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "165450 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8104,
            "unit": "ns/op",
            "extra": "165450 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "165450 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "165450 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.2,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10426255 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.2,
            "unit": "ns/op",
            "extra": "10426255 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10426255 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10426255 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 354066,
            "unit": "ns/op\t  297986 B/op\t    3435 allocs/op",
            "extra": "4056 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 354066,
            "unit": "ns/op",
            "extra": "4056 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 297986,
            "unit": "B/op",
            "extra": "4056 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3435,
            "unit": "allocs/op",
            "extra": "4056 times\n4 procs"
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
          "id": "0c2b9a9c7cd1d70d822c7cb219b3c02158855a56",
          "message": "Fix token permission code scanning alert (#425)",
          "timestamp": "2024-12-20T15:28:40Z",
          "tree_id": "66cd1729ed81c23b2f5d45a28b355dbabb2546ba",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/0c2b9a9c7cd1d70d822c7cb219b3c02158855a56"
        },
        "date": 1734708560532,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4824268,
            "unit": "ns/op\t  701033 B/op\t   19679 allocs/op",
            "extra": "243 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4824268,
            "unit": "ns/op",
            "extra": "243 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 701033,
            "unit": "B/op",
            "extra": "243 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19679,
            "unit": "allocs/op",
            "extra": "243 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6846,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "189625 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6846,
            "unit": "ns/op",
            "extra": "189625 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "189625 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "189625 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 119.6,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10367259 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 119.6,
            "unit": "ns/op",
            "extra": "10367259 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10367259 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10367259 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 377494,
            "unit": "ns/op\t  298681 B/op\t    3440 allocs/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 377494,
            "unit": "ns/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298681,
            "unit": "B/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3440,
            "unit": "allocs/op",
            "extra": "3954 times\n4 procs"
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
          "id": "d6fc43c6df3d30869f2869d3d166d14dcb7b57a2",
          "message": "Bump the all-gha-deps group with 2 updates (#430)\n\nBumps the all-gha-deps group with 2 updates: [github/codeql-action](https://github.com/github/codeql-action) and [actions/upload-artifact](https://github.com/actions/upload-artifact).\n\n\nUpdates `github/codeql-action` from 3.27.9 to 3.28.0\n- [Release notes](https://github.com/github/codeql-action/releases)\n- [Changelog](https://github.com/github/codeql-action/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/github/codeql-action/compare/df409f7d9260372bd5f19e5b04e83cb3c43714ae...48ab28a6f5dbc2a99bf1e0131198dd8f1df78169)\n\nUpdates `actions/upload-artifact` from 4.4.3 to 4.5.0\n- [Release notes](https://github.com/actions/upload-artifact/releases)\n- [Commits](https://github.com/actions/upload-artifact/compare/b4b15b8c7c6ac21ea08fcf65892d2ee8f75cf882...6f51ac03b9356f520e9adb1b1b7802705f340c2b)\n\n---\nupdated-dependencies:\n- dependency-name: github/codeql-action\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-gha-deps\n- dependency-name: actions/upload-artifact\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-01-06T03:47:09Z",
          "tree_id": "35ab06a74a2344ba552ed478db3f1c139e95d630",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/d6fc43c6df3d30869f2869d3d166d14dcb7b57a2"
        },
        "date": 1736135306681,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4842123,
            "unit": "ns/op\t  704036 B/op\t   19709 allocs/op",
            "extra": "280 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4842123,
            "unit": "ns/op",
            "extra": "280 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 704036,
            "unit": "B/op",
            "extra": "280 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19709,
            "unit": "allocs/op",
            "extra": "280 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8453,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "124407 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8453,
            "unit": "ns/op",
            "extra": "124407 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "124407 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "124407 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.9,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10279794 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.9,
            "unit": "ns/op",
            "extra": "10279794 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10279794 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10279794 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 357857,
            "unit": "ns/op\t  297686 B/op\t    3431 allocs/op",
            "extra": "3529 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 357857,
            "unit": "ns/op",
            "extra": "3529 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 297686,
            "unit": "B/op",
            "extra": "3529 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3431,
            "unit": "allocs/op",
            "extra": "3529 times\n4 procs"
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
          "id": "58184b977d8e954597fc44967915980bc61f1d00",
          "message": "Bump the all-go-deps group across 1 directory with 7 updates (#432)\n\nBumps the all-go-deps group with 5 updates in the / directory:\r\n\r\n| Package | From | To |\r\n| --- | --- | --- |\r\n| [cloud.google.com/go/storage](https://github.com/googleapis/google-cloud-go) | `1.48.0` | `1.49.0` |\r\n| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.32.6` | `1.32.7` |\r\n| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.28.6` | `1.28.7` |\r\n| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.71.0` | `1.72.0` |\r\n| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.68.1` | `1.69.2` |\r\n\r\n\r\n\r\nUpdates `cloud.google.com/go/storage` from 1.48.0 to 1.49.0\r\n- [Release notes](https://github.com/googleapis/google-cloud-go/releases)\r\n- [Changelog](https://github.com/googleapis/google-cloud-go/blob/main/CHANGES.md)\r\n- [Commits](https://github.com/googleapis/google-cloud-go/compare/spanner/v1.48.0...spanner/v1.49.0)\r\n\r\nUpdates `github.com/aws/aws-sdk-go-v2` from 1.32.6 to 1.32.7\r\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\r\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/v1.32.6...v1.32.7)\r\n\r\nUpdates `github.com/aws/aws-sdk-go-v2/config` from 1.28.6 to 1.28.7\r\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\r\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.28.6...config/v1.28.7)\r\n\r\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.71.0 to 1.72.0\r\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\r\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.71.0...service/s3/v1.72.0)\r\n\r\nUpdates `google.golang.org/api` from 0.210.0 to 0.214.0\r\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\r\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\r\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.210.0...v0.214.0)\r\n\r\nUpdates `google.golang.org/grpc` from 1.68.1 to 1.69.2\r\n- [Release notes](https://github.com/grpc/grpc-go/releases)\r\n- [Commits](https://github.com/grpc/grpc-go/compare/v1.68.1...v1.69.2)\r\n\r\nUpdates `golang.org/x/net` from 0.32.0 to 0.33.0\r\n- [Commits](https://github.com/golang/net/compare/v0.32.0...v0.33.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: cloud.google.com/go/storage\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n  dependency-group: all-go-deps\r\n- dependency-name: github.com/aws/aws-sdk-go-v2\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n  dependency-group: all-go-deps\r\n- dependency-name: github.com/aws/aws-sdk-go-v2/config\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n  dependency-group: all-go-deps\r\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n  dependency-group: all-go-deps\r\n- dependency-name: google.golang.org/api\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n  dependency-group: all-go-deps\r\n- dependency-name: google.golang.org/grpc\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n  dependency-group: all-go-deps\r\n- dependency-name: golang.org/x/net\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n  dependency-group: all-go-deps\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-01-06T11:10:28Z",
          "tree_id": "5cfca2560cbb6c845a35e6d8e59787e4234b4a44",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/58184b977d8e954597fc44967915980bc61f1d00"
        },
        "date": 1736161903608,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4537455,
            "unit": "ns/op\t  702897 B/op\t   19695 allocs/op",
            "extra": "274 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4537455,
            "unit": "ns/op",
            "extra": "274 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702897,
            "unit": "B/op",
            "extra": "274 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19695,
            "unit": "allocs/op",
            "extra": "274 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7683,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "166917 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7683,
            "unit": "ns/op",
            "extra": "166917 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "166917 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "166917 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.5,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10398736 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.5,
            "unit": "ns/op",
            "extra": "10398736 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10398736 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10398736 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 358401,
            "unit": "ns/op\t  298686 B/op\t    3440 allocs/op",
            "extra": "3955 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 358401,
            "unit": "ns/op",
            "extra": "3955 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298686,
            "unit": "B/op",
            "extra": "3955 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3440,
            "unit": "allocs/op",
            "extra": "3955 times\n4 procs"
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
          "id": "3af7a0ee376f033780dacff6d44cba08faf87cf5",
          "message": "Bump the all-go-deps group with 2 updates (#434)\n\nBumps the all-go-deps group with 2 updates: [golang.org/x/crypto](https://github.com/golang/crypto) and [golang.org/x/net](https://github.com/golang/net).\r\n\r\n\r\nUpdates `golang.org/x/crypto` from 0.31.0 to 0.32.0\r\n- [Commits](https://github.com/golang/crypto/compare/v0.31.0...v0.32.0)\r\n\r\nUpdates `golang.org/x/net` from 0.33.0 to 0.34.0\r\n- [Commits](https://github.com/golang/net/compare/v0.33.0...v0.34.0)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: golang.org/x/crypto\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n  dependency-group: all-go-deps\r\n- dependency-name: golang.org/x/net\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n  dependency-group: all-go-deps\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-01-07T13:29:31Z",
          "tree_id": "215199fa6cd72ee4bc792bfffd56a137d5439eca",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/3af7a0ee376f033780dacff6d44cba08faf87cf5"
        },
        "date": 1736256655089,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4447925,
            "unit": "ns/op\t  700881 B/op\t   19678 allocs/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4447925,
            "unit": "ns/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700881,
            "unit": "B/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19678,
            "unit": "allocs/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7400,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "202981 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7400,
            "unit": "ns/op",
            "extra": "202981 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "202981 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "202981 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 120,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10208613 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 120,
            "unit": "ns/op",
            "extra": "10208613 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10208613 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10208613 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 355078,
            "unit": "ns/op\t  298309 B/op\t    3438 allocs/op",
            "extra": "3890 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 355078,
            "unit": "ns/op",
            "extra": "3890 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298309,
            "unit": "B/op",
            "extra": "3890 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3438,
            "unit": "allocs/op",
            "extra": "3890 times\n4 procs"
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
          "id": "2edef5836737da9f2fc707061f6dfdee91a01173",
          "message": "[Architecture] Switch over to a common impl with separate drivers (#433)\n\nThe functionality is the same, but now all personalities have a common wrapper around a driver instead of having separate implementations directly in their hand.\r\n\r\nThe big advantage of the main API being this shape is that Tessera now \"feels\" like a single API, with drivers to bind it to different environments. This is more similar to canonical usage of sql/db package, which many go developers are familiar with.\r\n\r\nAnother advantage realized in this change is that deduplication is now folded into the core Appender object, instead of functionally floating around alongside it and needing to be handled by the personalities.",
          "timestamp": "2025-01-08T13:02:49Z",
          "tree_id": "6b8598bc7d612b005281abeb867626b92ffa3302",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/2edef5836737da9f2fc707061f6dfdee91a01173"
        },
        "date": 1736341451807,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4432382,
            "unit": "ns/op\t  699657 B/op\t   19665 allocs/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4432382,
            "unit": "ns/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 699657,
            "unit": "B/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19665,
            "unit": "allocs/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8088,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "140977 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8088,
            "unit": "ns/op",
            "extra": "140977 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "140977 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "140977 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 118.6,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10272558 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 118.6,
            "unit": "ns/op",
            "extra": "10272558 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10272558 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10272558 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 366946,
            "unit": "ns/op\t  297615 B/op\t    3432 allocs/op",
            "extra": "3751 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 366946,
            "unit": "ns/op",
            "extra": "3751 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 297615,
            "unit": "B/op",
            "extra": "3751 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3432,
            "unit": "allocs/op",
            "extra": "3751 times\n4 procs"
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
          "id": "f6a7b6b80167eeedf0b00c61ba8841215b6f1d29",
          "message": "Bumped grpc deps (#435)\n\nThis is the result of running the commands in this comment, along with go mod tidy: https://github.com/googleapis/google-cloud-go/issues/11283\\#issuecomment-2558566621",
          "timestamp": "2025-01-08T13:16:53Z",
          "tree_id": "dfdde84db3bf4d253df581723431dfde6f85a101",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/f6a7b6b80167eeedf0b00c61ba8841215b6f1d29"
        },
        "date": 1736342290252,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 5288389,
            "unit": "ns/op\t  703953 B/op\t   19710 allocs/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 5288389,
            "unit": "ns/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 703953,
            "unit": "B/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19710,
            "unit": "allocs/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8012,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "126756 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8012,
            "unit": "ns/op",
            "extra": "126756 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "126756 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "126756 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 118,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10132042 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 118,
            "unit": "ns/op",
            "extra": "10132042 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10132042 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10132042 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 363203,
            "unit": "ns/op\t  298830 B/op\t    3441 allocs/op",
            "extra": "3943 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 363203,
            "unit": "ns/op",
            "extra": "3943 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 298830,
            "unit": "B/op",
            "extra": "3943 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3441,
            "unit": "allocs/op",
            "extra": "3943 times\n4 procs"
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
          "id": "60243a5f7f9b454d190471e264e336b992697959",
          "message": "Integrate operates in terms of hashes (#437)",
          "timestamp": "2025-01-09T16:32:47Z",
          "tree_id": "01ba58337a4b40b16b398bd203681160bb1984fb",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/60243a5f7f9b454d190471e264e336b992697959"
        },
        "date": 1736440402213,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1166842,
            "unit": "ns/op\t  689396 B/op\t   19559 allocs/op",
            "extra": "985 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1166842,
            "unit": "ns/op",
            "extra": "985 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689396,
            "unit": "B/op",
            "extra": "985 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19559,
            "unit": "allocs/op",
            "extra": "985 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 2039,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "693253 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 2039,
            "unit": "ns/op",
            "extra": "693253 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "693253 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "693253 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 118.1,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10391788 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 118.1,
            "unit": "ns/op",
            "extra": "10391788 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10391788 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10391788 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 351615,
            "unit": "ns/op\t  290431 B/op\t    3036 allocs/op",
            "extra": "4148 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 351615,
            "unit": "ns/op",
            "extra": "4148 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290431,
            "unit": "B/op",
            "extra": "4148 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "4148 times\n4 procs"
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
          "id": "ddac6ea1fa91dc330423377b3ccb114a359a554e",
          "message": "[POSIX] Tweaks (#439)",
          "timestamp": "2025-01-09T17:21:47Z",
          "tree_id": "5b3b84882f6c853bf9d68975f3f224b4641989f8",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/ddac6ea1fa91dc330423377b3ccb114a359a554e"
        },
        "date": 1736443338568,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1138777,
            "unit": "ns/op\t  689161 B/op\t   19557 allocs/op",
            "extra": "1074 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1138777,
            "unit": "ns/op",
            "extra": "1074 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689161,
            "unit": "B/op",
            "extra": "1074 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19557,
            "unit": "allocs/op",
            "extra": "1074 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1964,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "643021 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1964,
            "unit": "ns/op",
            "extra": "643021 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "643021 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "643021 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 122.5,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "9370680 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 122.5,
            "unit": "ns/op",
            "extra": "9370680 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "9370680 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9370680 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 343430,
            "unit": "ns/op\t  291334 B/op\t    3043 allocs/op",
            "extra": "4275 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 343430,
            "unit": "ns/op",
            "extra": "4275 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291334,
            "unit": "B/op",
            "extra": "4275 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3043,
            "unit": "allocs/op",
            "extra": "4275 times\n4 procs"
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
          "id": "7914bd2c94c23abd1a60f10826965e3906319673",
          "message": "Extract tighter-scoped integrate funcs in storage impls (#440)\n\nThis PR pulls the act of updating the merkle tree resources out into separate funcs.\r\n\r\nThis helps with readability by reducing the size/tightening the scope of some of the funcs in storage implementations (MySQL in particular was getting quite large), but also aids in supporting other lifecycle modes (e.g. for #414).\r\n\r\nNo functional changes.",
          "timestamp": "2025-01-10T12:11:41Z",
          "tree_id": "65db3431fd90d6929ebd3d4cd97b1bbd36bb6c62",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/7914bd2c94c23abd1a60f10826965e3906319673"
        },
        "date": 1736511134054,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1154956,
            "unit": "ns/op\t  689261 B/op\t   19557 allocs/op",
            "extra": "1033 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1154956,
            "unit": "ns/op",
            "extra": "1033 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689261,
            "unit": "B/op",
            "extra": "1033 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19557,
            "unit": "allocs/op",
            "extra": "1033 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1981,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "632551 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1981,
            "unit": "ns/op",
            "extra": "632551 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "632551 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "632551 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 123.3,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "9560186 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 123.3,
            "unit": "ns/op",
            "extra": "9560186 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "9560186 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9560186 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 356932,
            "unit": "ns/op\t  290271 B/op\t    3035 allocs/op",
            "extra": "4065 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 356932,
            "unit": "ns/op",
            "extra": "4065 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290271,
            "unit": "B/op",
            "extra": "4065 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3035,
            "unit": "allocs/op",
            "extra": "4065 times\n4 procs"
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
          "id": "75a11eeee8c8d1ade1fc82b1f516de16bd14ddbd",
          "message": "[Docs] Glanceable roadmap features in main README (#442)\n\nAlso fixed some out of date docs around alpha and AWS",
          "timestamp": "2025-01-13T15:03:03Z",
          "tree_id": "11b6ee4f6d4882702fa1e04d0fe180fde7ee80e7",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/75a11eeee8c8d1ade1fc82b1f516de16bd14ddbd"
        },
        "date": 1736780618616,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1158255,
            "unit": "ns/op\t  689200 B/op\t   19557 allocs/op",
            "extra": "1021 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1158255,
            "unit": "ns/op",
            "extra": "1021 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689200,
            "unit": "B/op",
            "extra": "1021 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19557,
            "unit": "allocs/op",
            "extra": "1021 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1936,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "697722 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1936,
            "unit": "ns/op",
            "extra": "697722 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "697722 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "697722 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 123.4,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "9695916 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 123.4,
            "unit": "ns/op",
            "extra": "9695916 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "9695916 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9695916 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 351700,
            "unit": "ns/op\t  290466 B/op\t    3036 allocs/op",
            "extra": "4030 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 351700,
            "unit": "ns/op",
            "extra": "4030 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290466,
            "unit": "B/op",
            "extra": "4030 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "4030 times\n4 procs"
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
          "id": "c098a25f33f72064b20be3f99b5bef06245a3d29",
          "message": "Bump workflows in prep for Go 1.23 (#444)",
          "timestamp": "2025-01-13T17:29:09Z",
          "tree_id": "cec0daba9a614db963b1a3f49769051f43cf4822",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/c098a25f33f72064b20be3f99b5bef06245a3d29"
        },
        "date": 1736789382156,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1222237,
            "unit": "ns/op\t  689619 B/op\t   19561 allocs/op",
            "extra": "1022 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1222237,
            "unit": "ns/op",
            "extra": "1022 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689619,
            "unit": "B/op",
            "extra": "1022 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19561,
            "unit": "allocs/op",
            "extra": "1022 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1991,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "599902 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1991,
            "unit": "ns/op",
            "extra": "599902 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "599902 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "599902 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.8,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10021544 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.8,
            "unit": "ns/op",
            "extra": "10021544 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10021544 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10021544 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 347853,
            "unit": "ns/op\t  290624 B/op\t    3038 allocs/op",
            "extra": "4180 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 347853,
            "unit": "ns/op",
            "extra": "4180 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290624,
            "unit": "B/op",
            "extra": "4180 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3038,
            "unit": "allocs/op",
            "extra": "4180 times\n4 procs"
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
          "id": "902ee7220616a7c728b524b458beb384d1ac05bd",
          "message": "Bump to Go 1.23 (#443)",
          "timestamp": "2025-01-13T17:46:19Z",
          "tree_id": "b9f595cac4a58af74ab2c4eafb83caf98b761072",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/902ee7220616a7c728b524b458beb384d1ac05bd"
        },
        "date": 1736790464704,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4751429,
            "unit": "ns/op\t  701214 B/op\t   19679 allocs/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4751429,
            "unit": "ns/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 701214,
            "unit": "B/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19679,
            "unit": "allocs/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7724,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "218325 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7724,
            "unit": "ns/op",
            "extra": "218325 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "218325 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "218325 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.5,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10509267 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.5,
            "unit": "ns/op",
            "extra": "10509267 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10509267 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10509267 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 353916,
            "unit": "ns/op\t  291004 B/op\t    3040 allocs/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 353916,
            "unit": "ns/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291004,
            "unit": "B/op",
            "extra": "3954 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3040,
            "unit": "allocs/op",
            "extra": "3954 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "paflynn@google.com",
            "name": "Patrick Flynn",
            "username": "patflynn"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "4ef73e9f3d656bf07997261eb0840ca286196cdb",
          "message": "Add codeowners so that PRs get auto-assigned (#427)\n\n* Add codeowners so that PRs get auto-assigned\r\n\r\n* fix whitespace",
          "timestamp": "2025-01-13T21:20:40Z",
          "tree_id": "2303dc67faeeebf0ecf6c5c0697ebb92ea988b70",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/4ef73e9f3d656bf07997261eb0840ca286196cdb"
        },
        "date": 1736803324977,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4557401,
            "unit": "ns/op\t  702251 B/op\t   19693 allocs/op",
            "extra": "284 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4557401,
            "unit": "ns/op",
            "extra": "284 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702251,
            "unit": "B/op",
            "extra": "284 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19693,
            "unit": "allocs/op",
            "extra": "284 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7796,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "155685 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7796,
            "unit": "ns/op",
            "extra": "155685 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "155685 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "155685 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 118.5,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10087470 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 118.5,
            "unit": "ns/op",
            "extra": "10087470 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10087470 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10087470 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 354240,
            "unit": "ns/op\t  290731 B/op\t    3039 allocs/op",
            "extra": "3896 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 354240,
            "unit": "ns/op",
            "extra": "3896 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290731,
            "unit": "B/op",
            "extra": "3896 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3039,
            "unit": "allocs/op",
            "extra": "3896 times\n4 procs"
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
          "id": "7176fa98e72b265e25933087961d928870cdbc24",
          "message": "Bump the all-go-deps group with 8 updates (#445)\n\nBumps the all-go-deps group with 8 updates:\n\n| Package | From | To |\n| --- | --- | --- |\n| [cloud.google.com/go/storage](https://github.com/googleapis/google-cloud-go) | `1.49.0` | `1.50.0` |\n| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.32.7` | `1.32.8` |\n| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.28.7` | `1.28.10` |\n| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.48` | `1.17.51` |\n| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.72.0` | `1.72.2` |\n| [github.com/gdamore/tcell/v2](https://github.com/gdamore/tcell) | `2.7.4` | `2.8.1` |\n| [google.golang.org/api](https://github.com/googleapis/google-api-go-client) | `0.215.0` | `0.216.0` |\n| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.69.2` | `1.69.4` |\n\n\nUpdates `cloud.google.com/go/storage` from 1.49.0 to 1.50.0\n- [Release notes](https://github.com/googleapis/google-cloud-go/releases)\n- [Changelog](https://github.com/googleapis/google-cloud-go/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-cloud-go/compare/spanner/v1.49.0...spanner/v1.50.0)\n\nUpdates `github.com/aws/aws-sdk-go-v2` from 1.32.7 to 1.32.8\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/v1.32.7...v1.32.8)\n\nUpdates `github.com/aws/aws-sdk-go-v2/config` from 1.28.7 to 1.28.10\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.28.7...config/v1.28.10)\n\nUpdates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.48 to 1.17.51\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.48...credentials/v1.17.51)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.72.0 to 1.72.2\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.72.0...service/s3/v1.72.2)\n\nUpdates `github.com/gdamore/tcell/v2` from 2.7.4 to 2.8.1\n- [Release notes](https://github.com/gdamore/tcell/releases)\n- [Changelog](https://github.com/gdamore/tcell/blob/main/CHANGESv2.md)\n- [Commits](https://github.com/gdamore/tcell/compare/v2.7.4...v2.8.1)\n\nUpdates `google.golang.org/api` from 0.215.0 to 0.216.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.215.0...v0.216.0)\n\nUpdates `google.golang.org/grpc` from 1.69.2 to 1.69.4\n- [Release notes](https://github.com/grpc/grpc-go/releases)\n- [Commits](https://github.com/grpc/grpc-go/compare/v1.69.2...v1.69.4)\n\n---\nupdated-dependencies:\n- dependency-name: cloud.google.com/go/storage\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/config\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/credentials\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/gdamore/tcell/v2\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/grpc\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-01-14T10:58:07Z",
          "tree_id": "c6ce60c5b25e8dbfbb5d3b7c97a740c8f9b3c7da",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/7176fa98e72b265e25933087961d928870cdbc24"
        },
        "date": 1736852370850,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4542912,
            "unit": "ns/op\t  699526 B/op\t   19663 allocs/op",
            "extra": "278 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4542912,
            "unit": "ns/op",
            "extra": "278 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 699526,
            "unit": "B/op",
            "extra": "278 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19663,
            "unit": "allocs/op",
            "extra": "278 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7733,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "155791 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7733,
            "unit": "ns/op",
            "extra": "155791 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "155791 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "155791 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 118,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "8983794 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 118,
            "unit": "ns/op",
            "extra": "8983794 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "8983794 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8983794 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 347763,
            "unit": "ns/op\t  290395 B/op\t    3036 allocs/op",
            "extra": "4134 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 347763,
            "unit": "ns/op",
            "extra": "4134 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290395,
            "unit": "B/op",
            "extra": "4134 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "4134 times\n4 procs"
          }
        ]
      }
    ]
  }
}