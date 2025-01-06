# SaniLoader

SaniLoader is a custom load balancer written in [go](https://go.dev/). It's just for practicing.

## Features

- Dynamically load balance between Docker Containers.

## Usage/Examples

```bash
saniloader start      #scan all running Docker Containers and run loadbalancer on them
saniloader start -d       #run loadbalancer in dynamic mode
saniloader start -c path/to/config.json   #run loadbalancer with config file and to all running docker containers
saniloader start -c path/to/config.json --only    #run loadbalancer only with config file
```

To get metrics for each container on host machine use this:

```bash
curl localhost:9191/[name-of-container]
```

You can use this code to run a metrics server on each container:

```bash
#!/bin/bash

while true; do { echo -e 'HTTP/1.1 200 OK\r\n'; curl [your-docker-host-ip]:9191/[name-of-container] } | nc -l 9090; done &
```

- You should have netcat installed on your docker container for script above.
- By default [your-docker-host-ip] is 172.0.0.1 but in case you've customized it you should change it to your config.

## Contact

- [@saniyar.dev (github)](https://github.com/saniyar-dev)
- [@saniyar.dev (hamgit)](https://hamgit.ir/saniyar.dev)
- [@saniyar.dev (telegram)](https://t.me/saniyar_dev)
