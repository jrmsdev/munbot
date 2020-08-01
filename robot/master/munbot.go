// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

type Munbot interface {
	Start() error
	Stop() error
	Running() bool
}
