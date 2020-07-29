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
	lastStatus // for testing purposes
)

var stMap = map[Status]string{
	OK:    "OK",
	DONE:  "DONE",
	EXIT:  "EXIT",
	ERROR: "ERROR",
}

func (s Status) String() string {
	return stMap[s]
}

type State interface {
	Error() error
	Run(context.Context) (context.Context, Status)
	String() string
}

type StateID int

const (
	__zero StateID = iota // as int are init as 0 we need to start at 1
	Init
	Configure
	Start
	lastStateID // for testing purposes
)

var stidMap = map[StateID]string{
	Init:      "Init",
	Configure: "Configure",
	Start:     "Start",
}

func (s StateID) String() string {
	return stidMap[s]
}
