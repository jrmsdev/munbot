// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"errors"
	"fmt"
)

var ErrInit error = errors.New("can not run Init in this state")
var ErrConfigure error = errors.New("can not run Configure in this state")
var ErrStart error = errors.New("can not run Start in this state")
var ErrStop error = errors.New("can not run Stop in this state")

type State interface {
	Init() error
	Configure() error
	Start() error
	Stop() error
}

type StateID int

const (
	Dead StateID = iota
	Init
	Run
)

var sidMap map[StateID]string = map[StateID]string{
	Dead:      "Dead",
	Init:      "Init",
	Run:       "Run",
}

func StateName(sid StateID) string {
	n, ok := sidMap[sid]
	if !ok {
		return fmt.Sprintf("invalid state id: %v", sid)
	}
	return n
}
