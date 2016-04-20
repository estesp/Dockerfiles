#!/bin/bash

docker run -d --name unprivsockfwd \
     -v /var/run/docker.sock:/var/run/docker.sock \
     --privileged --userns=host unprivsockfwd
