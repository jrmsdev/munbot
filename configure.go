// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jrmsdev/munbot/config"
	"github.com/jrmsdev/munbot/config/flags"
	"github.com/jrmsdev/munbot/log"
)

var Config *config.Munbot
var fileOpen func(string) (*os.File, error)

func init() {
	Config = config.New()
	config.SetDefaults(Config)
	fileOpen = os.Open
}

func Configure(fs *flag.FlagSet) error {
	log.Debugf("configure %s ...", fs.Name())
	if err := tryDir(flags.ConfigDistDir); err != nil {
		return err
	}
	if err := tryDir(flags.ConfigSysDir); err != nil {
		return err
	}
	if err := tryDir(flags.ConfigDir); err != nil {
		return err
	}
	config.Flags(Config, fs)
	return setup()
}

func tryDir(dn string) error {
	fn := filepath.Join(dn, flags.ConfigFile)
	fh, err := fileOpen(fn)
	if err != nil {
		if os.IsNotExist(err) {
			log.Debug(err)
		} else {
			return fmt.Errorf("%s: %s", fn, err)
		}
	} else {
		if err := config.Read(Config, fh); err != nil {
			return fmt.Errorf("%s: %s", fn, err)
		}
		log.Printf("Config loaded %s", fn)
	}
	return nil
}
