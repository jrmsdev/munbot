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
	Name          string = "master"
	ConfigDir     string = filepath.FromSlash("~/.config/munbot")
	configDirErr  error
	ConfigDistDir string = filepath.FromSlash("/etc/munbot")
	ConfigSysDir  string = filepath.FromSlash("/usr/local/etc/munbot")
	ConfigFile    string = "config.json"
	CacheDir      string = filepath.FromSlash("~/.cache/munbot")
	cacheDirErr   error
	DataDir       string = filepath.FromSlash("~/.munbot")
	dataDirErr    error
)

var fs *flag.FlagSet

func init() {
	ConfigDir, configDirErr = os.UserConfigDir()
	CacheDir, cacheDirErr = os.UserCacheDir()
	DataDir, dataDirErr = os.UserHomeDir()
}

func Init(program string) *flag.FlagSet {
	progname = program
	fs = flag.NewFlagSet(progname, flag.ExitOnError)

	if configDirErr != nil {
		log.Panic(configDirErr)
	}
	ConfigDir = filepath.Join(ConfigDir, "munbot")

	if cacheDirErr != nil {
		log.Panic(cacheDirErr)
	}
	CacheDir = filepath.Join(CacheDir, "munbot")

	if dataDirErr != nil {
		log.Panic(dataDirErr)
	}
	DataDir = filepath.Join(DataDir, ".munbot")

	fs.BoolVar(&Debug, "debug", false, "enable debug")
	fs.BoolVar(&Version, "version", false, "show version info and exit")
	fs.StringVar(&Name, "name", Name, "profile name")

	fs.StringVar(&ConfigDir, "cfg.dir", ConfigDir,
		"config dir `path`")
	fs.StringVar(&ConfigDistDir, "cfg.distdir", ConfigDistDir,
		"dist config dir `path`")
	fs.StringVar(&ConfigSysDir, "cfg.sysdir", ConfigSysDir,
		"system config dir `path`")
	fs.StringVar(&ConfigFile, "cfg", ConfigFile, "config file `name`")

	fs.StringVar(&CacheDir, "cache.dir", CacheDir, "cache dir `path`")
	fs.StringVar(&DataDir, "data.dir", DataDir, "data dir `path`")

	return fs
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
	ConfigDir = filepath.Clean(filepath.Join(ConfigDir, Name))
	CacheDir = filepath.Clean(filepath.Join(CacheDir, Name))
	DataDir = filepath.Clean(filepath.Join(DataDir, Name))
	// TODO: check and clean the other Config*Dirs and ConfigFile
}

func showVersion() {
	version.Print(progname)
	os.Exit(2)
}
