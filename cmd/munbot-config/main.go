// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

var (
	listAll bool
	newUserName string
)

func main() {
	fs := flags.Init("munbot-config")
	fs.BoolVar(&listAll, "a", false, "list all options including default values")
	fs.StringVar(&newUserName, "new.user", "", "create a new user `name`")
	flags.Parse(os.Args[1:])

	log.Debug("start")
	cfg := munbot.Configure()

	if newUserName != "" {
		newUser(newUserName)
	} else {
		filter := fs.Arg(0)
		args := fs.Arg(1)
		if args != "" {
			edit(cfg, filter, args)
		} else {
			dump(cfg, os.Stdout, listAll, filter)
		}
	}

	log.Debug("end")
}
