// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package worker defines and implements the worker robot interface.
package worker

var _ Munbot = &Robot{}

type Robot struct {
}

func New() *Robot {
	return &Robot{}
}
