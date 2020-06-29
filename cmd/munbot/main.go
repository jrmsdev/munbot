// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/flags"
)

func main() {
	flags.Init("munbot")
	flags.Parse(os.Args[1:])
	master := munbot.New()
	master.Main()
}
