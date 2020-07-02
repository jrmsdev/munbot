#!/bin/sh
set -eu
cd /var/empty
export GOPATH=/godoc
exec go doc $@
