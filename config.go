// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jrmsdev/munbot/config2"
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

var Config *config2.Munbot
var fileOpen func(string) (*os.File, error)

func init() {
	Config = config2.New()
	config2.SetDefaults(Config)
	fileOpen = os.Open
}

func Configure() error {
	dirs := []string{
		flags.ConfigDistDir,
		flags.ConfigSysDir,
		flags.ConfigDir,
	}
	log.Debugf("configure %s %v", flags.ConfigFile, dirs)
	for _, dn := range dirs {
		fn := filepath.Join(dn, flags.ConfigFile)
		fh, err := fileOpen(fn)
		if err != nil {
			if os.IsNotExist(err) {
				log.Debug(err)
			} else {
				return fmt.Errorf("%s: %s", fn, err)
			}
		} else {
			log.Debugf("read %s", fn)
			if err := config2.Read(Config, fh); err != nil {
				return fmt.Errorf("%s: %s", fn, err)
			}
		}
	}
	return setup()
}
