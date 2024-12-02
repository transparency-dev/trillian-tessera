# Conformance MySQL log

This binary runs an HTTP web server that accepts POST HTTP requests to an `/add` endpoint.
This endpoint takes arbitrary data and adds it to a MySQL based Tessera log.

> [!WARNING]
> - This is an example and is not fit for production use, but demonstrates a way of using the Tessera Log with MySQL storage backend.
> - This example is built on the [tlog tiles API](https://c2sp.org/tlog-tiles) for read endpoints and exposes a /add endpoint that allows any POSTed data to be added to the log.

## Bring up a log

This will help you bring up a MySQL database to store a Tessera log, and start a personality
binary that can add entries to it.

You can run this personality using Docker Compose or manually with `go run`.

Note that all the commands are executed at the root directory of this repository.

### Docker Compose

#### Prerequisites

Install [Docker Compose](https://docs.docker.com/compose/install/).

#### Start the log

```sh
docker compose -f ./cmd/conformance/mysql/docker/compose.yaml up
```

Add `-d` if you want to run the log in detached mode.

#### Stop the log

```sh
docker compose -f ./cmd/conformance/mysql/docker/compose.yaml down
```

### Manual 

#### Prerequisites

You need to have a MySQL database running on port `3306`.

You can start one using [Docker](https://docs.docker.com/engine/install/).

```sh
docker run --name test-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test_tessera -d mysql:8.4
```

#### Start the log

```sh
go run ./cmd/conformance/mysql --mysql_uri="root:root@tcp(localhost:3306)/test_tessera" --init_schema_path="./storage/mysql/schema.sql" --private_key_path="./cmd/conformance/mysql/docker/testdata/key"
```

#### Stop the log

<kbd>Ctrl</kbd> <kbd>C</kbd>

## Add entries to the log

### Manually

Head over to the [codelab](../#codelab) to manually add entries to the log, and inspect the log.

### Using the hammer

In this example, we're running 256 writers against the log to add 1024 new leaves within 1 minute.

Note that the writes are sent to the HTTP server we brought up in the previous step, but reads are sent directly to the file system.

```shell
go run ./internal/hammer \
  --log_public_key=transparency.dev/tessera/example+ae330e15+ASf4/L1zE859VqlfQgGzKy34l91Gl8W6wfwp+vKP62DW \
  --log_url=http://localhost:2024/ \
  --max_read_ops=0 \
  --num_writers=256 \
  --max_write_ops=256 \
  --max_runtime=1m \
  --leaf_write_goal=1024 \
  --show_ui=false
```

Optionally, inspect the log using the woodpecker tool to see the contents:

```shell
go run github.com/mhutchinson/woodpecker@main --custom_log_type=tiles --custom_log_url=http://localhost:2024/ --custom_log_vkey=transparency.dev/tessera/example+ae330e15+ASf4/L1zE859VqlfQgGzKy34l91Gl8W6wfwp+vKP62DW
```
