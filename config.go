// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

type Config struct {
	Name string `json:"name,omitempty"`
}

func newConfig() *Config {
	return &Config{Name: "munbot"}
}

func (c *Config) String() string {
	return c.Name
}

func (c *Config) Read(fh io.ReadCloser) error {
	defer fh.Close()
	blob, err := ioutil.ReadAll(fh)
	if err != nil {
		return log.Error(err)
	}
	if err := json.Unmarshal(blob, c); err != nil {
		return log.Error(err)
	}
	return nil
}

var fileOpen = os.Open

func Configure() *Config {
	dirs := []string{
		flags.ConfigDistDir,
		flags.ConfigSysDir,
		flags.ConfigDir,
	}
	log.Debugf("configure %s %v", flags.ConfigFile, dirs)
	cfg := newConfig()
	for _, dn := range dirs {
		fn := filepath.Join(dn, flags.ConfigFile)
		fh, err := fileOpen(fn)
		if err != nil {
			if os.IsNotExist(err) {
				log.Debug(err)
			} else {
				log.Panicf("%s: %s", fn, err)
			}
		} else {
			log.Debugf("read %s", fn)
			if err := cfg.Read(fh); err != nil {
				log.Panicf("%s: %s", fn, err)
			}
		}
	}
	return cfg
}
