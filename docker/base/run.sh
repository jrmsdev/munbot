#!/bin/sh
set -eu

docker run -it --rm --network none --name munbot-base --hostname munbot-base \
	-u munbot jrmsdev/munbot:base

exit 0
