#!/bin/sh
set -eu
go mod vendor
exec go mod tidy
