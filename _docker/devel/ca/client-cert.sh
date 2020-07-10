#!/bin/sh
set -eu

CERT=${1:?'cert id?'}
CN=${2:-'munbot.devel'}

newcert=${PWD}/docker/devel/ca/newcert.sh

${newcert} ${CERT} ${CN}

cd docker/devel/ca/files/${CERT}

openssl pkcs12 -export -out cert.p12 -inkey key.rsa -in cert.pem

exit 0
