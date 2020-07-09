// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"

	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
)

func dump(cfg *config.Munbot, out io.Writer, jsonFormat bool, filter string) error {
	log.Debugf("dump all=%v json=%v", listAll, jsonFormat)
	if filter != "" {
		if jsonFormat {
			return jsonFilter(cfg, out, filter)
		}
		return parseFilter(cfg, out, filter)
	}
	if jsonFormat {
		return jsonDump(cfg, out)
	}
	return parseDump(cfg, out)
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
	if blob, err := json.MarshalIndent(cfg, "", "\t"); err != nil {
		return err
	} else {
		if _, err := out.Write(blob); err != nil {
			return err
		}
		out.Write([]byte("\n"))
	}
	return nil
}

func jsonFilter(cfg *config.Munbot, out io.Writer, filter string) error {
	m, err := config.ParseJSON(cfg, filter)
	if err != nil {
		return err
	}
	if blob, err := json.MarshalIndent(m, "", "\t"); err != nil {
		return err
	} else {
		if _, err := out.Write(blob); err != nil {
			return err
		}
		out.Write([]byte("\n"))
	}
	return nil
}

func jsonDiff(def *config.Munbot, cfg *config.Munbot, out io.Writer) error {
	return nil
}

func parseDump(cfg *config.Munbot, out io.Writer) error {
	m, err := config.Parse(cfg, "")
	if err != nil {
		return err
	}
	for _, k := range list(m) {
		if _, err := fmt.Fprintf(out, "%s=%v\n", k, m[k]); err != nil {
			return err
		}
	}
	return nil
}

func parseFilter(cfg *config.Munbot, out io.Writer, filter string) error {
	m, err := config.Parse(cfg, filter)
	if err != nil {
		return err
	}
	if _, found := m[filter]; found {
		fmt.Fprintf(out, "%s\n", m[filter])
	} else {
		for _, k := range list(m) {
			if _, err := fmt.Fprintf(out, "%s=%v\n", k, m[k]); err != nil {
				return err
			}
		}
	}
	return nil
}

func parseDiff(def *config.Munbot, cfg *config.Munbot, out io.Writer) error {
	defm, deferr := config.Parse(def, "")
	if deferr != nil {
		return deferr
	}
	m, err := config.Parse(cfg, "")
	if err != nil {
		return err
	}
	for _, k := range list(m) {
		defv, ok := defm[k]
		v := m[k]
		if !ok || v != defv {
			if _, err := fmt.Fprintf(out, "%s=%v\n", k, v); err != nil {
				return err
			}
		}
	}
	return nil
}
