// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
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

func Configure() error {
	log.Debug("configure...")
	if err := tryDir(flags.ConfigDistDir); err != nil {
		return err
	}
	if err := tryDir(flags.ConfigSysDir); err != nil {
		return err
	}
	if err := tryDir(flags.ConfigDir); err != nil {
		return err
	}
	return setup()
}

func tryDir(dn string) error {
	fn := filepath.Join(dn, flags.ConfigFile)
	log.Debugf("try config file %s", fn)
	fh, err := fileOpen(fn)
	if err != nil {
		if os.IsNotExist(err) {
			log.Debug(err)
		} else {
			return fmt.Errorf("%s: %s", fn, err)
		}
	} else {
		log.Debugf("read %s", fn)
		if err := config.Read(Config, fh); err != nil {
			return fmt.Errorf("%s: %s", fn, err)
		}
		log.Printf("Config loaded %s", fn)
	}
	return nil
}
