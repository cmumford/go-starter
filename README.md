# Go Web Service with Docker Starter Project

## Using Docker Compose

To build and run a local Docker container. The following will watch for,
and immediately rebuild/restart the container on any changed files.

```sh
make dev
```

The following will run a production configuration:

```sh
make prod
```

## Using Docker

### Building a Docker Image

```sh
docker build -t go-starter:latest .
```

### Running a Docker Image

```sh
docker run -p 80:8080 go-starter:latest
```

### With Docker-Compose

```sh
docker-compose up --build --watch
```

## Misc. Docker Stuff

Delete all containers
```sh
docker rm -f $(docker ps -a -q)
```

Delete all images
```sh
docker rmi -f $(docker images -aq)
```

Tagging local image

```sh
docker tag go-starter:latest cmumford/go-starter:latest
```

Pushing to a registry (dockerhub)

```sh
docker push cmumford/go-starter:latest
```
