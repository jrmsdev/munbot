// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"fmt"
	"os"
	"strings"
)

// colors and escape stolen from: golang.org/x/crypto/ssh/terminal/terminal.go
const escape = 27

var (
	reset = "0"
	black = "30"
	red = "31"
	green = "32"
	yellow = "33"
	blue = "34"
	magenta = "35"
	cyan = "36"
	white = "37"
	grey = "1;30"
)

var levelColor = map[Level]string{
	PANIC: string(red),
	FATAL: string(red),
	ERROR: string(red),
	WARN:  string(yellow),
	MSG:   string(magenta),
	INFO:  string(cyan),
	DEBUG: string(green),
}

func (l *Logger) Colors() bool {
	return l.colored
}

func (l *Logger) color(lvl Level, msg string) string {
	esc := []byte{escape}
	col := levelColor[lvl]
	return fmt.Sprintf("%s[%sm%s%s[%sm", esc, col, msg, esc, reset)
}

func (l *Logger) SetColors(cfg string) {
	l.Lock()
	defer l.Unlock()
	cfg = strings.TrimSpace(cfg)
	l.colored = false
	switch cfg {
	case "":
		return
	case "off":
		return
	case "on":
		l.colored = true
	case "auto":
		if istty(os.Stdout) && istty(os.Stderr) {
			l.colored = true
		}
	}
	if l.colored {
		setColors(cfg)
	}
}

func istty(fh *os.File) bool {
	if st, err := fh.Stat(); err == nil {
		m := st.Mode()
		if m&os.ModeDevice != 0 && m&os.ModeCharDevice != 0 {
			return true
		}
	}
	return false
}

func setColors(cfg string) {
}
