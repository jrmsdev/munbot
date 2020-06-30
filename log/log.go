// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"errors"
	"fmt"
	gol "log"
)

var (
	cdepth     int  = 2
	debug      bool = false
	debugFlags int  = gol.Ldate | gol.Ltime | gol.Lmicroseconds | gol.Llongfile
)

func DebugEnable() {
	debug = true
	gol.SetFlags(debugFlags)
}

func Panic(v ...interface{}) {
	gol.Output(cdepth, fmt.Sprintf("[PANIC] %s", fmt.Sprint(v...)))
	panic("oops!!")
}

func Panicf(format string, v ...interface{}) {
	gol.Output(cdepth, fmt.Sprintf("[PANIC] %s", fmt.Sprintf(format, v...)))
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

func Error(v ...interface{}) error {
	err := errors.New(fmt.Sprint(v...))
	gol.Output(cdepth, fmt.Sprintf("[ERROR] %s", err))
	return err
}

func Errorf(format string, v ...interface{}) error {
	err := errors.New(fmt.Sprintf(format, v...))
	gol.Output(cdepth, fmt.Sprintf("[ERROR] %s", err))
	return err
}
