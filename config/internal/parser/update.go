// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"fmt"
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
