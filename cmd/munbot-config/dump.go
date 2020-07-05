// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"io"

	"github.com/jrmsdev/munbot/config"
	"github.com/jrmsdev/munbot/log"
)

func dump(cfg *config.Munbot, out io.Writer, listAll bool, filter string) {
	log.Debugf("dump all=%v", listAll)
	//~ cfg.Dump(out, listAll, filter)
}
