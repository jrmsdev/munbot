// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/munbot"
)

func main() {
	cfg := munbot.NewConfig("munbot")
	println(cfg.String())
}
