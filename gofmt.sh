#!/bin/sh
exec gofmt -w -l -s \
	cmd \
	config \
	log \
	mb \
	state \
	testing \
	version \
	vfs
