// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"flag"

	"github.com/munbot/master"
	"github.com/munbot/master/cmd"
)

type Cmd struct {
}

func New() *Cmd {
	return &Cmd{}
}

func (c *Cmd) FlagSet(fs *flag.FlagSet) {
}

func (c *Cmd) Command(flags *master.Flags) cmd.Command {
	return &Main{}
}

type Main struct {
}

func (m *Main) Run(args []string) error {
	return nil
}
