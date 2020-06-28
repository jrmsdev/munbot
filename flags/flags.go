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
	ConfigDir string = filepath.FromSlash("~/.config/munbot")
	configDirErr error
	ConfigDistDir string = filepath.FromSlash("/etc/munbot")
	ConfigSysDir string = filepath.FromSlash("/usr/local/etc/munbot")
	MasterName string = "munbot"
)

func init() {
	ConfigDir, configDirErr = os.UserConfigDir()
}

func Init() {
	if configDirErr != nil {
		log.Panic(configDirErr)
	}
	ConfigDir = filepath.Join(ConfigDir, MasterName)

	flag.StringVar(&ConfigDir, "cfgdir", ConfigDir,
		"config dir")
	flag.StringVar(&ConfigDistDir, "cfgdistdir", ConfigDistDir,
		"dist config dir")
	flag.StringVar(&ConfigSysDir, "cfgsysdir", ConfigSysDir,
		"system config dir")
	flag.StringVar(&MasterName, "name", MasterName,
		"master robot name")
}

func Parse() {
	flag.Parse()
}
