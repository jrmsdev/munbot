// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/munbot/master/cmd"
	"github.com/munbot/master/cmd/config"
)

func main() {
	m := cmd.New("mbcfg", config.New())
	m.AddCommand("config", config.New())
	m.Main(os.Args[1:])
}
