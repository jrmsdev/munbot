#!/bin/sh
set -eu
cd /var/empty

echo "--- munbot godoc: http://localhost:9090/"
GOPATH=/godoc godoc -http=:9090 &>/tmp/godoc.log &
echo $! >/tmp/godoc.pid

echo "--- vendor godoc: http://localhost:6060/"
GOPATH=/godoc/vendor godoc -http=:6060 &>/tmp/godoc-vendor.log &
echo $! >/tmp/godoc-vendor.pid

exit 0
