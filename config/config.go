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
	"github.com/munbot/master/config/value"
	"github.com/munbot/master/log"
	"github.com/munbot/master/profile"
	"github.com/munbot/master/vfs"
)

// Defaults contains the values that will be used by Config.SetDefaults().
var Defaults value.DB = value.DB{
	"default": value.Map{
		"enable": "false",
		"name":   "munbot",
	},
	"master": value.Map{
		"name": "${default.name}",
	},
	"master.api": value.Map{
		"enable": "true",
		"addr":   "0.0.0.0",
		"port":   "6492",
	},
}

var handler *parser.Config

func init() {
	handler = parser.New()
}

// Parse returns a map with section.option as keys with their respective values
// from the global parser object. The filter string can be empty "" or contain
// the prefix of the values to filter. In example, if filter is "master", only
// values from master section ("master.*") will be returned.
func Parse(filter string) map[string]string {
	return parser.Parse(handler, filter)
}

// Update updates section.option on the global parser object with the new
// provided value. If section.option does not exists already, an error is
// returned.
func Update(option, newval string) error {
	return parser.Update(handler, option, newval)
}

// Set sets section.option with provided value. It's an error if the option
// already exists.
func Set(option, val string) error {
	return parser.Set(handler, option, val)
}

// SetOrUpdate sets config section.option with provided value or updates it if
// already exists.
func SetOrUpdate(option, val string) {
	parser.SetOrUpdate(handler, option, val)
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
	return &Config{h: handler, dump: handler.Dump}
}

// Copy creates a new Config object with a copy of the source handler, so
// changes in the new copy object will not affect the global parser.
func Copy() *Config {
	h := handler.Copy()
	return &Config{h: h, dump: h.Dump}
}

// SetDefaults set the values from the Defaults global variable. If a section
// already exists, it will be overriden.
func (c *Config) SetDefaults() {
	handler.SetDefaults(Defaults)
}

// Load reads the configuration files from the provided profile.
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

// Read reads config content from reader.
func (c *Config) Read(r io.Reader) error {
	blob, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return handler.Load(blob)
}

// Save writes configuration to the provided profile.
func (c *Config) Save(p *profile.Profile) error {
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
	return handler.HasOption(section, option)
}

// HasSection checks if the section exists in the global parser.
func (c *Config) HasSection(name string) bool {
	return handler.HasSection(name)
}

// Section creates a new Section object with its named section data attached to
// it. If the section name does not exists, "default" is used.
func (c *Config) Section(name string) *Section {
	if !handler.HasSection(name) {
		// TODO: debug log about missing section?
		name = "default"
	}
	return &Section{name, handler}
}
