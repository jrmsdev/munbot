#!/bin/sh
set -eu
./docker/base/build.sh $@
./docker/build/build.sh
./docker/devel/build.sh
exit 0
