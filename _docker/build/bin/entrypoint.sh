#!/bin/sh
set -eu
SRC=${1:-'munbot'}
BUILD=${2:-'default'}
./clean.sh
go env
echo '***'
echo "*** build $@"
echo '***'
exec sh -x ./build.sh ${SRC} ${BUILD} -pkgdir ./_build/pkg
