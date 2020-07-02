#!/bin/sh
set -eu
cd /var/empty
export GOPATH=/godoc/vendor
exec go doc $@
