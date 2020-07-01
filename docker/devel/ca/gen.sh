#!/bin/sh
set -eu

workdir=${PWD}
ca=${PWD}/docker/devel/ca/run.sh
cert=${PWD}/docker/devel/ca/newcert.sh

mkdir -vp /var/tmp/munbot/ca
cp -va docker/devel/ca/openssl.cnf /var/tmp/munbot/ca/
cd /var/tmp/munbot/ca

${ca} -newca

cp -va cacert.pem ${workdir}/docker/devel/ca/files
cp -va private/cakey.pem ${workdir}/docker/devel/ca/files/cakey.rsa
openssl rsa -in ${workdir}/docker/devel/ca/files/cakey.rsa \
	-out ${workdir}/docker/devel/ca/files/cakey.pem

cd ${workdir}
${cert} api master.munbot.devel

exit 0
