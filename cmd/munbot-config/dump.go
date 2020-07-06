// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"io"

	"github.com/jrmsdev/munbot/config"
	"github.com/jrmsdev/munbot/log"
)

func dump(cfg *config.Munbot, out io.Writer, listAll, jsonFormat bool, filter string) error {
	log.Debugf("dump all=%v json=%v", listAll, jsonFormat)
	if jsonFormat {
		return jsonDump(cfg, out)
	}
	return nil
}

func jsonDump(cfg *config.Munbot, out io.Writer) error {
	return config.Write(cfg, out)
}
