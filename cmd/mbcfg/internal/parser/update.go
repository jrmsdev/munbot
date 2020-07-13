// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"fmt"
	"strings"

	"github.com/munbot/master/cmd/mbcfg/internal/config"
)

// {"master":{"name":"testing"}}
// {"master.api":{"addr":"0.0.0.0"}}

func Update(c *config.Config, option, newval string) error {
	i := strings.Split(option, ".")
	ilen := len(i)
	opt := i[ilen-1]
	sect := checkSection(c, i[0:ilen-1])
	if sect == "" {
		for j := ilen-2; j > 0; j-- {
			sect = checkSection(c, i[0:j])
			opt = fmt.Sprintf("%s.%s", i[j], opt)
			if sect != "" {
				break
			}
		}
	}
	if !c.HasSection(sect) || !c.HasOption(sect, opt) {
		return fmt.Errorf("invalid option: %s", option)
	}
	s := fmt.Sprintf(`{"%s":{"%s":"%s"}}`, sect, opt, newval)
	return c.Load([]byte(s))
}

func checkSection(c *config.Config, args []string) string {
	n := strings.Join(args, ".")
	if c.HasSection(n) {
		return n
	}
	return ""
}
