#!/bin/sh

export ADDR_PORT=:8210

export DOCKER_HOST=unix://$HOME/docker.sock

./containers2ql
