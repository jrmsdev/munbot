// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"fmt"
	"sort"
	"strings"
)

func Parse(c *Config, filter string) map[string]string {
	filter = strings.TrimSpace(filter)
	dst := make(map[string]string)
	for _, s := range listSections(c) {
		for _, k := range listOptions(c, s) {
			opt := fmt.Sprintf("%s.%s", s, k)
			if filter == "" || strings.HasPrefix(opt, filter) {
				dst[opt] = c.db[s][k]
			}
		}
	}
	return dst
}

func listSections(c *Config) []string {
	l := make([]string, 0, len(c.db))
	for n := range c.db {
		l = append(l, n)
	}
	sort.Strings(l)
	return l
}

func listOptions(c *Config, section string) []string {
	l := make([]string, 0, len(c.db[section]))
	for n := range c.db[section] {
		l = append(l, n)
	}
	sort.Strings(l)
	return l
}
