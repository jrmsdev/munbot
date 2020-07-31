// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"errors"
)

var ErrInit error = errors.New("Init")
var ErrConfigure error = errors.New("Configure")
var ErrStart error = errors.New("Start")
var ErrStop error = errors.New("Stop")

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
)

var sidMap map[StateID]string = map[StateID]string{
	Dead:  "Dead",
	Init: "Init",
}

func StateName(sid StateID) string {
	n, ok := sidMap[sid]
	if !ok {
		return "invalid state: " + string(sid)
	}
	return n
}
