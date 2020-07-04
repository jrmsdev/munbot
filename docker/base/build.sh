#!/bin/sh
set -eu
docker build $@ --rm -t jrmsdev/munbot:base \
	--build-arg MUNBOT_UID=$(id -u) \
	--build-arg MUNBOT_GID=$(id -g) \
	--build-arg MUNBOT_UMASK=$(umask) \
	./docker/base
exit 0
