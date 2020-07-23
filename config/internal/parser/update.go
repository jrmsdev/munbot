// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"fmt"

	"github.com/munbot/master/config/value"
)

func Update(c *Config, option, newval string) error {
	sect, opt := c.getSectOpt(option)
	if !c.HasSection(sect) {
		return fmt.Errorf("invalid section: %s", sect)
	}
	if !c.HasOption(sect, opt) {
		return fmt.Errorf("%s section invalid option: %s", sect, opt)
	}
	c.db[sect][opt] = newval
	return nil
}

func Set(c *Config, option, val string) error {
	sect, opt := c.getSectOpt(option)
	if !c.HasSection(sect) {
		c.db[sect] = value.Map{}
	} else if c.HasOption(sect, opt) {
		return fmt.Errorf("%s.%s option already exists", sect, opt)
	}
	c.db[sect][opt] = val
	return nil
}

func SetOrUpdate(c *Config, option, val string) {
	sect, opt := c.getSectOpt(option)
	if !c.HasOption(sect, opt) {
		Set(c, option, val)
	}
	Update(c, option, val)
}
