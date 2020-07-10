#!/bin/sh
set -eu
mkdir -vp ./_docker/devel/tmp
cp -va ./go.mod ./go.sum ./_docker/devel/tmp
docker build --rm -t munbot/master:devel ./_docker/devel
exit 0
