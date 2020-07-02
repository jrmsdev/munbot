// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
	"github.com/jrmsdev/munbot/version"
)

func main() {
	fs := flags.Init("munbot")
	fs.BoolVar(&flags.DebugApi, "debug.api", false, "enable api debug")
	flags.Parse(os.Args[1:])
	log.Printf("munbot version %s", version.String())
	cfg := munbot.Configure()
	master := munbot.New()
	master.Main(cfg.Master)
}
