#!/bin/sh
set -eu
if test -s /tmp/godoc.pid; then
	pid=$(cat /tmp/godoc.pid)
	echo "--- kill godoc ${pid}"
	if ps | grep godoc | grep "${pid} munbot" >/dev/null; then
		kill ${pid}
	fi
fi
exit 0
