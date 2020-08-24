#!/bin/sh
set -eu
exec docker run -it --rm --name munbot-devel --hostname devel.munbot.local \
	-e 'MBAPI_ADDR=devel.munbot.local' \
	-e 'MBSSHD_ADDR=devel.munbot.local' \
	-p 127.0.0.1:6060:6060 \
	-p 127.0.0.1:6490:6490 \
	-p 127.0.0.1:6492:6492 \
	-v ${PWD}:/munbot/src/master \
	-u munbot munbot/master:devel $@
