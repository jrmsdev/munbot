#!/bin/sh
set -eu
SRC=${1:-''}
if test '' = "${SRC}"; then
	SRC='munbot'
else
	SRC="munbot-${SRC}"
	shift
fi
exec go build -mod=vendor -i -o ./_build/cmd/${SRC}.bin ./cmd/${SRC}
