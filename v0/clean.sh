#!/bin/sh
set -eu
PKGS=${1:-'./...'}
rm -vrf ./_build ./_testing
exec go clean -mod vendor -cache -testcache ${PKGS}
