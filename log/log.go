// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package log provides the logger functionalities.
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
	stdFlags   int  = gol.Ldate | gol.Ltime | gol.Lmicroseconds
)

var Output func(int, string) error = gol.Output

var setFlags func(int) = gol.SetFlags
var setPrefix func(string) = gol.SetPrefix

func init() {
	setFlags(stdFlags)
}

func SetDebug() {
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

func SetMode(lvl string) {
	switch lvl {
	case "quiet":
		SetQuiet()
	case "debug":
		SetDebug()
	default:
		SetVerbose()
	}
}

func SetPrefix(name string) {
	setPrefix(fmt.Sprintf("[%s:%d] ", name, os.Getpid()))
}

func Panic(v ...interface{}) {
	err := errors.New(fmt.Sprint(v...))
	Output(cdepth, fmt.Sprintf("[PANIC] %s", err))
	panic(err)
}

func Panicf(format string, v ...interface{}) {
	err := errors.New(fmt.Sprintf(format, v...))
	Output(cdepth, fmt.Sprintf("[PANIC] %s", err))
	panic(err)
}

func Print(v ...interface{}) {
	if verbose {
		Output(cdepth, fmt.Sprint(v...))
	}
}

func Printf(format string, v ...interface{}) {
	if verbose {
		Output(cdepth, fmt.Sprintf(format, v...))
	}
}

func Debug(v ...interface{}) {
	if debug {
		Output(cdepth, fmt.Sprint(v...))
	}
}

func Debugf(format string, v ...interface{}) {
	if debug {
		Output(cdepth, fmt.Sprintf(format, v...))
	}
}

func Error(v ...interface{}) error {
	err := errors.New(fmt.Sprint(v...))
	Output(cdepth, fmt.Sprintf("[ERROR] %s", err))
	return err
}

func Errorf(format string, v ...interface{}) error {
	err := errors.New(fmt.Sprintf(format, v...))
	Output(cdepth, fmt.Sprintf("[ERROR] %s", err))
	return err
}

var osExit func(int) = os.Exit

func Fatal(v ...interface{}) {
	Output(cdepth, fmt.Sprintf("[FATAL] %s", fmt.Sprint(v...)))
	osExit(2)
}

func Fatalf(format string, v ...interface{}) {
	Output(cdepth, fmt.Sprintf("[FATAL] %s", fmt.Sprintf(format, v...)))
	osExit(2)
}

func Warn(v ...interface{}) {
	if verbose {
		Output(cdepth, fmt.Sprintf("[WARNING] %s", fmt.Sprint(v...)))
	}
}

func Warnf(format string, v ...interface{}) {
	if verbose {
		Output(cdepth, fmt.Sprintf("[WARNING] %s", fmt.Sprintf(format, v...)))
	}
}
