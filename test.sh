#!/bin/sh
set -eu
ARGS=${@:-'./...'}
go clean -mod vendor -testcache ./env
export MBENV='test'
export MBENV_CFGDIR=${PWD}/env
exec go test -mod vendor ${ARGS}
