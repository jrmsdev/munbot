#!/bin/sh
set -eu
exec docker run -it --rm --name munbot-devel --hostname devel.munbot.local \
	-p 127.0.0.1:6060:6060 \
	-p 127.0.0.1:9090:9090 \
	-p 127.0.0.1:6492:6492 \
	-v ${PWD}/_docker/devel/ca/files:/home/munbot/.config/munbot/master/ssl \
	-v ${PWD}:/godoc/src/github.com/munbot/master \
	-v ${PWD}/vendor:/godoc/vendor/src \
	-v ${PWD}:/munbot/src/master \
	-u munbot munbot/master:devel $@
