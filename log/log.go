// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"errors"
	"fmt"
	"os"

	gol "log"
)

var (
	cdepth     int  = 2
	debug      bool = false
	debugFlags int  = gol.Ldate | gol.Ltime | gol.Lmicroseconds | gol.Llongfile
	verbose    bool = true
)

var output func(int, string) error = gol.Output
var setFlags func(int) = gol.SetFlags
var setPrefix func(string) = gol.SetPrefix

func DebugEnable() {
	debug = true
	verbose = true
	setFlags(debugFlags)
}

func SetQuiet() {
	if !debug {
		verbose = false
	}
}

func SetVerbose() {
	verbose = true
}

func SetPrefix(name string) {
	setPrefix(fmt.Sprintf("[%s:%d] ", name, os.Getpid()))
}

func Panic(v ...interface{}) {
	output(cdepth, fmt.Sprintf("[PANIC] %s", fmt.Sprint(v...)))
	panic("oops!!")
}

func Panicf(format string, v ...interface{}) {
	output(cdepth, fmt.Sprintf("[PANIC] %s", fmt.Sprintf(format, v...)))
	panic("oops!!")
}

func Print(v ...interface{}) {
	if verbose {
		output(cdepth, fmt.Sprint(v...))
	}
}

func Printf(format string, v ...interface{}) {
	if verbose {
		output(cdepth, fmt.Sprintf(format, v...))
	}
}

func Debug(v ...interface{}) {
	if debug {
		output(cdepth, fmt.Sprint(v...))
	}
}

func Debugf(format string, v ...interface{}) {
	if debug {
		output(cdepth, fmt.Sprintf(format, v...))
	}
}

func Error(v ...interface{}) error {
	err := errors.New(fmt.Sprint(v...))
	output(cdepth, fmt.Sprintf("[ERROR] %s", err))
	return err
}

func Errorf(format string, v ...interface{}) error {
	err := errors.New(fmt.Sprintf(format, v...))
	output(cdepth, fmt.Sprintf("[ERROR] %s", err))
	return err
}

var osExit func(int) = os.Exit

func Fatal(v ...interface{}) {
	output(cdepth, fmt.Sprintf("[FATAL] %s", fmt.Sprint(v...)))
	osExit(2)
}

func Fatalf(format string, v ...interface{}) {
	output(cdepth, fmt.Sprintf("[FATAL] %s", fmt.Sprintf(format, v...)))
	osExit(2)
}

func Warn(v ...interface{}) {
	if verbose {
		output(cdepth, fmt.Sprintf("[WARNING] %s", fmt.Sprint(v...)))
	}
}

func Warnf(format string, v ...interface{}) {
	if verbose {
		output(cdepth, fmt.Sprintf("[WARNING] %s", fmt.Sprintf(format, v...)))
	}
}
