#!/bin/sh
set -eu
SRC=${1:-'munbot'}
BUILD=${2:-'default'}
./clean.sh
go env
sh -x ./build.sh ${SRC} ${BUILD}
echo "$(ls ./_build/cmd/${SRC}.bin) created"
exit 0
