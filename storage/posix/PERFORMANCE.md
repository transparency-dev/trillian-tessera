# POSIX Performance

## Local system VM + local disk

An LXC container with 8 cores and 8GB RAM, using a local ZFS filesystem built from two 6TB SAS disks configured as a mirrored device.

This configuration is likely to continue scaling, but 40000 writes/s seemed _sufficient_ for most usecases.

```
┌───────────────────────────────────────────────────────────────────────────────┐
│Read (8 workers): Current max: 20/s. Oversupply in last second: 0              │
│Write (2000 workers): Current max: 40000/s. Oversupply in last second: 0       │
│TreeSize: 2360352 (Δ 40021qps over 30s)                                        │
│Time-in-queue: 212ms/228ms/273ms (min/avg/max)                                 │
│Observed-time-to-integrate: 5409ms/5492ms/5537ms (min/avg/max)                 │
├───────────────────────────────────────────────────────────────────────────────┤
```

```
top - 12:46:53 up 8 days, 17:51,  1 user,  load average: 23.96, 15.37, 7.22
Tasks:  35 total,   1 running,  34 sleeping,   0 stopped,   0 zombie
%Cpu(s): 13.7 us,  6.1 sy,  0.0 ni, 78.5 id,  0.5 wa,  0.0 hi,  1.2 si,  0.0 st
MiB Mem :   8192.0 total,   5552.3 free,   1844.5 used,    795.1 buff/cache
MiB Swap:    512.0 total,    512.0 free,      0.0 used.   6346.3 avail Mem

    PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
 710077 al        20   0 1371244 132940   5632 S 325.3   1.6  21:22.60 hammer
 709742 al        20   0 1304500 105820   4096 S 278.3   1.3  19:03.57 posix
```

## GCP Free Tier VM Instance

**e2-micro**

- vCPU: 0.25-2 vCPU (1 shared core)
- Memory: 1 GB
- OS: Debian GNU/Linux 12 (bookworm)

> [!NOTE]
> Virtual CPUs (vCPUs) in virtualized environments often share physical CPU cores with other vCPUs and introduce variability and potential performance impacts.


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
