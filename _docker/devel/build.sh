#!/bin/sh
set -eu
tmpdir=./_docker/devel/tmp
rm -rvf ${tmpdir}
mkdir -vp ${tmpdir}
cp -va ./go.env ${tmpdir}/
for srcd in "./v0"; do
	dstd=${tmpdir}/${srcd}
	mkdir -vp ${dstd}
	cp -va ${srcd}/go.mod ${srcd}/go.sum ${dstd}/
	install -v -p -d ${srcd}/vendor ${dstd}/vendor
done
docker build --rm -t munbot/master:devel ./_docker/devel
rm -vrf ${tmpdir}
exit 0
