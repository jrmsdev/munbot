#!/bin/sh
set -eu
test -x ./docker/base/build.sh && ./docker/base/build.sh
docker build --rm -t jrmsdev/munbot:devel \
	--build-arg MUNBOT_UMASK=$(umask) \
	./docker/devel
exit 0
