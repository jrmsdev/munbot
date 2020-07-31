// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package robot defines and implements the master robot interface.
package master

var _ Munbot = &Robot{}

type Robot struct {
}

func New() Munbot {
	return &Robot{}
}
