// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/jrmsdev/munbot/log"
	"github.com/jrmsdev/munbot/config/flags"
	"github.com/jrmsdev/munbot/version"
)

var (
	progname      string = "munbot"
	showVersion   bool   = false
	configDirErr  error
	cacheDirErr   error
	dataDirErr    error
)

var fs *flag.FlagSet

func init() {
	flags.ConfigDir, configDirErr = os.UserConfigDir()
	flags.CacheDir, cacheDirErr = os.UserCacheDir()
	flags.DataDir, dataDirErr = os.UserHomeDir()
}

func Init(program string) *flag.FlagSet {
	progname = program
	fs = flag.NewFlagSet(progname, flag.ExitOnError)

	if configDirErr != nil {
		log.Panic(configDirErr)
	}
	flags.ConfigDir = filepath.Join(flags.ConfigDir, "munbot")

	if cacheDirErr != nil {
		log.Panic(cacheDirErr)
	}
	flags.CacheDir = filepath.Join(flags.CacheDir, "munbot")

	if dataDirErr != nil {
		log.Panic(dataDirErr)
	}
	flags.DataDir = filepath.Join(flags.DataDir, ".munbot")

	fs.BoolVar(&flags.Debug, "debug", false, "enable debug")

	fs.BoolVar(&showVersion, "version", false, "show version info and exit")
	fs.StringVar(&flags.Name, "name", flags.Name, "master robot name")

	fs.StringVar(&flags.ConfigDir, "cfg.dir", flags.ConfigDir,
		"config dir `path`")
	fs.StringVar(&flags.ConfigDistDir, "cfg.distdir", flags.ConfigDistDir,
		"dist config dir `path`")
	fs.StringVar(&flags.ConfigSysDir, "cfg.sysdir", flags.ConfigSysDir,
		"system config dir `path`")
	fs.StringVar(&flags.ConfigFile, "cfg", flags.ConfigFile, "config file `name`")

	fs.StringVar(&flags.CacheDir, "cache.dir", flags.CacheDir, "cache dir `path`")
	fs.StringVar(&flags.DataDir, "data.dir", flags.DataDir, "data dir `path`")

	return fs
}

func Parse(args []string) {
	err := fs.Parse(args)
	if err != nil {
		log.Panic(err)
	}
	if showVersion {
		printVersion()
	}
	if flags.Debug {
		log.DebugEnable()
	}
	log.SetPrefix(flags.Name)
	flags.ConfigDir = filepath.Clean(filepath.Join(flags.ConfigDir, flags.Name))
	flags.CacheDir = filepath.Clean(filepath.Join(flags.CacheDir, flags.Name))
	flags.DataDir = filepath.Clean(filepath.Join(flags.DataDir, flags.Name))
	// TODO: check and clean the other Config*Dirs and ConfigFile
}

func printVersion() {
	version.Print(progname)
	os.Exit(2)
}
