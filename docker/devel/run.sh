#!/bin/sh
set -eu

# main port should be 6492 (pioji messages)

docker run -it --rm --name munbot-devel --hostname munbot-devel -u munbot \
	-p 127.0.0.1:6060:6060 \
	-p 127.0.0.1:9090:9090 \
	-p 127.0.0.1:3000:3000 \
	-v ${PWD}/docker/devel/ca/files:/home/munbot/.config/munbot/master/ssl \
	-v ${PWD}:/godoc/src/github.com/jrmsdev/munbot \
	-v ${PWD}/vendor:/godoc/vendor/src \
	-v ${PWD}:/go/src/munbot jrmsdev/munbot:devel $@

exit 0
