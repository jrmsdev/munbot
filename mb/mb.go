// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package mb implements main cmd util.
package mb

import (
	"context"
	"flag"

	"github.com/munbot/master"
	"github.com/munbot/master/cmd"
	"github.com/munbot/master/config"
	"github.com/munbot/master/core"
	"github.com/munbot/master/log"
)

type Cmd struct {
	flags *core.Flags
}

func New() *Cmd {
	return &Cmd{flags: core.NewFlags()}
}

func (c *Cmd) FlagSet(fs *flag.FlagSet) {
	c.flags.Set(fs)
}

func (c *Cmd) Command(flags *config.Flags) cmd.Command {
	return newMain(c.flags, flags)
}

type Main struct {
	kf  *core.Flags
	cf  *config.Flags
	rt  core.Runtime
	cfg *config.Config
}

func newMain(kf *core.Flags, cf *config.Flags) *Main {
	return &Main{
		kf:  kf,
		cf:  cf,
		rt:  core.NewRuntime(),
		cfg: config.New(),
	}
}

func (m *Main) Run(args []string) int {
	log.Debugf("munbot version %s", master.Version())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mbot := master.NewMaster(m.rt)
	if _, err := mbot.Init(ctx); err != nil {
		return 10
	}
	if err := mbot.Configure(m.kf, m.cf, m.cfg); err != nil {
		return 11
	}
	if err := mbot.Start(); err != nil {
		return 12
	}
	return 0
}
