// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// master state machine.
package state

type Status int

const (
	OK Status = iota
	EXIT
	ERROR
	PANIC
)

var stMap = map[Status]string{
	OK: "OK",
	EXIT: "EXIT",
	ERROR: "ERROR",
	PANIC: "PANIC",
}

type State interface {
	Error() error
	Run() Status
	String() string
}
