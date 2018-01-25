# Community Test Net

Spinning up a factom node on the community testnet as well as monitoring tools. The docker image for factomd will enable Factom developers to have ssh access to docker container.

## Basic usage

This setup uses *docker* and *docker-compose* to run the factomd network and
all supporting services in docker containers. For more details about the usage,
consult with [docker](https://docs.docker.com/) and
[*docker-compose*](https://docs.docker.com/compose/) documentation.

### Prerequisites

To create the environment install the following components first:
 - [docker](https://www.docker.com/community-edition)
 - [docker-compose](https://docs.docker.com/compose/install/)

### Creating the environment

First create the volume for factomd
```
docker volume create factomd_volume
```

Place the config you have inside
```
docker run --rm -v ${PWD}/factomd.conf:/source -v factomd_volume:/destination busybox /bin/cp /source /destination/factomd.conf
```

Pull the factomd image

```
docker pull emyrk/factomd_testnet_community:v1
```

Launch factomd & monitoring

```
docker-compose up -d
```

This command:
 - builds all container images if necessary
 - starts all the containers
 - starts printing all *stdout* logs from the containers to the foreground.

If you prefer to run this in the foreground, do not use the `-d` option:

```
docker-compose up
```

After *docker-compose* finishes the startup, be sure to check the output of
`docker ps` command to make sure all the services are running.

Now you can start factomd
```
docker exec factomd_node bash /root/start.sh
```

### Starting/stopping the environment

```
docker-compose stop
```

This command stops all the containers without removing them. Note that all
non-persistent storage will be wiped out as the containers are restarted.

The containers can be started again with:

```
docker-compose start
```

### Tearing down the environment

```
docker-compose down
```

This command:
 - removes all service containers
 - removes all created networks

### Updating the enviroment

Get latest updates, then update docker containers

```
git pull
docker-compose build
```

## Using services

Once the environment is properly built and started you can start using it for
monitoring the running factomd network.

### Factomd instances

The Control Panel web UI for all 3 instances are mapped to the ports on the
host machine, so they are available at the following addresses:
 * *factomd_1* - http://localhost:8090

### Metrics

Metrics are pulled from *factomd* instances into *Prometheus*. To view the
collected metrics open the *Prometheus* web UI at: http://localhost:9090/,
select one of the metrics from the dropdown and hit the *Execute* button. The
metrics for all instances are labeled using the hostname/port that was used for
scraping the metrics, e.g. `instance=factomd_node:9876`.

# Useful commands

### Listing containers

```
docker ps
```
```
docker-compose ps
```

### Viewing the stdout logs

Single container:

```
docker logs <container_name>
```
```
docker logs <container_name> -f
```

All containers:

```
docker-compose logs
```
```
docker-compose logs -f
```

### Logging into a container

```
docker exec -it <container_name> bash
```

### Stopping a single container

To gracefully shutdown a container (the application receives a SIGTERM signal):

```
docker stop <container_name>
```

To kill the container:

```
docker kill <container_name>
```

### Getting IP addresses

Display all containers with all networks they belong to and their static /
assigned IP addresses in a network:

```
docker inspect -f '{{.Name}} - {{range $name, $net := .NetworkSettings.Networks}}{{$name}}:{{$net.IPAddress}} {{end}}' $(docker ps -aq)
```

### Prometheus

*Prometheus* periodically gathers metrics from a all *factomd* instances. The
built-in web UI is exposed to the local host using port `9090`.

Note that *Prometheus* is pull-based, so it fetches metrics from *factomd*
instances, not the other way around, so to add monitoring for your local nodes,
you'll need to modify its configuration and rebuild the container.

The configuration for *Prometheus* is copied during the build from
`prometheus/config/prometheus.yml`. Currently it pull metrics from all 3
instances and labels them using the instance name.

## Known issues

* If you are running this setup in a *Docker for Mac* or a *Docker for Windows*
  environment you might want to adjust the CPU and Memory settings
  (*Preferences...* -> *Advanced*), since the services started in this setup
  may require more than the defaults. It may also be useful to bump the default
  *docker-compose* HTTP timeout setting if you see some problems during
  startup, e.g.:

```
COMPOSE_HTTP_TIMEOUT=120 docker-compose up
```

* The `factomd` instances sometimes exit after the first build, another
  `docker-compose up` command should bring start them correctly.

* There is an issue when using the environment on Mac OS:

```
ERROR: for kibana  Cannot start service kibana: driver failed programming
external connectivity on endpoint kibana
(7e6b3eaddf72eb60f384edff6d5c0bbac759af0cf5c24cfaff646ec558815cd5): Timed out
proxy starting the userland proxy
```

Unfortunately this is an unresolved *Docker for Mac* issue for which there is
no good solution (restarting does not help), you'll need to retry until it
succeeds.