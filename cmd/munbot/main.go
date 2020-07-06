// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/cmd/internal/flags"
	"github.com/jrmsdev/munbot/log"
	"github.com/jrmsdev/munbot/version"

	cf "github.com/jrmsdev/munbot/config/flags"
)

func main() {
	fs := flags.Init("munbot")
	fs.BoolVar(&cf.DebugApi, "debug.api", false, "enable api debug")
	flags.Parse(os.Args[1:])
	log.Printf("munbot version %s", version.String())
	err := munbot.Configure()
	if err != nil {
		log.Panic(err)
	}
	master := munbot.New()
	master.Main(munbot.Config.Master)
}
