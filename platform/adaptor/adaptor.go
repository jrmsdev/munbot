// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package adaptor implements the munbot gobot.Adaptor interface.
package adaptor

import (
	"gobot.io/x/gobot"
)

var _ Munbot = &Adaptor{}

type Munbot interface {
	gobot.Adaptor
	Ping() string
}

type Adaptor struct {
	name string
}

func New() *Adaptor {
	return &Adaptor{name: "Adaptor"}
}

func (m *Adaptor) Name() string {
	return m.name
}

func (m *Adaptor) SetName(name string) {
	m.name = name
}

func (m *Adaptor) Connect() error {
	return nil
}

func (m *Adaptor) Finalize() error {
	return nil
}

func (m *Adaptor) Ping() string {
	return "pong"
}
