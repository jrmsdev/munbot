// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package mb implements main cmd util.
package mb

import (
	"flag"
	//~ "fmt"
	//~ "sort"

	"github.com/munbot/master/cmd"
	"github.com/munbot/master/config"
	//~ "github.com/munbot/master/config/profile"
	//~ "github.com/munbot/master/log"
)

type Flags struct {
}

func (f *Flags) set(fs *flag.FlagSet) {
	//~ fs.BoolVar(&f.ListAll, "a", false, "list all options")
	//~ fs.BoolVar(&f.Set, "set", false, "set option instead of updating it")
	//~ fs.BoolVar(&f.Unset, "unset", false, "unset option from configuration file")
}

type Cmd struct {
	flags *Flags
}

func New() *Cmd {
	return &Cmd{flags: &Flags{}}
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
	return 0
}
