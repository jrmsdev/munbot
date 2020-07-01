#!/bin/sh
set -eu
ca=/go/src/munbot/docker/devel/ca/run.sh
${ca} -newca
exit 0
