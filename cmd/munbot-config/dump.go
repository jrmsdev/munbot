// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"io"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/log"
)

func dump(cfg *munbot.Config, out io.Writer, listAll bool, filter string) {
	log.Debugf("dump all=%v", listAll)
	cfg.Dump(out, listAll, filter)
}
