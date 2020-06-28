// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/flags"
)

func main() {
	flags.Init()
	flags.Parse()
	cfg := munbot.NewConfig(flags.MasterName)
	println(cfg.String())
}
