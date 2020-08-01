#!/bin/sh
set -eu
rm -rvf ./_docker/devel/tmp
mkdir -vp ./_docker/devel/tmp
cp -va ./go.mod ./go.sum ./_docker/devel/tmp/
install -v -p -d ./vendor ./_docker/devel/tmp/vendor
docker build --rm -t munbot/master:devel ./_docker/devel
rm -vrf ./_docker/devel/tmp
exit 0
