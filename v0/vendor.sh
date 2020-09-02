#!/bin/sh
set -eu
if test '--upgrade' = "${1:-''}"; then
	mkdir -p ./_build
	cp -a go.mod ./_build/go.mod
	cat ./_build/go.mod | sed 's#^\t\([^ ]*\) .*#   \1 latest#' >go.mod
fi
go mod vendor
exec go mod tidy
