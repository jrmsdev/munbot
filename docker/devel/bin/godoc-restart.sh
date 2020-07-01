#!/bin/sh
set -eu
cd /var/empty
if test -s /tmp/godoc.pid; then
	pid=$(cat /tmp/godoc.pid)
	if ps | grep godoc | grep "${pid} munbot" >/dev/null; then
		kill ${pid}
	fi
fi
exec /usr/local/bin/godoc-start.sh
