// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/jrmsdev/munbot/version"
)

var (
	progname      string = "munbot"
	Debug         bool   = false
	Version       bool   = false
	ConfigDir     string = filepath.FromSlash("~/.config/munbot")
	configDirErr  error
	ConfigDistDir string = filepath.FromSlash("/etc/munbot")
	ConfigSysDir  string = filepath.FromSlash("/usr/local/etc/munbot")
	MasterName    string = "munbot"
)

var fs *flag.FlagSet

func init() {
	ConfigDir, configDirErr = os.UserConfigDir()
}

func Init(program string) {
	progname = program
	fs = flag.NewFlagSet(progname, flag.ExitOnError)

	if configDirErr != nil {
		log.Panic(configDirErr)
	}
	ConfigDir = filepath.Join(ConfigDir, MasterName)

	fs.BoolVar(&Debug, "debug", false, "enable debug")
	fs.BoolVar(&Version, "version", false, "show version info and exit")

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
	if Version {
		showVersion()
	}
}

func showVersion() {
	version.Print(progname)
	os.Exit(2)
}
