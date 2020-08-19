// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cmd

import (
	"flag"
	"os"

	"github.com/munbot/master/v0/config"
	"github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/log"
	"github.com/munbot/master/v0/version"
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
var osExit func(int)

func init() {
	flagsErrorHandler = flag.ExitOnError
	osExit = os.Exit
}

func New(name string, main Builder) *Main {
	return &Main{
		name:   name,
		main:   main,
		subcmd: make(map[string]Builder),
	}
}

func (m *Main) AddCommand(name string, b Builder) {
	// TODO: panic if command already exists?
	m.subcmd[name] = b
}

func (m *Main) Main(args []string) {
	var (
		action   string
		build    Builder
		cmdargs  []string
		progname string
	)
	if len(args) >= 1 {
		action = args[0]
	}
	if action == "" {
		action = "__NOACTION__"
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
	var showVersion bool
	fs := flag.NewFlagSet(progname, flagsErrorHandler)
	fs.BoolVar(&showVersion, "version", false, "show version info and exit")
	flags := config.NewFlags(fs)
	build.FlagSet(fs)
	fs.Parse(cmdargs)
	flags.Parse()
	if showVersion {
		version.Print(progname)
		osExit(0)
	}
	if flags.Debug {
		log.DebugFlags(env.Get("MB_LOG_DEBUG"))
	}
	log.SetMode(env.Get("MB_LOG"))
	log.SetColors(env.Get("MB_LOG_COLORS"))
	log.SetPrefix(env.Get("MUNBOT"))
	cmd := build.Command(flags)
	if cmd == nil {
		osExit(9)
	}
	rc := cmd.Run(fs.Args())
	osExit(rc)
}
