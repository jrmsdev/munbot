#!/bin/sh
set -eu
/usr/local/bin/godoc-stop.sh
exec /usr/local/bin/godoc-start.sh
