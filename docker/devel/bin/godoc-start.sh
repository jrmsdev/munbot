#!/bin/sh
set -eu
echo "--- godoc: http://localhost:6060/"
cd /var/empty
godoc -http=:6060 &>/tmp/godoc.log &
echo $! >/tmp/godoc.pid
exit 0
