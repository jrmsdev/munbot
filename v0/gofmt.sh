#!/bin/sh
if test "X${1}" = 'Xall'; then
	exec gofmt -w -l -s .
fi
srclist() {
	for src in $(ls ./*.go); do
		echo ${src}
	done
	for srcd in $(ls -d ./* | grep -v vendor | grep -vE '^\./_'); do
		if test -d ${srcd}; then
			echo ${srcd}
		fi
	done
}
exec gofmt -w -l -s $(srclist)
