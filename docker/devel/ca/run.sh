#!/bin/sh
set -eu
export OPENSSL_CATOP=/var/tmp/munbot/ca
export OPENSSL_CONFIG=./openssl.cnf
exec /etc/ssl/misc/CA.pl $@
