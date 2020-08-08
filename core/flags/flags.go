// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"

	"github.com/munbot/master/config"
)

// Flags holds main flags settings.
type Flags struct {
	ApiEnable      bool
	apiDisable     bool
	ApiDebug       bool
	ApiAddr        string
	ApiPort        uint
	ConsoleEnable  bool
	consoleDisable bool
	ConsoleAddr    string
	ConsolePort    uint
}

func New() *Flags {
	return &Flags{
		ApiEnable:     true,
		ConsoleEnable: true,
	}
}

// Set sets the flags to the provided handler.
func (f *Flags) Set(fs *flag.FlagSet) {
	fs.BoolVar(&f.apiDisable, "api.disable", false, "disable api")
	fs.BoolVar(&f.ApiDebug, "api.debug", false, "debug api")
	fs.StringVar(&f.ApiAddr, "api.addr", "", "api tcp `address` to bind to")
	fs.UintVar(&f.ApiPort, "api.port", 0, "api tcp `port` to bind to")
	fs.BoolVar(&f.consoleDisable, "console.disable", false, "disable console")
	fs.StringVar(&f.ConsoleAddr, "console.addr", "", "console tcp `address` to bind to")
	fs.UintVar(&f.ConsolePort, "console.port", 0, "console tcp `port` to bind to")
}

// Parse parses the flags that were not set via the flags handler (cmd args
// usually) and sets them with their respective values from the configuration.
func (f *Flags) Parse(c *config.Config) {
	f.parseApi(c.Section("master.api"))
	f.parseConsole(c.Section("master.console"))
}

func (f *Flags) parseApi(s *config.Section) {
	f.ApiEnable = s.GetBool("enable")
	if f.apiDisable {
		f.ApiEnable = false
	}
	if f.ApiAddr == "" {
		f.ApiAddr = s.Get("netaddr")
	}
	if f.ApiPort == 0 {
		f.ApiPort = s.GetUint("netport")
	}
}

func (f *Flags) parseConsole(s *config.Section) {
	f.ConsoleEnable = s.GetBool("enable")
	if f.consoleDisable {
		f.ConsoleEnable = false
	}
	if f.ConsoleAddr == "" {
		f.ConsoleAddr = s.Get("netaddr")
	}
	if f.ConsolePort == 0 {
		f.ConsolePort = s.GetUint("netport")
	}
}
