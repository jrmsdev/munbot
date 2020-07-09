// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/munbot/master"
	"github.com/munbot/master/config"
	"github.com/munbot/master/cmd/internal/flags"
	"github.com/munbot/master/log"
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
	if err := munbot.Setup(); err != nil {
		log.Fatal(err)
	}

	filter := fs.Arg(0)
	args := fs.Arg(1)

	// use a config with no defaults set
	cfg := config.New()
	if listAll || filter != "" {
		// unless list all is requested
		config.SetDefaults(cfg)
	}
	if err := config.ReadFiles(cfg); err != nil {
		log.Fatal(err)
	}

	var err error
	if newUserName != "" {
		err = newUser(cfg, newUserName)
	} else {
		if args != "" {
			err = edit(cfg, filter, args)
		} else {
			dump(cfg, os.Stdout, jsonFormat, filter)
		}
	}
	if err != nil {
		log.Debugf("exit error: %s", err)
		os.Exit(2)
	}
	log.Debug("end")
}
