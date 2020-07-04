#!/bin/sh
set -eu
SRC=${1:-''}
NAME='munbot'
if test '' = "${SRC}"; then
	SRC='munbot'
elif test 'munbot' = "${SRC}"; then
	SRC='munbot'
	shift
else
	NAME=${SRC}
	SRC="munbot-${SRC}"
	shift
fi
set -x
./clean.sh
go env
sh -x ./build.sh ${NAME} $@
set +x
echo "$(ls ./_build/cmd/${SRC}.bin) created"
exit 0
