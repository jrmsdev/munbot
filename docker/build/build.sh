#!/bin/sh
set -eu
docker build --rm -t munbot/master:build ./docker/build
exit 0
