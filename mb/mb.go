// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package mb implements main cmd util.
package mb

import (
	"context"
	"flag"
	"os"

	"github.com/munbot/master/cmd"
	"github.com/munbot/master/config"
	"github.com/munbot/master/internal/core"
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

func (c *Cmd) Command(cf *config.Flags) cmd.Command {
	return newMain(c.flags, cf)
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
	if len(args) > 0 {
		log.Errorf("invalid args: %v; check %s -help", args, os.Args[0])
		return 9
	}
	m.kf.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if _, err := m.rt.Init(ctx); err != nil {
		return 10
	}
	if err := m.rt.Configure(m.cf, m.cfg); err != nil {
		return 11
	}
	if err := m.rt.Start(); err != nil {
		return 12
	}
	return 0
}
