#!/bin/sh
set -eu
docker run -it --rm --network none --name munbot-build --hostname munbot-build \
	-u munbot -v ${PWD}:/go/src/munbot jrmsdev/munbot:build
exit 0
