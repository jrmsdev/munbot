// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package log provides the logger functionalities.
package log

import (
	"errors"
	"fmt"
	"io"
	"os"

	gol "log"

	"github.com/munbot/master/log/internal/logger"
)

var (
	cdepth     int  = 1
	debug      bool = false
	debugFlags int  = gol.Ldate | gol.Ltime | gol.Lmicroseconds | gol.Llongfile
	verbose    bool = true
	stdFlags   int  = gol.Ldate | gol.Ltime | gol.Lmicroseconds
)

var l *logger.Logger

func init() {
	l = logger.New()
	l.SetDepth(cdepth)
	l.SetFlags(stdFlags)
	gol.SetFlags(stdFlags)
}

func SetDebug() {
	l.SetFlags(debugFlags)
	l.SetDebug(true)
	l.Lock()
	defer l.Unlock()
	debug = true
	verbose = true
	gol.SetFlags(debugFlags)
}

func SetQuiet() {
	if !debug {
		l.SetDebug(false)
		l.SetFlags(stdFlags)
		l.Lock()
		defer l.Unlock()
		verbose = false
		gol.SetFlags(stdFlags)
	}
}

func SetVerbose() {
	l.SetDebug(false)
	l.SetFlags(stdFlags)
	l.Lock()
	defer l.Unlock()
	debug = false
	verbose = true
	gol.SetFlags(stdFlags)
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

func SetColors(cfg string) {
	l.SetColors(cfg)
}

func SetPrefix(name string) {
	p := fmt.Sprintf("[%s:%d] ", name, os.Getpid())
	l.SetPrefix(p)
	l.Lock()
	defer l.Unlock()
	gol.SetPrefix(p)
}

func SetOutput(out io.Writer) {
	l.SetOutput(out)
}

func Output(calldepth int, s string) error {
	return l.Output(calldepth, s)
}

func Panic(v ...interface{}) {
	err := errors.New(fmt.Sprint(v...))
	l.Print(logger.PANIC, v...)
	panic(err)
}

func Panicf(format string, v ...interface{}) {
	err := errors.New(fmt.Sprintf(format, v...))
	l.Printf(logger.PANIC, format, v...)
	panic(err)
}

func Print(v ...interface{}) {
	if verbose {
		l.Print(logger.MSG, v...)
	}
}

func Printf(format string, v ...interface{}) {
	if verbose {
		l.Printf(logger.MSG, format, v...)
	}
}

func Debug(v ...interface{}) {
	if debug {
		l.Print(logger.DEBUG, v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if debug {
		l.Printf(logger.DEBUG, format, v...)
	}
}

func Error(v ...interface{}) error {
	err := errors.New(fmt.Sprint(v...))
	l.Print(logger.ERROR, v...)
	return err
}

func Errorf(format string, v ...interface{}) error {
	err := errors.New(fmt.Sprintf(format, v...))
	l.Printf(logger.ERROR, format, v...)
	return err
}

var osExit func(int) = os.Exit

func Fatal(v ...interface{}) {
	l.Print(logger.FATAL, v...)
	osExit(2)
}

func Fatalf(format string, v ...interface{}) {
	l.Printf(logger.FATAL, format, v...)
	osExit(2)
}

func Warn(v ...interface{}) {
	if verbose {
		l.Print(logger.WARN, v...)
	}
}

func Warnf(format string, v ...interface{}) {
	if verbose {
		l.Printf(logger.WARN, format, v...)
	}
}
