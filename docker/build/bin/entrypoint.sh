#!/bin/sh
set -eu
./clean.sh
go env
exec sh -x ./build.sh $@
