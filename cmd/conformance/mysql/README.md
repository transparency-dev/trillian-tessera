# How to Run a Tessera Log (MySQL)

> [!WARNING]
> - This is an example and is not fit for production use, but demonstrates a way of using the Tessera Log with MySQL storage backend.
> - This example is built on the [tlog tiles API](https://c2sp.org/tlog-tiles) for read endpoints and exposes a /add endpoint that allows any POSTed data to be added to the log.

The tessera log with the MySQL storage implementation can be started with either Docker Compose or manual `go run`.

Note that all the commands are executed at the root directory of this repository.

## Docker Compose

### Prerequisites

- [Docker Compose](https://docs.docker.com/compose/install/)

### Starting

```sh
docker compose -f ./cmd/conformance/mysql/docker/compose.yaml up
```

Add `-d` if you want to run the log in detached mode.

### Stopping

```sh
docker compose -f ./cmd/conformance/mysql/docker/compose.yaml down
```

## Manual 

### Prerequisites

Assume you have the MySQL database ready. An alternative way is to run a MySQL database via [Docker](https://docs.docker.com/engine/install/).

```sh
docker run --name test-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test_tessera -d mysql:8.4
```

### Starting

```sh
go run ./cmd/conformance/mysql --mysql_uri="root:root@tcp(localhost:3306)/test_tessera" --init_schema_path="./storage/mysql/schema.sql" --private_key_path="./cmd/conformance/mysql/docker/testdata/key" --public_key_path="./cmd/conformance/mysql/docker/testdata/key.pub"
```

### Stopping

<kbd>Ctrl</kbd> <kbd>C</kbd>

## Using the log

In this example, we're running 256 writers against the log to add 1024 new leaves within 1 minute.

Note that the writes are sent to the HTTP server we brought up in the previous step, but reads are sent directly to the file system.

```shell
go run ./hammer \
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
