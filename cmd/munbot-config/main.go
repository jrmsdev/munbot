// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"

	"github.com/jrmsdev/munbot/config2"
)

var (
	listAll     bool
	newUserName string
)

func main() {
	fs := flags.Init("munbot-config")
	fs.BoolVar(&listAll, "a", false, "list all options including default values")
	fs.StringVar(&newUserName, "new.user", "", "create a new user `name`")
	flags.Parse(os.Args[1:])

	cfg2 := config2.New()
	config2.SetDefaults(cfg2)
	config2.Save(cfg2)
	os.Exit(9)

	log.Debug("start")
	cfg := munbot.Configure()

	if newUserName != "" {
		newUser(cfg, newUserName)
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
