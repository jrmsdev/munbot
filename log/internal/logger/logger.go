// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"io"
	"log"

	gfmt "fmt"
)

type Level int

const (
	PANIC Level = iota
	FATAL
	ERROR
	WARN
	MSG
	INFO
	DEBUG
)

var levelTag = map[Level]string{
	PANIC: "[PANIC] ",
	FATAL: "[FATAL] ",
	ERROR: "[ERROR] ",
	WARN:  "[WARNING] ",
	MSG:   "",
	INFO:  "[INFO] ",
	DEBUG: "[DEBUG] ",
}

type Logger struct {
	depth   int
	colored bool
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) SetDepth(n int) {
	l.depth = n
}

func (l *Logger) Print(lvl Level, args ...interface{}) {
	msg := gfmt.Sprint(args...)
	if l.colored {
		clr := levelColor[lvl]
		log.Output(l.depth, clr+msg+reset)
	} else {
		tag := levelTag[lvl]
		log.Output(l.depth, tag+msg)
	}
}

func (l *Logger) Printf(lvl Level, fmt string, args ...interface{}) {
	msg := gfmt.Sprintf(fmt, args...)
	if l.colored {
		clr := levelColor[lvl]
		log.Output(l.depth, clr+msg+reset)
	} else {
		tag := levelTag[lvl]
		log.Output(l.depth, tag+msg)
	}
}

func (l *Logger) SetOutput(out io.Writer) {
	log.SetOutput(out)
}

func (l *Logger) SetFlags(f int) {
	log.SetFlags(f)
}
