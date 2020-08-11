// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"flag"

	"github.com/munbot/master/env"
)

// Flags holds main flags settings.
type Flags struct {
	authDisable    bool
	apiDisable     bool
	apiDebug       bool
	apiAddr        string
	apiPort        uint
	consoleDisable bool
	consoleAddr    string
	consolePort    uint
}

func NewFlags() *Flags {
	return &Flags{}
}

// Set sets the flags to the provided handler.
func (f *Flags) Set(fs *flag.FlagSet) {
	fs.BoolVar(&f.authDisable, "auth.disable", false, "disable auth")
	fs.BoolVar(&f.apiDisable, "api.disable", false, "disable api")
	fs.BoolVar(&f.apiDebug, "api.debug", false, "debug api")
	fs.StringVar(&f.apiAddr, "api.addr", "", "api tcp `address` to bind to")
	fs.UintVar(&f.apiPort, "api.port", 0, "api tcp `port` to bind to")
	fs.BoolVar(&f.consoleDisable, "console.disable", false, "disable console")
	fs.StringVar(&f.consoleAddr, "console.addr", "", "console tcp `address` to bind to")
	fs.UintVar(&f.consolePort, "console.port", 0, "console tcp `port` to bind to")
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
	if f.apiAddr != "" {
		env.Set("MBAPI_ADDR", f.apiAddr)
	}
	if f.apiPort != 0 {
		env.SetUint("MBAPI_PORT", f.apiPort)
	}
}

func (f *Flags) parseConsole() {
	if f.consoleDisable {
		env.Set("MBCONSOLE", "false")
	}
	if f.consoleAddr != "" {
		env.Set("MBCONSOLE_ADDR", f.consoleAddr)
	}
	if f.consolePort != 0 {
		env.SetUint("MBCONSOLE_PORT", f.consolePort)
	}
}
