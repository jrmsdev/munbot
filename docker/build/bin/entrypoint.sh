#!/bin/sh
set -eu
export GOPATH=${PWD}/_build
./clean.sh
mkdir -vp ${GOPATH}
echo '--- munbot build'
go env
exec ./build.sh
