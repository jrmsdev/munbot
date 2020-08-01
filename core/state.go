// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"errors"
	"fmt"
)

var ErrInit error = errors.New("can not Init in this state")
var ErrConfigure error = errors.New("can not Configure in this state")
var ErrStart error = errors.New("can not Start in this state")
var ErrRun error = errors.New("can not Run in this state")
var ErrStop error = errors.New("can not Stop in this state")
var ErrHalt error = errors.New("can not Halt in this state")

type State interface {
	Init() error
	Configure() error
	Start() error
	Run() error
	Stop() error
	Halt() error
}

type StateID int

const (
	Dead StateID = iota
	Init
	Run
	Halt
)

var sidMap map[StateID]string = map[StateID]string{
	Dead: "Dead",
	Init: "Init",
	Run:  "Run",
	Halt: "Halt",
}

func StateName(sid StateID) string {
	n, ok := sidMap[sid]
	if !ok {
		return fmt.Sprintf("invalid state id: %v", sid)
	}
	return n
}
