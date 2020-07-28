#!/bin/sh
set -eu
for pkg in $(cat ./vendor-upgrade.deps); do
	go get -v -u ${pkg}
done
exec ./vendor.sh
