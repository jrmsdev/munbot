#!/bin/sh
set -eu
exec docker build $@ --rm -t munbot/master:base \
	--build-arg MUNBOT_UID=$(id -u) \
	--build-arg MUNBOT_GID=$(id -g) \
	--build-arg MUNBOT_UMASK=$(umask) \
	./_docker/base
