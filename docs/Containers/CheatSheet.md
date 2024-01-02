# My Own Docker Check Sheet

## Containers

### Running a container

```base
docker run --name <my-container-name> -d <image:tag> -e KEY=VALUE
```

Where:

* --name: lets you specify a name for your container
* -d: instructs Docker to run the container as a service in the background.
* <image:tag>: is where you specify the image you want to run
* -e: lets you create environment variables

### Getting into a container

```bash
docker exec -it <container-name> command (e.g. /bin/bash)
```

Where:

* -it: stands for `interacting`, which means that you will be in the containers terminal.

### Listing all containers:

```bash
docker ps -a
```

### Listing Running Containers

```bash
docker ps
```

### Removing all containers

```bash
docker container rm $(docker container ls -aq)
```

## Images

### Listing Docker Images

```bash
docker image ls
```

OR

```base
docker images
```

### Building a Docker Image without Cache

```bash
docker build --no-cache -t <image-name-tab> .
```

### Deleting Docker Images

```base
docker image rm <id>
```

## System

### Completely clear all cache

```bash
docker system prune
```
