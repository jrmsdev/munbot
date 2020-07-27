#!/bin/sh
exec gofmt -w -l -s \
	cmd \
	config \
	core \
	log \
	mb \
	state \
	testing \
	version \
	vfs
