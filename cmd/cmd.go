// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cmd

import (
	"flag"
	"os"

	"github.com/munbot/master/config"
)

type Command interface {
	Run(args []string) int
}

type Builder interface {
	FlagSet(fs *flag.FlagSet)
	Command(flags *config.Flags) Command
}

type Main struct {
	name   string
	main   Builder
	subcmd map[string]Builder
}

var flagsErrorHandler flag.ErrorHandling

func init() {
	flagsErrorHandler = flag.ExitOnError
}

func New(name string, main Builder) *Main {
	return &Main{
		name:   name,
		main:  main,
		subcmd: make(map[string]Builder),
	}
}

func (m *Main) AddCommand(name string, b Builder) {
	// TODO: panic if command already exists?
	m.subcmd[name] = b
}

func (m *Main) Main(args []string) {
	var (
		action string
		build Builder
		cmdargs []string
		progname string
	)
	if len(args) >= 1 {
		action = args[0]
	}
	if b, ok := m.subcmd[action]; ok {
		build = b
		cmdargs = args[1:]
		progname = m.name + "-" + action
	} else {
		build = m.main
		cmdargs = args
		progname = m.name
	}
	fs := flag.NewFlagSet(progname, flagsErrorHandler)
	flags := new(config.Flags)
	flags.Set(fs)
	build.FlagSet(fs)
	fs.Parse(cmdargs)
	if err := flags.Parse(); err != nil {
		os.Exit(1)
	}
	cmd := build.Command(flags)
	rc := cmd.Run(fs.Args())
	os.Exit(rc)
}
