#!/bin/sh
set -eu

CERT=${1:?'cert id?'}
CN=${2:-'munbot.devel'}

ca=${PWD}/docker/devel/ca/run.sh
cnf=${PWD}/docker/devel/ca/openssl.cnf

cd docker/devel/ca/files
mkdir -vp ./new ./${CERT}

cp -va ${cnf} new/
cd new

echo "**"
echo "** new cert ${CERT} ${CN}"
echo "**"

${ca} -newcert
mv -vf ./newcert.pem ../${CERT}/cert.pem
mv -vf ./newkey.pem ../${CERT}/key.rsa

cd ../${CERT}
openssl rsa -in key.rsa -out key.pem

exit 0
