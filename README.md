# SaniLoader

SaniLoader is a custom load balancer written in [go](https://go.dev/).

## Features

- Dynamically load balance between Docker Containers.

## Usage/Examples

```bash
saniloader run      #scan all running Docker Containers
saniloader run --containers container1 container2 ...       #run loadbalancer for container1, container2, ...
saniloader run -d       #run loadbalancer in dynamic mode

saniloader list     #list Docker Container names which saniloader is running on

saniloader metrics       #give all metrics for all containers
saniloader metrics --containers container1 container2 ...      #give metrics for container1, contianer2, ...

saniloader stop     #stop loadbalancer
```

## Contact

- [@saniyar.dev (github)](https://github.com/saniyar-dev)
- [@saniyar.dev (hamgit)](https://hamgit.ir/saniyar.dev)
- [@saniyar.dev (telegram)](https://t.me/saniyar_dev)
