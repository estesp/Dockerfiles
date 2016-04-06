## Dockersocat

When user namespaces are enabled in a Docker daemon, mounting
the Docker API UNIX socket into the container is not directly
useful. Without munging the ownership of the UNIX socket, the
container will have no access to the socket for either read or
write.

This `Dockerfile` builds a simple container that can use the
new `--privileged` capability in Docker 1.11 (a privileged container
even while the daemon has user namespaces enabled) to pass
traffic from a TCP endpoint to the UNIX socket using socat.

### Build & Run

Building can be performed with a simple `docker build`:

```
docker build -t dockersocat .
```

> **NOTE**: This **requires** Docker 1.11 for the `--privileged`
> support for user namespaced-enabled daemon.

Use the `run.sh` script or run the container as follows:

```
docker run -d --name dockersock \
    -v /var/run/docker.sock:/var/run/docker.sock \
	--privileged --userns=host dockersocat
```

Note that I am not portmapping the TCP listener to the host as the
expectation is that inter-container communication is on and other
user namespaced containers are the consumers of this service and
will connect to the container IP at port :2375 for `DOCKER_HOST`.
Of course you could portmap this to the host, but this exposes your
Docker engine endpoint to a broad audience with all the usual concerns
and risks for doing so.

### Use from other (unprivileged) containers

To use the socket TCP->UNIX forwarding container from a user
namespaced container, I will show an example using "links". If you
are using modern libnetwork/overlay networking, using the embedded
DNS will be the future-proof path, given "links" are a deprecated
feature. For a basic example, however, it's easy to show using a
simple container that has a Docker client installed:

```
docker run -ti --rm --link dockersock:dockersock dockerclient
/ # export DOCKER_HOST=tcp://dockersock:2375
/ # docker version
Client:
 Version:      1.10.0-dev
 API version:  1.22
 Go version:   go1.5.2
 Git commit:   8ed14c2
 Built:        Wed Dec  9 01:04:08 2015
 OS/Arch:      linux/amd64
 Experimental: true

Server:
 Version:      1.11.0-rc3
 API version:  1.23
 Go version:   go1.5.3
 Git commit:   eabf97a
 Built:        Fri Apr  1 22:26:46 2016
 OS/Arch:      linux/amd64
/ #
```
