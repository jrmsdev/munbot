// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/flags"
)

func main() {
	flags.Parse()
	cfg := munbot.NewConfig("munbot")
	println(cfg.String())
}
