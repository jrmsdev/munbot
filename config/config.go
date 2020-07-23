// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package config handles the configuration files.
package config

import (
	"flag"
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
		"netaddr":   "127.0.0.1",
		"netport":   "6492",
	},
	"master": value.Map{
		"name": "munbot",
	},
	"master.api": value.Map{
		"enable": "true",
		"netaddr":   "0.0.0.0",
	},
}

var __handler *parser.Config

func init() {
	__handler = parser.New()
}

// Parse returns a map with section.option as keys with their respective values
// from the global parser object. The filter string can be empty "" or contain
// the prefix of the values to filter. In example, if filter is "master", only
// values from master section ("master.*") will be returned.
func Parse(filter string) map[string]string {
	return parser.Parse(__handler, filter)
}

// Update updates section.option on the global parser object with the new
// provided value. If section.option does not exists already, an error is
// returned.
func Update(option, newval string) error {
	return parser.Update(__handler, option, newval)
}

// Set sets section.option with provided value. It's an error if the option
// already exists.
func Set(option, val string) error {
	return parser.Set(__handler, option, val)
}

// SetOrUpdate sets config section.option with provided value or updates it if
// already exists.
func SetOrUpdate(option, val string) {
	parser.SetOrUpdate(__handler, option, val)
}

type dumpFunc func() ([]byte, error)

// Config is the main configuration manager.
type Config struct {
	h    *parser.Config
	dump dumpFunc
	flags *Flags
}

// New creates a new Config object with the global handler attached to it. So
// _ALL_ instances will work on the same data.
func New() *Config {
	return &Config{h: __handler, dump: __handler.Dump}
}

// Copy creates a new Config object with a copy of the source handler, so
// changes in the new copy object will not affect the global parser.
func (c *Config) Copy() *Config {
	h := c.h.Copy()
	// FIXME: should c.flags.Copy()
	return &Config{h: h, dump: h.Dump}
}

// SetDefaults set the values from the Defaults global variable. If a section
// already exists, it will be overriden.
func (c *Config) SetDefaults() {
	c.h.SetDefaults(Defaults)
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
	return c.h.Load(blob)
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

// FlagSet sets the configurable flags to the provided flags handler.
func (c *Config) FlagSet(fs *flag.FlagSet) {
	c.flags = nil
	c.flags = new(Flags)
	c.flags.set(fs)
}

// Flags parses the flags and returns a pointer to them. Flags that were not
// set via the flags handler (cmd args usually) are set with their respective
// values from the configuration. Only if the parser was previously set of
// course. Otherwise flags will have their (Go) default values.
func (c *Config) Flags() *Flags {
	if c.flags == nil {
		return new(Flags)
	}
	c.flags.parse(c)
	return c.flags
}
