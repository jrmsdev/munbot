// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"flag"

	"github.com/munbot/master/v0/env"
)

// Flags holds main flags settings.
type Flags struct {
	apiDisable     bool
	apiDebug       bool
	apiAddr        string
	apiPort        uint
	authDisable    bool
	consoleDisable bool
	consoleAddr    string
	consolePort    uint
}

func NewFlags() *Flags {
	return &Flags{}
}

// Set sets the flags to the provided handler.
func (f *Flags) Set(fs *flag.FlagSet) {
	fs.BoolVar(&f.apiDisable, "api.disable", false, "disable api server")
	fs.BoolVar(&f.apiDebug, "api.debug", false, "debug api")
	fs.StringVar(&f.apiAddr, "api.addr", "", "api tcp network `address`")
	fs.UintVar(&f.apiPort, "api.port", 0, "api tcp port `number`")
	fs.BoolVar(&f.authDisable, "auth.disable", false, "disable auth")
	fs.BoolVar(&f.consoleDisable, "console.disable", false, "disable console server")
	fs.StringVar(&f.consoleAddr, "console.addr", "", "console tcp network `address`")
	fs.UintVar(&f.consolePort, "console.port", 0, "console tcp port `number`")
}

// Parse parses the flags that were not set via the flags handler (cmd args
// usually) and sets them with their respective values from the configuration.
func (f *Flags) Parse() {
	f.parseApi()
	f.parseAuth()
	f.parseConsole()
}

func (f *Flags) parseApi() {
	if f.apiDisable {
		env.Set("MBAPI", "false")
	}
	if f.apiDebug {
		env.Set("MBAPI_DEBUG", "true")
	}
	if f.apiAddr != "" {
		env.Set("MBAPI_ADDR", f.apiAddr)
	}
	if f.apiPort != 0 {
		env.SetUint("MBAPI_PORT", f.apiPort)
	}
}

func (f *Flags) parseAuth() {
	if f.authDisable {
		env.Set("MBAUTH", "false")
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
