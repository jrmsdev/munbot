// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"io"
	"os"
	"path/filepath"

	"github.com/jrmsdev/munbot/config"
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

var fileOpen func(string) (*os.File, error) = os.Open

type Config struct {
	*config.Manager
	db map[string]config.Value
	reg map[string]*string
	Name string `json:"name,omitempty"`
	Test config.Value `json:"test,omitempty"`
	TestInt *config.IntValue `json:"testint,omitempty"`
	TestBool *config.BoolValue `json:"testbool,omitempty"`
}

func newConfig() *Config {
	c := config.New()
	return &Config{
		Manager: c,
		db: make(map[string]config.Value),
		reg: make(map[string]*string),
		Name: "munbot",
		Test: c.NewString("test", ""),
		TestInt: c.NewInt("testint", 0),
		TestBool: c.NewBool("testbool", false),
	}
}

func (c *Config) String() string {
	return c.Name
}

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

func (c *Config) Read(fh io.ReadCloser) error {
	defer fh.Close()
	return c.Manager.Read(c, fh)
}

func (c *Config) Write(fh io.WriteCloser) error {
	defer fh.Close()
	return c.Manager.Write(c, fh)
}
