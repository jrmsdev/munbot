// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"fmt"
	"strings"
)

func Update(c *Config, option, newval string) error {
	i := strings.Split(option, ".")
	ilen := len(i)
	opt := i[ilen-1]
	sect := checkSection(c, i[0:ilen-1])
	if sect == "" {
		for j := ilen - 2; j > 0; j-- {
			sect = checkSection(c, i[0:j])
			opt = fmt.Sprintf("%s.%s", i[j], opt)
			if sect != "" {
				break
			}
		}
	}
	return doUpdate(c, sect, opt, newval)
}

func checkSection(c *Config, args []string) string {
	n := strings.Join(args, ".")
	if c.HasSection(n) {
		return n
	}
	return ""
}

func doUpdate(c *Config, section, option, newval string) error {
	if !c.HasSection(section) {
		return fmt.Errorf("invalid section: %s", section)
	}
	if !c.HasOption(section, option) {
		return fmt.Errorf("%s section invalid option: %s", section, option)
	}
	c.db[section][option] = newval
	return nil
}
