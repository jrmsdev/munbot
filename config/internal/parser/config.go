// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/munbot/master/config/value"
)

type Config struct {
	db value.DB
}

func New() *Config {
	return &Config{db: make(value.DB)}
}

func (c *Config) Copy() *Config {
	db := value.DB{}
	for s, l := range c.db {
		db[s] = value.Map{}
		for k, v := range l {
			db[s][k] = v
		}
	}
	return &Config{db: db}
}

func (c *Config) SetDefaults(src value.DB) {
	for k, v := range src {
		c.db[k] = v
	}
}

func (c *Config) Dump() ([]byte, error) {
	return json.Marshal(c.db)
}

func (c *Config) Load(b []byte) error {
	return json.Unmarshal(b, &c.db)
}

func (c *Config) HasOption(section, option string) bool {
	if !c.HasSection(section) {
		return false
	}
	_, found := c.db[section][option]
	return found
}

func (c *Config) HasSection(name string) bool {
	_, found := c.db[name]
	return found
}

func (c *Config) Get(sect, opt string) string {
	if sect == "" || !c.HasSection(sect) {
		return fmt.Sprintf("ECFGMISS:%s.%s", sect, opt)
	}
	if opt == "" || !c.HasOption(sect, opt) {
		return fmt.Sprintf("ECFGMISS:%s.%s", sect, opt)
	}
	return c.eval(c.db[sect][opt])
}

func (c *Config) eval(value string) string {
	return os.Expand(value, c.expand())
}

func (c *Config) expand() func(string) string {
	var loop func(value string) string
	done := make(map[string]bool)
	fn := func(option string) string {
		if done[option] {
			return fmt.Sprintf("ECFGLOOP:%s", option)
		}
		done[option] = true
		sect, opt := c.getSectOpt(option)
		if sect == "" || !c.HasSection(sect) {
			return fmt.Sprintf("ECFGMISS:%s", option)
		}
		if opt == "" || !c.HasOption(sect, opt) {
			return fmt.Sprintf("ECFGMISS:%s", option)
		}
		return loop(c.db[sect][opt])
	}
	loop = func(value string) string {
		return os.Expand(value, fn)
	}
	return fn
}

func (c *Config) getSectOpt(option string) (string, string) {
	i := strings.Split(option, ".")
	ilen := len(i)
	sect := option
	opt := ""
	if ilen > 1 {
		sect = strings.Join(i[0:ilen-1], ".")
		opt = i[ilen-1]
	}
	return sect, opt
}
