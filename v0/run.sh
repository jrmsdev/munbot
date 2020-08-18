#!/bin/sh
set -eu
SRC=${1:-''}
if test '' = "${SRC}"; then
	SRC='mb'
else
	shift
fi
./build.sh ${SRC}
export MBENV='devel'
export MBENV_CONFIG=${PWD}/env
export MB_CONFIG=${PWD}/_devel/etc
export MB_RUN=${PWD}/_devel/run
exec ./_build/cmd/main/${SRC}.bin $@
