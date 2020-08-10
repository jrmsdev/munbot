#!/bin/sh
set -eu
SRC=${1:-''}
if test '' = "${SRC}"; then
	SRC='munbot'
else
	shift
fi
./build.sh ${SRC}
export MBENV='devel'
export MB_CONFIG=${PWD}/_devel/etc
exec ./_build/cmd/${SRC}.bin $@
