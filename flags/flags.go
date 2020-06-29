// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var (
	Debug bool = false
	ConfigDir string = filepath.FromSlash("~/.config/munbot")
	configDirErr error
	ConfigDistDir string = filepath.FromSlash("/etc/munbot")
	ConfigSysDir string = filepath.FromSlash("/usr/local/etc/munbot")
	MasterName string = "munbot"
)

var fs *flag.FlagSet

func init() {
	ConfigDir, configDirErr = os.UserConfigDir()
}

func Init(program string) {
	fs = flag.NewFlagSet(program, flag.ExitOnError)

	if configDirErr != nil {
		log.Panic(configDirErr)
	}
	ConfigDir = filepath.Join(ConfigDir, MasterName)

	fs.BoolVar(&Debug, "debug", Debug, "enable debug")

	fs.StringVar(&ConfigDir, "cfgdir", ConfigDir, "config dir")
	fs.StringVar(&ConfigDistDir, "cfgdistdir", ConfigDistDir, "dist config dir")
	fs.StringVar(&ConfigSysDir, "cfgsysdir", ConfigSysDir, "system config dir")
	fs.StringVar(&MasterName, "name", MasterName, "master robot name")
}

func Parse(args []string) {
	err := fs.Parse(args)
	if err != nil {
		log.Panic(err)
	}
}
