#!/bin/bash

docker run -d --name dockersock \
     -v /var/run/docker.sock:/var/run/docker.sock \
     --privileged --userns=host dockersocat
