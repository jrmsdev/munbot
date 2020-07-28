#!/bin/sh
set -eu
cd ${GOPATH}/src/master

echo "--- godoc: http://localhost:6060/"
godoc -http=:6060 &>/tmp/godoc.log &
echo $! >/tmp/godoc.pid

exit 0
