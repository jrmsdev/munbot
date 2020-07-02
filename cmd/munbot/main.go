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
	flags.Init("munbot")
	flags.Parse(os.Args[1:])
	log.Printf("munbot version %s", version.String())
	cfg := munbot.Configure()
	munbot.SetupInfo()
	master := munbot.New()
	master.Main(cfg.Master)
}
