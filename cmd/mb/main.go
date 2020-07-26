// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// mb command is the main munbot application.
package main

import (
	"os"

	"github.com/munbot/master/cmd"
	"github.com/munbot/master/mb"
)

func main() {
	m := cmd.New("mb", mb.New())
	m.Main(os.Args[1:])
}
