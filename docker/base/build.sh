#!/bin/sh
set -eu
docker build --pull --rm -t jrmsdev/munbot:base \
	--build-arg MUNBOT_UID=$(id -u) \
	--build-arg MUNBOT_GID=$(id -g) \
	./docker/base
exit 0
