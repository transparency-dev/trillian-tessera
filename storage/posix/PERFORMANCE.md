# POSIX Performance

## Local system VM + local disk

An LXC container with 12 cores and 8GB RAM, using a local ZFS filesystem built from two 6TB SAS disks configured as a mirrored device.

The limiting factor is storage latency; writes are `fsync`'d for durability. NVME storage would likely improve throughput.

Without antispam, it was able to sustain around 2900 writes/s.

```
┌────────────────────────────────────────────────────────────────────────────────────┐
│Read (8 workers): Current max: 20/s. Oversupply in last second: 0                   │
│Write (3000 workers): Current max: 3000/s. Oversupply in last second: 0             │
│TreeSize: 1470460 (Δ 2927qps over 30s)                                              │
│Time-in-queue: 136ms/1110ms/1356ms (min/avg/max)                                    │
│Observed-time-to-integrate: 583ms/6019ms/6594ms (min/avg/max)                       │
├────────────────────────────────────────────────────────────────────────────────────┤
```

With antispam enabled (badger), it was able to sustain around 1600 writes/s.

```
┌────────────────────────────────────────────────────────────────────────────────────┐
│Read (8 workers): Current max: 20/s. Oversupply in last second: 0                   │
│Write (1800 workers): Current max: 1800/s. Oversupply in last second: 0             │
│TreeSize: 2041087 (Δ 1664qps over 30s)                                              │
│Time-in-queue: 0ms/112ms/448ms (min/avg/max)                                        │
│Observed-time-to-integrate: 593ms/3232ms/5754ms (min/avg/max)                       │
├────────────────────────────────────────────────────────────────────────────────────┤
```

```
top - 16:03:40 up 13 days,  7:14,  3 users,  load average: 1.47, 1.94, 1.97
Tasks:  72 total,   1 running,  71 sleeping,   0 stopped,   0 zombie
%Cpu(s):  1.7 us,  1.2 sy,  0.0 ni, 97.0 id,  0.0 wa,  0.0 hi,  0.1 si,  0.0 st
MiB Mem :   8192.0 total,   5552.3 free,   1844.5 used,    795.1 buff/cache
MiB Swap:    512.0 total,    512.0 free,      0.0 used.  28332.6 avail Mem

    PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
2141681 al        20   0 4258436 114576   5120 S  50.5   0.3   6:26.34 hammer
2140012 al        20   0 7480052 735092 143196 S  29.9   2.2   4:55.74 posix
```

## GCP Free Tier VM Instance

**e2-micro**

- vCPU: 0.25-2 vCPU (1 shared core)
- Memory: 1 GB
- OS: Debian GNU/Linux 12 (bookworm)

> [!NOTE]
> Virtual CPUs (vCPUs) in virtualized environments often share physical CPU cores with other vCPUs and introduce variability and potential performance impacts.

This is able to sustain around 600 write/s with antispam enabled.


```
┌─────────────────────────────────────────────────────────────────────────┐
│Read (8 workers): Current max: 20/s. Oversupply in last second: 0        │
│Write (1000 workers): Current max: 1000/s. Oversupply in last second: 538│
│TreeSize: 96827 (Δ 683qps over 30s)                                      │
│Time-in-queue: 167ms/1264ms/2000ms (min/avg/max)                         │
│Observed-time-to-integrate: 416ms/7877ms/10990ms (min/avg/max)           │
├─────────────────────────────────────────────────────────────────────────┤
```

```
top - 15:37:30 up 22 min,  4 users,  load average: 6.56, 4.83, 2.68
Tasks: 101 total,   1 running, 100 sleeping,   0 stopped,   0 zombie
Cpu(s): 56.8 us, 21.6 sy,  0.0 ni,  8.0 id,  9.1 wa,  0.0 hi,  4.5 si,  0.0 st
MiB Mem :    970.0 total,     62.2 free,    839.5 used,    227.0 buff/cache
0.0 total,      0.0 free,      0.0 used.    130.5 avail Mem
    PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
   4716 al        20   0 3671328 196420  15700 S  96.0  19.8   5:01.70 posix
   5244 al        20   0 1300984  85548   6160 S  74.4   8.6   3:44.09 hammer
```

### Steps

1. Create a [GCP free tier](https://cloud.google.com/free/docs/free-cloud-features#free-tier) e2-micro VM instance in us-central1 (Iowa).

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

1. Clone the Tessera repository

   ```sh
   instance:~$ git clone https://github.com/transparency-dev/tessera.git
   Cloning into 'tessera'...
   ```

1. Run `cmd/conformance/posix` via Docker compose

   ```sh
   instance:~/tessera$ sudo docker compose -f ./cmd/conformance/posix/docker/compose.yaml up
   ```

1. Run `hammer` and get performance metrics

   ```sh
   hammer:~/tessera$ go run ./internal/hammer --log_public_key=example.com/log/testdata+33d7b496+AeHTu4Q3hEIMHNqc6fASMsq3rKNx280NI+oO5xCFkkSx --log_url=http://localhost:2025 --max_read_ops=0 --num_writers=512 --max_write_ops=512
   ```
