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
# https://github.com/golang/go/issues/26492#issuecomment-635563222
# STATIC="-tags 'osusergo netgo'"
TAGS='munbot'
if test 'static' = "${1:-'default'}"; then
	TAGS='munbot osusergo netgo static_build'
fi
exec go build -v -mod vendor -i -o ./_build/cmd/${SRC}.bin -tags "${TAGS}" ./cmd/${SRC}
