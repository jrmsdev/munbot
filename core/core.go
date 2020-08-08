// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package core implements the runtime state machine.
package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/munbot/master/config"
	"github.com/munbot/master/core/flags"
	"github.com/munbot/master/log"
	"github.com/munbot/master/utils/uuid"
	"github.com/munbot/master/version"
)

var _ Runtime = &Core{}
var _ Machine = &Core{}

type Core struct {
	rt    *Mem
	ctx   context.Context
	uuid  string
	cfg   *config.Config
	cfl   *config.Flags
	kfl   *flags.Flags
	state State
	stid  StateID
	sInit State
	sRun  State
	sHalt State
}

var mem *Mem

func init() {
	mem = newMem()
}

func New(m *Mem) *Core {
	log.Printf("Munbot version %s", version.String())
	k := &Core{rt: m, uuid: uuid.Rand()}
	k.sInit = newInit(k, k.rt)
	k.sRun = newRun(k, k.rt)
	k.sHalt = newHalt(k, k.rt)
	// init state
	k.state = k.sInit
	k.stid = Init
	return k
}

func NewRuntime() Runtime {
	return New(mem)
}

func (k *Core) String() string {
	return "Core:" + k.uuid
}

func (k *Core) UUID() string {
	return k.uuid
}

func (k *Core) State() string {
	return StateName(k.stid)
}

func (k *Core) StateID() StateID {
	return k.stid
}

func (k *Core) error(err error) error {
	log.Output(2, fmt.Sprintf("[ERROR] core: %s", err))
	return err
}

func (k *Core) errorf(f string, args ...interface{}) error {
	err := errors.New(fmt.Sprintf(f, args...))
	log.Output(2, fmt.Sprintf("[ERROR] core %s: %s", StateName(k.stid), err))
	return err
}
