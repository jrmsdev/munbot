#!/bin/sh
set -eu

docker run -it --rm --name munbot-devel --hostname munbot-devel -u munbot \
	-p 127.0.0.1:6060:6060 \
	-p 127.0.0.1:3000:3000 \
	-v ${PWD}/docker/devel/ca/db:/etc/ssl/munbot \
	-v ${PWD}/vendor:/go/src/vendor_deps \
	-v ${PWD}:/go/src/munbot jrmsdev/munbot:devel

exit 0
