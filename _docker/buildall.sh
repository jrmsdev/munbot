#!/bin/sh
set -eu
./_docker/base/build.sh $@
./_docker/build/build.sh
./_docker/devel/build.sh
exit 0
