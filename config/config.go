// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/munbot/master/config/parser"
	"github.com/munbot/master/log"
	"github.com/munbot/master/profile"
	"github.com/munbot/master/vfs"
)

var Defaults parser.DB = parser.DB{
	"master": parser.Map{
		"name": "munbot",
	},
	"master.api": parser.Map{
		"enable": "true",
		"addr":   "0.0.0.0",
		"port":   "6492",
	},
}

var handler *parser.Config

func init() {
	handler = parser.New()
}

type dumpFunc func() ([]byte, error)

type Config struct {
	dump dumpFunc
}

func New() *Config {
	return &Config{dump: handler.Dump}
}

func (c *Config) SetDefaults() {
	handler.SetDefaults(Defaults)
}

func (c *Config) Load(p *profile.Profile) error {
	for _, fn := range p.ListConfigFiles() {
		if err := c.readFile(fn); err != nil {
			return fmt.Errorf("%s: %s", fn, err)
		}
	}
	return nil
}

func (c *Config) readFile(name string) error {
	fh, err := vfs.Open(name)
	if err != nil {
		if os.IsNotExist(err) {
			log.Debug(err)
			return nil
		} else {
			return err
		}
	}
	defer fh.Close()
	return c.Read(fh)
}

func (c *Config) Read(r io.Reader) error {
	blob, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return handler.Load(blob)
}

func (c *Config) Save(p *profile.Profile) error {
	fn := p.GetConfigFile()
	fh, err := vfs.Create(fn)
	if err != nil {
		return err
	}
	defer fh.Close()
	return c.Write(fh)
}

func (c *Config) Write(w io.Writer) error {
	blob, err := c.dump()
	if err != nil {
		return err
	}
	if _, err := w.Write(blob); err != nil {
		return err
	}
	return nil
}

func (c *Config) HasOption(section, option string) bool {
	return handler.HasOption(section, option)
}

func (c *Config) HasSection(name string) bool {
	return handler.HasSection(name)
}

func (c *Config) Section(name string) *Section {
	if !handler.HasSection(name) {
		// TODO: debug log about missing section, maybe panic?
		name = fmt.Sprintf("ECFGSECT:%s", name)
	}
	return &Section{name, handler}
}
