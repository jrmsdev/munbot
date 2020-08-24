#!/bin/sh
set -eu
for srcd in "v0"; do
	pidfn=/tmp/godoc-${srcd}.pid
	if test -s ${pidfn}; then
		pid=$(cat ${pidfn})
		echo "--- kill godoc ${srcd} ${pid}"
		if ps | grep godoc | grep "${pid} munbot" >/dev/null; then
			kill ${pid}
		fi
	fi
done
exit 0
