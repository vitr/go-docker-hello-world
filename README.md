# go-docker-hello-world
Golang "Hello, World!" modern web server in a docker image hosted on GitHub

Docker base image: alpine

Router: gorilla/mux

GitHub Container Registry: ghcr.io

Github action: matootie/github-docker@v3.1.0

### Usage

```shell
docker run --rm -t -i -p 8080:80 ghcr.io/vitr/go-docker-hello-world
```

` --rm` Automatically remove the container
when it exits

` -i, --interactive` Keep STDIN open even if not attached

`-t, --tty` Allocate a pseudo-TTY

Why do we need both `-t` and `-i` options, explained here https://github.com/moby/moby/issues/2838

`-p`  Publish the container's port 80 to
the host port 8080

Don't be confused by the message

`GO web server is running on http://localhost:80`

if you publish to a host port other than 80, as inside the container it always runs on 80.

_"From inside a container you cannot figure out to which docker host port a container port is mapped to."_ explained here https://stackoverflow.com/questions/32444612/how-to-get-the-mapped-port-on-host-from-a-docker-container#:~:text=From%20inside%20a%20container%20you,container%20port%20is%20mapped%20to.