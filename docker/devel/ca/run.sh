#!/bin/sh
set -eu
export OPENSSL_CATOP=/go/src/munbot/docker/devel/ca/db
export OPENSSL_CONFIG=/go/src/munbot/docker/devel/ca/openssl.cnf
exec /etc/ssl/misc/CA.pl $@
