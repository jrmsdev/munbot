#!/bin/sh
if test "X${1}" = 'X--all'; then
	exec gofmt -w -l -s .
fi
exec gofmt -w -l -s \
	cmd \
	config \
	core \
	log \
	mb \
	robot \
	testing \
	utils \
	version \
	vfs
