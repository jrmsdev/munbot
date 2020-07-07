#!/bin/sh
set -eu
mkdir -vp ./docker/devel/tmp
cp -va ./go.mod ./go.sum ./docker/devel/tmp
docker build --rm -t munbot/master:devel ./docker/devel
exit 0
