#!/bin/sh
set -eu
rm -vrf ./_build ./_testing
exec go clean -mod vendor -cache -testcache ./...
