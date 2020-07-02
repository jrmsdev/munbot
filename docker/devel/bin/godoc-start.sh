#!/bin/sh
set -eu
echo "--- godoc: http://localhost:6060/"
cd /var/empty
export GOPATH=/godoc
godoc -http=:6060 &>/tmp/godoc.log &
echo $! >/tmp/godoc.pid
exit 0
