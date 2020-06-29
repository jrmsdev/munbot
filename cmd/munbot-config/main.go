// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

func main() {
	flags.Init("munbot-config")
	flags.Parse(os.Args[1:])
	log.Debug("start")
	cfg := munbot.Configure()
	println(cfg.String())
	log.Debug("end")
}
