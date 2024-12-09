window.BENCHMARK_DATA = {
  "lastUpdate": 1733767978938,
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
      }
    ]
  }
}