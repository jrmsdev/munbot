// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/munbot/master/config/internal/parser"
)

type Parser struct {
	cfg *parser.Config
}

func NewParser(c *Config) *Parser {
	return &Parser{cfg: c.h}
}

// Map returns a map with section.option as keys with their respective values
// from the global parser object. The filter string can be empty "" or contain
// the prefix of the values to filter. In example, if filter is "master", only
// values from master section ("master.*") will be returned.
func (p *Parser) Map(filter string) map[string]string {
	return parser.Parse(p.cfg, filter)
}

// Update updates section.option on the global parser object with the new
// provided value. If section.option does not exists already, an error is
// returned.
func (p *Parser) Update(option, newval string) error {
	return parser.Update(p.cfg, option, newval)
}

// Set sets section.option with provided value. It's an error if the option
// already exists.
func (p *Parser) Set(option, val string) error {
	return parser.Set(p.cfg, option, val)
}

// Unset removes section.option.
func (p *Parser) Unset(option string) error {
	return parser.Unset(p.cfg, option)
}
