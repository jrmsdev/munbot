#!/bin/sh
set -eu
ca=/go/src/munbot/docker/devel/ca/run.sh
${ca} -newca

cd /go/src/munbot/docker/devel/ca/db
mkdir -vp new certs private
cd new

${ca} -newcert
mv -v ./newcert.pem ../certs/munbot.devel.pem
mv -v ./newkey.pem ../private/munbot.devel.pem

exit 0
