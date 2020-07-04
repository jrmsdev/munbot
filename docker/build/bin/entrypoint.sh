#!/bin/sh
set -eu
./clean.sh
go env
echo '***'
echo "*** build $@"
echo '***'
exec sh -x ./build.sh $@
