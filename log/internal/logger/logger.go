// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"io"
	"log"
	"sync"

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
	INFO:  "",
	DEBUG: "",
}

type Logger struct {
	*sync.Mutex
	depth   int
	colored bool
}

func New() *Logger {
	return &Logger{Mutex: new(sync.Mutex)}
}

func (l *Logger) SetDepth(n int) {
	l.Lock()
	defer l.Unlock()
	l.depth = n
}

func (l *Logger) SetOutput(out io.Writer) {
	l.Lock()
	defer l.Unlock()
	log.SetOutput(out)
}

func (l *Logger) SetFlags(f int) {
	l.Lock()
	defer l.Unlock()
	log.SetFlags(f)
}

func (l *Logger) tag(lvl Level, msg string) string {
	tag := levelTag[lvl]
	return tag+msg
}

func (l *Logger) Print(lvl Level, args ...interface{}) {
	msg := gfmt.Sprint(args...)
	if l.colored {
		log.Output(l.depth, l.color(lvl, msg))
	} else {
		log.Output(l.depth, l.tag(lvl, msg))
	}
}

func (l *Logger) Printf(lvl Level, fmt string, args ...interface{}) {
	msg := gfmt.Sprintf(fmt, args...)
	if l.colored {
		log.Output(l.depth, l.color(lvl, msg))
	} else {
		log.Output(l.depth, l.tag(lvl, msg))
	}
}
