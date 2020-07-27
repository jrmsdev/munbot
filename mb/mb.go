// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package mb implements main cmd util.
package mb

import (
	"flag"

	"github.com/munbot/master"
	"github.com/munbot/master/cmd"
	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
)

type Flags struct {
	Master *master.Flags
}

func (f *Flags) set(fs *flag.FlagSet) {
	f.Master.Set(fs)
}

type Cmd struct {
	flags *Flags
}

func New() *Cmd {
	return &Cmd{flags: &Flags{Master: master.NewFlags()}}
}

func (c *Cmd) FlagSet(fs *flag.FlagSet) {
	c.flags.set(fs)
}

func (c *Cmd) Command(flags *config.Flags) cmd.Command {
	return &Main{flags: c.flags, cf: flags}
}

type Main struct {
	flags *Flags
	cf    *config.Flags
}

func (m *Main) Run(args []string) int {
	log.Debugf("munbot version %s", master.Version())
	mbot := master.New()
	if err := mbot.Init(m.cf, m.flags.Master); err != nil {
		log.Error(err)
		return 10
	}
	if err := mbot.Run(); err != nil {
		log.Error(err)
		return 11
	}
	return 0
}
