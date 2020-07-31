// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

type State interface {
	Init() error
	Configure() error
	Start() error
	Stop() error
}
