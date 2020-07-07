#!/bin/sh
set -eu
SRC=${1:-''}
BUILD=${2:-''}
if test '' = "${SRC}"; then
	SRC='munbot'
else
	shift
fi
if test '' = "${BUILD}"; then
	BUILD='default'
else
	shift
fi
TAGS='munbot'
if test 'static' = "${BUILD}"; then
	# https://github.com/golang/go/issues/26492#issuecomment-635563222
	TAGS='munbot,static,osusergo,netgo'
fi
imp="github.com/munbot/master/version"
BUILD_DATE="-X ${imp}.buildDate=$(date -u '+%Y%m%d.%H%M%S')"
BUILD_INFO="-X ${imp}.buildOS=$(go env GOOS)"
BUILD_INFO="${BUILD_INFO} -X ${imp}.buildArch=$(go env GOARCH)"
BUILD_INFO="${BUILD_INFO} -X ${imp}.buildTags=${TAGS}"
build_cmds=${SRC}
if test 'all' = "${build_cmds}"; then
	build_cmds='munbot munbot-config'
fi
for cmd in ${build_cmds}; do
	go build -v -mod vendor -i -o ./_build/cmd/${cmd}.bin $@ -tags "${TAGS}" \
		-ldflags "${BUILD_DATE} ${BUILD_INFO}" ./cmd/${cmd} || exit 1
done
exit 0
