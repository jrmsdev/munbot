#!/bin/sh
set -eu
port=6060
for srcd in "v0"; do
	cd ${GOPATH}/src/master/${srcd}
	echo "--- godoc ${srcd}: http://localhost:${port}/"
	godoc -http=:${port} &>/tmp/godoc-${srcd}.log &
	echo $! >/tmp/godoc-${srcd}.pid
	port=$(expr ${port} + 1)
done
exit 0
