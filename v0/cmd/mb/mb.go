// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package mb implements main cmd util.
package mb

import (
	"flag"
	"os"

	"github.com/munbot/master/v0"
	"github.com/munbot/master/v0/cmd"
	"github.com/munbot/master/v0/config"
	"github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/log"
)

type Cmd struct {
	flags *master.Flags
}

func New() *Cmd {
	return &Cmd{flags: master.NewFlags()}
}

func (c *Cmd) FlagSet(fs *flag.FlagSet) {
	c.flags.Set(fs)
}

func (c *Cmd) Command(cf *config.Flags) cmd.Command {
	return newMain(c.flags, cf)
}

type Main struct {
	fl  *master.Flags
	cf  *config.Flags
}

func newMain(fl *master.Flags, cf *config.Flags) *Main {
	if cf.Debug {
		log.DebugFlags(env.Get("MB_LOG_DEBUG"))
	}
	log.SetMode(env.Get("MB_LOG"))
	log.SetColors(env.Get("MB_LOG_COLORS"))
	log.SetPrefix(env.Get("MUNBOT"))
	return &Main{
		fl:  fl,
		cf:  cf,
	}
}

func (m *Main) Run(args []string) int {
	if len(args) > 0 {
		log.Errorf("invalid args: %v; check %s -help", args, os.Args[0])
		return 9
	}
	m.fl.Parse()
	return 0
}
