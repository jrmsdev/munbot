// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"flag"

	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
)

var Config *config.Munbot

func init() {
	Config = config.New()
	config.SetDefaults(Config)
}

func Configure(fs *flag.FlagSet) error {
	if err := Setup(); err != nil {
		return log.Error(err)
	}
	log.Debugf("configure %s ...", fs.Name())
	if err := config.ReadFiles(Config); err != nil {
		return err
	}
	config.Flags(Config, fs)
	return nil
}
