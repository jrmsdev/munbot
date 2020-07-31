// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"errors"
	"fmt"
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
	Configure
	Run
)

var sidMap map[StateID]string = map[StateID]string{
	Dead:  "Dead",
	Init: "Init",
	Configure: "Configure",
	Run: "Run",
}

func StateName(sid StateID) string {
	n, ok := sidMap[sid]
	if !ok {
		return fmt.Sprintf("invalid state id: %v", sid)
	}
	return n
}
