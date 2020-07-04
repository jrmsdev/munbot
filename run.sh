#!/bin/sh
set -eu
SRC=${1:-''}
if test '' = "${SRC}"; then
	SRC='munbot'
else
	shift
fi
./build.sh ${SRC}
exec ./_build/cmd/${SRC}.bin $@
