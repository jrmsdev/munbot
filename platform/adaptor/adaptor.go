// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package adaptor implements the munbot gobot.Adaptor interface.
package adaptor

import (
	"gobot.io/x/gobot"
)

type Adaptor interface {
	gobot.Adaptor
	Ping() string
}

type Munbot struct {
	name string
}

func New() *Munbot {
	return &Munbot{name: "Munbot"}
}

func (m *Munbot) Name() string {
	return m.name
}

func (m *Munbot) SetName(name string) {
	m.name = name
}

func (m *Munbot) Connect() error {
	return nil
}

func (m *Munbot) Finalize() error {
	return nil
}

func (m *Munbot) Ping() string {
	return "pong"
}
