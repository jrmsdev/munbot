// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/jrmsdev/munbot/config"
	"github.com/jrmsdev/munbot/log"
)

func dump(cfg *config.Munbot, out io.Writer, listAll, jsonFormat bool, filter string) error {
	log.Debugf("dump all=%v json=%v", listAll, jsonFormat)
	if listAll {
		if jsonFormat {
			return jsonDump(cfg, out)
		}
		return parseDump(cfg, out)
	} else if filter != "" {
		if jsonFormat {
			return jsonFilter(cfg, out, filter)
		}
		return parseFilter(cfg, out, filter)
	}
	def := config.New()
	config.SetDefaults(def)
	if jsonFormat {
		return jsonDiff(def, cfg, out)
	}
	return parseDiff(def, cfg, out)
}

func list(m map[string]string) []string {
	l := make([]string, 0)
	for k := range m {
		l = append(l, k)
	}
	sort.Strings(l)
	return l
}

func jsonDump(cfg *config.Munbot, out io.Writer) error {
	return config.Write(cfg, out)
}

func jsonFilter(cfg *config.Munbot, out io.Writer, filter string) error {
	// FIXME: implement jsonFilter
	return nil
}

func jsonDiff(def *config.Munbot, cfg *config.Munbot, out io.Writer) error {
	// FIXME: implement jsonDiff
	return nil
}

func parseDump(cfg *config.Munbot, out io.Writer) error {
	m, err := config.Parse(cfg)
	if err != nil {
		return err
	}
	for _, k := range list(m) {
		fmt.Fprintf(out, "%s=%v\n", k, m[k])
	}
	return nil
}

func parseFilter(cfg *config.Munbot, out io.Writer, filter string) error {
	m, err := config.Parse(cfg)
	if err != nil {
		return err
	}
	if _, found := m[filter]; found {
		fmt.Fprintf(out, "%s\n", m[filter])
	} else {
		for _, k := range list(m) {
			if strings.HasPrefix(k, filter) {
				fmt.Fprintf(out, "%s=%v\n", k, m[k])
			}
		}
	}
	return nil
}

func parseDiff(def *config.Munbot, cfg *config.Munbot, out io.Writer) error {
	defm, deferr := config.Parse(def)
	if deferr != nil {
		return deferr
	}
	m, err := config.Parse(cfg)
	if err != nil {
		return err
	}
	for _, k := range list(m) {
		defv, ok := defm[k]
		v := m[k]
		if !ok || v != defv {
			fmt.Fprintf(out, "%s=%v\n", k, v)
		}
	}
	return nil
}
