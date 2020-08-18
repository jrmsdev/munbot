// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// mbcfg command manages munbot configuration files.
package main

import (
	"os"

	"github.com/munbot/master/v0/cmd"
	"github.com/munbot/master/v0/cmd/mbcfg"
)

func main() {
	m := cmd.New("mbcfg", mbcfg.New())
	m.Main(os.Args[1:])
}
