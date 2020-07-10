#!/bin/sh
set -eu
ARGS=${@:-'./...'}
exec go test -mod vendor ${ARGS}
