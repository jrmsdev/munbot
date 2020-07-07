#!/bin/sh
set -eu
docker build $@ --rm -t munbot/master:base \
	--build-arg MUNBOT_UID=$(id -u) \
	--build-arg MUNBOT_GID=$(id -g) \
	--build-arg MUNBOT_UMASK=$(umask) \
	./docker/base
exit 0
