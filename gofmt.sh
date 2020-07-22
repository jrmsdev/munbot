#!/bin/sh
exec gofmt -w -l -s \
	cmd \
	config \
	log \
	profile \
	testing \
	version \
	vfs
