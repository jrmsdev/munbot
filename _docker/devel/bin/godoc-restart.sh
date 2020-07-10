#!/bin/sh
set -eu
cd /var/empty

echo "--- kill munbot godoc"
if test -s /tmp/godoc.pid; then
	pid=$(cat /tmp/godoc.pid)
	if ps | grep godoc | grep "${pid} munbot" >/dev/null; then
		kill ${pid}
	fi
fi

echo "--- kill vendor godoc"
if test -s /tmp/godoc-vendor.pid; then
	pid=$(cat /tmp/godoc-vendor.pid)
	if ps | grep godoc | grep "${pid} munbot" >/dev/null; then
		kill ${pid}
	fi
fi

exec /usr/local/bin/godoc-start.sh
