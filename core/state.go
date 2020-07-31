// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

type State interface {
	Init() error
	Configure() error
	Start() error
	Stop() error
}

type StateID int

const (
	New StateID = iota
	Init
)

var sidMap map[StateID]string = map[StateID]string{
	New:  "New",
	Init: "Init",
}

func StateName(sid StateID) string {
	n, ok := sidMap[sid]
	if !ok {
		return "invalid state: " + string(sid)
	}
	return n
}
