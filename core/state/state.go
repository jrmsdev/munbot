// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

type State interface {
	Init() error
	Configure() error
	Start() error
	Stop() error
}

type ID int

const (
	New ID = iota
	Init
)

var idMap map[ID]string = map[ID]string{
	New:  "New",
	Init: "Init",
}

func Name(s ID) string {
	n, ok := idMap[s]
	if !ok {
		return "invalid state: " + string(s)
	}
	return n
}
