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
	build  Builder
	subcmd map[string]Builder
	fs     *flag.FlagSet
	subfs  map[string]*flag.FlagSet
	flags  *config.Flags
}

var flagsErrorHandler flag.ErrorHandling

func init() {
	flagsErrorHandler = flag.ExitOnError
}

func New(name string, main Builder) *Main {
	fs := flag.NewFlagSet(name, flagsErrorHandler)
	flags := new(config.Flags)
	flags.Set(fs)
	main.FlagSet(fs)
	return &Main{
		name:   name,
		build:  main,
		subcmd: make(map[string]Builder),
		fs:     fs,
		subfs:  make(map[string]*flag.FlagSet),
		flags:  flags,
	}
}

func (m *Main) AddCommand(name string, b Builder) {
	// TODO: panic if command already exists?
	fs := flag.NewFlagSet(name, flagsErrorHandler)
	m.flags.Set(fs)
	b.FlagSet(fs)
	m.subfs[name] = fs
	m.subcmd[name] = b
}

func (m *Main) Main(args []string) {
	var build Builder
	var action string
	var fs *flag.FlagSet
	var cmdargs []string
	if len(args) >= 1 {
		action = args[0]
	}
	if b, ok := m.subcmd[action]; ok {
		build = b
		fs = m.subfs[action]
		cmdargs = args[1:]
	} else {
		build = m.build
		fs = m.fs
		cmdargs = args
	}
	fs.Parse(cmdargs)
	if err := m.flags.Parse(); err != nil {
		os.Exit(1)
	}
	cmd := build.Command(m.flags)
	rc := cmd.Run(cmdargs)
	os.Exit(rc)
}
