#!/bin/sh
set -eu
SRC=${1:-''}
if test '' = "${SRC}"; then
	SRC='munbot'
elif test 'munbot' = "${SRC}"; then
	SRC='munbot'
	shift
else
	SRC="munbot-${SRC}"
	shift
fi
exec go build -v -mod vendor -i -o ./_build/cmd/${SRC}.bin ./cmd/${SRC}
