#!/bin/sh
if test "X${1}" = 'Xall'; then
	exec gofmt -w -l -s .
fi
exec gofmt -w -l -s \
	cmd \
	config \
	env \
	log \
	testing \
	version \
	vfs
