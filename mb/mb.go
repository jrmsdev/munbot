// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package mb implements main cmd util.
package mb

import (
	"flag"

	"github.com/munbot/master"
	"github.com/munbot/master/cmd"
	"github.com/munbot/master/core"
	"github.com/munbot/master/config"
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
	return &Main{flags: c.flags, cf: flags}
}

type Main struct {
	flags *core.Flags
	cf    *config.Flags
}

func (m *Main) Run(args []string) int {
	log.Debugf("munbot version %s", master.Version())
	mbot := master.New()
	if err := mbot.Init(m.cf, m.flags); err != nil {
		return 10
	}
	if err := mbot.Run(); err != nil {
		return 11
	}
	return 0
}
