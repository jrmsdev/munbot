// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/flags"
)

func main() {
	flags.Init("munbot-config")
	flags.Parse(os.Args[1:])
	cfg := munbot.NewConfig(flags.MasterName)
	println(cfg.String())
}
