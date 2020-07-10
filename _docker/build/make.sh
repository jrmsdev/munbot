#!/bin/sh
set -eu
GOOS=${GOOS:-''}
GOARCH=${GOARCH:-''}
exec docker run -it --rm --network none --name munbot-build \
	--hostname build.munbot.local -u munbot \
	-e "GOOS=${GOOS}" -e "GOARCH=${GOARCH}" \
	-v ${PWD}:/munbot/src/master munbot/master:build $@
