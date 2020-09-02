#!/bin/sh
exec docker run -it --rm --network none --name munbot-base \
	--hostname base.munbot.local -u munbot munbot/master:base
