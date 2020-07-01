#!/bin/sh
set -eu
cd /var/empty
exec go doc $@
