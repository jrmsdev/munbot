// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package state

import (
	"context"
	"time"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core"
	"github.com/munbot/master/log"
)

var _ Machine = &sm{}

type Machine interface {
	Init(*config.Flags, *core.Flags) error
	Run(context.Context) error
}

type hist struct {
	Time time.Time `json:"time"`
	State string `json:"state"`
}

type sm struct {
	hist        []*hist
	st          State
	init        State
	configure   State
	start       State
	Config      *config.Config
	ConfigFlags *config.Flags
	CoreFlags   *core.Flags
	Runtime     *core.Runtime
}

func NewMachine() Machine {
	m := &sm{}
	m.hist = make([]*hist, 0)
	m.init = newInit(m)
	m.configure = newConfigure(m)
	m.start = newStart(m)
	m.setState(m.init)
	return m
}

func (m *sm) setState(s State) {
	log.Debugf("%v set %s", m.st, s)
	m.st = s
	m.hist = append(m.hist, &hist{time.Now(), s.String()})
}

func (m *sm) Init(cf *config.Flags, fl *core.Flags) error {
	m.Config = config.New()
	m.ConfigFlags = cf
	m.CoreFlags = fl
	return nil
}

func (m *sm) Run(ctx context.Context) error {
	var err error
	rc := OK
	for rc == OK {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			rc = DONE
		default:
			n := m.st.String()
			log.Debugf("%s run", n)
			rc = m.st.Run(ctx)
			log.Debugf("%s status %s", n, stMap[rc])
		}
	}
	if rc == ERROR {
		err = m.st.Error()
	} else if rc == PANIC {
		log.Panic(m.st.Error())
	}
	return err
}
