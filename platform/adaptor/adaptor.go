// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package adaptor implements the munbot gobot.Adaptor interface.
package adaptor

import (
	"time"

	"gobot.io/x/gobot"
)

type Adaptor interface {
	gobot.Adaptor
	Interval() time.Duration
	SetInterval(time.Duration)
}

type Munbot struct {
	name     string
	interval time.Duration
}

func New() *Munbot {
	return &Munbot{
		name: "munbot",
		interval: 500 * time.Millisecond,
	}
}

// gobot interface

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

// munbot interface

func (m *Munbot) Interval() time.Duration {
	return m.interval
}

func (m *Munbot) SetInterval(d time.Duration) {
	m.interval = d
}
