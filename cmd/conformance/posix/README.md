# conformance-posix

This command runs an HTTP web server that accepts POST HTTP requests to a `/add` endpoint.
This endpoint takes arbitrary data and adds it to a file-based log.

## Using in the hammer

First bring up a new log in one terminal:
```shell
export LOG_PRIVATE_KEY="PRIVATE+KEY+example.com/log/testdata+33d7b496+AeymY/SZAX0jZcJ8enZ5FY1Dz+wTML2yWSkK+9DSF3eg"
export LOG_PUBLIC_KEY="example.com/log/testdata+33d7b496+AeHTu4Q3hEIMHNqc6fASMsq3rKNx280NI+oO5xCFkkSx"
export LOG_DIR=/tmp/mylog2

# Initialize a new log
go run ./cmd/conformance/posix \
  --storage_dir=${LOG_DIR} \
  --initialise \
  --listen=:2025 \
  --v=2
```

In another terminal, run the hammer against the log.
In this example, we're running 32 writers against the log to add 128 new leaves within 1 minute.

Note that the writes are sent to the HTTP server we brought up in the previous step, but reads are sent directly to the file system.

```shell
go run ./hammer \
  --log_public_key=example.com/log/testdata+33d7b496+AeHTu4Q3hEIMHNqc6fASMsq3rKNx280NI+oO5xCFkkSx \
  --write_log_url=http://localhost:2025 \
  --log_url=file:///tmp/mylog2/ \
  --max_read_ops=0 \
  --num_writers=32 \
  --max_write_ops=64 \
  --max_runtime=1m \
  --leaf_write_goal=128 \
  --show_ui=false
```

Optionally, inspect the log using the woodpecker tool to see the contents:

```shell
go run github.com/mhutchinson/woodpecker@main --custom_log_type=tiles --custom_log_url=file:///${LOG_DIR}/ --custom_log_origin=example.com/log/testdata --custom_log_vkey=${LOG_PUBLIC_KEY}
```

