// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/jrmsdev/munbot/log"
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
	ConfigFile    string = "config.json"
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
	ConfigDir = filepath.Join(ConfigDir, "munbot")

	fs.BoolVar(&Debug, "debug", false, "enable debug")
	fs.BoolVar(&Version, "version", false, "show version info and exit")

	fs.StringVar(&ConfigDir, "cfgdir", ConfigDir,
		"config dir `path`")
	fs.StringVar(&ConfigDistDir, "cfgdistdir", ConfigDistDir,
		"dist config dir `path`")
	fs.StringVar(&ConfigSysDir, "cfgsysdir", ConfigSysDir,
		"system config dir `path`")
	fs.StringVar(&ConfigFile, "cfg", ConfigFile, "config file `name`")
}

func Parse(args []string) {
	err := fs.Parse(args)
	if err != nil {
		log.Panic(err)
	}
	if Version {
		showVersion()
	}
	if Debug {
		log.DebugEnable()
	}
}

func showVersion() {
	version.Print(progname)
	os.Exit(2)
}
