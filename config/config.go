// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package config handles the configuration files.
package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/munbot/master/config/internal/parser"
	"github.com/munbot/master/config/profile"
	"github.com/munbot/master/config/value"
	"github.com/munbot/master/log"
	"github.com/munbot/master/vfs"
)

// Defaults contains some default values.
var Defaults value.DB = value.DB{}

var __handler *parser.Config

func init() {
	__handler = parser.New()
}

type dumpFunc func() ([]byte, error)

// Config is the main configuration manager.
type Config struct {
	h    *parser.Config
	dump dumpFunc
}

// New creates a new Config object with the global handler attached to it. So
// _ALL_ instances will work on the same data.
func New() *Config {
	return &Config{h: __handler, dump: __handler.Dump}
}

// Copy creates a new Config object with a copy of the data handler, so
// we can detach from the global parser.
func (c *Config) Copy() *Config {
	h := c.h.Copy()
	return &Config{h: h, dump: h.Dump}
}

// SetDefaults set the values from the Defaults global variable. If a section
// already exists, it will be overriden.
func (c *Config) SetDefaults(v value.DB) {
	c.h.SetDefaults(v)
}

// Load reads the configuration files from the provided profile.
func (c *Config) Load() error {
	p := profile.New()
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

// Read reads config content from reader.
func (c *Config) Read(r io.Reader) error {
	blob, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return c.h.Load(blob)
}

// Save writes configuration to the provided profile.
func (c *Config) Save() error {
	p := profile.New()
	fn := p.GetConfigFile()
	fh, err := vfs.Create(fn)
	if err != nil {
		return err
	}
	defer fh.Close()
	return c.Write(fh)
}

// Write writes config content to writer.
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

// HasOption checks if option exists in section.
func (c *Config) HasOption(section, option string) bool {
	return c.h.HasOption(section, option)
}

// HasSection checks if the section exists in the global parser.
func (c *Config) HasSection(name string) bool {
	return c.h.HasSection(name)
}

// Section creates a new Section object with its named section data attached to
// it. If the section name does not exists, "default" is used.
func (c *Config) Section(name string) *Section {
	if !c.h.HasSection(name) {
		// TODO: debug log about missing section?
		name = "default"
	}
	return &Section{name, c.h}
}
