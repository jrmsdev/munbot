// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// master state machine.
package state

import (
	"context"
)

type Status int

const (
	OK Status = iota
	DONE
	EXIT
	ERROR
	PANIC
	lastStatus // for testing purposes
)

var stMap = map[Status]string{
	OK:    "OK",
	DONE:  "DONE",
	EXIT:  "EXIT",
	ERROR: "ERROR",
	PANIC: "PANIC",
}

type State interface {
	Error() error
	Run(context.Context) Status
	String() string
}
