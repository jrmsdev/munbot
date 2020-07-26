#!/bin/sh
exec gofmt -w -l -s \
	cmd \
	config \
	log \
	testing \
	version \
	vfs
