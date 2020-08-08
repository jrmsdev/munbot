#!/bin/sh
if test "X${1}" = 'X--all'; then
	exec gofmt -w -l -s .
fi
exec gofmt -w -l -s \
	api \
	cmd \
	config \
	console \
	core \
	log \
	mb \
	platform \
	robot \
	testing \
	utils \
	version \
	vfs
