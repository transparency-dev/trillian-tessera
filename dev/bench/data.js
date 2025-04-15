window.BENCHMARK_DATA = {
  "lastUpdate": 1744735333302,
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
          "id": "c8604b6db53c11ab380c5b277aeb56bb6b9dc2a5",
          "message": "Experimental migration support for POSIX (#441)",
          "timestamp": "2025-01-14T14:02:12Z",
          "tree_id": "900dc9aaaf01aab2142f640c8f9e5ceece68ea2f",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/c8604b6db53c11ab380c5b277aeb56bb6b9dc2a5"
        },
        "date": 1736863417740,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4502555,
            "unit": "ns/op\t  700490 B/op\t   19672 allocs/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4502555,
            "unit": "ns/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700490,
            "unit": "B/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19672,
            "unit": "allocs/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8097,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "171025 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8097,
            "unit": "ns/op",
            "extra": "171025 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "171025 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "171025 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10300754 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117,
            "unit": "ns/op",
            "extra": "10300754 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10300754 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10300754 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 351934,
            "unit": "ns/op\t  290571 B/op\t    3037 allocs/op",
            "extra": "4003 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 351934,
            "unit": "ns/op",
            "extra": "4003 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290571,
            "unit": "B/op",
            "extra": "4003 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3037,
            "unit": "allocs/op",
            "extra": "4003 times\n4 procs"
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
          "id": "c883591a4cf040738bd0f58ace3d1d8bcfd677dc",
          "message": "Bump the all-gha-deps group with 2 updates (#446)\n\nBumps the all-gha-deps group with 2 updates: [github/codeql-action](https://github.com/github/codeql-action) and [actions/upload-artifact](https://github.com/actions/upload-artifact).\r\n\r\n\r\nUpdates `github/codeql-action` from 3.28.0 to 3.28.1\r\n- [Release notes](https://github.com/github/codeql-action/releases)\r\n- [Changelog](https://github.com/github/codeql-action/blob/main/CHANGELOG.md)\r\n- [Commits](https://github.com/github/codeql-action/compare/48ab28a6f5dbc2a99bf1e0131198dd8f1df78169...b6a472f63d85b9c78a3ac5e89422239fc15e9b3c)\r\n\r\nUpdates `actions/upload-artifact` from 4.5.0 to 4.6.0\r\n- [Release notes](https://github.com/actions/upload-artifact/releases)\r\n- [Commits](https://github.com/actions/upload-artifact/compare/6f51ac03b9356f520e9adb1b1b7802705f340c2b...65c4c4a1ddee5b72f698fdd19549f0f0fb45cf08)\r\n\r\n---\r\nupdated-dependencies:\r\n- dependency-name: github/codeql-action\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-patch\r\n  dependency-group: all-gha-deps\r\n- dependency-name: actions/upload-artifact\r\n  dependency-type: direct:production\r\n  update-type: version-update:semver-minor\r\n  dependency-group: all-gha-deps\r\n...\r\n\r\nSigned-off-by: dependabot[bot] <support@github.com>\r\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-01-14T16:00:39Z",
          "tree_id": "61e1c821034597a9e866ba46d812239f292143d0",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/c883591a4cf040738bd0f58ace3d1d8bcfd677dc"
        },
        "date": 1736870490252,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 5329169,
            "unit": "ns/op\t  701226 B/op\t   19680 allocs/op",
            "extra": "232 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 5329169,
            "unit": "ns/op",
            "extra": "232 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 701226,
            "unit": "B/op",
            "extra": "232 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19680,
            "unit": "allocs/op",
            "extra": "232 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7192,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "210364 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7192,
            "unit": "ns/op",
            "extra": "210364 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "210364 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "210364 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.9,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10254630 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.9,
            "unit": "ns/op",
            "extra": "10254630 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10254630 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10254630 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 344561,
            "unit": "ns/op\t  290216 B/op\t    3033 allocs/op",
            "extra": "3685 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 344561,
            "unit": "ns/op",
            "extra": "3685 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290216,
            "unit": "B/op",
            "extra": "3685 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3033,
            "unit": "allocs/op",
            "extra": "3685 times\n4 procs"
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
          "id": "770517a21045a1dcaa893a01c226440a70905c50",
          "message": "GCP Migrate tool (#447)",
          "timestamp": "2025-01-14T17:05:17Z",
          "tree_id": "6e0f96bd3734f5cc170c37faaea9bdfcad544d9e",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/770517a21045a1dcaa893a01c226440a70905c50"
        },
        "date": 1736874366918,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 5150162,
            "unit": "ns/op\t  705236 B/op\t   19719 allocs/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 5150162,
            "unit": "ns/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 705236,
            "unit": "B/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19719,
            "unit": "allocs/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6762,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "213811 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6762,
            "unit": "ns/op",
            "extra": "213811 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "213811 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "213811 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.8,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10291172 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.8,
            "unit": "ns/op",
            "extra": "10291172 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10291172 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10291172 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 358634,
            "unit": "ns/op\t  290909 B/op\t    3040 allocs/op",
            "extra": "3909 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 358634,
            "unit": "ns/op",
            "extra": "3909 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290909,
            "unit": "B/op",
            "extra": "3909 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3040,
            "unit": "allocs/op",
            "extra": "3909 times\n4 procs"
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
          "id": "1b545d033945e4c85e943c11371d00ab10e57b03",
          "message": "[Hammer] Refactor to allow code reuse with CT (#448)\n\nThis isn't the full job, as the core of the hammer needs extracting from\r\nhammer.go. This is a good start and sets the right direction for making\r\nthis a general purpose library for use in true tlog-tiles and the static\r\nCT variation.",
          "timestamp": "2025-01-15T14:57:21Z",
          "tree_id": "121994839d74238ac84f813db87acb3ab1af74ee",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/1b545d033945e4c85e943c11371d00ab10e57b03"
        },
        "date": 1736953084580,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4536496,
            "unit": "ns/op\t  698747 B/op\t   19655 allocs/op",
            "extra": "248 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4536496,
            "unit": "ns/op",
            "extra": "248 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 698747,
            "unit": "B/op",
            "extra": "248 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19655,
            "unit": "allocs/op",
            "extra": "248 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6964,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "218204 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6964,
            "unit": "ns/op",
            "extra": "218204 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "218204 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "218204 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.8,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "9331662 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.8,
            "unit": "ns/op",
            "extra": "9331662 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "9331662 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9331662 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 359312,
            "unit": "ns/op\t  290952 B/op\t    3040 allocs/op",
            "extra": "3912 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 359312,
            "unit": "ns/op",
            "extra": "3912 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290952,
            "unit": "B/op",
            "extra": "3912 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3040,
            "unit": "allocs/op",
            "extra": "3912 times\n4 procs"
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
          "id": "94e075991c8531d759483afe49710194519a0853",
          "message": "Add schema compatibility version checking to GCP (#451)",
          "timestamp": "2025-01-16T16:56:36Z",
          "tree_id": "d9efef0ae9247c3980544561d2660ba00770ecc1",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/94e075991c8531d759483afe49710194519a0853"
        },
        "date": 1737046641914,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4784672,
            "unit": "ns/op\t  696400 B/op\t   19631 allocs/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4784672,
            "unit": "ns/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 696400,
            "unit": "B/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19631,
            "unit": "allocs/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7151,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "146677 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7151,
            "unit": "ns/op",
            "extra": "146677 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "146677 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "146677 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.9,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10272903 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.9,
            "unit": "ns/op",
            "extra": "10272903 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10272903 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10272903 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 350042,
            "unit": "ns/op\t  290031 B/op\t    3032 allocs/op",
            "extra": "3787 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 350042,
            "unit": "ns/op",
            "extra": "3787 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290031,
            "unit": "B/op",
            "extra": "3787 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3787 times\n4 procs"
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
          "id": "ecacf9c0a91fb59e4c08d17d6aa5e79758389e58",
          "message": "Rando suffix for CI (#452)",
          "timestamp": "2025-01-16T17:54:42Z",
          "tree_id": "4ce346ba9a2fe66d5e7d8cd1ce629c9b9d176487",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/ecacf9c0a91fb59e4c08d17d6aa5e79758389e58"
        },
        "date": 1737050128978,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4740038,
            "unit": "ns/op\t  702305 B/op\t   19692 allocs/op",
            "extra": "279 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4740038,
            "unit": "ns/op",
            "extra": "279 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702305,
            "unit": "B/op",
            "extra": "279 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19692,
            "unit": "allocs/op",
            "extra": "279 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7384,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "188605 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7384,
            "unit": "ns/op",
            "extra": "188605 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "188605 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "188605 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.2,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10126574 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.2,
            "unit": "ns/op",
            "extra": "10126574 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10126574 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10126574 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 352019,
            "unit": "ns/op\t  290950 B/op\t    3040 allocs/op",
            "extra": "3914 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 352019,
            "unit": "ns/op",
            "extra": "3914 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290950,
            "unit": "B/op",
            "extra": "3914 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3040,
            "unit": "allocs/op",
            "extra": "3914 times\n4 procs"
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
          "id": "ee70a897f836eb7fdecbb56cb666a42a3b5f5f75",
          "message": "Output conformance bucket and use in GCB (#453)",
          "timestamp": "2025-01-17T10:38:01Z",
          "tree_id": "5a56f585187f39c858bb7acd975682645ad1c87e",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/ee70a897f836eb7fdecbb56cb666a42a3b5f5f75"
        },
        "date": 1737110324351,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4116715,
            "unit": "ns/op\t  697448 B/op\t   19641 allocs/op",
            "extra": "248 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4116715,
            "unit": "ns/op",
            "extra": "248 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 697448,
            "unit": "B/op",
            "extra": "248 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19641,
            "unit": "allocs/op",
            "extra": "248 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6985,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "236356 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6985,
            "unit": "ns/op",
            "extra": "236356 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "236356 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "236356 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 118.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10298379 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 118.7,
            "unit": "ns/op",
            "extra": "10298379 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10298379 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10298379 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 347126,
            "unit": "ns/op\t  289952 B/op\t    3032 allocs/op",
            "extra": "3747 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 347126,
            "unit": "ns/op",
            "extra": "3747 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 289952,
            "unit": "B/op",
            "extra": "3747 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3747 times\n4 procs"
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
          "id": "7f25479b4992301875b960d1a0de4f546e6a5dab",
          "message": "Rename GCB trigger to better match the purpose (#454)",
          "timestamp": "2025-01-17T12:06:26Z",
          "tree_id": "29b8da3a3a703e4bddda735bff27d54383600817",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/7f25479b4992301875b960d1a0de4f546e6a5dab"
        },
        "date": 1737115634691,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4849373,
            "unit": "ns/op\t  700222 B/op\t   19671 allocs/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4849373,
            "unit": "ns/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700222,
            "unit": "B/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19671,
            "unit": "allocs/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7777,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "163528 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7777,
            "unit": "ns/op",
            "extra": "163528 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "163528 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "163528 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 118.8,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10204651 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 118.8,
            "unit": "ns/op",
            "extra": "10204651 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10204651 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10204651 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 386835,
            "unit": "ns/op\t  290028 B/op\t    3032 allocs/op",
            "extra": "3796 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 386835,
            "unit": "ns/op",
            "extra": "3796 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290028,
            "unit": "B/op",
            "extra": "3796 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3796 times\n4 procs"
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
          "id": "2613db51460955a5edf32860cfedb5c26b532c73",
          "message": "Bump the all-go-deps group with 5 updates (#455)\n\nBumps the all-go-deps group with 5 updates:\n\n| Package | From | To |\n| --- | --- | --- |\n| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.32.8` | `1.33.0` |\n| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.28.10` | `1.29.1` |\n| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.51` | `1.17.54` |\n| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.72.2` | `1.73.2` |\n| [google.golang.org/api](https://github.com/googleapis/google-api-go-client) | `0.216.0` | `0.217.0` |\n\n\nUpdates `github.com/aws/aws-sdk-go-v2` from 1.32.8 to 1.33.0\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/v1.32.8...v1.33.0)\n\nUpdates `github.com/aws/aws-sdk-go-v2/config` from 1.28.10 to 1.29.1\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.28.10...config/v1.29.1)\n\nUpdates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.51 to 1.17.54\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.51...credentials/v1.17.54)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.72.2 to 1.73.2\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.72.2...service/s3/v1.73.2)\n\nUpdates `google.golang.org/api` from 0.216.0 to 0.217.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.216.0...v0.217.0)\n\n---\nupdated-dependencies:\n- dependency-name: github.com/aws/aws-sdk-go-v2\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/config\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/credentials\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-01-20T20:30:36Z",
          "tree_id": "4d612004eda49b735b778b7ab31f3b71b3a33463",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/2613db51460955a5edf32860cfedb5c26b532c73"
        },
        "date": 1737405129300,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4686726,
            "unit": "ns/op\t  703027 B/op\t   19699 allocs/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4686726,
            "unit": "ns/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 703027,
            "unit": "B/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19699,
            "unit": "allocs/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8126,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "219658 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8126,
            "unit": "ns/op",
            "extra": "219658 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "219658 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "219658 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.3,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10110645 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.3,
            "unit": "ns/op",
            "extra": "10110645 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10110645 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10110645 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 356325,
            "unit": "ns/op\t  290645 B/op\t    3038 allocs/op",
            "extra": "3891 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 356325,
            "unit": "ns/op",
            "extra": "3891 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290645,
            "unit": "B/op",
            "extra": "3891 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3038,
            "unit": "allocs/op",
            "extra": "3891 times\n4 procs"
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
          "id": "20a5dcfc29be5230aaaae736c4742ebd245a1abe",
          "message": "Bump golangci/golangci-lint-action in the all-gha-deps group (#456)\n\nBumps the all-gha-deps group with 1 update: [golangci/golangci-lint-action](https://github.com/golangci/golangci-lint-action).\n\n\nUpdates `golangci/golangci-lint-action` from 6.1.1 to 6.2.0\n- [Release notes](https://github.com/golangci/golangci-lint-action/releases)\n- [Commits](https://github.com/golangci/golangci-lint-action/compare/971e284b6050e8a5849b72094c50ab08da042db8...ec5d18412c0aeab7936cb16880d708ba2a64e1ae)\n\n---\nupdated-dependencies:\n- dependency-name: golangci/golangci-lint-action\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-01-20T22:02:03Z",
          "tree_id": "cc1d15c54a8a64739c859a0bb69607f72d0baee3",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/20a5dcfc29be5230aaaae736c4742ebd245a1abe"
        },
        "date": 1737410568703,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 5076326,
            "unit": "ns/op\t  702449 B/op\t   19693 allocs/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 5076326,
            "unit": "ns/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702449,
            "unit": "B/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19693,
            "unit": "allocs/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7177,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "201294 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7177,
            "unit": "ns/op",
            "extra": "201294 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "201294 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "201294 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 119.3,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10330394 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 119.3,
            "unit": "ns/op",
            "extra": "10330394 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10330394 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10330394 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 345459,
            "unit": "ns/op\t  290103 B/op\t    3033 allocs/op",
            "extra": "3813 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 345459,
            "unit": "ns/op",
            "extra": "3813 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290103,
            "unit": "B/op",
            "extra": "3813 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3033,
            "unit": "allocs/op",
            "extra": "3813 times\n4 procs"
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
          "id": "5c626d7328e2740466625bc8ab21eaed85eb47af",
          "message": "[Hammer] Moved more common functionality into loadtest package (#449)\n\nThis is a rough approximation of what a library version of the hammer\r\ncould look like. The remaining code in the outer `main` package is now\r\nthe tlog-tiles specific stuff. The code in the `loadtest` package should\r\nbe configurable to use with Static CT.",
          "timestamp": "2025-01-22T10:38:53Z",
          "tree_id": "98faebb542dcb1203aa5422ed711b096c493774a",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/5c626d7328e2740466625bc8ab21eaed85eb47af"
        },
        "date": 1737542375776,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4612599,
            "unit": "ns/op\t  699745 B/op\t   19665 allocs/op",
            "extra": "225 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4612599,
            "unit": "ns/op",
            "extra": "225 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 699745,
            "unit": "B/op",
            "extra": "225 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19665,
            "unit": "allocs/op",
            "extra": "225 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6866,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "228860 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6866,
            "unit": "ns/op",
            "extra": "228860 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "228860 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "228860 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10155724 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117,
            "unit": "ns/op",
            "extra": "10155724 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10155724 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10155724 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 355818,
            "unit": "ns/op\t  290301 B/op\t    3035 allocs/op",
            "extra": "4102 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 355818,
            "unit": "ns/op",
            "extra": "4102 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290301,
            "unit": "B/op",
            "extra": "4102 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3035,
            "unit": "allocs/op",
            "extra": "4102 times\n4 procs"
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
          "id": "cd469009461352f08965a3ec9946f6c708667386",
          "message": "[POSIX] Ensure compatibility version (#457)\n\nThis will upgrade any existing unversioned logs to have the current version. Towards #450\r\n\r\nRegenerated the testdata to confirm this works and checked in the .state\r\ndirectory, with lock files and all. We may want to remove lock files\r\nfrom disk when they are unlocked, but that's beyond the scope of this\r\nPR.",
          "timestamp": "2025-01-22T14:53:46Z",
          "tree_id": "7af50abe29b2c2604c3681fd2d6707259d07557c",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/cd469009461352f08965a3ec9946f6c708667386"
        },
        "date": 1737557712492,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4249160,
            "unit": "ns/op\t  700875 B/op\t   19676 allocs/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4249160,
            "unit": "ns/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700875,
            "unit": "B/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19676,
            "unit": "allocs/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6938,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "163443 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6938,
            "unit": "ns/op",
            "extra": "163443 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "163443 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "163443 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10338648 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.7,
            "unit": "ns/op",
            "extra": "10338648 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10338648 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10338648 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 351106,
            "unit": "ns/op\t  291274 B/op\t    3042 allocs/op",
            "extra": "3933 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 351106,
            "unit": "ns/op",
            "extra": "3933 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291274,
            "unit": "B/op",
            "extra": "3933 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3042,
            "unit": "allocs/op",
            "extra": "3933 times\n4 procs"
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
          "id": "754eacdeec0af1e1e9c9f6bfe4b9589358bbe1e7",
          "message": "[MySQL] Check that the DB schema is compatible (#458)\n\nThe version is written at schema creation, and checked on startup. If they are different, it fails. Towards #450.",
          "timestamp": "2025-01-23T10:59:44Z",
          "tree_id": "2800f2ed1f03a8a9d1efb82d4eabadf502b1a355",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/754eacdeec0af1e1e9c9f6bfe4b9589358bbe1e7"
        },
        "date": 1737630026436,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4606535,
            "unit": "ns/op\t  700810 B/op\t   19677 allocs/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4606535,
            "unit": "ns/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700810,
            "unit": "B/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19677,
            "unit": "allocs/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6935,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "183234 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6935,
            "unit": "ns/op",
            "extra": "183234 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "183234 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "183234 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 113.6,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10681069 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 113.6,
            "unit": "ns/op",
            "extra": "10681069 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10681069 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10681069 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 330178,
            "unit": "ns/op\t  291048 B/op\t    3041 allocs/op",
            "extra": "4308 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 330178,
            "unit": "ns/op",
            "extra": "4308 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291048,
            "unit": "B/op",
            "extra": "4308 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3041,
            "unit": "allocs/op",
            "extra": "4308 times\n4 procs"
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
          "id": "cc87631916c95732771ac6eaa7ae5e12893e3bc0",
          "message": "[Cleanup] Rename benchmark actions (#460)\n\nThese two actions had the same name in the GH actions tab and that makes it hard to glance at where the problem is, and communicate about it. Unique naming is good.",
          "timestamp": "2025-01-23T14:04:03Z",
          "tree_id": "7e0c05ebeccafc3314b12331b416c23b308f8777",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/cc87631916c95732771ac6eaa7ae5e12893e3bc0"
        },
        "date": 1737641088342,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4712542,
            "unit": "ns/op\t  699791 B/op\t   19664 allocs/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4712542,
            "unit": "ns/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 699791,
            "unit": "B/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19664,
            "unit": "allocs/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6921,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "157605 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6921,
            "unit": "ns/op",
            "extra": "157605 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "157605 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "157605 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 114.1,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10176067 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 114.1,
            "unit": "ns/op",
            "extra": "10176067 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10176067 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10176067 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 355638,
            "unit": "ns/op\t  290009 B/op\t    3032 allocs/op",
            "extra": "3784 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 355638,
            "unit": "ns/op",
            "extra": "3784 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290009,
            "unit": "B/op",
            "extra": "3784 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3784 times\n4 procs"
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
          "id": "c78d54b58ba7f595ac59dfa409bd30139b4b6b58",
          "message": "[AWS] Check compatibility version of the DB schema on startup (#459)\n\nThe version is written at schema creation, and checked on startup. If they are different, it fails. Towards #450.",
          "timestamp": "2025-01-23T14:08:09Z",
          "tree_id": "a36e424d4d0942f0c4d3aa5692bf44b91bb79d17",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/c78d54b58ba7f595ac59dfa409bd30139b4b6b58"
        },
        "date": 1737641332919,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 5035778,
            "unit": "ns/op\t  702469 B/op\t   19694 allocs/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 5035778,
            "unit": "ns/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702469,
            "unit": "B/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19694,
            "unit": "allocs/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7105,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "167548 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7105,
            "unit": "ns/op",
            "extra": "167548 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "167548 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "167548 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.5,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10331688 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.5,
            "unit": "ns/op",
            "extra": "10331688 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10331688 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10331688 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 360998,
            "unit": "ns/op\t  290146 B/op\t    3034 allocs/op",
            "extra": "3824 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 360998,
            "unit": "ns/op",
            "extra": "3824 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290146,
            "unit": "B/op",
            "extra": "3824 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3034,
            "unit": "allocs/op",
            "extra": "3824 times\n4 procs"
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
          "id": "d8c3474bbc7675a7260c465c78facba28ba8755a",
          "message": "Fix decorators (#462)",
          "timestamp": "2025-01-27T10:39:31Z",
          "tree_id": "ed409a61654388cc2129a2d2e2ea8ac581d4d521",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/d8c3474bbc7675a7260c465c78facba28ba8755a"
        },
        "date": 1737974903164,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4781487,
            "unit": "ns/op\t  700699 B/op\t   19673 allocs/op",
            "extra": "924 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4781487,
            "unit": "ns/op",
            "extra": "924 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700699,
            "unit": "B/op",
            "extra": "924 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19673,
            "unit": "allocs/op",
            "extra": "924 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8401,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "134785 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8401,
            "unit": "ns/op",
            "extra": "134785 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "134785 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "134785 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 113.9,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10586583 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 113.9,
            "unit": "ns/op",
            "extra": "10586583 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10586583 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10586583 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 354745,
            "unit": "ns/op\t  291167 B/op\t    3041 allocs/op",
            "extra": "3925 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 354745,
            "unit": "ns/op",
            "extra": "3925 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291167,
            "unit": "B/op",
            "extra": "3925 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3041,
            "unit": "allocs/op",
            "extra": "3925 times\n4 procs"
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
          "id": "838cd0cff1d3ae84108d68938da690f263e0a829",
          "message": "Bump the all-go-deps group with 7 updates (#465)\n\nBumps the all-go-deps group with 7 updates:\n\n| Package | From | To |\n| --- | --- | --- |\n| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.33.0` | `1.34.0` |\n| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.1` | `1.29.2` |\n| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.54` | `1.17.55` |\n| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.73.2` | `1.74.1` |\n| [github.com/aws/smithy-go](https://github.com/aws/smithy-go) | `1.22.1` | `1.22.2` |\n| [google.golang.org/api](https://github.com/googleapis/google-api-go-client) | `0.217.0` | `0.218.0` |\n| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.69.4` | `1.70.0` |\n\n\nUpdates `github.com/aws/aws-sdk-go-v2` from 1.33.0 to 1.34.0\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/v1.33.0...v1.34.0)\n\nUpdates `github.com/aws/aws-sdk-go-v2/config` from 1.29.1 to 1.29.2\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.1...config/v1.29.2)\n\nUpdates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.54 to 1.17.55\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.54...credentials/v1.17.55)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.73.2 to 1.74.1\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.73.2...service/s3/v1.74.1)\n\nUpdates `github.com/aws/smithy-go` from 1.22.1 to 1.22.2\n- [Release notes](https://github.com/aws/smithy-go/releases)\n- [Changelog](https://github.com/aws/smithy-go/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/aws/smithy-go/compare/v1.22.1...v1.22.2)\n\nUpdates `google.golang.org/api` from 0.217.0 to 0.218.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.217.0...v0.218.0)\n\nUpdates `google.golang.org/grpc` from 1.69.4 to 1.70.0\n- [Release notes](https://github.com/grpc/grpc-go/releases)\n- [Commits](https://github.com/grpc/grpc-go/compare/v1.69.4...v1.70.0)\n\n---\nupdated-dependencies:\n- dependency-name: github.com/aws/aws-sdk-go-v2\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/config\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/credentials\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/smithy-go\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/grpc\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-01-27T21:55:50Z",
          "tree_id": "70b984152845b6691f0f214e1fd89a52d1e93d61",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/838cd0cff1d3ae84108d68938da690f263e0a829"
        },
        "date": 1738015035844,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4437006,
            "unit": "ns/op\t  702463 B/op\t   19691 allocs/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4437006,
            "unit": "ns/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702463,
            "unit": "B/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19691,
            "unit": "allocs/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8343,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "160027 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8343,
            "unit": "ns/op",
            "extra": "160027 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "160027 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "160027 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.3,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10401956 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.3,
            "unit": "ns/op",
            "extra": "10401956 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10401956 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10401956 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 351381,
            "unit": "ns/op\t  290713 B/op\t    3038 allocs/op",
            "extra": "3895 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 351381,
            "unit": "ns/op",
            "extra": "3895 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290713,
            "unit": "B/op",
            "extra": "3895 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3038,
            "unit": "allocs/op",
            "extra": "3895 times\n4 procs"
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
          "id": "2e2f0aa400227d887bb666a3159f01a34a9e8fae",
          "message": "Bump the all-gha-deps group across 1 directory with 2 updates (#466)\n\n* Bump the all-gha-deps group across 1 directory with 2 updates\n\nBumps the all-gha-deps group with 2 updates in the / directory: [actions/setup-go](https://github.com/actions/setup-go) and [github/codeql-action](https://github.com/github/codeql-action).\n\n\nUpdates `actions/setup-go` from 5.2.0 to 5.3.0\n- [Release notes](https://github.com/actions/setup-go/releases)\n- [Commits](https://github.com/actions/setup-go/compare/3041bf56c941b39c61721a86cd11f3bb1338122a...f111f3307d8850f501ac008e886eec1fd1932a34)\n\nUpdates `github/codeql-action` from 3.28.1 to 3.28.6\n- [Release notes](https://github.com/github/codeql-action/releases)\n- [Changelog](https://github.com/github/codeql-action/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/github/codeql-action/compare/b6a472f63d85b9c78a3ac5e89422239fc15e9b3c...17a820bf2e43b47be2c72b39cc905417bc1ab6d0)\n\n---\nupdated-dependencies:\n- dependency-name: actions/setup-go\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-gha-deps\n- dependency-name: github/codeql-action\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\n\n* update comments\n\n---------\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>\nCo-authored-by: Philippe Boneff <phb@google.com>",
          "timestamp": "2025-02-03T17:36:54Z",
          "tree_id": "96e8e29a06c457315c2d502a9f5193e50c9198b8",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/2e2f0aa400227d887bb666a3159f01a34a9e8fae"
        },
        "date": 1738604298252,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4490851,
            "unit": "ns/op\t  700402 B/op\t   19675 allocs/op",
            "extra": "282 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4490851,
            "unit": "ns/op",
            "extra": "282 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700402,
            "unit": "B/op",
            "extra": "282 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19675,
            "unit": "allocs/op",
            "extra": "282 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7879,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "177154 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7879,
            "unit": "ns/op",
            "extra": "177154 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "177154 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "177154 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 113.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10608285 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 113.7,
            "unit": "ns/op",
            "extra": "10608285 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10608285 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10608285 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 347616,
            "unit": "ns/op\t  290169 B/op\t    3034 allocs/op",
            "extra": "3831 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 347616,
            "unit": "ns/op",
            "extra": "3831 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290169,
            "unit": "B/op",
            "extra": "3831 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3034,
            "unit": "allocs/op",
            "extra": "3831 times\n4 procs"
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
          "id": "ac56ce55cabcd4cecbb9c2a4f96239edab4921f2",
          "message": "Bump the all-go-deps group with 6 updates (#468)\n\nBumps the all-go-deps group with 6 updates:\n\n| Package | From | To |\n| --- | --- | --- |\n| [cloud.google.com/go/spanner](https://github.com/googleapis/google-cloud-go) | `1.73.0` | `1.75.0` |\n| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.34.0` | `1.36.0` |\n| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.2` | `1.29.4` |\n| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.55` | `1.17.57` |\n| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.74.1` | `1.75.2` |\n| [google.golang.org/api](https://github.com/googleapis/google-api-go-client) | `0.218.0` | `0.219.0` |\n\n\nUpdates `cloud.google.com/go/spanner` from 1.73.0 to 1.75.0\n- [Release notes](https://github.com/googleapis/google-cloud-go/releases)\n- [Changelog](https://github.com/googleapis/google-cloud-go/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-cloud-go/compare/spanner/v1.73.0...spanner/v1.75.0)\n\nUpdates `github.com/aws/aws-sdk-go-v2` from 1.34.0 to 1.36.0\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/v1.34.0...v1.36.0)\n\nUpdates `github.com/aws/aws-sdk-go-v2/config` from 1.29.2 to 1.29.4\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.2...config/v1.29.4)\n\nUpdates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.55 to 1.17.57\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.55...credentials/v1.17.57)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.74.1 to 1.75.2\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.74.1...service/s3/v1.75.2)\n\nUpdates `google.golang.org/api` from 0.218.0 to 0.219.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.218.0...v0.219.0)\n\n---\nupdated-dependencies:\n- dependency-name: cloud.google.com/go/spanner\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/config\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/credentials\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-02-03T21:15:53Z",
          "tree_id": "91d0cda294a08506ce6890746f0cb6f49cab5a17",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/ac56ce55cabcd4cecbb9c2a4f96239edab4921f2"
        },
        "date": 1738617436906,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4568048,
            "unit": "ns/op\t  700880 B/op\t   19679 allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4568048,
            "unit": "ns/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700880,
            "unit": "B/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19679,
            "unit": "allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8016,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "211338 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8016,
            "unit": "ns/op",
            "extra": "211338 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "211338 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "211338 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 115.1,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10412548 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 115.1,
            "unit": "ns/op",
            "extra": "10412548 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10412548 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10412548 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 349710,
            "unit": "ns/op\t  289956 B/op\t    3032 allocs/op",
            "extra": "3732 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 349710,
            "unit": "ns/op",
            "extra": "3732 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 289956,
            "unit": "B/op",
            "extra": "3732 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3732 times\n4 procs"
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
          "id": "2d6f5e360efa12726eb11947a9ed12853e336ab4",
          "message": "Bump the all-gha-deps group with 3 updates (#467)\n\nBumps the all-gha-deps group with 3 updates: [aws-actions/configure-aws-credentials](https://github.com/aws-actions/configure-aws-credentials), [gruntwork-io/terragrunt-action](https://github.com/gruntwork-io/terragrunt-action) and [github/codeql-action](https://github.com/github/codeql-action).\n\n\nUpdates `aws-actions/configure-aws-credentials` from 4.0.2 to 4.0.3\n- [Release notes](https://github.com/aws-actions/configure-aws-credentials/releases)\n- [Changelog](https://github.com/aws-actions/configure-aws-credentials/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/aws-actions/configure-aws-credentials/compare/e3dd6a429d7300a6a4c196c26e071d42e0343502...4fc4975a852c8cd99761e2de1f4ba73402e44dd9)\n\nUpdates `gruntwork-io/terragrunt-action` from 2.1.4 to 2.1.5\n- [Release notes](https://github.com/gruntwork-io/terragrunt-action/releases)\n- [Commits](https://github.com/gruntwork-io/terragrunt-action/compare/cf355638dbd066b6a70eb01f31ad9293870a0e34...9559e51d05873b0ea467c42bbabcb5c067642ccc)\n\nUpdates `github/codeql-action` from 3.28.6 to 3.28.8\n- [Release notes](https://github.com/github/codeql-action/releases)\n- [Changelog](https://github.com/github/codeql-action/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/github/codeql-action/compare/17a820bf2e43b47be2c72b39cc905417bc1ab6d0...dd746615b3b9d728a6a37ca2045b68ca76d4841a)\n\n---\nupdated-dependencies:\n- dependency-name: aws-actions/configure-aws-credentials\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n- dependency-name: gruntwork-io/terragrunt-action\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n- dependency-name: github/codeql-action\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-02-03T22:14:26Z",
          "tree_id": "e4a1e7aa209955c2788162e145e5c46e0e5a44aa",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/2d6f5e360efa12726eb11947a9ed12853e336ab4"
        },
        "date": 1738620909830,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4936651,
            "unit": "ns/op\t  698986 B/op\t   19657 allocs/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4936651,
            "unit": "ns/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 698986,
            "unit": "B/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19657,
            "unit": "allocs/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6979,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "188101 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6979,
            "unit": "ns/op",
            "extra": "188101 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "188101 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "188101 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 113.9,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10524630 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 113.9,
            "unit": "ns/op",
            "extra": "10524630 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10524630 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10524630 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 349534,
            "unit": "ns/op\t  290407 B/op\t    3036 allocs/op",
            "extra": "4144 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 349534,
            "unit": "ns/op",
            "extra": "4144 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290407,
            "unit": "B/op",
            "extra": "4144 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "4144 times\n4 procs"
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
          "id": "572101c67119005c5cf70d1ec903adb63903fb84",
          "message": "Factor out GCP DB init for later reuse (#477)",
          "timestamp": "2025-02-10T15:48:36Z",
          "tree_id": "c852c56081013bbdc7b761acf62a7e0345d6db44",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/572101c67119005c5cf70d1ec903adb63903fb84"
        },
        "date": 1739202704452,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 5071999,
            "unit": "ns/op\t  704682 B/op\t   19713 allocs/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 5071999,
            "unit": "ns/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 704682,
            "unit": "B/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19713,
            "unit": "allocs/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7598,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "266889 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7598,
            "unit": "ns/op",
            "extra": "266889 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "266889 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "266889 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.5,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10259131 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.5,
            "unit": "ns/op",
            "extra": "10259131 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10259131 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10259131 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 350432,
            "unit": "ns/op\t  290293 B/op\t    3034 allocs/op",
            "extra": "3674 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 350432,
            "unit": "ns/op",
            "extra": "3674 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290293,
            "unit": "B/op",
            "extra": "3674 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3034,
            "unit": "allocs/op",
            "extra": "3674 times\n4 procs"
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
          "id": "55e717a9efcd39322f41f089a5eed52debe79876",
          "message": "Bump the all-gha-deps group with 2 updates (#478)\n\nBumps the all-gha-deps group with 2 updates: [github/codeql-action](https://github.com/github/codeql-action) and [golangci/golangci-lint-action](https://github.com/golangci/golangci-lint-action).\n\n\nUpdates `github/codeql-action` from 3.28.8 to 3.28.9\n- [Release notes](https://github.com/github/codeql-action/releases)\n- [Changelog](https://github.com/github/codeql-action/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/github/codeql-action/compare/dd746615b3b9d728a6a37ca2045b68ca76d4841a...9e8d0789d4a0fa9ceb6b1738f7e269594bdd67f0)\n\nUpdates `golangci/golangci-lint-action` from 6.2.0 to 6.3.2\n- [Release notes](https://github.com/golangci/golangci-lint-action/releases)\n- [Commits](https://github.com/golangci/golangci-lint-action/compare/ec5d18412c0aeab7936cb16880d708ba2a64e1ae...051d91933864810ecd5e2ea2cfd98f6a5bca5347)\n\n---\nupdated-dependencies:\n- dependency-name: github/codeql-action\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n- dependency-name: golangci/golangci-lint-action\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-02-10T20:40:50Z",
          "tree_id": "9eeeefd0f8a2271a3c03a8c9441449260aa13be4",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/55e717a9efcd39322f41f089a5eed52debe79876"
        },
        "date": 1739220102384,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4835073,
            "unit": "ns/op\t  702655 B/op\t   19694 allocs/op",
            "extra": "254 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4835073,
            "unit": "ns/op",
            "extra": "254 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702655,
            "unit": "B/op",
            "extra": "254 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19694,
            "unit": "allocs/op",
            "extra": "254 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7156,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "230330 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7156,
            "unit": "ns/op",
            "extra": "230330 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "230330 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "230330 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.9,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10418798 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.9,
            "unit": "ns/op",
            "extra": "10418798 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10418798 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10418798 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 349842,
            "unit": "ns/op\t  290659 B/op\t    3038 allocs/op",
            "extra": "3890 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 349842,
            "unit": "ns/op",
            "extra": "3890 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290659,
            "unit": "B/op",
            "extra": "3890 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3038,
            "unit": "allocs/op",
            "extra": "3890 times\n4 procs"
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
          "id": "c8b81c87ed45c587300c750db30cc68a6b2c7ad1",
          "message": "Bump the all-go-deps group with 9 updates (#479)\n\nBumps the all-go-deps group with 9 updates:\n\n| Package | From | To |\n| --- | --- | --- |\n| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.36.0` | `1.36.1` |\n| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.4` | `1.29.6` |\n| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.57` | `1.17.59` |\n| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.75.2` | `1.76.1` |\n| [golang.org/x/mod](https://github.com/golang/mod) | `0.22.0` | `0.23.0` |\n| [google.golang.org/api](https://github.com/googleapis/google-api-go-client) | `0.219.0` | `0.220.0` |\n| [golang.org/x/crypto](https://github.com/golang/crypto) | `0.32.0` | `0.33.0` |\n| [golang.org/x/net](https://github.com/golang/net) | `0.34.0` | `0.35.0` |\n| [golang.org/x/sync](https://github.com/golang/sync) | `0.10.0` | `0.11.0` |\n\n\nUpdates `github.com/aws/aws-sdk-go-v2` from 1.36.0 to 1.36.1\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/v1.36.0...v1.36.1)\n\nUpdates `github.com/aws/aws-sdk-go-v2/config` from 1.29.4 to 1.29.6\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.4...config/v1.29.6)\n\nUpdates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.57 to 1.17.59\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.57...credentials/v1.17.59)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.75.2 to 1.76.1\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.75.2...service/s3/v1.76.1)\n\nUpdates `golang.org/x/mod` from 0.22.0 to 0.23.0\n- [Commits](https://github.com/golang/mod/compare/v0.22.0...v0.23.0)\n\nUpdates `google.golang.org/api` from 0.219.0 to 0.220.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.219.0...v0.220.0)\n\nUpdates `golang.org/x/crypto` from 0.32.0 to 0.33.0\n- [Commits](https://github.com/golang/crypto/compare/v0.32.0...v0.33.0)\n\nUpdates `golang.org/x/net` from 0.34.0 to 0.35.0\n- [Commits](https://github.com/golang/net/compare/v0.34.0...v0.35.0)\n\nUpdates `golang.org/x/sync` from 0.10.0 to 0.11.0\n- [Commits](https://github.com/golang/sync/compare/v0.10.0...v0.11.0)\n\n---\nupdated-dependencies:\n- dependency-name: github.com/aws/aws-sdk-go-v2\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/config\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/credentials\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/mod\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/crypto\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/net\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/sync\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-02-11T10:01:17Z",
          "tree_id": "a1144578b982304ee65700bdaff0cc6769694bcd",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/c8b81c87ed45c587300c750db30cc68a6b2c7ad1"
        },
        "date": 1739268162532,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4552582,
            "unit": "ns/op\t  700192 B/op\t   19670 allocs/op",
            "extra": "289 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4552582,
            "unit": "ns/op",
            "extra": "289 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700192,
            "unit": "B/op",
            "extra": "289 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19670,
            "unit": "allocs/op",
            "extra": "289 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7782,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "157586 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7782,
            "unit": "ns/op",
            "extra": "157586 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "157586 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "157586 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 119.3,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10139617 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 119.3,
            "unit": "ns/op",
            "extra": "10139617 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10139617 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10139617 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 358281,
            "unit": "ns/op\t  289970 B/op\t    3032 allocs/op",
            "extra": "3727 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 358281,
            "unit": "ns/op",
            "extra": "3727 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 289970,
            "unit": "B/op",
            "extra": "3727 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3727 times\n4 procs"
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
          "id": "5f70186ff794c8293c04006d6fe9ce9811ddcc67",
          "message": "Replace GCP dedup with follower-style implementation (#475)",
          "timestamp": "2025-02-11T17:29:05Z",
          "tree_id": "5ef628e44443fd54531ea33c3aceb93ca8c25468",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/5f70186ff794c8293c04006d6fe9ce9811ddcc67"
        },
        "date": 1739295383502,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1182702,
            "unit": "ns/op\t  689426 B/op\t   19560 allocs/op",
            "extra": "1006 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1182702,
            "unit": "ns/op",
            "extra": "1006 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689426,
            "unit": "B/op",
            "extra": "1006 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19560,
            "unit": "allocs/op",
            "extra": "1006 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1818,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "616456 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1818,
            "unit": "ns/op",
            "extra": "616456 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "616456 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "616456 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 115.5,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10261608 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 115.5,
            "unit": "ns/op",
            "extra": "10261608 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10261608 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10261608 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 348666,
            "unit": "ns/op\t  290521 B/op\t    3037 allocs/op",
            "extra": "3880 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 348666,
            "unit": "ns/op",
            "extra": "3880 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290521,
            "unit": "B/op",
            "extra": "3880 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3037,
            "unit": "allocs/op",
            "extra": "3880 times\n4 procs"
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
          "id": "f90de013bf71a4edc29c9df726b275b017c9bff4",
          "message": "Tighten up POSIX file operations (#481)",
          "timestamp": "2025-02-11T18:11:18Z",
          "tree_id": "9347f014ee9b12d75c7e2b779dc709ca75bab66e",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/f90de013bf71a4edc29c9df726b275b017c9bff4"
        },
        "date": 1739297525074,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1198352,
            "unit": "ns/op\t  690056 B/op\t   19566 allocs/op",
            "extra": "997 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1198352,
            "unit": "ns/op",
            "extra": "997 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 690056,
            "unit": "B/op",
            "extra": "997 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19566,
            "unit": "allocs/op",
            "extra": "997 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1803,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "676288 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1803,
            "unit": "ns/op",
            "extra": "676288 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "676288 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "676288 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 119.4,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10196937 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 119.4,
            "unit": "ns/op",
            "extra": "10196937 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10196937 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10196937 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 358443,
            "unit": "ns/op\t  290949 B/op\t    3040 allocs/op",
            "extra": "3957 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 358443,
            "unit": "ns/op",
            "extra": "3957 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290949,
            "unit": "B/op",
            "extra": "3957 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3040,
            "unit": "allocs/op",
            "extra": "3957 times\n4 procs"
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
          "id": "84fabd2d9612204fbff713d5f898ac3b9c1cd82e",
          "message": "[API] Refactoring to make appender API future-proof (#461)\n\nThe inflexibility of this was noticed while looking at options to fix #341. This PR includes a comment indicating that in future we want to add a Shutdown method, but more interesting are the other changes that:\r\n  a) move the dedupe decorators behind an options, and\r\n  b) making Appender more than just a function to allow new methods (such as Shutdown) to be added in the future.",
          "timestamp": "2025-02-12T11:32:40Z",
          "tree_id": "493b4fa88105b3cabd7952cc4fa1f386141f7e77",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/84fabd2d9612204fbff713d5f898ac3b9c1cd82e"
        },
        "date": 1739360005157,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1197611,
            "unit": "ns/op\t  689964 B/op\t   19565 allocs/op",
            "extra": "1008 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1197611,
            "unit": "ns/op",
            "extra": "1008 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689964,
            "unit": "B/op",
            "extra": "1008 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19565,
            "unit": "allocs/op",
            "extra": "1008 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1796,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "616708 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1796,
            "unit": "ns/op",
            "extra": "616708 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "616708 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "616708 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.2,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10173142 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.2,
            "unit": "ns/op",
            "extra": "10173142 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10173142 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10173142 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 345646,
            "unit": "ns/op\t  290125 B/op\t    3033 allocs/op",
            "extra": "3820 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 345646,
            "unit": "ns/op",
            "extra": "3820 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290125,
            "unit": "B/op",
            "extra": "3820 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3033,
            "unit": "allocs/op",
            "extra": "3820 times\n4 procs"
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
          "id": "8c5dff73b2725e5d16ddd360421307dbcd33ecc4",
          "message": "Migrate storage implementations to Appender lifecycle model (#482)",
          "timestamp": "2025-02-13T16:44:07Z",
          "tree_id": "30e2f1dd9985486bbf33f289b59a839383fba57c",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/8c5dff73b2725e5d16ddd360421307dbcd33ecc4"
        },
        "date": 1739465134617,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4657077,
            "unit": "ns/op\t  703836 B/op\t   19708 allocs/op",
            "extra": "265 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4657077,
            "unit": "ns/op",
            "extra": "265 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 703836,
            "unit": "B/op",
            "extra": "265 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19708,
            "unit": "allocs/op",
            "extra": "265 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7933,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "229640 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7933,
            "unit": "ns/op",
            "extra": "229640 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "229640 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "229640 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 118.4,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10002728 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 118.4,
            "unit": "ns/op",
            "extra": "10002728 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10002728 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10002728 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 356274,
            "unit": "ns/op\t  290158 B/op\t    3034 allocs/op",
            "extra": "3829 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 356274,
            "unit": "ns/op",
            "extra": "3829 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290158,
            "unit": "B/op",
            "extra": "3829 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3034,
            "unit": "allocs/op",
            "extra": "3829 times\n4 procs"
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
          "id": "5f7ca097f44a91c0db90ca1dc71f2a70574fc6f0",
          "message": "Update CT support to match lifecycle change (#485)",
          "timestamp": "2025-02-13T17:17:19Z",
          "tree_id": "941bc8a3147661a046e071a6861c429cc890f53e",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/5f7ca097f44a91c0db90ca1dc71f2a70574fc6f0"
        },
        "date": 1739467090348,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4695451,
            "unit": "ns/op\t  703379 B/op\t   19703 allocs/op",
            "extra": "241 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4695451,
            "unit": "ns/op",
            "extra": "241 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 703379,
            "unit": "B/op",
            "extra": "241 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19703,
            "unit": "allocs/op",
            "extra": "241 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7119,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "180948 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7119,
            "unit": "ns/op",
            "extra": "180948 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "180948 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "180948 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.6,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10461619 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.6,
            "unit": "ns/op",
            "extra": "10461619 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10461619 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10461619 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 354645,
            "unit": "ns/op\t  290314 B/op\t    3035 allocs/op",
            "extra": "4084 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 354645,
            "unit": "ns/op",
            "extra": "4084 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290314,
            "unit": "B/op",
            "extra": "4084 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3035,
            "unit": "allocs/op",
            "extra": "4084 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "hblauzvern@google.com",
            "name": "Hayden B",
            "username": "haydentherapper"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6d36a7e7714cda791de0fb1ef5f54f2cfc1e0fb0",
          "message": "Lock initialise to prevent concurrent initialisation  (#480)\n\n* Lock initialise to prevent concurrent initialisation\r\n\r\nIf two tasks attempted to initialise an empty checkpoint, and one task\r\nstalled while the other began accepting traffic, then the tree state and\r\ncheckpoint would be invalidated by the stalled task. This guarantees\r\ninitialisation only once per log.",
          "timestamp": "2025-02-14T18:30:37Z",
          "tree_id": "6718f6654db7fe9dbedca7ed7eaa8fa943a2f63b",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/6d36a7e7714cda791de0fb1ef5f54f2cfc1e0fb0"
        },
        "date": 1739557883290,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4760041,
            "unit": "ns/op\t  703657 B/op\t   19706 allocs/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4760041,
            "unit": "ns/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 703657,
            "unit": "B/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19706,
            "unit": "allocs/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7196,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "167580 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7196,
            "unit": "ns/op",
            "extra": "167580 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "167580 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "167580 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.8,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10194333 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.8,
            "unit": "ns/op",
            "extra": "10194333 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10194333 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10194333 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 351204,
            "unit": "ns/op\t  290135 B/op\t    3033 allocs/op",
            "extra": "3819 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 351204,
            "unit": "ns/op",
            "extra": "3819 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290135,
            "unit": "B/op",
            "extra": "3819 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3033,
            "unit": "allocs/op",
            "extra": "3819 times\n4 procs"
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
          "id": "778e52bd003e983f1f7d63551154cb6dd4eba6d3",
          "message": "Fix missing error message for errors occuring every second (#487)",
          "timestamp": "2025-02-17T16:09:47Z",
          "tree_id": "f34d827da846b06122c9cff4c9ef2b9a78f167ea",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/778e52bd003e983f1f7d63551154cb6dd4eba6d3"
        },
        "date": 1739808637543,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4452855,
            "unit": "ns/op\t  700729 B/op\t   19672 allocs/op",
            "extra": "289 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4452855,
            "unit": "ns/op",
            "extra": "289 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700729,
            "unit": "B/op",
            "extra": "289 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19672,
            "unit": "allocs/op",
            "extra": "289 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7512,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "226735 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7512,
            "unit": "ns/op",
            "extra": "226735 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "226735 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "226735 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.3,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10377444 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.3,
            "unit": "ns/op",
            "extra": "10377444 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10377444 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10377444 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 352127,
            "unit": "ns/op\t  289978 B/op\t    3032 allocs/op",
            "extra": "3775 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 352127,
            "unit": "ns/op",
            "extra": "3775 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 289978,
            "unit": "B/op",
            "extra": "3775 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3775 times\n4 procs"
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
          "id": "33aa456c6641e79eb14ed167163454586ec5e752",
          "message": "Bump the all-gha-deps group with 2 updates (#489)\n\nBumps the all-gha-deps group with 2 updates: [aws-actions/configure-aws-credentials](https://github.com/aws-actions/configure-aws-credentials) and [golangci/golangci-lint-action](https://github.com/golangci/golangci-lint-action).\n\n\nUpdates `aws-actions/configure-aws-credentials` from 4.0.3 to 4.1.0\n- [Release notes](https://github.com/aws-actions/configure-aws-credentials/releases)\n- [Changelog](https://github.com/aws-actions/configure-aws-credentials/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/aws-actions/configure-aws-credentials/compare/4fc4975a852c8cd99761e2de1f4ba73402e44dd9...ececac1a45f3b08a01d2dd070d28d111c5fe6722)\n\nUpdates `golangci/golangci-lint-action` from 6.3.2 to 6.5.0\n- [Release notes](https://github.com/golangci/golangci-lint-action/releases)\n- [Commits](https://github.com/golangci/golangci-lint-action/compare/051d91933864810ecd5e2ea2cfd98f6a5bca5347...2226d7cb06a077cd73e56eedd38eecad18e5d837)\n\n---\nupdated-dependencies:\n- dependency-name: aws-actions/configure-aws-credentials\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-gha-deps\n- dependency-name: golangci/golangci-lint-action\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-02-17T21:48:08Z",
          "tree_id": "04c5eb772d2cf8f1985390cd0501cfc984d8fdc0",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/33aa456c6641e79eb14ed167163454586ec5e752"
        },
        "date": 1739828931698,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4516975,
            "unit": "ns/op\t  700597 B/op\t   19674 allocs/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4516975,
            "unit": "ns/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700597,
            "unit": "B/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19674,
            "unit": "allocs/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7288,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "172834 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7288,
            "unit": "ns/op",
            "extra": "172834 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "172834 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "172834 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 115.4,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10554400 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 115.4,
            "unit": "ns/op",
            "extra": "10554400 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10554400 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10554400 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 347177,
            "unit": "ns/op\t  290438 B/op\t    3036 allocs/op",
            "extra": "4147 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 347177,
            "unit": "ns/op",
            "extra": "4147 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290438,
            "unit": "B/op",
            "extra": "4147 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "4147 times\n4 procs"
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
          "id": "298a76863b4b876c0309f35db6f5eefc01acd10f",
          "message": "Bump the all-go-deps group with 2 updates (#490)\n\nBumps the all-go-deps group with 2 updates: [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) and [google.golang.org/api](https://github.com/googleapis/google-api-go-client).\n\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.76.1 to 1.77.0\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.76.1...service/s3/v1.77.0)\n\nUpdates `google.golang.org/api` from 0.220.0 to 0.221.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.220.0...v0.221.0)\n\n---\nupdated-dependencies:\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-02-17T21:53:37Z",
          "tree_id": "0e2cb03365dc5ac4c0adaddeccd28b111c1ac7c9",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/298a76863b4b876c0309f35db6f5eefc01acd10f"
        },
        "date": 1739829301423,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4446225,
            "unit": "ns/op\t  702134 B/op\t   19689 allocs/op",
            "extra": "289 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4446225,
            "unit": "ns/op",
            "extra": "289 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702134,
            "unit": "B/op",
            "extra": "289 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19689,
            "unit": "allocs/op",
            "extra": "289 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7664,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "141405 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7664,
            "unit": "ns/op",
            "extra": "141405 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "141405 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "141405 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 114.9,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10273159 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 114.9,
            "unit": "ns/op",
            "extra": "10273159 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10273159 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10273159 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 343848,
            "unit": "ns/op\t  290566 B/op\t    3038 allocs/op",
            "extra": "4173 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 343848,
            "unit": "ns/op",
            "extra": "4173 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290566,
            "unit": "B/op",
            "extra": "4173 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3038,
            "unit": "allocs/op",
            "extra": "4173 times\n4 procs"
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
          "id": "5165141dc2107d8a72f4f43ca4c588b29b67b9de",
          "message": "Defined witness policy configuration (#488)\n\nThis allows the required witnesses to be defined and the theshold\r\npolicies that apply within each group. Arbitrarily nested structures can\r\nbe built, each with different numbers of signatures.\r\n\r\nEach WitnessGroup provides the URLs at which the witness can be reached\r\nto perform witnessing, and a function that determines if the group is\r\nsatisfied.\r\n\r\nThis format is consistent with the only other known witness policy\r\nconfiguration format out there:\r\nhttps://git.glasklar.is/sigsum/core/sigsum-go/-/blob/main/doc/policy.md\r\n\r\nTowards #309.",
          "timestamp": "2025-02-18T17:05:12Z",
          "tree_id": "0502237e5a2727bf6ae72d8a646b8eadb17d602c",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/5165141dc2107d8a72f4f43ca4c588b29b67b9de"
        },
        "date": 1739898366916,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4447211,
            "unit": "ns/op\t  699190 B/op\t   19661 allocs/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4447211,
            "unit": "ns/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 699190,
            "unit": "B/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19661,
            "unit": "allocs/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 232975,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5430 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 232975,
            "unit": "ns/op",
            "extra": "5430 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5430 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5430 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8320,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "163221 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8320,
            "unit": "ns/op",
            "extra": "163221 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "163221 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "163221 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 114.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10394272 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 114.7,
            "unit": "ns/op",
            "extra": "10394272 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10394272 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10394272 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 347806,
            "unit": "ns/op\t  290346 B/op\t    3036 allocs/op",
            "extra": "4120 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 347806,
            "unit": "ns/op",
            "extra": "4120 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290346,
            "unit": "B/op",
            "extra": "4120 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "4120 times\n4 procs"
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
          "id": "709b4d3176df20415395e129cca9a77f46a95628",
          "message": "Bump the all-go-deps group with 11 updates (#492)\n\nBumps the all-go-deps group with 11 updates:\n\n| Package | From | To |\n| --- | --- | --- |\n| [cloud.google.com/go/spanner](https://github.com/googleapis/google-cloud-go) | `1.75.0` | `1.76.1` |\n| [github.com/avast/retry-go/v4](https://github.com/avast/retry-go) | `4.6.0` | `4.6.1` |\n| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.36.1` | `1.36.2` |\n| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.6` | `1.29.7` |\n| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.59` | `1.17.60` |\n| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.77.0` | `1.77.1` |\n| [github.com/aws/smithy-go](https://github.com/aws/smithy-go) | `1.22.2` | `1.22.3` |\n| [github.com/google/go-cmp](https://github.com/google/go-cmp) | `0.6.0` | `0.7.0` |\n| [google.golang.org/api](https://github.com/googleapis/google-api-go-client) | `0.221.0` | `0.222.0` |\n| [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) | `1.8.1` | `1.9.0` |\n| [golang.org/x/crypto](https://github.com/golang/crypto) | `0.33.0` | `0.35.0` |\n\n\nUpdates `cloud.google.com/go/spanner` from 1.75.0 to 1.76.1\n- [Release notes](https://github.com/googleapis/google-cloud-go/releases)\n- [Changelog](https://github.com/googleapis/google-cloud-go/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-cloud-go/compare/spanner/v1.75.0...spanner/v1.76.1)\n\nUpdates `github.com/avast/retry-go/v4` from 4.6.0 to 4.6.1\n- [Release notes](https://github.com/avast/retry-go/releases)\n- [Commits](https://github.com/avast/retry-go/compare/4.6.0...4.6.1)\n\nUpdates `github.com/aws/aws-sdk-go-v2` from 1.36.1 to 1.36.2\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/v1.36.1...v1.36.2)\n\nUpdates `github.com/aws/aws-sdk-go-v2/config` from 1.29.6 to 1.29.7\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.6...config/v1.29.7)\n\nUpdates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.59 to 1.17.60\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.59...credentials/v1.17.60)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.77.0 to 1.77.1\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.77.0...service/s3/v1.77.1)\n\nUpdates `github.com/aws/smithy-go` from 1.22.2 to 1.22.3\n- [Release notes](https://github.com/aws/smithy-go/releases)\n- [Changelog](https://github.com/aws/smithy-go/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/aws/smithy-go/compare/v1.22.2...v1.22.3)\n\nUpdates `github.com/google/go-cmp` from 0.6.0 to 0.7.0\n- [Release notes](https://github.com/google/go-cmp/releases)\n- [Commits](https://github.com/google/go-cmp/compare/v0.6.0...v0.7.0)\n\nUpdates `google.golang.org/api` from 0.221.0 to 0.222.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.221.0...v0.222.0)\n\nUpdates `github.com/go-sql-driver/mysql` from 1.8.1 to 1.9.0\n- [Release notes](https://github.com/go-sql-driver/mysql/releases)\n- [Changelog](https://github.com/go-sql-driver/mysql/blob/master/CHANGELOG.md)\n- [Commits](https://github.com/go-sql-driver/mysql/compare/v1.8.1...v1.9.0)\n\nUpdates `golang.org/x/crypto` from 0.33.0 to 0.35.0\n- [Commits](https://github.com/golang/crypto/compare/v0.33.0...v0.35.0)\n\n---\nupdated-dependencies:\n- dependency-name: cloud.google.com/go/spanner\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/avast/retry-go/v4\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/config\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/credentials\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/smithy-go\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/google/go-cmp\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/go-sql-driver/mysql\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/crypto\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-02-24T21:18:26Z",
          "tree_id": "69dcc473353370d1d73be95afd02b6fbf1f2d9be",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/709b4d3176df20415395e129cca9a77f46a95628"
        },
        "date": 1740431993911,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4708400,
            "unit": "ns/op\t  698197 B/op\t   19651 allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4708400,
            "unit": "ns/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 698197,
            "unit": "B/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19651,
            "unit": "allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 226559,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5847 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 226559,
            "unit": "ns/op",
            "extra": "5847 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5847 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5847 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6546,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "186492 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6546,
            "unit": "ns/op",
            "extra": "186492 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "186492 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "186492 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.2,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10337833 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.2,
            "unit": "ns/op",
            "extra": "10337833 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10337833 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10337833 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 352080,
            "unit": "ns/op\t  290213 B/op\t    3033 allocs/op",
            "extra": "3685 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 352080,
            "unit": "ns/op",
            "extra": "3685 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290213,
            "unit": "B/op",
            "extra": "3685 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3033,
            "unit": "allocs/op",
            "extra": "3685 times\n4 procs"
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
          "id": "0be5795e2bb761d85e84405d61d6ed4c5d1e18b8",
          "message": "Bump the all-gha-deps group with 3 updates (#493)\n\nBumps the all-gha-deps group with 3 updates: [github/codeql-action](https://github.com/github/codeql-action), [ossf/scorecard-action](https://github.com/ossf/scorecard-action) and [actions/upload-artifact](https://github.com/actions/upload-artifact).\n\n\nUpdates `github/codeql-action` from 3.28.9 to 3.28.10\n- [Release notes](https://github.com/github/codeql-action/releases)\n- [Changelog](https://github.com/github/codeql-action/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/github/codeql-action/compare/9e8d0789d4a0fa9ceb6b1738f7e269594bdd67f0...b56ba49b26e50535fa1e7f7db0f4f7b4bf65d80d)\n\nUpdates `ossf/scorecard-action` from 2.4.0 to 2.4.1\n- [Release notes](https://github.com/ossf/scorecard-action/releases)\n- [Changelog](https://github.com/ossf/scorecard-action/blob/main/RELEASE.md)\n- [Commits](https://github.com/ossf/scorecard-action/compare/62b2cac7ed8198b15735ed49ab1e5cf35480ba46...f49aabe0b5af0936a0987cfb85d86b75731b0186)\n\nUpdates `actions/upload-artifact` from 4.6.0 to 4.6.1\n- [Release notes](https://github.com/actions/upload-artifact/releases)\n- [Commits](https://github.com/actions/upload-artifact/compare/65c4c4a1ddee5b72f698fdd19549f0f0fb45cf08...4cec3d8aa04e39d1a68397de0c4cd6fb9dce8ec1)\n\n---\nupdated-dependencies:\n- dependency-name: github/codeql-action\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n- dependency-name: ossf/scorecard-action\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n- dependency-name: actions/upload-artifact\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-02-24T21:45:24Z",
          "tree_id": "3b3b1300677cdb27c4911bae53118093c90059db",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/0be5795e2bb761d85e84405d61d6ed4c5d1e18b8"
        },
        "date": 1740433571536,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 5139495,
            "unit": "ns/op\t  704702 B/op\t   19715 allocs/op",
            "extra": "237 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 5139495,
            "unit": "ns/op",
            "extra": "237 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 704702,
            "unit": "B/op",
            "extra": "237 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19715,
            "unit": "allocs/op",
            "extra": "237 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 237227,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5956 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 237227,
            "unit": "ns/op",
            "extra": "5956 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5956 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5956 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7406,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "205569 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7406,
            "unit": "ns/op",
            "extra": "205569 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "205569 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "205569 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 114.9,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10405788 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 114.9,
            "unit": "ns/op",
            "extra": "10405788 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10405788 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10405788 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 352612,
            "unit": "ns/op\t  290822 B/op\t    3039 allocs/op",
            "extra": "3973 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 352612,
            "unit": "ns/op",
            "extra": "3973 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290822,
            "unit": "B/op",
            "extra": "3973 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3039,
            "unit": "allocs/op",
            "extra": "3973 times\n4 procs"
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
          "id": "62f7bf5d259665543752d6506b998380fbd42902",
          "message": "Antispam WithOption (#491)",
          "timestamp": "2025-02-25T11:09:56Z",
          "tree_id": "f7fc246c7b34d3a5da720a639cd3b5dbc90673c1",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/62f7bf5d259665543752d6506b998380fbd42902"
        },
        "date": 1740481843622,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4545884,
            "unit": "ns/op\t  702893 B/op\t   19699 allocs/op",
            "extra": "274 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4545884,
            "unit": "ns/op",
            "extra": "274 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702893,
            "unit": "B/op",
            "extra": "274 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19699,
            "unit": "allocs/op",
            "extra": "274 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 232194,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5757 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 232194,
            "unit": "ns/op",
            "extra": "5757 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5757 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5757 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8266,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "197857 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8266,
            "unit": "ns/op",
            "extra": "197857 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "197857 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "197857 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10269508 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.7,
            "unit": "ns/op",
            "extra": "10269508 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10269508 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10269508 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 347520,
            "unit": "ns/op\t  291150 B/op\t    3041 allocs/op",
            "extra": "3927 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 347520,
            "unit": "ns/op",
            "extra": "3927 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291150,
            "unit": "B/op",
            "extra": "3927 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3041,
            "unit": "allocs/op",
            "extra": "3927 times\n4 procs"
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
          "id": "69b429d6a35e916d098ca06421823cc89d26ae20",
          "message": "Implementing witnessing code and API (#494)\n\nTowards #309.",
          "timestamp": "2025-02-26T10:28:08Z",
          "tree_id": "6deeef004f29d1e689e8c04b94bb475eacbfacf7",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/69b429d6a35e916d098ca06421823cc89d26ae20"
        },
        "date": 1740565737522,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4684109,
            "unit": "ns/op\t  700525 B/op\t   19673 allocs/op",
            "extra": "241 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4684109,
            "unit": "ns/op",
            "extra": "241 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700525,
            "unit": "B/op",
            "extra": "241 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19673,
            "unit": "allocs/op",
            "extra": "241 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 235107,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5529 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 235107,
            "unit": "ns/op",
            "extra": "5529 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5529 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5529 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7264,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "161364 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7264,
            "unit": "ns/op",
            "extra": "161364 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "161364 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "161364 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10329147 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.7,
            "unit": "ns/op",
            "extra": "10329147 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10329147 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10329147 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 348085,
            "unit": "ns/op\t  290716 B/op\t    3038 allocs/op",
            "extra": "3986 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 348085,
            "unit": "ns/op",
            "extra": "3986 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290716,
            "unit": "B/op",
            "extra": "3986 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3038,
            "unit": "allocs/op",
            "extra": "3986 times\n4 procs"
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
          "id": "ba5db800d23c420b6d3491755704a9bd87d1d8e8",
          "message": "Basic validation of witness signatures (#496)\n\nThis ensures that we always return a non-nil error if the policy was not satisfied, and adds some tests.\n\nTowards #309.",
          "timestamp": "2025-02-26T12:13:39Z",
          "tree_id": "2830937ee2db84bacabac77f21d6938c9280934b",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/ba5db800d23c420b6d3491755704a9bd87d1d8e8"
        },
        "date": 1740572064538,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4542258,
            "unit": "ns/op\t  700933 B/op\t   19678 allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4542258,
            "unit": "ns/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700933,
            "unit": "B/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19678,
            "unit": "allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 227413,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "4446 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 227413,
            "unit": "ns/op",
            "extra": "4446 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "4446 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "4446 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7644,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "176178 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7644,
            "unit": "ns/op",
            "extra": "176178 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "176178 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "176178 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 119,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10238550 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 119,
            "unit": "ns/op",
            "extra": "10238550 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10238550 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10238550 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 341451,
            "unit": "ns/op\t  290357 B/op\t    3036 allocs/op",
            "extra": "3864 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 341451,
            "unit": "ns/op",
            "extra": "3864 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290357,
            "unit": "B/op",
            "extra": "3864 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "3864 times\n4 procs"
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
          "id": "4c294a9a09d440bc22d438814aee75149e17af36",
          "message": "Add Migration lifecycle support (#495)",
          "timestamp": "2025-02-26T12:37:55Z",
          "tree_id": "a7026bee7a79b306f9dd97a8f5ed0d0316814b2c",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/4c294a9a09d440bc22d438814aee75149e17af36"
        },
        "date": 1740573525033,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4563113,
            "unit": "ns/op\t  701434 B/op\t   19684 allocs/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4563113,
            "unit": "ns/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 701434,
            "unit": "B/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19684,
            "unit": "allocs/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 225434,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "6110 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 225434,
            "unit": "ns/op",
            "extra": "6110 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "6110 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "6110 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6829,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "160518 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6829,
            "unit": "ns/op",
            "extra": "160518 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "160518 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "160518 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 120.1,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10341463 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 120.1,
            "unit": "ns/op",
            "extra": "10341463 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10341463 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10341463 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 353754,
            "unit": "ns/op\t  290081 B/op\t    3032 allocs/op",
            "extra": "3712 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 353754,
            "unit": "ns/op",
            "extra": "3712 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290081,
            "unit": "B/op",
            "extra": "3712 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3712 times\n4 procs"
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
          "id": "04da9c83c3848add0ee9515f9f825adec7192910",
          "message": "[Witnessing] Check responses for valid signatures (#500)\n\nThis now verifies the body of 200 responses. It checks that the note can\nbe verified using the signature, and then returns only the signature\nthat the log has a verifier for.\n\nThis means that witnesses that return a valid signature and then a load\nof other signatures will not be able to pollute the checkpoint with\nthese other signatures. On the other hand, it means we will need to\nconsider how to support witness key rotation in Tessera in the future.\nThere are a few ways to solve this, but I don't believe this approach\nblocks any of them.",
          "timestamp": "2025-02-26T15:20:23Z",
          "tree_id": "a288c34df7e0710c108e78211210a957a2358ef3",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/04da9c83c3848add0ee9515f9f825adec7192910"
        },
        "date": 1740583268535,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4568836,
            "unit": "ns/op\t  700623 B/op\t   19675 allocs/op",
            "extra": "244 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4568836,
            "unit": "ns/op",
            "extra": "244 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700623,
            "unit": "B/op",
            "extra": "244 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19675,
            "unit": "allocs/op",
            "extra": "244 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 225976,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5044 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 225976,
            "unit": "ns/op",
            "extra": "5044 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5044 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5044 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7807,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "155664 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7807,
            "unit": "ns/op",
            "extra": "155664 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "155664 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "155664 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 115.1,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10398758 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 115.1,
            "unit": "ns/op",
            "extra": "10398758 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10398758 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10398758 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 355555,
            "unit": "ns/op\t  290436 B/op\t    3037 allocs/op",
            "extra": "3873 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 355555,
            "unit": "ns/op",
            "extra": "3873 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290436,
            "unit": "B/op",
            "extra": "3873 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3037,
            "unit": "allocs/op",
            "extra": "3873 times\n4 procs"
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
          "id": "cd604aa607bd745affeca1cd5f35a3f4561d9694",
          "message": "Tidy gcp.go (#499)",
          "timestamp": "2025-02-26T15:27:31Z",
          "tree_id": "93fd4b14588fdf61c725049347427d5145c9c634",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/cd604aa607bd745affeca1cd5f35a3f4561d9694"
        },
        "date": 1740583705954,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4842254,
            "unit": "ns/op\t  698784 B/op\t   19655 allocs/op",
            "extra": "223 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4842254,
            "unit": "ns/op",
            "extra": "223 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 698784,
            "unit": "B/op",
            "extra": "223 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19655,
            "unit": "allocs/op",
            "extra": "223 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 224341,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5160 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 224341,
            "unit": "ns/op",
            "extra": "5160 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5160 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5160 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7831,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "198250 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7831,
            "unit": "ns/op",
            "extra": "198250 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "198250 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "198250 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.4,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10312926 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.4,
            "unit": "ns/op",
            "extra": "10312926 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10312926 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10312926 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 358511,
            "unit": "ns/op\t  290002 B/op\t    3032 allocs/op",
            "extra": "3734 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 358511,
            "unit": "ns/op",
            "extra": "3734 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290002,
            "unit": "B/op",
            "extra": "3734 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3734 times\n4 procs"
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
          "id": "60d9efffcddfa73f4279a98dd33fca02a03505cf",
          "message": "Typed lifecycle options (#501)",
          "timestamp": "2025-02-26T16:39:29Z",
          "tree_id": "530bb3613ec2a94002f5eb071f1ccb4a7382e24a",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/60d9efffcddfa73f4279a98dd33fca02a03505cf"
        },
        "date": 1740588028181,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4570471,
            "unit": "ns/op\t  702484 B/op\t   19693 allocs/op",
            "extra": "225 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4570471,
            "unit": "ns/op",
            "extra": "225 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702484,
            "unit": "B/op",
            "extra": "225 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19693,
            "unit": "allocs/op",
            "extra": "225 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 229297,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5839 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 229297,
            "unit": "ns/op",
            "extra": "5839 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5839 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5839 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7584,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "163002 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7584,
            "unit": "ns/op",
            "extra": "163002 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "163002 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "163002 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.6,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10337000 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.6,
            "unit": "ns/op",
            "extra": "10337000 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10337000 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10337000 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 353702,
            "unit": "ns/op\t  289993 B/op\t    3032 allocs/op",
            "extra": "3770 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 353702,
            "unit": "ns/op",
            "extra": "3770 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 289993,
            "unit": "B/op",
            "extra": "3770 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3770 times\n4 procs"
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
          "id": "d423e02edceb5a81aa11d1dcd177fe5299ec679c",
          "message": "Migration lifecycle options (#502)",
          "timestamp": "2025-02-26T17:15:30Z",
          "tree_id": "cbc65ecff47a6cc1215d0285dcc686ac2b7e143b",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/d423e02edceb5a81aa11d1dcd177fe5299ec679c"
        },
        "date": 1740590180965,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4481004,
            "unit": "ns/op\t  700775 B/op\t   19678 allocs/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4481004,
            "unit": "ns/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700775,
            "unit": "B/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19678,
            "unit": "allocs/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 229509,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5336 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 229509,
            "unit": "ns/op",
            "extra": "5336 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5336 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5336 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8022,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "172604 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8022,
            "unit": "ns/op",
            "extra": "172604 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "172604 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "172604 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10263816 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.7,
            "unit": "ns/op",
            "extra": "10263816 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10263816 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10263816 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 348612,
            "unit": "ns/op\t  290099 B/op\t    3033 allocs/op",
            "extra": "3811 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 348612,
            "unit": "ns/op",
            "extra": "3811 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290099,
            "unit": "B/op",
            "extra": "3811 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3033,
            "unit": "allocs/op",
            "extra": "3811 times\n4 procs"
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
          "id": "e6151c5c069c348f0ef07d0402d685d21dc7d6ef",
          "message": "Split lifecycle code into separate files (#503)",
          "timestamp": "2025-02-27T10:41:04Z",
          "tree_id": "11c4fe50355b00e130787d66757693261530e3f6",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/e6151c5c069c348f0ef07d0402d685d21dc7d6ef"
        },
        "date": 1740652908920,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4357581,
            "unit": "ns/op\t  702917 B/op\t   19699 allocs/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4357581,
            "unit": "ns/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702917,
            "unit": "B/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19699,
            "unit": "allocs/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 226869,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5383 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 226869,
            "unit": "ns/op",
            "extra": "5383 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5383 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5383 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7707,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "166255 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7707,
            "unit": "ns/op",
            "extra": "166255 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "166255 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "166255 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116.8,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10465455 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116.8,
            "unit": "ns/op",
            "extra": "10465455 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10465455 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10465455 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 346611,
            "unit": "ns/op\t  290238 B/op\t    3034 allocs/op",
            "extra": "3837 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 346611,
            "unit": "ns/op",
            "extra": "3837 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290238,
            "unit": "B/op",
            "extra": "3837 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3034,
            "unit": "allocs/op",
            "extra": "3837 times\n4 procs"
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
          "id": "9ff9cb8873d9b5a4715d7dbd793242c0d3483cd7",
          "message": "Add antispam support to migration lifecycle (#504)",
          "timestamp": "2025-02-27T14:03:20Z",
          "tree_id": "26c981df4939e9592f57141387976a898c347a46",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/9ff9cb8873d9b5a4715d7dbd793242c0d3483cd7"
        },
        "date": 1740665048555,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4701186,
            "unit": "ns/op\t  706435 B/op\t   19733 allocs/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4701186,
            "unit": "ns/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 706435,
            "unit": "B/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19733,
            "unit": "allocs/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 240397,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5102 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 240397,
            "unit": "ns/op",
            "extra": "5102 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5102 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5102 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7400,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "173396 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7400,
            "unit": "ns/op",
            "extra": "173396 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "173396 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "173396 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 115.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10505916 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 115.7,
            "unit": "ns/op",
            "extra": "10505916 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10505916 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10505916 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 352173,
            "unit": "ns/op\t  290061 B/op\t    3033 allocs/op",
            "extra": "3802 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 352173,
            "unit": "ns/op",
            "extra": "3802 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290061,
            "unit": "B/op",
            "extra": "3802 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3033,
            "unit": "allocs/op",
            "extra": "3802 times\n4 procs"
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
          "id": "f275deb8d2ac21b243120a4da5db90a91f1189a0",
          "message": "[Hammer] support more verbose checkpoint logging (#505)\n\nIf the checkpoint size isn't updating, this logging will at least allow visibility that the main loop is still running.",
          "timestamp": "2025-02-27T15:41:02Z",
          "tree_id": "b76c0736921fc5629c4ce764ece7a8dc4c5d12d7",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/f275deb8d2ac21b243120a4da5db90a91f1189a0"
        },
        "date": 1740670906419,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4076134,
            "unit": "ns/op\t  697551 B/op\t   19643 allocs/op",
            "extra": "249 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4076134,
            "unit": "ns/op",
            "extra": "249 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 697551,
            "unit": "B/op",
            "extra": "249 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19643,
            "unit": "allocs/op",
            "extra": "249 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 224970,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5497 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 224970,
            "unit": "ns/op",
            "extra": "5497 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5497 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5497 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7848,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "193168 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7848,
            "unit": "ns/op",
            "extra": "193168 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "193168 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "193168 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.7,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10208935 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.7,
            "unit": "ns/op",
            "extra": "10208935 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10208935 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10208935 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 345342,
            "unit": "ns/op\t  290328 B/op\t    3036 allocs/op",
            "extra": "4117 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 345342,
            "unit": "ns/op",
            "extra": "4117 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290328,
            "unit": "B/op",
            "extra": "4117 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "4117 times\n4 procs"
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
          "id": "73e8a49557908e7774beed88fab996db4e5d8700",
          "message": "wrap errors from spanner calls inside txn (#498)",
          "timestamp": "2025-02-27T18:18:46Z",
          "tree_id": "c123fb9cae9decb47196c4d8a3ad87a05619d1f7",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/73e8a49557908e7774beed88fab996db4e5d8700"
        },
        "date": 1740680376129,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4583788,
            "unit": "ns/op\t  701051 B/op\t   19681 allocs/op",
            "extra": "242 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4583788,
            "unit": "ns/op",
            "extra": "242 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 701051,
            "unit": "B/op",
            "extra": "242 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19681,
            "unit": "allocs/op",
            "extra": "242 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 223961,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "6019 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 223961,
            "unit": "ns/op",
            "extra": "6019 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "6019 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "6019 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7668,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "171268 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7668,
            "unit": "ns/op",
            "extra": "171268 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "171268 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "171268 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117.4,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10167127 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117.4,
            "unit": "ns/op",
            "extra": "10167127 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10167127 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10167127 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 346942,
            "unit": "ns/op\t  289966 B/op\t    3032 allocs/op",
            "extra": "3744 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 346942,
            "unit": "ns/op",
            "extra": "3744 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 289966,
            "unit": "B/op",
            "extra": "3744 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3744 times\n4 procs"
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
          "id": "f55995cb1acaff9c64c95fec24ebccdecccd9c49",
          "message": "[MySQL] Add support for Migration lifecycle (#507)",
          "timestamp": "2025-02-28T17:12:09Z",
          "tree_id": "5c4f69aaa271b42b5fa385f0c74d5a5555ae13cd",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/f55995cb1acaff9c64c95fec24ebccdecccd9c49"
        },
        "date": 1740762780349,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4590048,
            "unit": "ns/op\t  696847 B/op\t   19637 allocs/op",
            "extra": "231 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4590048,
            "unit": "ns/op",
            "extra": "231 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 696847,
            "unit": "B/op",
            "extra": "231 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19637,
            "unit": "allocs/op",
            "extra": "231 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 220979,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5732 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 220979,
            "unit": "ns/op",
            "extra": "5732 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5732 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5732 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7650,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "193213 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7650,
            "unit": "ns/op",
            "extra": "193213 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "193213 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "193213 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 116,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10604286 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 116,
            "unit": "ns/op",
            "extra": "10604286 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10604286 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10604286 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 352300,
            "unit": "ns/op\t  290377 B/op\t    3034 allocs/op",
            "extra": "3663 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 352300,
            "unit": "ns/op",
            "extra": "3663 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290377,
            "unit": "B/op",
            "extra": "3663 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3034,
            "unit": "allocs/op",
            "extra": "3663 times\n4 procs"
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
          "id": "b56a8f74a821936b2a7a29026093e45ecffcf416",
          "message": "Fix witnessed size state (#510)\n\nPreviously the state was being lost. Even though this was unintentional, it worked, but involved constantly querying the witness as if it was at size 0.\n\nThis change makes the size sticky for the life of the process by correctly using pointers instead of copying structs. It also updates the size on success, which was missing earlier.",
          "timestamp": "2025-03-03T17:13:58Z",
          "tree_id": "cbba50a02dea62b867e050233df5ab4c1810923b",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/b56a8f74a821936b2a7a29026093e45ecffcf416"
        },
        "date": 1741022092052,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4683684,
            "unit": "ns/op\t  701510 B/op\t   19682 allocs/op",
            "extra": "242 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4683684,
            "unit": "ns/op",
            "extra": "242 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 701510,
            "unit": "B/op",
            "extra": "242 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19682,
            "unit": "allocs/op",
            "extra": "242 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 225345,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5709 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 225345,
            "unit": "ns/op",
            "extra": "5709 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5709 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5709 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7429,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "140793 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7429,
            "unit": "ns/op",
            "extra": "140793 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "140793 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "140793 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 118,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10140220 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 118,
            "unit": "ns/op",
            "extra": "10140220 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10140220 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10140220 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 353087,
            "unit": "ns/op\t  290039 B/op\t    3032 allocs/op",
            "extra": "3714 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 353087,
            "unit": "ns/op",
            "extra": "3714 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290039,
            "unit": "B/op",
            "extra": "3714 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3714 times\n4 procs"
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
          "id": "53c7e68c5269f03a81e92927452d7d0e15e0d71f",
          "message": "Bump the all-go-deps group with 5 updates (#513)\n\nBumps the all-go-deps group with 5 updates:\n\n| Package | From | To |\n| --- | --- | --- |\n| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.36.2` | `1.36.3` |\n| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.7` | `1.29.8` |\n| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.60` | `1.17.61` |\n| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.77.1` | `1.78.0` |\n| [google.golang.org/api](https://github.com/googleapis/google-api-go-client) | `0.222.0` | `0.223.0` |\n\n\nUpdates `github.com/aws/aws-sdk-go-v2` from 1.36.2 to 1.36.3\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/v1.36.2...v1.36.3)\n\nUpdates `github.com/aws/aws-sdk-go-v2/config` from 1.29.7 to 1.29.8\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.7...config/v1.29.8)\n\nUpdates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.60 to 1.17.61\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.60...credentials/v1.17.61)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.77.1 to 1.78.0\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.77.1...service/s3/v1.78.0)\n\nUpdates `google.golang.org/api` from 0.222.0 to 0.223.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.222.0...v0.223.0)\n\n---\nupdated-dependencies:\n- dependency-name: github.com/aws/aws-sdk-go-v2\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/config\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/credentials\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-03-03T20:46:47Z",
          "tree_id": "c41af6a4a496071d3a2c161ab278fd7fc2cc7e52",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/53c7e68c5269f03a81e92927452d7d0e15e0d71f"
        },
        "date": 1741034893139,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4740525,
            "unit": "ns/op\t  697916 B/op\t   19645 allocs/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4740525,
            "unit": "ns/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 697916,
            "unit": "B/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19645,
            "unit": "allocs/op",
            "extra": "261 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 238299,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5724 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 238299,
            "unit": "ns/op",
            "extra": "5724 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5724 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5724 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6835,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "172082 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6835,
            "unit": "ns/op",
            "extra": "172082 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "172082 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "172082 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 117,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10441986 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 117,
            "unit": "ns/op",
            "extra": "10441986 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10441986 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10441986 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 352446,
            "unit": "ns/op\t  290306 B/op\t    3035 allocs/op",
            "extra": "4060 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 352446,
            "unit": "ns/op",
            "extra": "4060 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290306,
            "unit": "B/op",
            "extra": "4060 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3035,
            "unit": "allocs/op",
            "extra": "4060 times\n4 procs"
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
          "id": "318ab44d27c1c4518b026616dc2b0eeb89b7ff57",
          "message": "[Client] First pass at tidying the tiles client (#506)\n\nThis PR cleans up the client implementation to make it more readable, maintainable, and safer to use.\n\n* Reordered methods\n\nGrouped types, functions, classes, and internal helpers in the more standard order.\n\n* Make LogStateTracker much more private\n\nOnly one of its many public properties was being used outside the class. Made a getter for this property, and made all fields private.\n\n* Make reads and writes to LST thread safe\n\nBefore there was a possibility of a race condition if multiple threads called Update. Such a race had the potential to leave the internal state inconsistent. It could have also allowed clients to see both sides of a split view from the tracker.\n\n* Removed CheckConsistency method\n\nThis method was encouraging people to do the wrong thing and build up a stack of checkpoints before checking them. This is bad; you should only ever have a trusted checkpoint and a potential new candidate for consistency. The method is deleted, and its usages removed and cleaned up. The integration test had a bug before where the array of checkpoints always contained the same checkpoint, so this is also a bug fix.",
          "timestamp": "2025-03-05T15:59:43Z",
          "tree_id": "4ab66b290ada139ee91a5ffdaf6b8f9ae916e180",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/318ab44d27c1c4518b026616dc2b0eeb89b7ff57"
        },
        "date": 1741190432759,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4409547,
            "unit": "ns/op\t  700626 B/op\t   19675 allocs/op",
            "extra": "294 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4409547,
            "unit": "ns/op",
            "extra": "294 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700626,
            "unit": "B/op",
            "extra": "294 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19675,
            "unit": "allocs/op",
            "extra": "294 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 228937,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "4951 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 228937,
            "unit": "ns/op",
            "extra": "4951 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "4951 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "4951 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8243,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "160330 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8243,
            "unit": "ns/op",
            "extra": "160330 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "160330 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "160330 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 118,
            "unit": "ns/op\t     112 B/op\t       3 allocs/op",
            "extra": "10320373 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 118,
            "unit": "ns/op",
            "extra": "10320373 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "10320373 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10320373 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 354527,
            "unit": "ns/op\t  289975 B/op\t    3032 allocs/op",
            "extra": "3738 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 354527,
            "unit": "ns/op",
            "extra": "3738 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 289975,
            "unit": "B/op",
            "extra": "3738 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3738 times\n4 procs"
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
          "id": "02f8df4e55f47069194a836852989afd021672f5",
          "message": "[Witness] construct proofs in witness (#514)\n\nThis switches away from the previous idea which was to pass in a proof\nfetcher. Now a tile fetcher is passed in, and the proofs are constructed\nin the witness code.\n\nNow that real proofs are being generated, we need to use a real log in\nthe witness tests. The POSIX log in testdata is used for this purpose.\n\nThe unsafe method for parsing checkpoints has also been extended to\nreturn hashes, which was needed to support this work.\n\nTowards #309.",
          "timestamp": "2025-03-10T11:50:37Z",
          "tree_id": "4b205ac33278f99e896e01908ccc49b367722483",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/02f8df4e55f47069194a836852989afd021672f5"
        },
        "date": 1741607482158,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4831950,
            "unit": "ns/op\t  704936 B/op\t   19716 allocs/op",
            "extra": "331 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4831950,
            "unit": "ns/op",
            "extra": "331 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 704936,
            "unit": "B/op",
            "extra": "331 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19716,
            "unit": "allocs/op",
            "extra": "331 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 225489,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5934 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 225489,
            "unit": "ns/op",
            "extra": "5934 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5934 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5934 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8128,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "203110 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8128,
            "unit": "ns/op",
            "extra": "203110 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "203110 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "203110 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 236.1,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5055319 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 236.1,
            "unit": "ns/op",
            "extra": "5055319 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5055319 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5055319 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 351500,
            "unit": "ns/op\t  290036 B/op\t    3032 allocs/op",
            "extra": "3780 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 351500,
            "unit": "ns/op",
            "extra": "3780 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290036,
            "unit": "B/op",
            "extra": "3780 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3780 times\n4 procs"
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
          "id": "83986cab7094aa99f7a01a4ed8a81baf5024b6e3",
          "message": "[Witnessing] Use witnessing in all drivers (#515)\n\nAll drivers now take set up witnessing as configured in the AppendOptions. If no options are provided then this is still set up, but does nothing. Towards #309.\n\nThis required changing the witness tests package to avoid a package cycle; the posix implementation now uses the witness, but the witness tests use posix.",
          "timestamp": "2025-03-10T15:27:07Z",
          "tree_id": "e21c795cfb03ac0573485329f650c2eebbd138cc",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/83986cab7094aa99f7a01a4ed8a81baf5024b6e3"
        },
        "date": 1741620480506,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4813277,
            "unit": "ns/op\t  700632 B/op\t   19676 allocs/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4813277,
            "unit": "ns/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700632,
            "unit": "B/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19676,
            "unit": "allocs/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 230850,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5732 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 230850,
            "unit": "ns/op",
            "extra": "5732 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5732 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5732 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7706,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "165181 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7706,
            "unit": "ns/op",
            "extra": "165181 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "165181 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "165181 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 237.5,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5109265 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 237.5,
            "unit": "ns/op",
            "extra": "5109265 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5109265 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5109265 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 361359,
            "unit": "ns/op\t  290667 B/op\t    3036 allocs/op",
            "extra": "3637 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 361359,
            "unit": "ns/op",
            "extra": "3637 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290667,
            "unit": "B/op",
            "extra": "3637 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "3637 times\n4 procs"
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
          "id": "4947c5becd8bbb0828afa333126f23864f9a7789",
          "message": "Bump github/codeql-action in the all-gha-deps group (#518)\n\nBumps the all-gha-deps group with 1 update: [github/codeql-action](https://github.com/github/codeql-action).\n\n\nUpdates `github/codeql-action` from 3.28.10 to 3.28.11\n- [Release notes](https://github.com/github/codeql-action/releases)\n- [Changelog](https://github.com/github/codeql-action/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/github/codeql-action/compare/b56ba49b26e50535fa1e7f7db0f4f7b4bf65d80d...6bb031afdd8eb862ea3fc1848194185e076637e5)\n\n---\nupdated-dependencies:\n- dependency-name: github/codeql-action\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-03-10T20:42:12Z",
          "tree_id": "9533731e4377a1e4494e24cc86f4568ca138004d",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/4947c5becd8bbb0828afa333126f23864f9a7789"
        },
        "date": 1741639376990,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4330920,
            "unit": "ns/op\t  700552 B/op\t   19674 allocs/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4330920,
            "unit": "ns/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700552,
            "unit": "B/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19674,
            "unit": "allocs/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 224752,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5839 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 224752,
            "unit": "ns/op",
            "extra": "5839 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5839 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5839 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7767,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "179340 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7767,
            "unit": "ns/op",
            "extra": "179340 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "179340 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "179340 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 231.7,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5158659 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 231.7,
            "unit": "ns/op",
            "extra": "5158659 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5158659 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5158659 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 345100,
            "unit": "ns/op\t  291132 B/op\t    3041 allocs/op",
            "extra": "3924 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 345100,
            "unit": "ns/op",
            "extra": "3924 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291132,
            "unit": "B/op",
            "extra": "3924 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3041,
            "unit": "allocs/op",
            "extra": "3924 times\n4 procs"
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
          "id": "aad20b4ed025ca51dc428c3765d622778af65989",
          "message": "Bump the all-go-deps group with 10 updates (#519)\n\nBumps the all-go-deps group with 10 updates:\n\n| Package | From | To |\n| --- | --- | --- |\n| [cloud.google.com/go/spanner](https://github.com/googleapis/google-cloud-go) | `1.76.1` | `1.77.0` |\n| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.8` | `1.29.9` |\n| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.61` | `1.17.62` |\n| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.78.0` | `1.78.1` |\n| [golang.org/x/mod](https://github.com/golang/mod) | `0.23.0` | `0.24.0` |\n| [google.golang.org/api](https://github.com/googleapis/google-api-go-client) | `0.223.0` | `0.224.0` |\n| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.70.0` | `1.71.0` |\n| [golang.org/x/crypto](https://github.com/golang/crypto) | `0.35.0` | `0.36.0` |\n| [golang.org/x/net](https://github.com/golang/net) | `0.35.0` | `0.37.0` |\n| [golang.org/x/sync](https://github.com/golang/sync) | `0.11.0` | `0.12.0` |\n\n\nUpdates `cloud.google.com/go/spanner` from 1.76.1 to 1.77.0\n- [Release notes](https://github.com/googleapis/google-cloud-go/releases)\n- [Changelog](https://github.com/googleapis/google-cloud-go/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-cloud-go/compare/spanner/v1.76.1...spanner/v1.77.0)\n\nUpdates `github.com/aws/aws-sdk-go-v2/config` from 1.29.8 to 1.29.9\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.8...config/v1.29.9)\n\nUpdates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.61 to 1.17.62\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.61...credentials/v1.17.62)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.78.0 to 1.78.1\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.78.0...service/s3/v1.78.1)\n\nUpdates `golang.org/x/mod` from 0.23.0 to 0.24.0\n- [Commits](https://github.com/golang/mod/compare/v0.23.0...v0.24.0)\n\nUpdates `google.golang.org/api` from 0.223.0 to 0.224.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.223.0...v0.224.0)\n\nUpdates `google.golang.org/grpc` from 1.70.0 to 1.71.0\n- [Release notes](https://github.com/grpc/grpc-go/releases)\n- [Commits](https://github.com/grpc/grpc-go/compare/v1.70.0...v1.71.0)\n\nUpdates `golang.org/x/crypto` from 0.35.0 to 0.36.0\n- [Commits](https://github.com/golang/crypto/compare/v0.35.0...v0.36.0)\n\nUpdates `golang.org/x/net` from 0.35.0 to 0.37.0\n- [Commits](https://github.com/golang/net/compare/v0.35.0...v0.37.0)\n\nUpdates `golang.org/x/sync` from 0.11.0 to 0.12.0\n- [Commits](https://github.com/golang/sync/compare/v0.11.0...v0.12.0)\n\n---\nupdated-dependencies:\n- dependency-name: cloud.google.com/go/spanner\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/config\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/credentials\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/mod\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/grpc\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/crypto\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/net\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/sync\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-03-10T21:42:01Z",
          "tree_id": "f774f0854a352bcbf4336976302337dedc1b0fbc",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/aad20b4ed025ca51dc428c3765d622778af65989"
        },
        "date": 1741643004740,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4719418,
            "unit": "ns/op\t  700359 B/op\t   19671 allocs/op",
            "extra": "278 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4719418,
            "unit": "ns/op",
            "extra": "278 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700359,
            "unit": "B/op",
            "extra": "278 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19671,
            "unit": "allocs/op",
            "extra": "278 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 237876,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5948 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 237876,
            "unit": "ns/op",
            "extra": "5948 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5948 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5948 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6086,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "189368 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6086,
            "unit": "ns/op",
            "extra": "189368 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "189368 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "189368 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 237.7,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5122686 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 237.7,
            "unit": "ns/op",
            "extra": "5122686 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5122686 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5122686 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 358426,
            "unit": "ns/op\t  291086 B/op\t    3041 allocs/op",
            "extra": "3919 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 358426,
            "unit": "ns/op",
            "extra": "3919 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291086,
            "unit": "B/op",
            "extra": "3919 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3041,
            "unit": "allocs/op",
            "extra": "3919 times\n4 procs"
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
          "id": "4674c5af678493a7602159b6c5812bea441cbdf9",
          "message": "Antispam tidy (#517)",
          "timestamp": "2025-03-11T10:54:34Z",
          "tree_id": "7e329cf3e80137ce10d09fbe27fdd242d156679b",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/4674c5af678493a7602159b6c5812bea441cbdf9"
        },
        "date": 1741690515954,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1164668,
            "unit": "ns/op\t  689773 B/op\t   19563 allocs/op",
            "extra": "997 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1164668,
            "unit": "ns/op",
            "extra": "997 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689773,
            "unit": "B/op",
            "extra": "997 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19563,
            "unit": "allocs/op",
            "extra": "997 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 135883,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8743 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 135883,
            "unit": "ns/op",
            "extra": "8743 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8743 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8743 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1846,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "543432 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1846,
            "unit": "ns/op",
            "extra": "543432 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "543432 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "543432 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 231.8,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "4818326 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 231.8,
            "unit": "ns/op",
            "extra": "4818326 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "4818326 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "4818326 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 348782,
            "unit": "ns/op\t  290402 B/op\t    3036 allocs/op",
            "extra": "3872 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 348782,
            "unit": "ns/op",
            "extra": "3872 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290402,
            "unit": "B/op",
            "extra": "3872 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "3872 times\n4 procs"
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
          "id": "ecabef06b8a186a552380a04f9ed4855e546af02",
          "message": "[Witnessing] Top level documentation (#520)\n\nThis fixes #309.",
          "timestamp": "2025-03-11T13:46:29Z",
          "tree_id": "211ceb4bfe604bbc469a2ccbbd01e6f19a7b4667",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/ecabef06b8a186a552380a04f9ed4855e546af02"
        },
        "date": 1741700831911,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1191248,
            "unit": "ns/op\t  690248 B/op\t   19568 allocs/op",
            "extra": "993 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1191248,
            "unit": "ns/op",
            "extra": "993 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 690248,
            "unit": "B/op",
            "extra": "993 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19568,
            "unit": "allocs/op",
            "extra": "993 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 141203,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8425 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 141203,
            "unit": "ns/op",
            "extra": "8425 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8425 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8425 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1856,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "601118 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1856,
            "unit": "ns/op",
            "extra": "601118 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "601118 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "601118 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 233.6,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5135782 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 233.6,
            "unit": "ns/op",
            "extra": "5135782 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5135782 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5135782 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 359717,
            "unit": "ns/op\t  291299 B/op\t    3042 allocs/op",
            "extra": "3933 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 359717,
            "unit": "ns/op",
            "extra": "3933 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291299,
            "unit": "B/op",
            "extra": "3933 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3042,
            "unit": "allocs/op",
            "extra": "3933 times\n4 procs"
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
          "id": "52781f91caa777eabf91d492c5636204bcf93219",
          "message": "[GCP] Add persistent anti-spam support to terraform (#521)",
          "timestamp": "2025-03-12T11:37:49Z",
          "tree_id": "cdaf95f9ef9b7233a541ef39e35a77691df60517",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/52781f91caa777eabf91d492c5636204bcf93219"
        },
        "date": 1741779512629,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1480283,
            "unit": "ns/op\t  691961 B/op\t   19586 allocs/op",
            "extra": "872 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1480283,
            "unit": "ns/op",
            "extra": "872 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 691961,
            "unit": "B/op",
            "extra": "872 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19586,
            "unit": "allocs/op",
            "extra": "872 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 144845,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8370 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 144845,
            "unit": "ns/op",
            "extra": "8370 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8370 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8370 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1873,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "605925 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1873,
            "unit": "ns/op",
            "extra": "605925 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "605925 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "605925 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 236.4,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "4922078 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 236.4,
            "unit": "ns/op",
            "extra": "4922078 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "4922078 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "4922078 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 352707,
            "unit": "ns/op\t  290459 B/op\t    3037 allocs/op",
            "extra": "3871 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 352707,
            "unit": "ns/op",
            "extra": "3871 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290459,
            "unit": "B/op",
            "extra": "3871 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3037,
            "unit": "allocs/op",
            "extra": "3871 times\n4 procs"
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
          "id": "0020add8158cb584509f1338df2a5749ff09636d",
          "message": "Modernize for loop using `range` over `int` (#522)",
          "timestamp": "2025-03-12T14:43:37Z",
          "tree_id": "0cd88e81490fdef1672084adbf1dd17c5dbba297",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/0020add8158cb584509f1338df2a5749ff09636d"
        },
        "date": 1741790660041,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1191844,
            "unit": "ns/op\t  689345 B/op\t   19559 allocs/op",
            "extra": "998 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1191844,
            "unit": "ns/op",
            "extra": "998 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689345,
            "unit": "B/op",
            "extra": "998 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19559,
            "unit": "allocs/op",
            "extra": "998 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136223,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8522 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136223,
            "unit": "ns/op",
            "extra": "8522 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8522 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8522 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1827,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "568006 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1827,
            "unit": "ns/op",
            "extra": "568006 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "568006 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "568006 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 236.2,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "4925504 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 236.2,
            "unit": "ns/op",
            "extra": "4925504 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "4925504 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "4925504 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 355044,
            "unit": "ns/op\t  290429 B/op\t    3036 allocs/op",
            "extra": "4020 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 355044,
            "unit": "ns/op",
            "extra": "4020 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290429,
            "unit": "B/op",
            "extra": "4020 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "4020 times\n4 procs"
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
          "id": "668a7bb74b377002fd56f58470723470ecb4745c",
          "message": "Add CT support for migration (#523)",
          "timestamp": "2025-03-13T13:37:09Z",
          "tree_id": "2b74b27231ad68f390df444f532d4a46815408c9",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/668a7bb74b377002fd56f58470723470ecb4745c"
        },
        "date": 1741873074325,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1305314,
            "unit": "ns/op\t  691284 B/op\t   19579 allocs/op",
            "extra": "996 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1305314,
            "unit": "ns/op",
            "extra": "996 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 691284,
            "unit": "B/op",
            "extra": "996 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19579,
            "unit": "allocs/op",
            "extra": "996 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 137831,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8328 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 137831,
            "unit": "ns/op",
            "extra": "8328 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8328 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8328 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1834,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "551756 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1834,
            "unit": "ns/op",
            "extra": "551756 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "551756 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "551756 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 236,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5138940 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 236,
            "unit": "ns/op",
            "extra": "5138940 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5138940 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5138940 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 352403,
            "unit": "ns/op\t  290731 B/op\t    3039 allocs/op",
            "extra": "3896 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 352403,
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
          "id": "3fc38900820bba64add66fc78c3f4c1cd28e5ba0",
          "message": "[Shutdown] Implement Shutdown in a single location (#525)\n\nThis is a revised attempt at Shutdown that is intended to replace #524.\nThis implementation doesn't have the granular integration with the\nchannels in each implementation. Instead, it is implemented at the top\nlevel of the appender lifecycle, and each driver doesn't need to have\nvery similar duplicated code for handling this.\n\nThe previous attempt was also trying to make it so that Shutdown would\nclose out channels and kill goroutines etc. This approach makes it clear\nthat the way to do this is to cancel the original context passed into\nthe appender.\n\nFixes #341.",
          "timestamp": "2025-03-13T14:49:55Z",
          "tree_id": "54719511c464a3785f82833a46c155d77f352bda",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/3fc38900820bba64add66fc78c3f4c1cd28e5ba0"
        },
        "date": 1741877440469,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1286986,
            "unit": "ns/op\t  690921 B/op\t   19575 allocs/op",
            "extra": "1002 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1286986,
            "unit": "ns/op",
            "extra": "1002 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 690921,
            "unit": "B/op",
            "extra": "1002 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19575,
            "unit": "allocs/op",
            "extra": "1002 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136424,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8733 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136424,
            "unit": "ns/op",
            "extra": "8733 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8733 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8733 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1860,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "642296 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1860,
            "unit": "ns/op",
            "extra": "642296 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "642296 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "642296 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 238.2,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5104987 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 238.2,
            "unit": "ns/op",
            "extra": "5104987 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5104987 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5104987 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 365641,
            "unit": "ns/op\t  290027 B/op\t    3032 allocs/op",
            "extra": "3717 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 365641,
            "unit": "ns/op",
            "extra": "3717 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290027,
            "unit": "B/op",
            "extra": "3717 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3717 times\n4 procs"
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
          "id": "692dc3bd607cd1647d94acf3b60fb88367cef124",
          "message": "Fix migration in POSIX (#528)",
          "timestamp": "2025-03-13T16:17:17Z",
          "tree_id": "e74d9389cd84b1454f3cafd4d17c5474ddad32fc",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/692dc3bd607cd1647d94acf3b60fb88367cef124"
        },
        "date": 1741882691417,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1328193,
            "unit": "ns/op\t  690815 B/op\t   19574 allocs/op",
            "extra": "998 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1328193,
            "unit": "ns/op",
            "extra": "998 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 690815,
            "unit": "B/op",
            "extra": "998 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19574,
            "unit": "allocs/op",
            "extra": "998 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136887,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8636 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136887,
            "unit": "ns/op",
            "extra": "8636 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8636 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8636 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1862,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "623043 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1862,
            "unit": "ns/op",
            "extra": "623043 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "623043 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "623043 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 233.2,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5064727 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 233.2,
            "unit": "ns/op",
            "extra": "5064727 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5064727 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5064727 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 355836,
            "unit": "ns/op\t  290775 B/op\t    3038 allocs/op",
            "extra": "3979 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 355836,
            "unit": "ns/op",
            "extra": "3979 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290775,
            "unit": "B/op",
            "extra": "3979 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3038,
            "unit": "allocs/op",
            "extra": "3979 times\n4 procs"
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
          "id": "0d95102bb45c180159b5f99b9048c703cb494a54",
          "message": "[Witnessing] Improve parsing response from witnesses (#529)\n\nParse straight to uint64. If #509 rebases on top of this then the CodeQL error will be easily fixable.",
          "timestamp": "2025-03-13T16:32:14Z",
          "tree_id": "892053a70e5f300cf694c6328dd1f79eb1a2dd77",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/0d95102bb45c180159b5f99b9048c703cb494a54"
        },
        "date": 1741883579762,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1289694,
            "unit": "ns/op\t  690423 B/op\t   19569 allocs/op",
            "extra": "1020 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1289694,
            "unit": "ns/op",
            "extra": "1020 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 690423,
            "unit": "B/op",
            "extra": "1020 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19569,
            "unit": "allocs/op",
            "extra": "1020 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136121,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8547 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136121,
            "unit": "ns/op",
            "extra": "8547 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8547 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8547 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1808,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "688423 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1808,
            "unit": "ns/op",
            "extra": "688423 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "688423 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "688423 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 239.8,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5100723 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 239.8,
            "unit": "ns/op",
            "extra": "5100723 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5100723 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5100723 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 347092,
            "unit": "ns/op\t  291251 B/op\t    3043 allocs/op",
            "extra": "4243 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 347092,
            "unit": "ns/op",
            "extra": "4243 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291251,
            "unit": "B/op",
            "extra": "4243 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3043,
            "unit": "allocs/op",
            "extra": "4243 times\n4 procs"
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
          "id": "cc9c3d2c618d9084352f93f494618992aa1b2f74",
          "message": "[Docs] Update code samples in docs (#527)\n\nAlso rename mix-ins to features, as things have changed a bit and 2/3 of these features are discoverable via the top-level API instead of being optional libraries.",
          "timestamp": "2025-03-13T16:37:55Z",
          "tree_id": "55e8f3d641e35db9948330ba2e56997d14c1bd3d",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/cc9c3d2c618d9084352f93f494618992aa1b2f74"
        },
        "date": 1741883918026,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1242256,
            "unit": "ns/op\t  689916 B/op\t   19564 allocs/op",
            "extra": "1004 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1242256,
            "unit": "ns/op",
            "extra": "1004 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689916,
            "unit": "B/op",
            "extra": "1004 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19564,
            "unit": "allocs/op",
            "extra": "1004 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 142757,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8572 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 142757,
            "unit": "ns/op",
            "extra": "8572 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8572 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8572 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1871,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "652857 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1871,
            "unit": "ns/op",
            "extra": "652857 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "652857 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "652857 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 232,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5120611 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 232,
            "unit": "ns/op",
            "extra": "5120611 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5120611 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5120611 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 353642,
            "unit": "ns/op\t  290503 B/op\t    3037 allocs/op",
            "extra": "4161 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 353642,
            "unit": "ns/op",
            "extra": "4161 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290503,
            "unit": "B/op",
            "extra": "4161 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3037,
            "unit": "allocs/op",
            "extra": "4161 times\n4 procs"
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
          "id": "588d4207038bce991bf29ede7b5f60b22191c84e",
          "message": "Expose migrate functionality to applications built with Tessera (#530)",
          "timestamp": "2025-03-13T16:49:38Z",
          "tree_id": "39d47f5b846c8f1e61ea75fe4bce4885047b08f0",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/588d4207038bce991bf29ede7b5f60b22191c84e"
        },
        "date": 1741884619910,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1191151,
            "unit": "ns/op\t  689581 B/op\t   19561 allocs/op",
            "extra": "961 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1191151,
            "unit": "ns/op",
            "extra": "961 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689581,
            "unit": "B/op",
            "extra": "961 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19561,
            "unit": "allocs/op",
            "extra": "961 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136528,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8467 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136528,
            "unit": "ns/op",
            "extra": "8467 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8467 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8467 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1616,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "712809 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1616,
            "unit": "ns/op",
            "extra": "712809 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "712809 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "712809 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 232.6,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "4836200 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 232.6,
            "unit": "ns/op",
            "extra": "4836200 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "4836200 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "4836200 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 345564,
            "unit": "ns/op\t  290151 B/op\t    3033 allocs/op",
            "extra": "3819 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 345564,
            "unit": "ns/op",
            "extra": "3819 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290151,
            "unit": "B/op",
            "extra": "3819 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3033,
            "unit": "allocs/op",
            "extra": "3819 times\n4 procs"
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
          "id": "645ecede88a44bd2e6239a6e7253fd7a583a4902",
          "message": "Replace m[k]=v loop with `maps.Copy` (#533)",
          "timestamp": "2025-03-13T17:22:27Z",
          "tree_id": "befc0a2242647447b5871a38359af8ec7cd91e30",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/645ecede88a44bd2e6239a6e7253fd7a583a4902"
        },
        "date": 1741886587886,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1183764,
            "unit": "ns/op\t  689881 B/op\t   19564 allocs/op",
            "extra": "1003 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1183764,
            "unit": "ns/op",
            "extra": "1003 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689881,
            "unit": "B/op",
            "extra": "1003 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19564,
            "unit": "allocs/op",
            "extra": "1003 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136270,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8595 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136270,
            "unit": "ns/op",
            "extra": "8595 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8595 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8595 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1593,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "724647 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1593,
            "unit": "ns/op",
            "extra": "724647 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "724647 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "724647 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 233.7,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5179096 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 233.7,
            "unit": "ns/op",
            "extra": "5179096 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5179096 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5179096 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 345754,
            "unit": "ns/op\t  290513 B/op\t    3037 allocs/op",
            "extra": "3880 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 345754,
            "unit": "ns/op",
            "extra": "3880 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290513,
            "unit": "B/op",
            "extra": "3880 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3037,
            "unit": "allocs/op",
            "extra": "3880 times\n4 procs"
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
          "id": "602b87625e2322e0a1d24803957fd934b40be930",
          "message": "Modernize if statement using `max` (#532)",
          "timestamp": "2025-03-13T19:56:27Z",
          "tree_id": "57d47c01db76aba1df314ee5dc869e52a011b3af",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/602b87625e2322e0a1d24803957fd934b40be930"
        },
        "date": 1741895828572,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1179046,
            "unit": "ns/op\t  689712 B/op\t   19563 allocs/op",
            "extra": "974 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1179046,
            "unit": "ns/op",
            "extra": "974 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689712,
            "unit": "B/op",
            "extra": "974 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19563,
            "unit": "allocs/op",
            "extra": "974 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136399,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8762 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136399,
            "unit": "ns/op",
            "extra": "8762 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8762 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8762 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1626,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "710235 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1626,
            "unit": "ns/op",
            "extra": "710235 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "710235 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "710235 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 233.6,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5069343 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 233.6,
            "unit": "ns/op",
            "extra": "5069343 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5069343 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5069343 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 352878,
            "unit": "ns/op\t  290319 B/op\t    3035 allocs/op",
            "extra": "4071 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 352878,
            "unit": "ns/op",
            "extra": "4071 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290319,
            "unit": "B/op",
            "extra": "4071 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3035,
            "unit": "allocs/op",
            "extra": "4071 times\n4 procs"
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
          "id": "be012b02d99ba34743242dc6ddf21e07a38a8c0d",
          "message": "Modernize if statement using `min` (#531)\n\n* Modernize if statement using `min`\n\n* Swap `min` arguments order\n\nCo-authored-by: Al Cutter <alcutter@google.com>\n\n* Swap `min` arguments order\n\nCo-authored-by: Al Cutter <alcutter@google.com>\n\n---------\n\nCo-authored-by: Al Cutter <alcutter@google.com>",
          "timestamp": "2025-03-13T20:20:05Z",
          "tree_id": "34d674a197c6c70ac7059c5d8950a881294b185b",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/be012b02d99ba34743242dc6ddf21e07a38a8c0d"
        },
        "date": 1741897244552,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1173784,
            "unit": "ns/op\t  689690 B/op\t   19562 allocs/op",
            "extra": "984 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1173784,
            "unit": "ns/op",
            "extra": "984 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689690,
            "unit": "B/op",
            "extra": "984 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19562,
            "unit": "allocs/op",
            "extra": "984 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136310,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8704 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136310,
            "unit": "ns/op",
            "extra": "8704 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8704 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8704 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1589,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "669342 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1589,
            "unit": "ns/op",
            "extra": "669342 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "669342 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "669342 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 237.2,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5173581 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 237.2,
            "unit": "ns/op",
            "extra": "5173581 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5173581 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5173581 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 348445,
            "unit": "ns/op\t  291201 B/op\t    3042 allocs/op",
            "extra": "3940 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 348445,
            "unit": "ns/op",
            "extra": "3940 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291201,
            "unit": "B/op",
            "extra": "3940 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3042,
            "unit": "allocs/op",
            "extra": "3940 times\n4 procs"
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
          "id": "b7c8fb6d44913d8243b645abdd4d6381d3075360",
          "message": "Improve main README file (#534)",
          "timestamp": "2025-03-14T14:37:07Z",
          "tree_id": "1809d433baa54186508a6eff34e5574ed421ce6c",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/b7c8fb6d44913d8243b645abdd4d6381d3075360"
        },
        "date": 1741963072600,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1302342,
            "unit": "ns/op\t  690082 B/op\t   19566 allocs/op",
            "extra": "1012 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1302342,
            "unit": "ns/op",
            "extra": "1012 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 690082,
            "unit": "B/op",
            "extra": "1012 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19566,
            "unit": "allocs/op",
            "extra": "1012 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 139615,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8608 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 139615,
            "unit": "ns/op",
            "extra": "8608 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8608 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8608 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1608,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "711596 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1608,
            "unit": "ns/op",
            "extra": "711596 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "711596 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "711596 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 235.8,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5160015 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 235.8,
            "unit": "ns/op",
            "extra": "5160015 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5160015 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5160015 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 355100,
            "unit": "ns/op\t  290462 B/op\t    3036 allocs/op",
            "extra": "4026 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 355100,
            "unit": "ns/op",
            "extra": "4026 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290462,
            "unit": "B/op",
            "extra": "4026 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "4026 times\n4 procs"
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
          "id": "a64a2b73859aa28c1691c3260d78733a54d5f8eb",
          "message": "Create security policy (#255)",
          "timestamp": "2025-03-17T12:23:32Z",
          "tree_id": "58c30d7dbaf78ccea314f4fe6dd9c74448a6430f",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/a64a2b73859aa28c1691c3260d78733a54d5f8eb"
        },
        "date": 1742214255458,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1255068,
            "unit": "ns/op\t  690439 B/op\t   19570 allocs/op",
            "extra": "1024 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1255068,
            "unit": "ns/op",
            "extra": "1024 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 690439,
            "unit": "B/op",
            "extra": "1024 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19570,
            "unit": "allocs/op",
            "extra": "1024 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 139507,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8628 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 139507,
            "unit": "ns/op",
            "extra": "8628 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8628 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8628 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1595,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "728458 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1595,
            "unit": "ns/op",
            "extra": "728458 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "728458 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "728458 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 232.7,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5139465 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 232.7,
            "unit": "ns/op",
            "extra": "5139465 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5139465 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5139465 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 357909,
            "unit": "ns/op\t  290407 B/op\t    3036 allocs/op",
            "extra": "4141 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 357909,
            "unit": "ns/op",
            "extra": "4141 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290407,
            "unit": "B/op",
            "extra": "4141 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "4141 times\n4 procs"
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
          "id": "ed0b4e51d70e45e8d218625730b34e97e18f6c2e",
          "message": "Remove unused things from API (#536)\n\nDuplicate docs instead of linking to private docs",
          "timestamp": "2025-03-17T14:22:35Z",
          "tree_id": "5adcba16bf3750b129f3150cd8e7966a11f45148",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/ed0b4e51d70e45e8d218625730b34e97e18f6c2e"
        },
        "date": 1742221403834,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1311715,
            "unit": "ns/op\t  690818 B/op\t   19574 allocs/op",
            "extra": "1022 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1311715,
            "unit": "ns/op",
            "extra": "1022 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 690818,
            "unit": "B/op",
            "extra": "1022 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19574,
            "unit": "allocs/op",
            "extra": "1022 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136436,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8276 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136436,
            "unit": "ns/op",
            "extra": "8276 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8276 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8276 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1602,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "749396 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1602,
            "unit": "ns/op",
            "extra": "749396 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "749396 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "749396 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 234.3,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5218366 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 234.3,
            "unit": "ns/op",
            "extra": "5218366 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5218366 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5218366 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 353886,
            "unit": "ns/op\t  290312 B/op\t    3035 allocs/op",
            "extra": "4065 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 353886,
            "unit": "ns/op",
            "extra": "4065 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290312,
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
          "id": "6f25ba6ca3a6d4fea45f2e14f2735fcb82d5774d",
          "message": "Use common implementations for hashing (#537)",
          "timestamp": "2025-03-17T14:39:07Z",
          "tree_id": "47a3a4b111a7c1f1914f7ed22c7a4c318735b41b",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/6f25ba6ca3a6d4fea45f2e14f2735fcb82d5774d"
        },
        "date": 1742222393973,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1134148,
            "unit": "ns/op\t  689738 B/op\t   19563 allocs/op",
            "extra": "1024 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1134148,
            "unit": "ns/op",
            "extra": "1024 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 689738,
            "unit": "B/op",
            "extra": "1024 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19563,
            "unit": "allocs/op",
            "extra": "1024 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 138250,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "7921 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 138250,
            "unit": "ns/op",
            "extra": "7921 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "7921 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "7921 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1598,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "707583 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1598,
            "unit": "ns/op",
            "extra": "707583 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "707583 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "707583 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 232.3,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5154974 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 232.3,
            "unit": "ns/op",
            "extra": "5154974 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5154974 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5154974 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 342577,
            "unit": "ns/op\t  290686 B/op\t    3039 allocs/op",
            "extra": "4198 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 342577,
            "unit": "ns/op",
            "extra": "4198 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290686,
            "unit": "B/op",
            "extra": "4198 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3039,
            "unit": "allocs/op",
            "extra": "4198 times\n4 procs"
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
          "id": "e61e2d86f685d143d65b69ff39fad0e66976927c",
          "message": "Minor updates in anticipation of alpha 2 release today (#535)",
          "timestamp": "2025-03-17T15:19:15Z",
          "tree_id": "d2bfd83abd6ca1787b56422b741b2edf6e89a6e2",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/e61e2d86f685d143d65b69ff39fad0e66976927c"
        },
        "date": 1742224803319,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1232838,
            "unit": "ns/op\t  690538 B/op\t   19571 allocs/op",
            "extra": "1065 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1232838,
            "unit": "ns/op",
            "extra": "1065 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 690538,
            "unit": "B/op",
            "extra": "1065 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19571,
            "unit": "allocs/op",
            "extra": "1065 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136050,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8698 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136050,
            "unit": "ns/op",
            "extra": "8698 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8698 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8698 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1612,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "645615 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1612,
            "unit": "ns/op",
            "extra": "645615 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "645615 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "645615 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 232,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5125250 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 232,
            "unit": "ns/op",
            "extra": "5125250 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5125250 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5125250 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 345266,
            "unit": "ns/op\t  290494 B/op\t    3037 allocs/op",
            "extra": "4160 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 345266,
            "unit": "ns/op",
            "extra": "4160 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290494,
            "unit": "B/op",
            "extra": "4160 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3037,
            "unit": "allocs/op",
            "extra": "4160 times\n4 procs"
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
          "id": "5bb36bbbd31459b7508b800da8c8139f1b34c358",
          "message": "Bump the all-gha-deps group with 2 updates (#539)\n\n* Bump the all-gha-deps group with 2 updates\n\nBumps the all-gha-deps group with 2 updates: [gruntwork-io/terragrunt-action](https://github.com/gruntwork-io/terragrunt-action) and [golangci/golangci-lint-action](https://github.com/golangci/golangci-lint-action).\n\n\nUpdates `gruntwork-io/terragrunt-action` from 2.1.5 to 2.1.8\n- [Release notes](https://github.com/gruntwork-io/terragrunt-action/releases)\n- [Commits](https://github.com/gruntwork-io/terragrunt-action/compare/9559e51d05873b0ea467c42bbabcb5c067642ccc...aee21a7df999be8b471c2a8564c6cd853cb674e1)\n\nUpdates `golangci/golangci-lint-action` from 6.5.0 to 6.5.1\n- [Release notes](https://github.com/golangci/golangci-lint-action/releases)\n- [Commits](https://github.com/golangci/golangci-lint-action/compare/2226d7cb06a077cd73e56eedd38eecad18e5d837...4696ba8babb6127d732c3c6dde519db15edab9ea)\n\n---\nupdated-dependencies:\n- dependency-name: gruntwork-io/terragrunt-action\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n- dependency-name: golangci/golangci-lint-action\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\n\n* Fix incorrect version comment\n\n---------\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>\nCo-authored-by: Roger Ng <rogerng@google.com>",
          "timestamp": "2025-03-17T22:27:25Z",
          "tree_id": "fb860b855d73bd5e8e3f47a64648590884561333",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/5bb36bbbd31459b7508b800da8c8139f1b34c358"
        },
        "date": 1742250487789,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 1237870,
            "unit": "ns/op\t  691022 B/op\t   19576 allocs/op",
            "extra": "1056 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 1237870,
            "unit": "ns/op",
            "extra": "1056 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 691022,
            "unit": "B/op",
            "extra": "1056 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19576,
            "unit": "allocs/op",
            "extra": "1056 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136235,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8724 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136235,
            "unit": "ns/op",
            "extra": "8724 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8724 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8724 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1598,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "775682 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1598,
            "unit": "ns/op",
            "extra": "775682 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "775682 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "775682 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 233.5,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5053491 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 233.5,
            "unit": "ns/op",
            "extra": "5053491 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5053491 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5053491 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 354564,
            "unit": "ns/op\t  290992 B/op\t    3040 allocs/op",
            "extra": "3918 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 354564,
            "unit": "ns/op",
            "extra": "3918 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290992,
            "unit": "B/op",
            "extra": "3918 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3040,
            "unit": "allocs/op",
            "extra": "3918 times\n4 procs"
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
          "id": "8f6c7333ffbc5f8ef667f4a385cfadea90290b5d",
          "message": "Bump the all-go-deps group with 3 updates (#538)\n\nBumps the all-go-deps group with 3 updates: [cloud.google.com/go/storage](https://github.com/googleapis/google-cloud-go), [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) and [google.golang.org/api](https://github.com/googleapis/google-api-go-client).\n\n\nUpdates `cloud.google.com/go/storage` from 1.50.0 to 1.51.0\n- [Release notes](https://github.com/googleapis/google-cloud-go/releases)\n- [Changelog](https://github.com/googleapis/google-cloud-go/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-cloud-go/compare/spanner/v1.50.0...spanner/v1.51.0)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.78.1 to 1.78.2\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.78.1...service/s3/v1.78.2)\n\nUpdates `google.golang.org/api` from 0.224.0 to 0.226.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.224.0...v0.226.0)\n\n---\nupdated-dependencies:\n- dependency-name: cloud.google.com/go/storage\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-03-18T11:40:19Z",
          "tree_id": "83b50b0d453b1e542f0b5b9cf90424ab592f7d29",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/8f6c7333ffbc5f8ef667f4a385cfadea90290b5d"
        },
        "date": 1742298104788,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 5005507,
            "unit": "ns/op\t  705139 B/op\t   19717 allocs/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 5005507,
            "unit": "ns/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 705139,
            "unit": "B/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19717,
            "unit": "allocs/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 228519,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "6117 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 228519,
            "unit": "ns/op",
            "extra": "6117 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "6117 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "6117 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1600,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "744141 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1600,
            "unit": "ns/op",
            "extra": "744141 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "744141 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "744141 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 233.5,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5072920 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 233.5,
            "unit": "ns/op",
            "extra": "5072920 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5072920 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5072920 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 352794,
            "unit": "ns/op\t  290553 B/op\t    3037 allocs/op",
            "extra": "4170 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 352794,
            "unit": "ns/op",
            "extra": "4170 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290553,
            "unit": "B/op",
            "extra": "4170 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3037,
            "unit": "allocs/op",
            "extra": "4170 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "haydentherapper@users.noreply.github.com",
            "name": "Hayden B",
            "username": "haydentherapper"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "0bb3cf48265764f466916fd48cd62433d3032ddd",
          "message": "Verify that the witness size is not larger than the log size (#509)\n\nThis prevents a malicious witness from updating the log witness state\nto an invalid size.\n\nCo-authored-by: Hayden B <8418760+haydentherapper@users.noreply.github.com>",
          "timestamp": "2025-03-18T12:15:10Z",
          "tree_id": "6f4d7b53b5c43efe8ebb4630200babf4a20a65d1",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/0bb3cf48265764f466916fd48cd62433d3032ddd"
        },
        "date": 1742300153047,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4311682,
            "unit": "ns/op\t  702469 B/op\t   19692 allocs/op",
            "extra": "262 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4311682,
            "unit": "ns/op",
            "extra": "262 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702469,
            "unit": "B/op",
            "extra": "262 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19692,
            "unit": "allocs/op",
            "extra": "262 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 233580,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5629 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 233580,
            "unit": "ns/op",
            "extra": "5629 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5629 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5629 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8032,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "155113 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8032,
            "unit": "ns/op",
            "extra": "155113 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "155113 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "155113 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 232.8,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5112458 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 232.8,
            "unit": "ns/op",
            "extra": "5112458 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5112458 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5112458 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 350040,
            "unit": "ns/op\t  290611 B/op\t    3037 allocs/op",
            "extra": "4000 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 350040,
            "unit": "ns/op",
            "extra": "4000 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290611,
            "unit": "B/op",
            "extra": "4000 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3037,
            "unit": "allocs/op",
            "extra": "4000 times\n4 procs"
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
          "id": "ddc67fb316c716b61517d86769cb53fbaf279d26",
          "message": "Return current state of tree even if no work done (#540)",
          "timestamp": "2025-03-18T16:21:35Z",
          "tree_id": "8638005f88dea7df132e52269f0fc17d7b26c63d",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/ddc67fb316c716b61517d86769cb53fbaf279d26"
        },
        "date": 1742314947805,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4172864,
            "unit": "ns/op\t  698949 B/op\t   19658 allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4172864,
            "unit": "ns/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 698949,
            "unit": "B/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19658,
            "unit": "allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 223419,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5722 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 223419,
            "unit": "ns/op",
            "extra": "5722 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5722 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5722 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7843,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "200420 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7843,
            "unit": "ns/op",
            "extra": "200420 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "200420 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "200420 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 231.8,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5200201 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 231.8,
            "unit": "ns/op",
            "extra": "5200201 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5200201 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5200201 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 351656,
            "unit": "ns/op\t  290668 B/op\t    3038 allocs/op",
            "extra": "3891 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 351656,
            "unit": "ns/op",
            "extra": "3891 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290668,
            "unit": "B/op",
            "extra": "3891 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3038,
            "unit": "allocs/op",
            "extra": "3891 times\n4 procs"
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
          "id": "cbe24fec9686211573102be352cdc6a2ac6f6e9b",
          "message": "Fix AWS doc comment typo (#543)",
          "timestamp": "2025-03-19T10:33:24Z",
          "tree_id": "7cebbfbd1eb9bf775bafa65593bc3ca507f5b355",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/cbe24fec9686211573102be352cdc6a2ac6f6e9b"
        },
        "date": 1742380464158,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 6262279,
            "unit": "ns/op\t  696070 B/op\t   19626 allocs/op",
            "extra": "164 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 6262279,
            "unit": "ns/op",
            "extra": "164 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 696070,
            "unit": "B/op",
            "extra": "164 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19626,
            "unit": "allocs/op",
            "extra": "164 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 246349,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5738 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 246349,
            "unit": "ns/op",
            "extra": "5738 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5738 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5738 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 9868,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "164124 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 9868,
            "unit": "ns/op",
            "extra": "164124 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "164124 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "164124 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 267.2,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "4593540 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 267.2,
            "unit": "ns/op",
            "extra": "4593540 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "4593540 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "4593540 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 417693,
            "unit": "ns/op\t  290328 B/op\t    3032 allocs/op",
            "extra": "3248 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 417693,
            "unit": "ns/op",
            "extra": "3248 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290328,
            "unit": "B/op",
            "extra": "3248 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3248 times\n4 procs"
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
          "id": "3846ada89f91d4fe9d4a7961a7455dd1fba07dd5",
          "message": "Move migration into lifecycle (#542)",
          "timestamp": "2025-03-19T11:59:13Z",
          "tree_id": "ca54fb50764c86c6edb18a760cd3be2dbb3fa9c6",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/3846ada89f91d4fe9d4a7961a7455dd1fba07dd5"
        },
        "date": 1742385602434,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4370323,
            "unit": "ns/op\t  704538 B/op\t   19710 allocs/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4370323,
            "unit": "ns/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 704538,
            "unit": "B/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19710,
            "unit": "allocs/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 224043,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5911 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 224043,
            "unit": "ns/op",
            "extra": "5911 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5911 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5911 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7980,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "180916 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7980,
            "unit": "ns/op",
            "extra": "180916 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "180916 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "180916 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 233.4,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5144672 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 233.4,
            "unit": "ns/op",
            "extra": "5144672 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5144672 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5144672 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 356112,
            "unit": "ns/op\t  290334 B/op\t    3035 allocs/op",
            "extra": "4094 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 356112,
            "unit": "ns/op",
            "extra": "4094 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290334,
            "unit": "B/op",
            "extra": "4094 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3035,
            "unit": "allocs/op",
            "extra": "4094 times\n4 procs"
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
          "id": "554e57dfcdcc50575f608d3211d8db41c042c3fd",
          "message": "[GCP] Add antispam tunables (#545)",
          "timestamp": "2025-03-19T13:31:03Z",
          "tree_id": "9223c013a888e3ce57be2ac915fd7b57ea8df7c7",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/554e57dfcdcc50575f608d3211d8db41c042c3fd"
        },
        "date": 1742391112639,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4395310,
            "unit": "ns/op\t  703123 B/op\t   19700 allocs/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4395310,
            "unit": "ns/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 703123,
            "unit": "B/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19700,
            "unit": "allocs/op",
            "extra": "256 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 230578,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5398 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 230578,
            "unit": "ns/op",
            "extra": "5398 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5398 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5398 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8146,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "230452 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8146,
            "unit": "ns/op",
            "extra": "230452 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "230452 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "230452 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 241.8,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "4919391 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 241.8,
            "unit": "ns/op",
            "extra": "4919391 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "4919391 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "4919391 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 356742,
            "unit": "ns/op\t  290609 B/op\t    3036 allocs/op",
            "extra": "3640 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 356742,
            "unit": "ns/op",
            "extra": "3640 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290609,
            "unit": "B/op",
            "extra": "3640 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "3640 times\n4 procs"
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
          "id": "9098674d0ba2c6f6fef143f46b9ea71ce98dfb10",
          "message": "[Docs] Fix deprecated flag (#546)\n\nWarning when running it with latest release of terraform:\n\n```\n14:38:16.962 WARN   The `--terragrunt-working-dir` flag is deprecated\n   and will be removed in a future version of Terragrunt. Use\n`--working-dir=deployment/live/aws/codelab/` instead.\n```",
          "timestamp": "2025-03-19T14:50:55Z",
          "tree_id": "89aa69b37b9605bd42de850eb2f01381cdc83240",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/9098674d0ba2c6f6fef143f46b9ea71ce98dfb10"
        },
        "date": 1742395909233,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4330481,
            "unit": "ns/op\t  701857 B/op\t   19687 allocs/op",
            "extra": "285 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4330481,
            "unit": "ns/op",
            "extra": "285 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 701857,
            "unit": "B/op",
            "extra": "285 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19687,
            "unit": "allocs/op",
            "extra": "285 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 222529,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5908 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 222529,
            "unit": "ns/op",
            "extra": "5908 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5908 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5908 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7469,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "220240 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7469,
            "unit": "ns/op",
            "extra": "220240 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "220240 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "220240 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 239.8,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5129518 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 239.8,
            "unit": "ns/op",
            "extra": "5129518 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5129518 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5129518 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 355231,
            "unit": "ns/op\t  290487 B/op\t    3037 allocs/op",
            "extra": "3882 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 355231,
            "unit": "ns/op",
            "extra": "3882 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290487,
            "unit": "B/op",
            "extra": "3882 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3037,
            "unit": "allocs/op",
            "extra": "3882 times\n4 procs"
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
          "id": "e18c9465361dd19c3e72d902d19284506f825bbf",
          "message": "AWS migration target support (#541)\n\nThis is based on the GCP migration target implementation.\n\nThe conformance cmd was updated to move the listen flag in order to keep\nall common AWS configuration flags together, and the flags for the\napplication logic built on top (e.g. web serving) in another block.\n\nTested in AWS and it migrated a POSIX log with 189k entries in 7s.\n\nFixes #473.",
          "timestamp": "2025-03-20T14:32:58Z",
          "tree_id": "cc944a17e8cbba40cd8a371dae9f5f62e9b2eaa6",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/e18c9465361dd19c3e72d902d19284506f825bbf"
        },
        "date": 1742481224494,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4317806,
            "unit": "ns/op\t  703894 B/op\t   19707 allocs/op",
            "extra": "279 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4317806,
            "unit": "ns/op",
            "extra": "279 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 703894,
            "unit": "B/op",
            "extra": "279 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19707,
            "unit": "allocs/op",
            "extra": "279 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 230325,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5694 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 230325,
            "unit": "ns/op",
            "extra": "5694 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5694 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5694 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7547,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "168069 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7547,
            "unit": "ns/op",
            "extra": "168069 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "168069 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "168069 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 232.2,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5151348 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 232.2,
            "unit": "ns/op",
            "extra": "5151348 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5151348 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5151348 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 345645,
            "unit": "ns/op\t  289965 B/op\t    3032 allocs/op",
            "extra": "3772 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 345645,
            "unit": "ns/op",
            "extra": "3772 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 289965,
            "unit": "B/op",
            "extra": "3772 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3772 times\n4 procs"
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
          "id": "3919e256bb5408f30dc9551b8a6b6890356e6d15",
          "message": "[AWS] Use library to create DNS string (#547)\n\nBetter than string formatting ourselves. Follows on from #541.",
          "timestamp": "2025-03-20T15:11:00Z",
          "tree_id": "ceb6a5ddaa14f3e951dc841a47a1b822105cd291",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/3919e256bb5408f30dc9551b8a6b6890356e6d15"
        },
        "date": 1742483511912,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4036267,
            "unit": "ns/op\t  700164 B/op\t   19671 allocs/op",
            "extra": "283 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4036267,
            "unit": "ns/op",
            "extra": "283 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 700164,
            "unit": "B/op",
            "extra": "283 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19671,
            "unit": "allocs/op",
            "extra": "283 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 236490,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5852 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 236490,
            "unit": "ns/op",
            "extra": "5852 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5852 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5852 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7618,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "212930 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7618,
            "unit": "ns/op",
            "extra": "212930 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "212930 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "212930 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 232.6,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5160686 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 232.6,
            "unit": "ns/op",
            "extra": "5160686 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5160686 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5160686 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 351763,
            "unit": "ns/op\t  291082 B/op\t    3041 allocs/op",
            "extra": "3946 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 351763,
            "unit": "ns/op",
            "extra": "3946 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291082,
            "unit": "B/op",
            "extra": "3946 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3041,
            "unit": "allocs/op",
            "extra": "3946 times\n4 procs"
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
          "id": "73534c56530549e3b7b3c320357f7d6f90854dae",
          "message": "Add richer Follower interface enabling clean shutdown & stats (#544)",
          "timestamp": "2025-03-20T15:14:33Z",
          "tree_id": "5d63e721362e9e19dcc716c93e9556b50318e8da",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/73534c56530549e3b7b3c320357f7d6f90854dae"
        },
        "date": 1742483726967,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4134623,
            "unit": "ns/op\t  702698 B/op\t   19697 allocs/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4134623,
            "unit": "ns/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 702698,
            "unit": "B/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 19697,
            "unit": "allocs/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 231328,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5818 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 231328,
            "unit": "ns/op",
            "extra": "5818 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5818 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5818 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7196,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "186790 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7196,
            "unit": "ns/op",
            "extra": "186790 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "186790 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "186790 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 233.6,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5130710 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 233.6,
            "unit": "ns/op",
            "extra": "5130710 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5130710 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5130710 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 349235,
            "unit": "ns/op\t  290124 B/op\t    3033 allocs/op",
            "extra": "3820 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 349235,
            "unit": "ns/op",
            "extra": "3820 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290124,
            "unit": "B/op",
            "extra": "3820 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3033,
            "unit": "allocs/op",
            "extra": "3820 times\n4 procs"
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
          "id": "ae724376e1ace4046767511c72c6006bde3ec87e",
          "message": "Antispam with lookup experiment (#508)",
          "timestamp": "2025-03-20T16:08:37Z",
          "tree_id": "93ad87036bc798cea7fb5f3adea6dd317472a4c7",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/ae724376e1ace4046767511c72c6006bde3ec87e"
        },
        "date": 1742486979192,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4337976,
            "unit": "ns/op\t  733052 B/op\t   21666 allocs/op",
            "extra": "277 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4337976,
            "unit": "ns/op",
            "extra": "277 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 733052,
            "unit": "B/op",
            "extra": "277 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21666,
            "unit": "allocs/op",
            "extra": "277 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 232246,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "4713 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 232246,
            "unit": "ns/op",
            "extra": "4713 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "4713 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "4713 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7548,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "176812 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7548,
            "unit": "ns/op",
            "extra": "176812 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "176812 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "176812 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 242.3,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "4777750 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 242.3,
            "unit": "ns/op",
            "extra": "4777750 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "4777750 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "4777750 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 360837,
            "unit": "ns/op\t  290128 B/op\t    3033 allocs/op",
            "extra": "3697 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 360837,
            "unit": "ns/op",
            "extra": "3697 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290128,
            "unit": "B/op",
            "extra": "3697 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3033,
            "unit": "allocs/op",
            "extra": "3697 times\n4 procs"
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
          "id": "6f93bc24dc8dfb201c9e9ee2eae9a15d276e0e5e",
          "message": "Remove id from IDSeq (#548)",
          "timestamp": "2025-03-24T14:21:14Z",
          "tree_id": "fbfa26f2d486c2569eb5edf019d39ad3809a7e11",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/6f93bc24dc8dfb201c9e9ee2eae9a15d276e0e5e"
        },
        "date": 1742826121427,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4437677,
            "unit": "ns/op\t  735459 B/op\t   21688 allocs/op",
            "extra": "247 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4437677,
            "unit": "ns/op",
            "extra": "247 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 735459,
            "unit": "B/op",
            "extra": "247 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21688,
            "unit": "allocs/op",
            "extra": "247 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 232144,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5010 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 232144,
            "unit": "ns/op",
            "extra": "5010 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5010 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5010 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7918,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "178809 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7918,
            "unit": "ns/op",
            "extra": "178809 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "178809 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "178809 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 232.3,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5117323 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 232.3,
            "unit": "ns/op",
            "extra": "5117323 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5117323 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5117323 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 355327,
            "unit": "ns/op\t  291282 B/op\t    3042 allocs/op",
            "extra": "3933 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 355327,
            "unit": "ns/op",
            "extra": "3933 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291282,
            "unit": "B/op",
            "extra": "3933 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3042,
            "unit": "allocs/op",
            "extra": "3933 times\n4 procs"
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
          "id": "9cba774fe136536b0940af62d5828250ccdcd60f",
          "message": "Because (#549)",
          "timestamp": "2025-03-24T15:54:51Z",
          "tree_id": "4248c2f56442377520c9f47b313f200076b3fe31",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/9cba774fe136536b0940af62d5828250ccdcd60f"
        },
        "date": 1742831744007,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4352920,
            "unit": "ns/op\t  735314 B/op\t   21690 allocs/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4352920,
            "unit": "ns/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 735314,
            "unit": "B/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21690,
            "unit": "allocs/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 227993,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5600 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 227993,
            "unit": "ns/op",
            "extra": "5600 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5600 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5600 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7513,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "181544 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7513,
            "unit": "ns/op",
            "extra": "181544 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "181544 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "181544 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 232,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5090844 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 232,
            "unit": "ns/op",
            "extra": "5090844 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5090844 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5090844 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 352987,
            "unit": "ns/op\t  290549 B/op\t    3037 allocs/op",
            "extra": "4168 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 352987,
            "unit": "ns/op",
            "extra": "4168 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290549,
            "unit": "B/op",
            "extra": "4168 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3037,
            "unit": "allocs/op",
            "extra": "4168 times\n4 procs"
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
          "id": "f59541bc3aaee13327d80ff4c4ee162dec56acc8",
          "message": "Bump the all-go-deps group with 4 updates (#551)\n\nBumps the all-go-deps group with 4 updates: [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2), [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2), [google.golang.org/api](https://github.com/googleapis/google-api-go-client) and [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql).\n\n\nUpdates `github.com/aws/aws-sdk-go-v2/config` from 1.29.9 to 1.29.10\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.9...config/v1.29.10)\n\nUpdates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.62 to 1.17.63\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.62...credentials/v1.17.63)\n\nUpdates `google.golang.org/api` from 0.226.0 to 0.227.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.226.0...v0.227.0)\n\nUpdates `github.com/go-sql-driver/mysql` from 1.9.0 to 1.9.1\n- [Release notes](https://github.com/go-sql-driver/mysql/releases)\n- [Changelog](https://github.com/go-sql-driver/mysql/blob/master/CHANGELOG.md)\n- [Commits](https://github.com/go-sql-driver/mysql/compare/v1.9.0...v1.9.1)\n\n---\nupdated-dependencies:\n- dependency-name: github.com/aws/aws-sdk-go-v2/config\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/credentials\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/go-sql-driver/mysql\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-03-24T21:21:27Z",
          "tree_id": "9131df2a85e205ba24e3e16edc7e9c6610821451",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/f59541bc3aaee13327d80ff4c4ee162dec56acc8"
        },
        "date": 1742851374167,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4884667,
            "unit": "ns/op\t  732928 B/op\t   21664 allocs/op",
            "extra": "254 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4884667,
            "unit": "ns/op",
            "extra": "254 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 732928,
            "unit": "B/op",
            "extra": "254 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21664,
            "unit": "allocs/op",
            "extra": "254 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 228748,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5368 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 228748,
            "unit": "ns/op",
            "extra": "5368 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5368 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5368 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1603,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "753711 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1603,
            "unit": "ns/op",
            "extra": "753711 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "753711 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "753711 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 234.8,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5120066 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 234.8,
            "unit": "ns/op",
            "extra": "5120066 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5120066 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5120066 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 344723,
            "unit": "ns/op\t  289989 B/op\t    3032 allocs/op",
            "extra": "3733 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 344723,
            "unit": "ns/op",
            "extra": "3733 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 289989,
            "unit": "B/op",
            "extra": "3733 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3733 times\n4 procs"
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
          "id": "9454d813f7fd380a0011b9553db7b087cc11ed48",
          "message": "Update lifecycle description doc (#550)",
          "timestamp": "2025-03-25T10:56:05Z",
          "tree_id": "3e49d3e01122081bb273a15dd84f284fb9911754",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/9454d813f7fd380a0011b9553db7b087cc11ed48"
        },
        "date": 1742900210159,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4323010,
            "unit": "ns/op\t  738350 B/op\t   21719 allocs/op",
            "extra": "247 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4323010,
            "unit": "ns/op",
            "extra": "247 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 738350,
            "unit": "B/op",
            "extra": "247 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21719,
            "unit": "allocs/op",
            "extra": "247 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 233191,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5829 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 233191,
            "unit": "ns/op",
            "extra": "5829 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5829 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5829 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 8045,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "149595 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 8045,
            "unit": "ns/op",
            "extra": "149595 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "149595 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "149595 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 231.7,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "4826727 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 231.7,
            "unit": "ns/op",
            "extra": "4826727 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "4826727 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "4826727 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 356251,
            "unit": "ns/op\t  289970 B/op\t    3032 allocs/op",
            "extra": "3739 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 356251,
            "unit": "ns/op",
            "extra": "3739 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 289970,
            "unit": "B/op",
            "extra": "3739 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3032,
            "unit": "allocs/op",
            "extra": "3739 times\n4 procs"
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
          "id": "f9d75d49174ba23ff61867a923e344bccb3ec1ee",
          "message": "Bump to golangci-lint@2.0.1 and fix new lint issues (#553)",
          "timestamp": "2025-03-25T12:09:58Z",
          "tree_id": "0f517e300d4ac4a0edf9363d4d6c704135a58648",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/f9d75d49174ba23ff61867a923e344bccb3ec1ee"
        },
        "date": 1742904644487,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4313262,
            "unit": "ns/op\t  731672 B/op\t   21652 allocs/op",
            "extra": "260 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4313262,
            "unit": "ns/op",
            "extra": "260 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 731672,
            "unit": "B/op",
            "extra": "260 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21652,
            "unit": "allocs/op",
            "extra": "260 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 232733,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "6021 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 232733,
            "unit": "ns/op",
            "extra": "6021 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "6021 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "6021 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7558,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "250604 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7558,
            "unit": "ns/op",
            "extra": "250604 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "250604 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "250604 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 231.8,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5180254 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 231.8,
            "unit": "ns/op",
            "extra": "5180254 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5180254 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5180254 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 354332,
            "unit": "ns/op\t  290381 B/op\t    3036 allocs/op",
            "extra": "4130 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 354332,
            "unit": "ns/op",
            "extra": "4130 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290381,
            "unit": "B/op",
            "extra": "4130 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "4130 times\n4 procs"
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
          "id": "a60c8d23e801e11eb7511a623d144d6a7da3d0d4",
          "message": "Bump the all-gha-deps group across 1 directory with 3 updates (#554)\n\nBumps the all-gha-deps group with 3 updates in the / directory: [actions/setup-go](https://github.com/actions/setup-go), [github/codeql-action](https://github.com/github/codeql-action) and [actions/upload-artifact](https://github.com/actions/upload-artifact).\n\n\nUpdates `actions/setup-go` from 5.3.0 to 5.4.0\n- [Release notes](https://github.com/actions/setup-go/releases)\n- [Commits](https://github.com/actions/setup-go/compare/f111f3307d8850f501ac008e886eec1fd1932a34...0aaccfd150d50ccaeb58ebd88d36e91967a5f35b)\n\nUpdates `github/codeql-action` from 3.28.11 to 3.28.13\n- [Release notes](https://github.com/github/codeql-action/releases)\n- [Changelog](https://github.com/github/codeql-action/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/github/codeql-action/compare/6bb031afdd8eb862ea3fc1848194185e076637e5...1b549b9259bda1cb5ddde3b41741a82a2d15a841)\n\nUpdates `actions/upload-artifact` from 4.6.1 to 4.6.2\n- [Release notes](https://github.com/actions/upload-artifact/releases)\n- [Commits](https://github.com/actions/upload-artifact/compare/4cec3d8aa04e39d1a68397de0c4cd6fb9dce8ec1...ea165f8d65b6e75b540449e92b4886f43607fa02)\n\n---\nupdated-dependencies:\n- dependency-name: actions/setup-go\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-gha-deps\n- dependency-name: github/codeql-action\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n- dependency-name: actions/upload-artifact\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-03-25T13:30:29Z",
          "tree_id": "bf841623539c7c552dcb98fb480784ff48af8b28",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/a60c8d23e801e11eb7511a623d144d6a7da3d0d4"
        },
        "date": 1742909483985,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4662341,
            "unit": "ns/op\t  734190 B/op\t   21679 allocs/op",
            "extra": "242 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4662341,
            "unit": "ns/op",
            "extra": "242 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 734190,
            "unit": "B/op",
            "extra": "242 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21679,
            "unit": "allocs/op",
            "extra": "242 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 225742,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5673 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 225742,
            "unit": "ns/op",
            "extra": "5673 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5673 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5673 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7844,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "264171 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7844,
            "unit": "ns/op",
            "extra": "264171 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "264171 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "264171 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 237.9,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5097708 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 237.9,
            "unit": "ns/op",
            "extra": "5097708 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5097708 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5097708 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 374203,
            "unit": "ns/op\t  290523 B/op\t    3035 allocs/op",
            "extra": "3572 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 374203,
            "unit": "ns/op",
            "extra": "3572 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290523,
            "unit": "B/op",
            "extra": "3572 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3035,
            "unit": "allocs/op",
            "extra": "3572 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "haydentherapper@users.noreply.github.com",
            "name": "Hayden B",
            "username": "haydentherapper"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "a95e1dcaf99650ef94228323451916e4defda587",
          "message": "Standardize logging when publishing a checkpoint (#555)\n\nAll storage backends will log when a new checkpoint is published.\n\nTiny change to standardize logging message for integration as well.\n\nSigned-off-by: Hayden B <8418760+haydentherapper@users.noreply.github.com>\nCo-authored-by: Hayden B <8418760+haydentherapper@users.noreply.github.com>",
          "timestamp": "2025-03-27T10:15:50Z",
          "tree_id": "aab8057e8cebb78e029ae4ae53f0fb0624c6e8c3",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/a95e1dcaf99650ef94228323451916e4defda587"
        },
        "date": 1743070606994,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4535867,
            "unit": "ns/op\t  733872 B/op\t   21673 allocs/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4535867,
            "unit": "ns/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 733872,
            "unit": "B/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21673,
            "unit": "allocs/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 225095,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5340 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 225095,
            "unit": "ns/op",
            "extra": "5340 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5340 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5340 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6953,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "158808 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6953,
            "unit": "ns/op",
            "extra": "158808 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "158808 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "158808 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 232.3,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5143842 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 232.3,
            "unit": "ns/op",
            "extra": "5143842 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5143842 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5143842 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 359411,
            "unit": "ns/op\t  291075 B/op\t    3041 allocs/op",
            "extra": "3922 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 359411,
            "unit": "ns/op",
            "extra": "3922 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 291075,
            "unit": "B/op",
            "extra": "3922 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3041,
            "unit": "allocs/op",
            "extra": "3922 times\n4 procs"
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
          "id": "c6564f46abba21f91a261ba80046213fecfac6d5",
          "message": "Add GCP antispam test (#556)",
          "timestamp": "2025-03-27T12:25:27Z",
          "tree_id": "3a60b5aad5133cc437942570ce815e0bd7e7407c",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/c6564f46abba21f91a261ba80046213fecfac6d5"
        },
        "date": 1743078378915,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4473279,
            "unit": "ns/op\t  733573 B/op\t   21669 allocs/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4473279,
            "unit": "ns/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 733573,
            "unit": "B/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21669,
            "unit": "allocs/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 225846,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5226 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 225846,
            "unit": "ns/op",
            "extra": "5226 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5226 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5226 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7364,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "182242 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7364,
            "unit": "ns/op",
            "extra": "182242 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "182242 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "182242 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 234,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5170838 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 234,
            "unit": "ns/op",
            "extra": "5170838 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5170838 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5170838 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 354756,
            "unit": "ns/op\t  290627 B/op\t    3038 allocs/op",
            "extra": "3888 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 354756,
            "unit": "ns/op",
            "extra": "3888 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 290627,
            "unit": "B/op",
            "extra": "3888 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3038,
            "unit": "allocs/op",
            "extra": "3888 times\n4 procs"
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
          "id": "887bee1aa49789d17e52c39d6e6fb03f716da954",
          "message": "Bump to Go 1.24 (#557)",
          "timestamp": "2025-03-27T14:53:21Z",
          "tree_id": "f52d7599662d69a2ced1bfd6bf2f10cb03748bc0",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/887bee1aa49789d17e52c39d6e6fb03f716da954"
        },
        "date": 1743087290700,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3524324,
            "unit": "ns/op\t  739871 B/op\t   21737 allocs/op",
            "extra": "338 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3524324,
            "unit": "ns/op",
            "extra": "338 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 739871,
            "unit": "B/op",
            "extra": "338 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21737,
            "unit": "allocs/op",
            "extra": "338 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 260143,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "4348 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 260143,
            "unit": "ns/op",
            "extra": "4348 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "4348 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "4348 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1708,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "721392 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1708,
            "unit": "ns/op",
            "extra": "721392 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "721392 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "721392 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 220.2,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5457030 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 220.2,
            "unit": "ns/op",
            "extra": "5457030 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5457030 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5457030 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 343268,
            "unit": "ns/op\t  274245 B/op\t    3035 allocs/op",
            "extra": "4036 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 343268,
            "unit": "ns/op",
            "extra": "4036 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 274245,
            "unit": "B/op",
            "extra": "4036 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3035,
            "unit": "allocs/op",
            "extra": "4036 times\n4 procs"
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
          "id": "1f8bbaa004c80811997262b6d5d24f222d7f0db4",
          "message": "Don't fail PR check for benchmarking (#560)",
          "timestamp": "2025-03-28T11:52:36Z",
          "tree_id": "cc79fbcd1464b205d58cc8681bedae6ac61466c2",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/1f8bbaa004c80811997262b6d5d24f222d7f0db4"
        },
        "date": 1743162802899,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 2957499,
            "unit": "ns/op\t  738171 B/op\t   21720 allocs/op",
            "extra": "346 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 2957499,
            "unit": "ns/op",
            "extra": "346 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 738171,
            "unit": "B/op",
            "extra": "346 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21720,
            "unit": "allocs/op",
            "extra": "346 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 228212,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "6037 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 228212,
            "unit": "ns/op",
            "extra": "6037 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "6037 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "6037 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7446,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "203894 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7446,
            "unit": "ns/op",
            "extra": "203894 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "203894 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "203894 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 222.7,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5350188 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 222.7,
            "unit": "ns/op",
            "extra": "5350188 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5350188 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5350188 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 333335,
            "unit": "ns/op\t  274658 B/op\t    3039 allocs/op",
            "extra": "3909 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 333335,
            "unit": "ns/op",
            "extra": "3909 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 274658,
            "unit": "B/op",
            "extra": "3909 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3039,
            "unit": "allocs/op",
            "extra": "3909 times\n4 procs"
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
          "id": "e649e5dd8a330f74b6b81aad66ef61d006daefb8",
          "message": "Implement StreamEntries on AWS and POSIX (#558)",
          "timestamp": "2025-03-28T14:57:52Z",
          "tree_id": "ace2e63e849526f9bc3047411cdd093584dadab6",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/e649e5dd8a330f74b6b81aad66ef61d006daefb8"
        },
        "date": 1743173921108,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3263672,
            "unit": "ns/op\t  744239 B/op\t   21773 allocs/op",
            "extra": "379 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3263672,
            "unit": "ns/op",
            "extra": "379 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 744239,
            "unit": "B/op",
            "extra": "379 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21773,
            "unit": "allocs/op",
            "extra": "379 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 222117,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5257 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 222117,
            "unit": "ns/op",
            "extra": "5257 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5257 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5257 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7366,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "211640 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7366,
            "unit": "ns/op",
            "extra": "211640 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "211640 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "211640 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 218.1,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5570312 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 218.1,
            "unit": "ns/op",
            "extra": "5570312 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5570312 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5570312 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 330517,
            "unit": "ns/op\t  274195 B/op\t    3034 allocs/op",
            "extra": "4054 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 330517,
            "unit": "ns/op",
            "extra": "4054 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 274195,
            "unit": "B/op",
            "extra": "4054 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3034,
            "unit": "allocs/op",
            "extra": "4054 times\n4 procs"
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
          "id": "61661a3e39541adec35d4eaaf02c983f2fff0e60",
          "message": "Add support for ephemeral testing logs, and use it in GCP antispam test. (#559)",
          "timestamp": "2025-03-28T15:55:38Z",
          "tree_id": "7029380c3b7daf6bf8dcb239b3df6bb103989a36",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/61661a3e39541adec35d4eaaf02c983f2fff0e60"
        },
        "date": 1743177392912,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3084341,
            "unit": "ns/op\t  735263 B/op\t   21690 allocs/op",
            "extra": "358 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3084341,
            "unit": "ns/op",
            "extra": "358 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 735263,
            "unit": "B/op",
            "extra": "358 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21690,
            "unit": "allocs/op",
            "extra": "358 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 228462,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5479 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 228462,
            "unit": "ns/op",
            "extra": "5479 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5479 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5479 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7401,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "181303 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7401,
            "unit": "ns/op",
            "extra": "181303 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "181303 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "181303 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 217.4,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5478313 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 217.4,
            "unit": "ns/op",
            "extra": "5478313 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5478313 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5478313 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 335859,
            "unit": "ns/op\t  274829 B/op\t    3040 allocs/op",
            "extra": "3951 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 335859,
            "unit": "ns/op",
            "extra": "3951 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 274829,
            "unit": "B/op",
            "extra": "3951 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3040,
            "unit": "allocs/op",
            "extra": "3951 times\n4 procs"
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
          "id": "59c55f4518798cd0265ba6f16f638f2ecc3afc2a",
          "message": "Bump the all-go-deps group with 6 updates (#569)\n\nBumps the all-go-deps group with 6 updates:\n\n| Package | From | To |\n| --- | --- | --- |\n| [cloud.google.com/go/spanner](https://github.com/googleapis/google-cloud-go) | `1.77.0` | `1.78.0` |\n| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.10` | `1.29.12` |\n| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.63` | `1.17.65` |\n| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.78.2` | `1.79.0` |\n| [google.golang.org/api](https://github.com/googleapis/google-api-go-client) | `0.227.0` | `0.228.0` |\n| [golang.org/x/net](https://github.com/golang/net) | `0.37.0` | `0.38.0` |\n\n\nUpdates `cloud.google.com/go/spanner` from 1.77.0 to 1.78.0\n- [Release notes](https://github.com/googleapis/google-cloud-go/releases)\n- [Changelog](https://github.com/googleapis/google-cloud-go/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-cloud-go/compare/spanner/v1.77.0...spanner/v1.78.0)\n\nUpdates `github.com/aws/aws-sdk-go-v2/config` from 1.29.10 to 1.29.12\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.10...config/v1.29.12)\n\nUpdates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.63 to 1.17.65\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.63...credentials/v1.17.65)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.78.2 to 1.79.0\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.78.2...service/s3/v1.79.0)\n\nUpdates `google.golang.org/api` from 0.227.0 to 0.228.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.227.0...v0.228.0)\n\nUpdates `golang.org/x/net` from 0.37.0 to 0.38.0\n- [Commits](https://github.com/golang/net/compare/v0.37.0...v0.38.0)\n\n---\nupdated-dependencies:\n- dependency-name: cloud.google.com/go/spanner\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/config\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/credentials\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/net\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-04-01T22:08:22+01:00",
          "tree_id": "6f33a0c3448887b2fc346b004966751660e4de6d",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/59c55f4518798cd0265ba6f16f638f2ecc3afc2a"
        },
        "date": 1743541786338,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3433687,
            "unit": "ns/op\t  739635 B/op\t   21734 allocs/op",
            "extra": "373 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3433687,
            "unit": "ns/op",
            "extra": "373 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 739635,
            "unit": "B/op",
            "extra": "373 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21734,
            "unit": "allocs/op",
            "extra": "373 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 246183,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5888 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 246183,
            "unit": "ns/op",
            "extra": "5888 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5888 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5888 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1835,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "559911 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1835,
            "unit": "ns/op",
            "extra": "559911 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "559911 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "559911 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 215.4,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5535955 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 215.4,
            "unit": "ns/op",
            "extra": "5535955 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5535955 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5535955 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 338997,
            "unit": "ns/op\t  274371 B/op\t    3036 allocs/op",
            "extra": "4161 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 338997,
            "unit": "ns/op",
            "extra": "4161 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 274371,
            "unit": "B/op",
            "extra": "4161 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3036,
            "unit": "allocs/op",
            "extra": "4161 times\n4 procs"
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
          "id": "43453cfcba954aeaaaadbc2c9f6f597767f85f52",
          "message": "[AWS] Antispam (#568)\n\nThis adds experimental support for antispam (dedupe) for AWS-hosted\nlogs. The hashes and indices of seen entries are stored in a MySQL\ndatabase.\n\nThis is based heavily on the GCP Spanner version, and still maintains\nits core way of operation (i.e. filling in the duplicates by using the\nFollower pattern).\n\nTowards #526.",
          "timestamp": "2025-04-02T15:53:51+01:00",
          "tree_id": "25b15c524aaa1c0c16d586e246d61720fb565b4b",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/43453cfcba954aeaaaadbc2c9f6f597767f85f52"
        },
        "date": 1743605671811,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 754204,
            "unit": "ns/op\t  721118 B/op\t   21548 allocs/op",
            "extra": "1428 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 754204,
            "unit": "ns/op",
            "extra": "1428 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 721118,
            "unit": "B/op",
            "extra": "1428 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21548,
            "unit": "allocs/op",
            "extra": "1428 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136678,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8437 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136678,
            "unit": "ns/op",
            "extra": "8437 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8437 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8437 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1672,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "648116 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1672,
            "unit": "ns/op",
            "extra": "648116 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "648116 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "648116 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 220.9,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5485570 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 220.9,
            "unit": "ns/op",
            "extra": "5485570 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5485570 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5485570 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 354024,
            "unit": "ns/op\t  273537 B/op\t    3027 allocs/op",
            "extra": "3442 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 354024,
            "unit": "ns/op",
            "extra": "3442 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 273537,
            "unit": "B/op",
            "extra": "3442 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3027,
            "unit": "allocs/op",
            "extra": "3442 times\n4 procs"
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
          "id": "adcd1fedb37f498016318abe431cfdcae4ecacf9",
          "message": "[AWS] Set cache headers for artifacts in s3 (#571)\n\nThis fixes #399 for AWS, and is the last storage implementation so fixes the whole issue.",
          "timestamp": "2025-04-02T15:59:36+01:00",
          "tree_id": "2aa294b6bd3e26970274553e6071325d59f2b9b6",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/adcd1fedb37f498016318abe431cfdcae4ecacf9"
        },
        "date": 1743606017822,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 842436,
            "unit": "ns/op\t  723188 B/op\t   21570 allocs/op",
            "extra": "1653 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 842436,
            "unit": "ns/op",
            "extra": "1653 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 723188,
            "unit": "B/op",
            "extra": "1653 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21570,
            "unit": "allocs/op",
            "extra": "1653 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 140616,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8413 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 140616,
            "unit": "ns/op",
            "extra": "8413 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8413 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8413 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1679,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "617508 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1679,
            "unit": "ns/op",
            "extra": "617508 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "617508 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "617508 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 224,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5491299 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 224,
            "unit": "ns/op",
            "extra": "5491299 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5491299 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5491299 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 334110,
            "unit": "ns/op\t  274877 B/op\t    3041 allocs/op",
            "extra": "4232 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 334110,
            "unit": "ns/op",
            "extra": "4232 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 274877,
            "unit": "B/op",
            "extra": "4232 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3041,
            "unit": "allocs/op",
            "extra": "4232 times\n4 procs"
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
          "id": "b53519faa0ba7f8a12ddcb1fd4b39f0ab14288f9",
          "message": "Fixed go.mod file to remove exclude (#572)\n\nExclude directives break running commands in the repository in the\nnormal go run way:\n\n```shell\ngo run github.com/transparency-dev/trillian-tessera/cmd/examples/posix-oneshot@main --storage_dir=/tmp/mylog --initialise\n\nThe go.mod file for the module providing named packages contains one or\nmore exclude directives. It must not contain directives that would cause\nit to be interpreted differently than if it were the main module.\n```\n\nI've run `go mod tidy` after this and it seems OK. We'll need to keep an\neye for dependabot making a mess.",
          "timestamp": "2025-04-02T16:42:37+01:00",
          "tree_id": "0dde22936edc0ca337dbcebf7da9d2f4d193c89c",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/b53519faa0ba7f8a12ddcb1fd4b39f0ab14288f9"
        },
        "date": 1743608592799,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 749167,
            "unit": "ns/op\t  721727 B/op\t   21554 allocs/op",
            "extra": "1616 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 749167,
            "unit": "ns/op",
            "extra": "1616 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 721727,
            "unit": "B/op",
            "extra": "1616 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21554,
            "unit": "allocs/op",
            "extra": "1616 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 135883,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8715 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 135883,
            "unit": "ns/op",
            "extra": "8715 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8715 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8715 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1650,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "702658 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1650,
            "unit": "ns/op",
            "extra": "702658 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "702658 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "702658 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 217,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5554809 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 217,
            "unit": "ns/op",
            "extra": "5554809 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5554809 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5554809 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 335200,
            "unit": "ns/op\t  274314 B/op\t    3035 allocs/op",
            "extra": "4144 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 335200,
            "unit": "ns/op",
            "extra": "4144 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 274314,
            "unit": "B/op",
            "extra": "4144 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3035,
            "unit": "allocs/op",
            "extra": "4144 times\n4 procs"
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
          "id": "aeae2c169b0713c03828438110af5e89179337ec",
          "message": "[Antispam] Faster tests (#573)\n\nThese tests were waiting for 10s before finishing. Publishing the checkpoint faster allows this to exit in 1.5s each, instead of 10s each.",
          "timestamp": "2025-04-03T10:06:30+01:00",
          "tree_id": "1b3f7dc13211970d3a0744e32124253411ab1cd2",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/aeae2c169b0713c03828438110af5e89179337ec"
        },
        "date": 1743671225368,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 737580,
            "unit": "ns/op\t  721362 B/op\t   21550 allocs/op",
            "extra": "1621 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 737580,
            "unit": "ns/op",
            "extra": "1621 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 721362,
            "unit": "B/op",
            "extra": "1621 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21550,
            "unit": "allocs/op",
            "extra": "1621 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 138321,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8698 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 138321,
            "unit": "ns/op",
            "extra": "8698 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8698 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8698 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1743,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "670432 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1743,
            "unit": "ns/op",
            "extra": "670432 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "670432 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "670432 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 221.4,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5032029 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 221.4,
            "unit": "ns/op",
            "extra": "5032029 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5032029 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5032029 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 340257,
            "unit": "ns/op\t  274132 B/op\t    3034 allocs/op",
            "extra": "3843 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 340257,
            "unit": "ns/op",
            "extra": "3843 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 274132,
            "unit": "B/op",
            "extra": "3843 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3034,
            "unit": "allocs/op",
            "extra": "3843 times\n4 procs"
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
          "id": "3f5fed6579750a9b86def76e8ee48406252e4c3a",
          "message": "Move witnessing fully into lifecycle layer (#578)",
          "timestamp": "2025-04-03T13:00:23+01:00",
          "tree_id": "d8b94c9cd9a9aaad9d9063167bee824ece8be846",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/3f5fed6579750a9b86def76e8ee48406252e4c3a"
        },
        "date": 1743681656145,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 748937,
            "unit": "ns/op\t  721878 B/op\t   21555 allocs/op",
            "extra": "1579 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 748937,
            "unit": "ns/op",
            "extra": "1579 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 721878,
            "unit": "B/op",
            "extra": "1579 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21555,
            "unit": "allocs/op",
            "extra": "1579 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 137632,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8744 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 137632,
            "unit": "ns/op",
            "extra": "8744 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8744 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8744 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1724,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "637129 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1724,
            "unit": "ns/op",
            "extra": "637129 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "637129 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "637129 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 229.4,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "4898850 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 229.4,
            "unit": "ns/op",
            "extra": "4898850 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "4898850 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "4898850 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 336081,
            "unit": "ns/op\t  274193 B/op\t    3033 allocs/op",
            "extra": "3662 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 336081,
            "unit": "ns/op",
            "extra": "3662 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 274193,
            "unit": "B/op",
            "extra": "3662 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3033,
            "unit": "allocs/op",
            "extra": "3662 times\n4 procs"
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
          "id": "0ab22ded2ba42427fb5c6ca101f635807a88c5e4",
          "message": "Add antispam overview (#579)",
          "timestamp": "2025-04-07T11:47:43+01:00",
          "tree_id": "08ea718a5de742d8ee4d31844c14620956b559b4",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/0ab22ded2ba42427fb5c6ca101f635807a88c5e4"
        },
        "date": 1744022896451,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 837345,
            "unit": "ns/op\t  723296 B/op\t   21571 allocs/op",
            "extra": "1364 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 837345,
            "unit": "ns/op",
            "extra": "1364 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 723296,
            "unit": "B/op",
            "extra": "1364 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21571,
            "unit": "allocs/op",
            "extra": "1364 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136875,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8548 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136875,
            "unit": "ns/op",
            "extra": "8548 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8548 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8548 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1656,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "668490 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1656,
            "unit": "ns/op",
            "extra": "668490 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "668490 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "668490 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 213.7,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "4679380 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 213.7,
            "unit": "ns/op",
            "extra": "4679380 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "4679380 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "4679380 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 332695,
            "unit": "ns/op\t  274679 B/op\t    3039 allocs/op",
            "extra": "4209 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 332695,
            "unit": "ns/op",
            "extra": "4209 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 274679,
            "unit": "B/op",
            "extra": "4209 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3039,
            "unit": "allocs/op",
            "extra": "4209 times\n4 procs"
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
          "id": "3666a679a8153eea7612c68ca8ed2146cdb2b7dc",
          "message": "Add initial OTel instrumentation to Tessera (#580)",
          "timestamp": "2025-04-07T17:30:10+01:00",
          "tree_id": "6de763a62ee592e3752fa0bec6ed91e5039aa78c",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/3666a679a8153eea7612c68ca8ed2146cdb2b7dc"
        },
        "date": 1744043494960,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3125490,
            "unit": "ns/op\t  741984 B/op\t   21757 allocs/op",
            "extra": "370 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3125490,
            "unit": "ns/op",
            "extra": "370 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 741984,
            "unit": "B/op",
            "extra": "370 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21757,
            "unit": "allocs/op",
            "extra": "370 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 235281,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5862 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 235281,
            "unit": "ns/op",
            "extra": "5862 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5862 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5862 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7452,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "186595 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7452,
            "unit": "ns/op",
            "extra": "186595 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "186595 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "186595 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 219,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5470986 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 219,
            "unit": "ns/op",
            "extra": "5470986 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5470986 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5470986 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 337061,
            "unit": "ns/op\t  279074 B/op\t    3081 allocs/op",
            "extra": "4094 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 337061,
            "unit": "ns/op",
            "extra": "4094 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279074,
            "unit": "B/op",
            "extra": "4094 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3081,
            "unit": "allocs/op",
            "extra": "4094 times\n4 procs"
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
          "id": "447e3d34f62bf4a5dbcf31d9c79e134fccb393a8",
          "message": "Bump github/codeql-action in the all-gha-deps group (#583)\n\nBumps the all-gha-deps group with 1 update: [github/codeql-action](https://github.com/github/codeql-action).\n\n\nUpdates `github/codeql-action` from 3.28.13 to 3.28.14\n- [Release notes](https://github.com/github/codeql-action/releases)\n- [Changelog](https://github.com/github/codeql-action/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/github/codeql-action/compare/1b549b9259bda1cb5ddde3b41741a82a2d15a841...fc7e4a0fa01c3cca5fd6a1fddec5c0740c977aa2)\n\n---\nupdated-dependencies:\n- dependency-name: github/codeql-action\n  dependency-version: 3.28.14\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-04-08T10:38:21+01:00",
          "tree_id": "7791af344168093bc19dcd23a6441f807e1672c0",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/447e3d34f62bf4a5dbcf31d9c79e134fccb393a8"
        },
        "date": 1744105139422,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3049284,
            "unit": "ns/op\t  736047 B/op\t   21696 allocs/op",
            "extra": "364 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3049284,
            "unit": "ns/op",
            "extra": "364 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 736047,
            "unit": "B/op",
            "extra": "364 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21696,
            "unit": "allocs/op",
            "extra": "364 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 229487,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5827 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 229487,
            "unit": "ns/op",
            "extra": "5827 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5827 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5827 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7367,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "142556 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7367,
            "unit": "ns/op",
            "extra": "142556 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "142556 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "142556 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 224.9,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5452322 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 224.9,
            "unit": "ns/op",
            "extra": "5452322 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5452322 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5452322 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 347612,
            "unit": "ns/op\t  278689 B/op\t    3077 allocs/op",
            "extra": "3753 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 347612,
            "unit": "ns/op",
            "extra": "3753 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 278689,
            "unit": "B/op",
            "extra": "3753 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3077,
            "unit": "allocs/op",
            "extra": "3753 times\n4 procs"
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
          "id": "4f07f5998dadb7a6b8691b825ad1e6104844dd7f",
          "message": "Add metrics around integration (#581)",
          "timestamp": "2025-04-08T11:02:38+01:00",
          "tree_id": "307cb791bf759dde4ea5821d7c3775401b116178",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/4f07f5998dadb7a6b8691b825ad1e6104844dd7f"
        },
        "date": 1744106601703,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3133846,
            "unit": "ns/op\t  740032 B/op\t   21737 allocs/op",
            "extra": "338 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3133846,
            "unit": "ns/op",
            "extra": "338 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 740032,
            "unit": "B/op",
            "extra": "338 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21737,
            "unit": "allocs/op",
            "extra": "338 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 249584,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "6294 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 249584,
            "unit": "ns/op",
            "extra": "6294 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "6294 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "6294 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7465,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "145269 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7465,
            "unit": "ns/op",
            "extra": "145269 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "145269 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "145269 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 217.7,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5482042 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 217.7,
            "unit": "ns/op",
            "extra": "5482042 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5482042 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5482042 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 351391,
            "unit": "ns/op\t  279268 B/op\t    3083 allocs/op",
            "extra": "4009 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 351391,
            "unit": "ns/op",
            "extra": "4009 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279268,
            "unit": "B/op",
            "extra": "4009 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3083,
            "unit": "allocs/op",
            "extra": "4009 times\n4 procs"
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
          "id": "ffe4098c824902d50abe835c7a94f2f94f176b02",
          "message": "GCP: push metrics & traces natively (#584)",
          "timestamp": "2025-04-08T12:00:34+01:00",
          "tree_id": "d5463b6d6885b639cda43d18b131ef2621ee2dd5",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/ffe4098c824902d50abe835c7a94f2f94f176b02"
        },
        "date": 1744110119518,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 4024745,
            "unit": "ns/op\t  744333 B/op\t   21783 allocs/op",
            "extra": "364 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 4024745,
            "unit": "ns/op",
            "extra": "364 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 744333,
            "unit": "B/op",
            "extra": "364 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21783,
            "unit": "allocs/op",
            "extra": "364 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 231271,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5558 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 231271,
            "unit": "ns/op",
            "extra": "5558 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5558 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5558 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1680,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "717597 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1680,
            "unit": "ns/op",
            "extra": "717597 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "717597 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "717597 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 219.7,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5409598 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 219.7,
            "unit": "ns/op",
            "extra": "5409598 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5409598 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5409598 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 350749,
            "unit": "ns/op\t  279608 B/op\t    3084 allocs/op",
            "extra": "3603 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 350749,
            "unit": "ns/op",
            "extra": "3603 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279608,
            "unit": "B/op",
            "extra": "3603 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3084,
            "unit": "allocs/op",
            "extra": "3603 times\n4 procs"
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
          "id": "a5ec4dd7371116c3b30134e1e5291c15c933e21d",
          "message": "[Docs] Improve docs for common lifecycles (#585)\n\nImproved docs primarily for appender and migrate lifecycle modes.\n\nTowards #574.",
          "timestamp": "2025-04-08T15:42:10+01:00",
          "tree_id": "d6878ec0d7d70b64829b7a5f8e95eee33b182361",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/a5ec4dd7371116c3b30134e1e5291c15c933e21d"
        },
        "date": 1744123369324,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3210211,
            "unit": "ns/op\t  736759 B/op\t   21698 allocs/op",
            "extra": "349 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3210211,
            "unit": "ns/op",
            "extra": "349 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 736759,
            "unit": "B/op",
            "extra": "349 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21698,
            "unit": "allocs/op",
            "extra": "349 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 240905,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5733 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 240905,
            "unit": "ns/op",
            "extra": "5733 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5733 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5733 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6982,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "168558 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6982,
            "unit": "ns/op",
            "extra": "168558 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "168558 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "168558 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 220.9,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5438185 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 220.9,
            "unit": "ns/op",
            "extra": "5438185 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5438185 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5438185 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 344870,
            "unit": "ns/op\t  279248 B/op\t    3083 allocs/op",
            "extra": "4012 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 344870,
            "unit": "ns/op",
            "extra": "4012 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279248,
            "unit": "B/op",
            "extra": "4012 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3083,
            "unit": "allocs/op",
            "extra": "4012 times\n4 procs"
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
          "id": "d1cc7c65696e186b60961543bd692e2fc8d1c346",
          "message": "Add follower metrics (#587)",
          "timestamp": "2025-04-08T16:31:34+01:00",
          "tree_id": "6778badc88b396e1207de40e7780f07e8e8e731a",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/d1cc7c65696e186b60961543bd692e2fc8d1c346"
        },
        "date": 1744126335160,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3291133,
            "unit": "ns/op\t  741766 B/op\t   21754 allocs/op",
            "extra": "384 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3291133,
            "unit": "ns/op",
            "extra": "384 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 741766,
            "unit": "B/op",
            "extra": "384 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21754,
            "unit": "allocs/op",
            "extra": "384 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 226617,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5685 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 226617,
            "unit": "ns/op",
            "extra": "5685 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5685 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5685 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7552,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "161106 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7552,
            "unit": "ns/op",
            "extra": "161106 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "161106 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "161106 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 221.5,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5575887 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 221.5,
            "unit": "ns/op",
            "extra": "5575887 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5575887 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5575887 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 343360,
            "unit": "ns/op\t  278796 B/op\t    3078 allocs/op",
            "extra": "3799 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 343360,
            "unit": "ns/op",
            "extra": "3799 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 278796,
            "unit": "B/op",
            "extra": "3799 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3078,
            "unit": "allocs/op",
            "extra": "3799 times\n4 procs"
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
          "id": "a650aa01f2a408ccbbe19bdba1c188cdb5bcd9e8",
          "message": "Bump the all-go-deps group across 1 directory with 9 updates (#586)\n\nBumps the all-go-deps group with 7 updates in the / directory:\n\n| Package | From | To |\n| --- | --- | --- |\n| [cloud.google.com/go/spanner](https://github.com/googleapis/google-cloud-go) | `1.78.0` | `1.79.0` |\n| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.12` | `1.29.13` |\n| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.79.0` | `1.79.1` |\n| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.71.0` | `1.71.1` |\n| [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) | `1.9.1` | `1.9.2` |\n| [golang.org/x/crypto](https://github.com/golang/crypto) | `0.36.0` | `0.37.0` |\n| [golang.org/x/net](https://github.com/golang/net) | `0.38.0` | `0.39.0` |\n\n\n\nUpdates `cloud.google.com/go/spanner` from 1.78.0 to 1.79.0\n- [Release notes](https://github.com/googleapis/google-cloud-go/releases)\n- [Changelog](https://github.com/googleapis/google-cloud-go/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-cloud-go/compare/spanner/v1.78.0...spanner/v1.79.0)\n\nUpdates `github.com/aws/aws-sdk-go-v2/config` from 1.29.12 to 1.29.13\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.12...config/v1.29.13)\n\nUpdates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.65 to 1.17.66\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.65...credentials/v1.17.66)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.79.0 to 1.79.1\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.79.0...service/s3/v1.79.1)\n\nUpdates `google.golang.org/grpc` from 1.71.0 to 1.71.1\n- [Release notes](https://github.com/grpc/grpc-go/releases)\n- [Commits](https://github.com/grpc/grpc-go/compare/v1.71.0...v1.71.1)\n\nUpdates `github.com/go-sql-driver/mysql` from 1.9.1 to 1.9.2\n- [Release notes](https://github.com/go-sql-driver/mysql/releases)\n- [Changelog](https://github.com/go-sql-driver/mysql/blob/master/CHANGELOG.md)\n- [Commits](https://github.com/go-sql-driver/mysql/compare/v1.9.1...v1.9.2)\n\nUpdates `golang.org/x/crypto` from 0.36.0 to 0.37.0\n- [Commits](https://github.com/golang/crypto/compare/v0.36.0...v0.37.0)\n\nUpdates `golang.org/x/net` from 0.38.0 to 0.39.0\n- [Commits](https://github.com/golang/net/compare/v0.38.0...v0.39.0)\n\nUpdates `golang.org/x/sync` from 0.12.0 to 0.13.0\n- [Commits](https://github.com/golang/sync/compare/v0.12.0...v0.13.0)\n\n---\nupdated-dependencies:\n- dependency-name: cloud.google.com/go/spanner\n  dependency-version: 1.79.0\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/config\n  dependency-version: 1.29.13\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/credentials\n  dependency-version: 1.17.66\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-version: 1.79.1\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/grpc\n  dependency-version: 1.71.1\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/go-sql-driver/mysql\n  dependency-version: 1.9.2\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/crypto\n  dependency-version: 0.37.0\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/net\n  dependency-version: 0.39.0\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n- dependency-name: golang.org/x/sync\n  dependency-version: 0.13.0\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-04-08T16:39:12+01:00",
          "tree_id": "7f20b1b848f73ab9949cab94a3e2e86d76b76851",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/a650aa01f2a408ccbbe19bdba1c188cdb5bcd9e8"
        },
        "date": 1744126834528,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3591020,
            "unit": "ns/op\t  739718 B/op\t   21735 allocs/op",
            "extra": "343 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3591020,
            "unit": "ns/op",
            "extra": "343 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 739718,
            "unit": "B/op",
            "extra": "343 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21735,
            "unit": "allocs/op",
            "extra": "343 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 234876,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "4884 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 234876,
            "unit": "ns/op",
            "extra": "4884 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "4884 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "4884 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1651,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "734707 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1651,
            "unit": "ns/op",
            "extra": "734707 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "734707 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "734707 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 216.7,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5406291 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 216.7,
            "unit": "ns/op",
            "extra": "5406291 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5406291 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5406291 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 335369,
            "unit": "ns/op\t  279922 B/op\t    3089 allocs/op",
            "extra": "4243 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 335369,
            "unit": "ns/op",
            "extra": "4243 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279922,
            "unit": "B/op",
            "extra": "4243 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3089,
            "unit": "allocs/op",
            "extra": "4243 times\n4 procs"
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
          "id": "1a98628ccc4f625b297be3d7d52a7e9443cf746e",
          "message": "[Docs] Freezing/Deleting/Sharded logs (#590)\n\nDocumentation for performing these actions. This fixes #574.",
          "timestamp": "2025-04-09T17:14:07+01:00",
          "tree_id": "40bc21cdb99904d61265a6320649d8b7c062f8cc",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/1a98628ccc4f625b297be3d7d52a7e9443cf746e"
        },
        "date": 1744215291943,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3297086,
            "unit": "ns/op\t  735158 B/op\t   21689 allocs/op",
            "extra": "331 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3297086,
            "unit": "ns/op",
            "extra": "331 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 735158,
            "unit": "B/op",
            "extra": "331 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21689,
            "unit": "allocs/op",
            "extra": "331 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 239601,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5017 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 239601,
            "unit": "ns/op",
            "extra": "5017 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5017 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5017 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7607,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "169916 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7607,
            "unit": "ns/op",
            "extra": "169916 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "169916 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "169916 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 218.6,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5408488 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 218.6,
            "unit": "ns/op",
            "extra": "5408488 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5408488 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5408488 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 343202,
            "unit": "ns/op\t  278907 B/op\t    3079 allocs/op",
            "extra": "3823 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 343202,
            "unit": "ns/op",
            "extra": "3823 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 278907,
            "unit": "B/op",
            "extra": "3823 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3079,
            "unit": "allocs/op",
            "extra": "3823 times\n4 procs"
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
          "id": "7d560b91fb445135b6ee67df7964249210638f9c",
          "message": "AWS OTel support (#591)\n\n* Add AWS ADOT support\n* Add terraform support for ADOT",
          "timestamp": "2025-04-14T15:04:28+01:00",
          "tree_id": "3e701cc9a5818f83653a811b420d6d8cc8f7f319",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/7d560b91fb445135b6ee67df7964249210638f9c"
        },
        "date": 1744639555573,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3207388,
            "unit": "ns/op\t  740039 B/op\t   21741 allocs/op",
            "extra": "406 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3207388,
            "unit": "ns/op",
            "extra": "406 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 740039,
            "unit": "B/op",
            "extra": "406 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21741,
            "unit": "allocs/op",
            "extra": "406 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 230764,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5406 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 230764,
            "unit": "ns/op",
            "extra": "5406 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5406 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5406 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7780,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "185046 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7780,
            "unit": "ns/op",
            "extra": "185046 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "185046 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "185046 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 220.2,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5156864 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 220.2,
            "unit": "ns/op",
            "extra": "5156864 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5156864 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5156864 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 346847,
            "unit": "ns/op\t  279064 B/op\t    3081 allocs/op",
            "extra": "4072 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 346847,
            "unit": "ns/op",
            "extra": "4072 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279064,
            "unit": "B/op",
            "extra": "4072 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3081,
            "unit": "allocs/op",
            "extra": "4072 times\n4 procs"
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
          "id": "15bd4550f8825490b4506b095dca6d98366b76c3",
          "message": "GCP Add OTel automatic resource detection (#593)",
          "timestamp": "2025-04-14T16:24:47+01:00",
          "tree_id": "e8c320b8d35e1ef6492c39d817d07800419c8446",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/15bd4550f8825490b4506b095dca6d98366b76c3"
        },
        "date": 1744644320480,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 743141,
            "unit": "ns/op\t  721676 B/op\t   21553 allocs/op",
            "extra": "1592 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 743141,
            "unit": "ns/op",
            "extra": "1592 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 721676,
            "unit": "B/op",
            "extra": "1592 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21553,
            "unit": "allocs/op",
            "extra": "1592 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 136825,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "8475 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 136825,
            "unit": "ns/op",
            "extra": "8475 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "8475 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "8475 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 1686,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "651393 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 1686,
            "unit": "ns/op",
            "extra": "651393 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "651393 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "651393 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 223.1,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5485663 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 223.1,
            "unit": "ns/op",
            "extra": "5485663 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5485663 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5485663 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 344946,
            "unit": "ns/op\t  279146 B/op\t    3082 allocs/op",
            "extra": "4036 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 344946,
            "unit": "ns/op",
            "extra": "4036 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279146,
            "unit": "B/op",
            "extra": "4036 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3082,
            "unit": "allocs/op",
            "extra": "4036 times\n4 procs"
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
          "id": "542b70112b76ecfde819a524e2c488a88c79ca93",
          "message": "Bump the all-go-deps group with 4 updates (#597)\n\nBumps the all-go-deps group with 4 updates: [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2), [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2), [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) and [google.golang.org/api](https://github.com/googleapis/google-api-go-client).\n\n\nUpdates `github.com/aws/aws-sdk-go-v2/config` from 1.29.13 to 1.29.14\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.13...config/v1.29.14)\n\nUpdates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.66 to 1.17.67\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.66...credentials/v1.17.67)\n\nUpdates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.79.1 to 1.79.2\n- [Release notes](https://github.com/aws/aws-sdk-go-v2/releases)\n- [Changelog](https://github.com/aws/aws-sdk-go-v2/blob/main/changelog-template.json)\n- [Commits](https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.79.1...service/s3/v1.79.2)\n\nUpdates `google.golang.org/api` from 0.228.0 to 0.229.0\n- [Release notes](https://github.com/googleapis/google-api-go-client/releases)\n- [Changelog](https://github.com/googleapis/google-api-go-client/blob/main/CHANGES.md)\n- [Commits](https://github.com/googleapis/google-api-go-client/compare/v0.228.0...v0.229.0)\n\n---\nupdated-dependencies:\n- dependency-name: github.com/aws/aws-sdk-go-v2/config\n  dependency-version: 1.29.14\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/credentials\n  dependency-version: 1.17.67\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: github.com/aws/aws-sdk-go-v2/service/s3\n  dependency-version: 1.79.2\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-go-deps\n- dependency-name: google.golang.org/api\n  dependency-version: 0.229.0\n  dependency-type: direct:production\n  update-type: version-update:semver-minor\n  dependency-group: all-go-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-04-15T11:20:56+01:00",
          "tree_id": "19abf03282a144dd5cc4312be7e3317cfc6e355e",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/542b70112b76ecfde819a524e2c488a88c79ca93"
        },
        "date": 1744712545006,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3201864,
            "unit": "ns/op\t  736789 B/op\t   21707 allocs/op",
            "extra": "339 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3201864,
            "unit": "ns/op",
            "extra": "339 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 736789,
            "unit": "B/op",
            "extra": "339 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21707,
            "unit": "allocs/op",
            "extra": "339 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 234311,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "4870 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 234311,
            "unit": "ns/op",
            "extra": "4870 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "4870 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "4870 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7557,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "156898 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7557,
            "unit": "ns/op",
            "extra": "156898 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "156898 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "156898 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 217.9,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5262888 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 217.9,
            "unit": "ns/op",
            "extra": "5262888 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5262888 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5262888 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 343589,
            "unit": "ns/op\t  279627 B/op\t    3085 allocs/op",
            "extra": "3963 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 343589,
            "unit": "ns/op",
            "extra": "3963 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279627,
            "unit": "B/op",
            "extra": "3963 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3085,
            "unit": "allocs/op",
            "extra": "3963 times\n4 procs"
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
          "id": "8463536cff45ad097fd2310b75e60a12df88ec57",
          "message": "Bump github/codeql-action in the all-gha-deps group (#596)\n\nBumps the all-gha-deps group with 1 update: [github/codeql-action](https://github.com/github/codeql-action).\n\n\nUpdates `github/codeql-action` from 3.28.14 to 3.28.15\n- [Release notes](https://github.com/github/codeql-action/releases)\n- [Changelog](https://github.com/github/codeql-action/blob/main/CHANGELOG.md)\n- [Commits](https://github.com/github/codeql-action/compare/fc7e4a0fa01c3cca5fd6a1fddec5c0740c977aa2...45775bd8235c68ba998cffa5171334d58593da47)\n\n---\nupdated-dependencies:\n- dependency-name: github/codeql-action\n  dependency-version: 3.28.15\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n  dependency-group: all-gha-deps\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>\nCo-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>",
          "timestamp": "2025-04-15T11:32:08+01:00",
          "tree_id": "c57500c1b99c8fccf969bec759bdcf63d1e1ef80",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/8463536cff45ad097fd2310b75e60a12df88ec57"
        },
        "date": 1744713173043,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3083274,
            "unit": "ns/op\t  739210 B/op\t   21729 allocs/op",
            "extra": "368 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3083274,
            "unit": "ns/op",
            "extra": "368 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 739210,
            "unit": "B/op",
            "extra": "368 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21729,
            "unit": "allocs/op",
            "extra": "368 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 227492,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5250 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 227492,
            "unit": "ns/op",
            "extra": "5250 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5250 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5250 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7236,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "169554 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7236,
            "unit": "ns/op",
            "extra": "169554 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "169554 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "169554 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 226.5,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5514166 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 226.5,
            "unit": "ns/op",
            "extra": "5514166 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5514166 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5514166 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 340910,
            "unit": "ns/op\t  279385 B/op\t    3084 allocs/op",
            "extra": "3992 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 340910,
            "unit": "ns/op",
            "extra": "3992 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279385,
            "unit": "B/op",
            "extra": "3992 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3084,
            "unit": "allocs/op",
            "extra": "3992 times\n4 procs"
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
          "id": "51891c3d908585af143dfb20ddca5db50c9de01a",
          "message": "Set dedup and pushback labels on all add metrics (#598)",
          "timestamp": "2025-04-15T12:42:34+01:00",
          "tree_id": "fa506b638bea2325baab6d78bc7a02b63e87f8c2",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/51891c3d908585af143dfb20ddca5db50c9de01a"
        },
        "date": 1744717401369,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3435873,
            "unit": "ns/op\t  739940 B/op\t   21736 allocs/op",
            "extra": "409 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3435873,
            "unit": "ns/op",
            "extra": "409 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 739940,
            "unit": "B/op",
            "extra": "409 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21736,
            "unit": "allocs/op",
            "extra": "409 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 232321,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5562 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 232321,
            "unit": "ns/op",
            "extra": "5562 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5562 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5562 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7622,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "172605 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7622,
            "unit": "ns/op",
            "extra": "172605 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "172605 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "172605 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 219.1,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5471200 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 219.1,
            "unit": "ns/op",
            "extra": "5471200 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5471200 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5471200 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 340265,
            "unit": "ns/op\t  278757 B/op\t    3077 allocs/op",
            "extra": "3721 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 340265,
            "unit": "ns/op",
            "extra": "3721 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 278757,
            "unit": "B/op",
            "extra": "3721 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3077,
            "unit": "allocs/op",
            "extra": "3721 times\n4 procs"
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
          "id": "ad18efce880d8e7b904c4308c27f5d1720c29783",
          "message": "Tune AWS conformance (#595)",
          "timestamp": "2025-04-15T12:44:57+01:00",
          "tree_id": "84dcc1fc9ecb841c6d893a4d1fde5fd1a73ba977",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/ad18efce880d8e7b904c4308c27f5d1720c29783"
        },
        "date": 1744717536295,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3128150,
            "unit": "ns/op\t  741860 B/op\t   21758 allocs/op",
            "extra": "388 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3128150,
            "unit": "ns/op",
            "extra": "388 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 741860,
            "unit": "B/op",
            "extra": "388 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21758,
            "unit": "allocs/op",
            "extra": "388 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 227509,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5066 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 227509,
            "unit": "ns/op",
            "extra": "5066 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5066 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5066 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6990,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "222639 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6990,
            "unit": "ns/op",
            "extra": "222639 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "222639 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "222639 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 216.8,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5604588 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 216.8,
            "unit": "ns/op",
            "extra": "5604588 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5604588 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5604588 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 341389,
            "unit": "ns/op\t  279047 B/op\t    3081 allocs/op",
            "extra": "4075 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 341389,
            "unit": "ns/op",
            "extra": "4075 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279047,
            "unit": "B/op",
            "extra": "4075 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3081,
            "unit": "allocs/op",
            "extra": "4075 times\n4 procs"
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
          "id": "ccff4fe6c6fe500f15b9af6a1faa31faf21629aa",
          "message": "Tune for lower latency (#594)",
          "timestamp": "2025-04-15T12:48:42+01:00",
          "tree_id": "36ce02d40e1dbfe416804ac26acaaaa3cde9e717",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/ccff4fe6c6fe500f15b9af6a1faa31faf21629aa"
        },
        "date": 1744717762889,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3237346,
            "unit": "ns/op\t  739047 B/op\t   21727 allocs/op",
            "extra": "378 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3237346,
            "unit": "ns/op",
            "extra": "378 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 739047,
            "unit": "B/op",
            "extra": "378 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21727,
            "unit": "allocs/op",
            "extra": "378 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 229623,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5301 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 229623,
            "unit": "ns/op",
            "extra": "5301 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5301 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5301 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7003,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "184500 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7003,
            "unit": "ns/op",
            "extra": "184500 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "184500 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "184500 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 221.6,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5498576 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 221.6,
            "unit": "ns/op",
            "extra": "5498576 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5498576 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5498576 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 338218,
            "unit": "ns/op\t  279302 B/op\t    3083 allocs/op",
            "extra": "3885 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 338218,
            "unit": "ns/op",
            "extra": "3885 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279302,
            "unit": "B/op",
            "extra": "3885 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3083,
            "unit": "allocs/op",
            "extra": "3885 times\n4 procs"
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
          "id": "d0de2fade1c483fa834446c7b16157d135ebb33c",
          "message": "[AWS Antispam] End-to-end antispam for conformance (#592)\n\n* [AWS Antispam] End-to-end antispam for conformance\n\nAntispam DB can now be created via the Terraform, and the conformance\ncommand takes flags that allow antispam to be enabled and configured\nwith this DB.\n\nDocs have been added to the end of the codelab instructions to document\nhow to deploy this.\n\nSome performance testing was performed and results look good: with 1200\nwrite workers doing 1200 writes per second, with 30% dupe chance, the\nlog was keeping up, growing by around 800 leaves per second.\n\n```\nRead (8 workers): Current max: 20/s. Oversupply in last second: 0\nWrite (1200 workers): Current max: 1200/s. Oversupply in last second: 0\nTreeSize: 199416 (Δ 745qps over 30s)\nTime-in-queue: 6ms/411ms/1056ms (min/avg/max)\nObserved-time-to-integrate: 43ms/1291ms/4346ms (min/avg/max)\n```\n\nThere are a couple of fixes in here for state/transaction management to\nbe more resilient when following the DB. Previously TXs were sometimes\nnot closed, which would lock the table until TX expiry.\n\nFixes #526.\n\n* Enable antispam in CI",
          "timestamp": "2025-04-15T13:04:53+01:00",
          "tree_id": "cbc7bc6f2eb92472c62e3c3e2a9c7af9c9266197",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/d0de2fade1c483fa834446c7b16157d135ebb33c"
        },
        "date": 1744718740761,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3456577,
            "unit": "ns/op\t  741385 B/op\t   21749 allocs/op",
            "extra": "354 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3456577,
            "unit": "ns/op",
            "extra": "354 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 741385,
            "unit": "B/op",
            "extra": "354 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21749,
            "unit": "allocs/op",
            "extra": "354 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 224796,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "4518 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 224796,
            "unit": "ns/op",
            "extra": "4518 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "4518 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "4518 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7249,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "180889 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7249,
            "unit": "ns/op",
            "extra": "180889 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "180889 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "180889 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 218.6,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5251671 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 218.6,
            "unit": "ns/op",
            "extra": "5251671 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5251671 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5251671 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 343600,
            "unit": "ns/op\t  280172 B/op\t    3090 allocs/op",
            "extra": "4260 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 343600,
            "unit": "ns/op",
            "extra": "4260 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 280172,
            "unit": "B/op",
            "extra": "4260 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3090,
            "unit": "allocs/op",
            "extra": "4260 times\n4 procs"
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
          "id": "3141408f25541be0ab2bb404be31a269b72c74ef",
          "message": "[AWS Antispam] Fix CI (#599)\n\nThis is failing at the moment since being enabled in #592. This fixes it\nby not creating a second DB and putting the antispam tables in the main\ntessera DB. Not elegant, but better than not testing antispam in CI.\n\nThis does not change the codelab behaviour; it is still possible to\ncreate the antispam DB from that terraform and run the binaries with the\nflags to use it.",
          "timestamp": "2025-04-15T16:58:07+01:00",
          "tree_id": "9c0c03a0d8a1572d02ef18d15f5e24c3ec4870ed",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/3141408f25541be0ab2bb404be31a269b72c74ef"
        },
        "date": 1744732728885,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3275485,
            "unit": "ns/op\t  740242 B/op\t   21739 allocs/op",
            "extra": "336 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3275485,
            "unit": "ns/op",
            "extra": "336 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 740242,
            "unit": "B/op",
            "extra": "336 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21739,
            "unit": "allocs/op",
            "extra": "336 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 222809,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "4946 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 222809,
            "unit": "ns/op",
            "extra": "4946 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "4946 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "4946 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6876,
            "unit": "ns/op\t    6529 B/op\t       1 allocs/op",
            "extra": "167778 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6876,
            "unit": "ns/op",
            "extra": "167778 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6529,
            "unit": "B/op",
            "extra": "167778 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "167778 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 220.4,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5383803 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 220.4,
            "unit": "ns/op",
            "extra": "5383803 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5383803 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5383803 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 353392,
            "unit": "ns/op\t  279135 B/op\t    3082 allocs/op",
            "extra": "4124 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 353392,
            "unit": "ns/op",
            "extra": "4124 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279135,
            "unit": "B/op",
            "extra": "4124 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3082,
            "unit": "allocs/op",
            "extra": "4124 times\n4 procs"
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
          "id": "161930dbb5d312311e70f1502f5aee4994039d59",
          "message": "Add NextIndex metric (#600)",
          "timestamp": "2025-04-15T17:30:05+01:00",
          "tree_id": "abbadd441567f3a23ef803c3f43efbd42446d10c",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/161930dbb5d312311e70f1502f5aee4994039d59"
        },
        "date": 1744734647343,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3101821,
            "unit": "ns/op\t  742049 B/op\t   21756 allocs/op",
            "extra": "374 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3101821,
            "unit": "ns/op",
            "extra": "374 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 742049,
            "unit": "B/op",
            "extra": "374 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21756,
            "unit": "allocs/op",
            "extra": "374 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 228921,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5676 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 228921,
            "unit": "ns/op",
            "extra": "5676 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5676 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5676 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7436,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "208924 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7436,
            "unit": "ns/op",
            "extra": "208924 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "208924 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "208924 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 217.6,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5506418 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 217.6,
            "unit": "ns/op",
            "extra": "5506418 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5506418 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5506418 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 345096,
            "unit": "ns/op\t  278820 B/op\t    3078 allocs/op",
            "extra": "3804 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 345096,
            "unit": "ns/op",
            "extra": "3804 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 278820,
            "unit": "B/op",
            "extra": "3804 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3078,
            "unit": "allocs/op",
            "extra": "3804 times\n4 procs"
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
          "id": "3cc0c6513dc2a052c9a4e226d2657d5806b0f0c6",
          "message": "[Terraform] Add _db suffix to create_antispam. (#601)",
          "timestamp": "2025-04-15T17:33:44+01:00",
          "tree_id": "35861f1016e193d029ec8cdcf7e70d279a720124",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/3cc0c6513dc2a052c9a4e226d2657d5806b0f0c6"
        },
        "date": 1744734872260,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3190743,
            "unit": "ns/op\t  738503 B/op\t   21723 allocs/op",
            "extra": "327 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3190743,
            "unit": "ns/op",
            "extra": "327 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 738503,
            "unit": "B/op",
            "extra": "327 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21723,
            "unit": "allocs/op",
            "extra": "327 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 230583,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5647 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 230583,
            "unit": "ns/op",
            "extra": "5647 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5647 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5647 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 6925,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "186282 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 6925,
            "unit": "ns/op",
            "extra": "186282 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "186282 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "186282 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 217.6,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5391396 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 217.6,
            "unit": "ns/op",
            "extra": "5391396 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5391396 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5391396 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 347193,
            "unit": "ns/op\t  279208 B/op\t    3082 allocs/op",
            "extra": "4143 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 347193,
            "unit": "ns/op",
            "extra": "4143 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279208,
            "unit": "B/op",
            "extra": "4143 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3082,
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
          "id": "6bc6f18ae66277163fc77a66ac6e536c6461e920",
          "message": "Update freezing docs (#602)",
          "timestamp": "2025-04-15T17:41:26+01:00",
          "tree_id": "17b361b05a629ca0af6d888c3cc19f1cb281c97f",
          "url": "https://github.com/transparency-dev/trillian-tessera/commit/6bc6f18ae66277163fc77a66ac6e536c6461e920"
        },
        "date": 1744735332140,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDedupe",
            "value": 3085931,
            "unit": "ns/op\t  741530 B/op\t   21751 allocs/op",
            "extra": "343 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - ns/op",
            "value": 3085931,
            "unit": "ns/op",
            "extra": "343 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - B/op",
            "value": 741530,
            "unit": "B/op",
            "extra": "343 times\n4 procs"
          },
          {
            "name": "BenchmarkDedupe - allocs/op",
            "value": 21751,
            "unit": "allocs/op",
            "extra": "343 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction",
            "value": 225077,
            "unit": "ns/op\t    3152 B/op\t      36 allocs/op",
            "extra": "5390 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - ns/op",
            "value": 225077,
            "unit": "ns/op",
            "extra": "5390 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - B/op",
            "value": 3152,
            "unit": "B/op",
            "extra": "5390 times\n4 procs"
          },
          {
            "name": "BenchmarkWitnessGroupSatisfaction - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "5390 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText",
            "value": 7150,
            "unit": "ns/op\t    6528 B/op\t       1 allocs/op",
            "extra": "222409 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - ns/op",
            "value": 7150,
            "unit": "ns/op",
            "extra": "222409 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - B/op",
            "value": 6528,
            "unit": "B/op",
            "extra": "222409 times\n4 procs"
          },
          {
            "name": "BenchmarkLeafBundle_UnmarshalText - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "222409 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe",
            "value": 215.7,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "5526171 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - ns/op",
            "value": 215.7,
            "unit": "ns/op",
            "extra": "5526171 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5526171 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckpointUnsafe - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "5526171 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate",
            "value": 344883,
            "unit": "ns/op\t  279541 B/op\t    3086 allocs/op",
            "extra": "4339 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - ns/op",
            "value": 344883,
            "unit": "ns/op",
            "extra": "4339 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - B/op",
            "value": 279541,
            "unit": "B/op",
            "extra": "4339 times\n4 procs"
          },
          {
            "name": "BenchmarkIntegrate - allocs/op",
            "value": 3086,
            "unit": "allocs/op",
            "extra": "4339 times\n4 procs"
          }
        ]
      }
    ]
  }
}