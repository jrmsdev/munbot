#!/bin/sh
set -eu
/usr/local/bin/godoc-start.sh

mkdir -vp ~/.config/munbot/master/api
ln -vsf /etc/ssl/munbot/certs/munbot.devel.pem ~/.config/munbot/master/api/cert.pem
ln -vsf /etc/ssl/munbot/private/munbot.devel.pem ~/.config/munbot/master/api/key.pem

exec /bin/sh -l
