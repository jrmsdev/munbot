#!/bin/sh
set -eu
./clean.sh
exec go clean -i -modcache ./...
