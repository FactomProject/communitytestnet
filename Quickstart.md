# Quickstart

## Prerequisites

To create the environment install the following components first:
 - [docker](https://www.docker.com/community-edition)
 - [docker-compose](https://docs.docker.com/compose/install/)

## Running factomd

The quickstart commands

```
# Setup volume
cp factomd.conf.EXAMPLE factomd.conf
docker volume create factomd_volume
docker run --rm -v ${PWD}/factomd.conf:/source -v factomd_volume:/destination busybox /bin/cp /source /destination/factomd.conf

# Start docker containers
docker pull emyrk/factomd_testnet_community:v1
docker-compose up -d

# Start factomd
docker exec factomd_node bash /root/start.sh
```

You are now running on the community testnet. To check if your docker containers are indeed running:
```
docker ps
```

You now visit:
* localhost:8090 for the control panel
* localhost:3001 for Grafana
* localhost:9090 for Prometheus

## Cleanup/stopping

To stop factomd

```
docker exec factomd_node bash /root/start.sh
```

To stop all the containers

```
docker-compose down
```