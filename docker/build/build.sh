#!/bin/sh
exec docker build --rm -t munbot/master:build ./docker/build
