#!/bin/sh
exec gofmt -w -l -s \
	cmd \
	config \
	log \
	mb \
	testing \
	version \
	vfs
