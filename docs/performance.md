# Tessera Storage Performance

This document describes the performance of each Trillian Tessera storage implementation under different conditions.

> [!TIP]
> The performance test result shows that Trillian Tessera can run on the free tier VM instance on GCP.

## MySQL

### GCP Free Tier VM Instance

**tl;dr:** Tessera (MySQL) can run on the free tier VM instance on GCP with around **300 write QPS**. The bottleneck comes from the lack of memory which is consumed by the dockerized MySQL instance.

**e2-micro**

- vCPU: 0.25-2 vCPU (1 shared core)
- Memory: 1 GB
- OS: Debian GNU/Linux 12 (bookworm)

> [!NOTE]
> Virtual CPUs (vCPUs) in virtualized environments often share physical CPU cores with other vCPUs and introduce variability and potential performance impacts.

#### Result

```
┌──────────────────────────────────────────────────────────────────────┐
│Read (8 workers): Current max: 0/s. Oversupply in last second: 0      │
│Write (400 workers): Current max: 527/s. Oversupply in last second: 0 │
│TreeSize: 372394 (Δ 457qps over 30s)                                  │
│Time-in-queue: 57ms/400ms/1225ms (min/avg/max)                        │
│Observed-time-to-integrate: 895ms/1793ms/7094ms (min/avg/max)         │
└──────────────────────────────────────────────────────────────────────┘
```

The bottleneck is at the dockerized MySQL instance, which consumes around 50% of the memory. `kswapd0` started consuming 100% swapping the memory when pushing through the limit.

```
top - 18:15:31 up 18 min,  3 users,  load average: 0.12, 0.26, 0.13
Tasks: 103 total,   1 running, 101 sleeping,   0 stopped,   1 zombie
%Cpu(s):  2.7 us,  2.0 sy,  0.0 ni, 91.3 id,  0.8 wa,  0.0 hi,  3.0 si,  0.2 st 
MiB Mem :    970.0 total,     71.8 free,    924.5 used,     87.1 buff/cache     
MiB Swap:      0.0 total,      0.0 free,      0.0 used.     45.5 avail Mem 

    PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
   3021 root      20   0 1231580  24544   3476 S  10.3   2.5   0:17.89 conformance-mys
   2675 999       20   0 1842248 475440   3396 S   4.0  47.9   0:09.91 mysqld
```

#### Steps

1. Create a [GCP free tier](https://cloud.google.com/free/docs/free-cloud-features#free-tier) e2-micro VM instance in us-central1 (iowa).

1. [Install Go](https://go.dev/doc/install)
   
   ```sh
   instance:~$ wget https://go.dev/dl/go1.23.2.linux-amd64.tar.gz
   instance:~$ sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz
   instance:~$ export PATH=$PATH:/usr/local/go/bin
   instance:~$ go version
   go version go1.23.2 linux/amd64
   ```

1. [Install Docker using the `apt` repository](https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository)

1. Install Git

   ```sh
   instance:~$ sudo apt-get install git -y -q
   ...
   instance:~$ git version
   git version 2.39.5
   ```

1. Clone the Trillian Tessera repository

   ```sh
   instance:~$ git clone https://github.com/transparency-dev/trillian-tessera.git
   Cloning into 'trillian-tessera'...
   ```

1. Run `cmd/conformance/mysql` and MySQL database via Docker compose

   ```sh
   instance:~/trillian-tessera$ sudo docker compose -f ./cmd/conformance/mysql/docker/compose.yaml up
   ```

1. Run `hammer` and get performance metrics

   ```sh
   hammer:~/trillian-tessera$ go run ./internal/hammer --log_public_key=transparency.dev/tessera/example+ae330e15+ASf4/L1zE859VqlfQgGzKy34l91Gl8W6wfwp+vKP62DW --log_url=http://10.128.0.3:2024 --max_read_ops=0 --num_writers=512 --max_write_ops=512
   ```

### GCP Free Tier VM Instance + Cloud SQL (MySQL)

**tl;dr:** Tessera (MySQL) can run on the free tier VM instance on GCP with around **2000 write QPS** when the MySQL database is run on Cloud SQL.

**e2-micro**

- vCPU: 0.25-2 vCPU (1 shared core)
- Memory: 1 GB
- OS: Debian GNU/Linux 12 (bookworm)

> [!NOTE]
> Virtual CPUs (vCPUs) in virtualized environments often share physical CPU cores with other vCPUs and introduce variability and potential performance impacts.

**Cloud SQL (MySQL 8.0.31)**

- vCPUs: 4
- Memory: 7.5 GB
- SSD storage: 10 GB

#### Result

```
┌───────────────────────────────────────────────────────────────────────┐
│Read (66 workers): Current max: 1/s. Oversupply in last second: 0      │
│Write (541 workers): Current max: 4139/s. Oversupply in last second: 0 │
│TreeSize: 1087381 (Δ 3121qps over 30s)                                 │
│Time-in-queue: 71ms/339ms/1320ms (min/avg/max)                         │
│Observed-time-to-integrate: 887ms/2834ms/8510ms (min/avg/max)          │
└───────────────────────────────────────────────────────────────────────┘
```

The bottleneck comes from CPU usage of the `cmd/conformance/mysql` binary on the free tier VM instance. The Cloud SQL (MySQL) CPU usage is lower than 10%.

```
top - 00:57:43 up  7:00,  2 users,  load average: 0.13, 0.15, 0.07
Tasks:  91 total,   1 running,  90 sleeping,   0 stopped,   0 zombie
%Cpu(s): 10.6 us,  3.4 sy,  0.0 ni, 81.8 id,  0.0 wa,  0.0 hi,  4.1 si,  0.2 st 
MiB Mem :    970.0 total,    314.0 free,    529.5 used,    275.9 buff/cache     
MiB Swap:      0.0 total,      0.0 free,      0.0 used.    440.5 avail Mem 

    PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
  10240 rogerng+  20   0 1299188  77288   5464 S  32.9   7.8   1:04.62 mysql
    371 root      20   0  221788   3192      0 S   0.7   0.3   0:18.78 rsyslo
```

#### Steps

1. Create a MySQL instance on Cloud SQL.

1. Create a [GCP free tier](https://cloud.google.com/free/docs/free-cloud-features#free-tier) e2-micro VM instance in us-central1 (iowa).

1. Setup VPC peering.

1. [Install Go](https://go.dev/doc/install)
   
   ```sh
   instance:~$ wget https://go.dev/dl/go1.23.2.linux-amd64.tar.gz
   instance:~$ sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz
   instance:~$ export PATH=$PATH:/usr/local/go/bin
   instance:~$ go version
   go version go1.23.2 linux/amd64
   ```

1. Install Git

   ```sh
   instance:~$ sudo apt-get install git -y -q
   ...
   instance:~$ git version
   git version 2.39.5
   ```

1. Clone the Trillian Tessera repository

   ```sh
   instance:~$ git clone https://github.com/transparency-dev/trillian-tessera.git
   Cloning into 'trillian-tessera'...
   ```

1. Run [`cloud-sql-proxy`](https://cloud.google.com/sql/docs/mysql/sql-proxy#install)

   ```sh
   instance:~$ ./cloud-sql-proxy --port 3306 transparency-dev-playground:us-central1:mysql-dev-instance-1
   ```

1. Run `cmd/conformance/mysql`

   ```sh
   instance:~/trillian-tessera$ go run ./cmd/conformance/mysql --mysql_uri="root:root@tcp(127.0.0.1:3306)/test_tessera" --init_schema_path="./storage/mysql/schema.sql" --private_key_path="./cmd/conformance/mysql/docker/testdata/key" --public_key_path="./cmd/conformance/mysql/docker/testdata/key.pub" --db_max_open_conns=1024 --db_max_idle_conns=512
   ```

1. Run `hammer` and get performance metrics

   ```sh
   hammer:~/trillian-tessera$ go run ./internal/hammer --log_public_key=transparency.dev/tessera/example+ae330e15+ASf4/L1zE859VqlfQgGzKy34l91Gl8W6wfwp+vKP62DW --log_url=http://10.128.0.3:2024 --max_read_ops=0 --num_writers=512 --max_write_ops=512
   ```

## POSIX

### GCP Free Tier VM Instance

**e2-micro**

- vCPU: 0.25-2 vCPU (1 shared core)
- Memory: 1 GB
- OS: Debian GNU/Linux 12 (bookworm)

> [!NOTE]
> Virtual CPUs (vCPUs) in virtualized environments often share physical CPU cores with other vCPUs and introduce variability and potential performance impacts.

#### Result

```
┌───────────────────────────────────────────────────────────────────────┐
│Read (184 workers): Current max: 0/s. Oversupply in last second: 0     │
│Write (600 workers): Current max: 1758/s. Oversupply in last second: 0 │
│TreeSize: 1882477 (Δ 1587qps over 30s)                                 │
│Time-in-queue: 149ms/371ms/692ms (min/avg/max)                         │
│Observed-time-to-integrate: 569ms/1191ms/1878ms (min/avg/max)          │
└───────────────────────────────────────────────────────────────────────┘
```

```
top - 20:45:35 up 47 min,  3 users,  load average: 1.89, 0.88, 1.03
Tasks:  97 total,   1 running,  96 sleeping,   0 stopped,   0 zombie
%Cpu(s): 47.2 us, 24.7 sy,  0.0 ni,  0.0 id,  0.0 wa,  0.0 hi, 28.1 si,  0.0 st 
MiB Mem :    970.0 total,    158.7 free,    566.8 used,    409.3 buff/cache     
MiB Swap:      0.0 total,      0.0 free,      0.0 used.    403.2 avail Mem 

    PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
   8336 root      20   0 1231800  34784   5080 S 200.0   3.5   2:59.50 conformance-pos
    409 root      20   0 2442648  79112  26836 S   1.0   8.0   0:49.10 dockerd
   6848 root      20   0 1800176  34376  16940 S   0.7   3.5   0:12.57 docker-compose
```

#### Steps

1. Create a [GCP free tier](https://cloud.google.com/free/docs/free-cloud-features#free-tier) e2-micro VM instance in us-central1 (iowa).

1. [Install Go](https://go.dev/doc/install)
   
   ```sh
   instance:~$ wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
   instance:~$ sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
   instance:~$ export PATH=$PATH:/usr/local/go/bin
   instance:~$ go version
   go version go1.23.0 linux/amd64
   ```

1. [Install Docker using the `apt` repository](https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository)

1. Install Git

   ```sh
   instance:~$ sudo apt-get install git -y -q
   ...
   instance:~$ git version
   git version 2.39.2
   ```

1. Clone the Trillian Tessera repository

   ```sh
   instance:~$ git clone https://github.com/transparency-dev/trillian-tessera.git
   Cloning into 'trillian-tessera'...
   ```

1. Run `cmd/conformance/posix` via Docker compose

   ```sh
   instance:~/trillian-tessera$ sudo docker compose -f ./cmd/conformance/posix/docker/compose.yaml up
   ```

1. Run `hammer` and get performance metrics

   ```sh
   hammer:~/trillian-tessera$ go run ./internal/hammer --log_public_key=example.com/log/testdata+33d7b496+AeHTu4Q3hEIMHNqc6fASMsq3rKNx280NI+oO5xCFkkSx --log_url=http://localhost:2025 --max_read_ops=0 --num_writers=512 --max_write_ops=512
   ```

## GCP

Coming soon.
