// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	"io"

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
	}
	return nil
}

func jsonDump(cfg *config.Munbot, out io.Writer) error {
	return config.Write(cfg, out)
}

func parseDump(cfg *config.Munbot, out io.Writer) error {
	m, err := config.Parse(cfg)
	if err != nil {
		return err
	}
	for k, v := range m {
		fmt.Printf("%s=%v\n", k, v)
	}
	return nil
}
