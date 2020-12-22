# go-docker-hello-world
Modern Golang "Hello World!" web server in a docker image hosted on GitHub

Docker base image: alpine

Router: gorilla/mux

GitHub Container Registry: ghcr.io

Github action: matootie/github-docker@v3.1.0

### Usage

```shell
docker run --rm -p 80:8080 ghcr.io/vitr/go-docker-hello-world
```
