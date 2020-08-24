// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"flag"

	"github.com/munbot/master/v0/env"
)

// Flags holds main flags settings.
type Flags struct {
	apiDisable  bool
	apiDebug    bool
	apiNet      string
	apiAddr     string
	apiPort     uint
	authDisable bool
	sshdDisable bool
	sshdAddr    string
	sshdPort    uint
}

func NewFlags() *Flags {
	return &Flags{}
}

// Set sets the flags to the provided handler.
func (f *Flags) Set(fs *flag.FlagSet) {
	fs.BoolVar(&f.apiDisable, "api.disable", false, "disable api server")
	fs.BoolVar(&f.apiDebug, "api.debug", false, "debug api")
	fs.StringVar(&f.apiNet, "api.net", "", "api network `name`")
	fs.StringVar(&f.apiAddr, "api.addr", "", "api network `address`")
	fs.UintVar(&f.apiPort, "api.port", 0, "api tcp port `number`")
	fs.BoolVar(&f.authDisable, "auth.disable", false, "disable auth")
	fs.BoolVar(&f.sshdDisable, "sshd.disable", false, "disable ssh server")
	fs.StringVar(&f.sshdAddr, "sshd.addr", "", "ssh server tcp network `address`")
	fs.UintVar(&f.sshdPort, "sshd.port", 0, "ssh server tcp port `number`")
}

// Parse parses the flags that were not set via the flags handler (cmd args
// usually) and sets them with their respective values from the configuration.
func (f *Flags) Parse() {
	f.parseApi()
	f.parseAuth()
	f.parseSSHD()
}

func (f *Flags) parseApi() {
	if f.apiDisable {
		env.Set("MBAPI", "false")
	}
	if f.apiDebug {
		env.Set("MBAPI_DEBUG", "true")
	}
	if f.apiNet != "" {
		env.Set("MBAPI_NET", f.apiNet)
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

func (f *Flags) parseSSHD() {
	if f.sshdDisable {
		env.Set("MBSSHD", "false")
	}
	if f.sshdAddr != "" {
		env.Set("MBSSHD_ADDR", f.sshdAddr)
	}
	if f.sshdPort != 0 {
		env.SetUint("MBSSHD_PORT", f.sshdPort)
	}
}
