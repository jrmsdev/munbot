#!/bin/sh
set -eu

docker run -it --rm --name munbot-devel --hostname munbot-devel -u munbot \
	-p 127.0.0.1:6060:6060 -v ${PWD}:/go/src/munbot jrmsdev/munbot:devel

exit 0
