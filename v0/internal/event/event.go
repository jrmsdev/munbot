// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package event implements the internal core eventer and its shared events.
package event

import (
	"sync"

	"gobot.io/x/gobot"
	//~ "github.com/munbot/master/v0/log"
)

type Error struct {
	Type string
	Err  error
}

const (
	Fail      string = "fail"
	ApiStart         = "api.start"
	ApiStop          = "api.stop"
	SSHDStart        = "sshd.start"
	SSHDStop         = "sshd.stop"
)

type Eventer interface {
	gobot.Eventer
	Wait()
}

type evtr struct {
	wg *sync.WaitGroup
	gobot.Eventer
}

func NewEventer() Eventer {
	return &evtr{
		wg:      new(sync.WaitGroup),
		Eventer: gobot.NewEventer(),
	}
}

func (m *evtr) Wait() {
	m.wg.Wait()
}

func (m *evtr) On(name string, f func(d interface{})) error {
	m.wg.Add(1)
	wrapper := func(d interface{}) {
		defer m.wg.Done()
		f(d)
	}
	return m.Eventer.On(name, wrapper)
}

func (m *evtr) Once(name string, f func(d interface{})) error {
	m.wg.Add(1)
	wrapper := func(d interface{}) {
		defer m.wg.Done()
		f(d)
	}
	return m.Eventer.On(name, wrapper)
}
