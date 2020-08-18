#!/bin/sh
set -eu
SRC=${1:-''}
BUILD=${2:-''}
if test '' = "${SRC}"; then
	SRC='mb'
else
	shift
fi
if test '' = "${BUILD}"; then
	BUILD='default'
else
	shift
fi
TAGS='munbot'
STATIC=false
PKGDIR=''
if test 'static' = "${BUILD}"; then
	# https://github.com/golang/go/issues/26492#issuecomment-635563222
	TAGS='munbot,static,osusergo,netgo'
	STATIC=true
	PKGDIR='-pkgdir ./_build/pkg'
fi
imp="github.com/munbot/master/v0/version"
BUILD_DATE="-X ${imp}.buildDate=$(date -u '+%Y%m%d.%H%M%S')"
BUILD_INFO="-X ${imp}.buildOS=$(go env GOOS)"
BUILD_INFO="${BUILD_INFO} -X ${imp}.buildArch=$(go env GOARCH)"
BUILD_INFO="${BUILD_INFO} -X ${imp}.buildTags=${TAGS}"
build_cmds=${SRC}
if test 'all' = "${build_cmds}"; then
	build_cmds='mb mbcfg'
fi
for cmd in ${build_cmds}; do
	dst=${cmd}.bin
	if ${STATIC}; then
		dst=${cmd}-static.bin
	fi
	go build -v -mod vendor -i -o ./_build/cmd/${dst} ${PKGDIR} $@ \
		-tags "${TAGS}" -ldflags "${BUILD_DATE} ${BUILD_INFO}" \
		./cmd/${cmd} || exit 1
done
exit 0
