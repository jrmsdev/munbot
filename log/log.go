// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"fmt"
	gol "log"
)

var (
	cdepth int = 2
	defFlags int = gol.Flags()
	debug bool = false
	debugFlags int = gol.Ldate | gol.Ltime | gol.Lmicroseconds | gol.Llongfile
)

func DebugEnable() {
	debug = true
	gol.SetFlags(debugFlags)
}

func DebugDisable() {
	debug = false
	gol.SetFlags(defFlags)
}

func Panic(v ...interface{}) {
	gol.Output(cdepth, fmt.Sprint(v...))
	panic("oops!!")
}

func Panicf(format string, v ...interface{}) {
	gol.Output(cdepth, fmt.Sprintf(format, v...))
	panic("oops!!")
}

func Print(v ...interface{}) {
	gol.Output(cdepth, fmt.Sprint(v...))
}

func Printf(format string, v ...interface{}) {
	gol.Output(cdepth, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	if debug {
		gol.Output(cdepth, fmt.Sprint(v...))
	}
}

func Debugf(format string, v ...interface{}) {
	if debug {
		gol.Output(cdepth, fmt.Sprintf(format, v...))
	}
}
