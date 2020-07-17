// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Map map[string]string

type DB map[string]Map

type Config struct {
	db DB
}

func New() *Config {
	return &Config{db: make(DB)}
}

func (c *Config) SetDefaults(src DB) {
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
	s := c.Section(section)
	return s.HasOption(option)
}

func (c *Config) HasSection(name string) bool {
	_, found := c.db[name]
	return found
}

func (c *Config) Section(name string) *Section {
	m, found := c.db[name]
	if !found {
		// TODO: debug log about missing section
		m = Map{}
	}
	return &Section{name, m, c}
}

func (c *Config) getSectOpt(option string) (string, string) {
	i := strings.Split(option, ".")
	ilen := len(i)
	opt := i[ilen-1]
	sect := c.checkSection(i[0 : ilen-1])
	if sect == "" {
		for j := ilen - 2; j > 0; j-- {
			sect = c.checkSection(i[0:j])
			opt = fmt.Sprintf("%s.%s", i[j], opt)
			if sect != "" {
				break
			}
		}
	}
	return sect, opt
}

func (c *Config) checkSection(args []string) string {
	n := strings.Join(args, ".")
	if c.HasSection(n) {
		return n
	}
	return ""
}

func (c *Config) expand(option string) string {
	sect, opt := c.getSectOpt(option)
	if sect == "" || !c.HasSection(sect) {
		return fmt.Sprintf("ECFGMISS:%s", option)
	}
	if opt == "" || !c.HasOption(sect, opt) {
		return fmt.Sprintf("ECFGMISS:%s", option)
	}
	return c.db[sect][opt]
}
