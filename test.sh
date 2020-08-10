#!/bin/sh
set -eu
ARGS=${@:-'./...'}
export MBENV='test'
export MBENV_CONFIG=${PWD}/env
exec go test -mod vendor ${ARGS}
