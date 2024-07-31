# Hammer: A load testing tool for Tessera logs

This hammer sets up read and (optionally) write traffic to a log to test correctness and performance under load.
The read traffic is sent according to the [tlog-tiles](https://github.com/C2SP/C2SP/blob/main/tlog-tiles.md) spec, and thus could be used to load test any tiles-based log, not just Tessera logs.

If write traffic is enabled, then the target log must support `POST` requests to a `/add` path.
A successful request MUST return an ASCII decimal number representing the index that has been assigned to the new value.

## UI

The hammer runs using a text-based UI in the terminal that shows the current status, logs, and supports increasing/decreasing read and write traffic.
The process can be killed with `<Ctrl-C>`.
This TUI allows for a level of interactivity when probing a new configuration of a log in order to find any cliffs where performance degrades.

For real load-testing applications, especially headless runs as part of a CI pipeline, it is recommended to run the tool with `show_ui=false` in order to disable the UI.

## Usage

Example usage to test a deployment of `example-mysql`:

```shell
go run ./hammer \
  --log_public_key=Test-Betty+df84580a+AQQASqPUZoIHcJAF5mBOryctwFdTV1E0GRY4kEAtTzwB \
  --log_url=http://localhost:2024 \
  --num_writers=22 \
  --max_write_ops=42
```

