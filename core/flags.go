// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"flag"

	"github.com/munbot/master/env"
)

// Flags holds main flags settings.
type Flags struct {
	AuthEnable     bool
	authDisable    bool
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

func NewFlags() *Flags {
	return &Flags{
		AuthEnable:    true,
		ApiEnable:     true,
		ConsoleEnable: true,
	}
}

// Set sets the flags to the provided handler.
func (f *Flags) Set(fs *flag.FlagSet) {
	fs.BoolVar(&f.authDisable, "auth.disable", false, "disable auth")
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
func (f *Flags) Parse() {
	f.parseAuth()
	f.parseApi()
	f.parseConsole()
}

func (f *Flags) parseAuth() {
	if f.authDisable {
		env.Set("MBAUTH", "false")
	}
}

func (f *Flags) parseApi() {
	if f.apiDisable {
		env.Set("MBAPI", "false")
	}
	if f.ApiAddr != "" {
		env.Set("MBAPI_ADDR", f.ApiAddr)
	}
	if f.ApiPort != 0 {
		env.SetUint("MBAPI_PORT", f.ApiPort)
	}
}

func (f *Flags) parseConsole() {
	if f.consoleDisable {
		env.Set("MBCONSOLE", "false")
	}
	if f.ConsoleAddr != "" {
		env.Set("MBCONSOLE_ADDR", f.ConsoleAddr)
	}
	if f.ConsolePort != 0 {
		env.SetUint("MBCONSOLE_PORT", f.ConsolePort)
	}
}
