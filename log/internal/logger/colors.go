// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

// colors and escape stolen from: golang.org/x/crypto/ssh/terminal/terminal.go
const escape = 27

var (
	black = []byte{escape, '[', '3', '0', 'm'}
	red = []byte{escape, '[', '3', '1', 'm'}
	green = []byte{escape, '[', '3', '2', 'm'}
	yellow = []byte{escape, '[', '3', '3', 'm'}
	blue = []byte{escape, '[', '3', '4', 'm'}
	magenta = []byte{escape, '[', '3', '5', 'm'}
	cyan = []byte{escape, '[', '3', '6', 'm'}
	white = []byte{escape, '[', '3', '7', 'm'}
	grey = []byte{escape, '[', '1', ';', '3', '0', 'm'}
	reset = string([]byte{escape, '[', '0', 'm'})
)

//~ var (
	//~ no change? = "\033[0;0m"
	//~ grey   = "\033[1;30m"
	//~ red    = "\033[0;31m"
	//~ green  = "\033[0;32m"
	//~ yellow = "\033[0;33m"
	//~ blue = "\033[0;34m"
	//~ magenta?? = "\033[0;35m"
	//~ cyan  = "\033[0;36m"
	//~ reset  = "\033[0m"
//~ )

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

func (l *Logger) SetColors(enable bool) {
	l.colored = enable
}
