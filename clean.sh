#!/bin/sh
set -eu
rm -vrf ./_build ./_testing
exec go clean -cache -testcache ./...
