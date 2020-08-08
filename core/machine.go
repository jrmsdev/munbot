// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"github.com/munbot/master/config"
	"github.com/munbot/master/core/flags"
	"github.com/munbot/master/log"
)

type Machine interface {
	Abort() error
	SetState(StateID) error
	Config() *config.Config
	ConfigFlags() *config.Flags
	CoreFlags() *flags.Flags
}

func (k *Core) Abort() error {
	log.Debugf("[%s] abort...", k.State())
	log.Debug("state stop...")
	if err := k.state.Stop(); err != nil {
		return err
	}
	log.Debug("state halt...")
	if err := k.state.Halt(); err != nil {
		return err
	}
	return nil
}

func (k *Core) SetState(s StateID) error {
	log.Debugf("[%s] set state %s", k.State(), StateName(s))
	if s == k.stid {
		return k.errorf("core: state %s set twice", StateName(s))
	}
	switch s {
	case Init:
		k.state = k.sInit
	case Run:
		k.state = k.sRun
	case Halt:
		k.state = k.sHalt
	default:
		k.errorf("core: set %s", StateName(s))
	}
	k.stid = s
	k.rt.Master.CurrentState(k.State())
	return nil
}

func (k Core) Config() *config.Config {
	return k.cfg
}

func (k Core) ConfigFlags() *config.Flags {
	return k.cfl
}

func (k Core) CoreFlags() *flags.Flags {
	return k.kfl
}
