#!/bin/sh
set -eu
SRC=${1:-'munbot'}
./build.sh ${SRC}
exec ./_build/cmd/${SRC}.bin $@
