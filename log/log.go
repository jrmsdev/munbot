// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	gol "log"
)

func Panic(v ...interface{}) {
	gol.Panic(v)
}

func Printf(format string, v ...interface{}) {
	gol.Printf(format, v)
}
