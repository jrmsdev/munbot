// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func SetDefaults(c *Config) {
	c.db["master"] = Map{
		"enable": "true",
		"name": "munbot",
	}
	c.db["master.api"] = Map{
		"enable": "true",
		"addr": "0.0.0.0",
		"port": "6492",
	}
}

func Update(c *Config, section, option, newval string) error {
	if !c.HasSection(section) {
		return fmt.Errorf("invalid section: %s", section)
	}
	if !c.HasOption(section, option) {
		return fmt.Errorf("%s section invalid option: %s", section, option)
	}
	c.db[section][option] = newval
	return nil
}

type Map map[string]string

type DB map[string]Map

type Config struct {
	db DB
}

func New() *Config {
	return &Config{db: make(DB)}
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
	return &Section{name, m}
}

type Section struct {
	name string
	m Map
}

func (s *Section) Name() string {
	return s.name
}

func (s *Section) HasOption(name string) bool {
	_, found := s.m[name]
	return found
}

func (s *Section) Get(name string) string {
	v, found := s.m[name]
	if !found {
		// TODO: debug log about missing option
		return ""
	}
	return v
}

func (s *Section) GetBool(name string) bool {
	r, err := strconv.ParseBool(s.Get(name))
	if err != nil {
		// TODO: log about parsing error
		return false
	}
	return r
}
