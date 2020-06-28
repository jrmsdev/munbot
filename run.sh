#!/bin/sh
set -eu
SRC=${1:-''}
./build.sh ${SRC}
if test '' = "${SRC}"; then
	SRC='munbot'
elif test 'munbot' = "${SRC}"; then
	SRC='munbot'
	shift
else
	SRC="munbot-${SRC}"
	shift
fi
exec ./_build/cmd/${SRC}.bin $@
