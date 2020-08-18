#!/bin/sh
set -eu
PKGS=${1:-'./...'}
go clean -mod vendor -i -modcache ${PKGS}
exec ./clean.sh ${PKGS}
