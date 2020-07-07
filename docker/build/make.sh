#!/bin/sh
set -eu
GOOS=${GOOS:-''}
GOARCH=${GOARCH:-''}
exec docker run -it --rm --network none --name munbot-build --hostname munbot-build \
	-u munbot -e "GOOS=${GOOS}" -e "GOARCH=${GOARCH}" \
	-v ${PWD}:/go/src/munbot munbot/master:build $@
