#!/bin/sh
set -eu
docker build --rm -t jrmsdev/munbot:build ./docker/build
exit 0
