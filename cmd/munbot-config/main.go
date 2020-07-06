// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/cmd/internal/flags"
	"github.com/jrmsdev/munbot/log"
)

var (
	listAll     bool
	jsonFormat  bool
	newUserName string
)

func main() {
	log.SetQuiet()

	fs := flags.Init("munbot-config")
	fs.BoolVar(&listAll, "a", false, "list all options including default values")
	fs.BoolVar(&jsonFormat, "json", false, "json format")
	fs.StringVar(&newUserName, "new.user", "", "create a new user `name`")
	flags.Parse(os.Args[1:])

	log.Debug("start")
	if err := munbot.Configure(fs); err != nil {
		log.Fatal(err)
	}
	cfg := munbot.Config

	var err error
	if newUserName != "" {
		err = newUser(cfg, newUserName)
	} else {
		filter := fs.Arg(0)
		args := fs.Arg(1)
		if args != "" {
			err = edit(cfg, filter, args)
		} else {
			dump(cfg, os.Stdout, listAll, jsonFormat, filter)
		}
	}
	if err != nil {
		log.Debugf("exit error: %s", err)
		os.Exit(2)
	}
	log.Debug("end")
}
