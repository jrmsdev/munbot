// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

var fileOpen func(string) (*os.File, error) = os.Open

var (
	listAll bool
)

func main() {
	parser := flags.Init("munbot-config")
	parser.BoolVar(&listAll, "a", false, "list all config options including default values")
	flags.Parse(os.Args[1:])
	log.Debug("start")
	cfg := munbot.Configure()

	filter := parser.Arg(0)
	args := parser.Arg(1)
	if args != "" {
		edit(cfg, filter, args)
	} else {
		log.Debug("dump...")
		cfg.Dump(os.Stdout, listAll, filter)
	}
	log.Debug("end")
}

func edit(cfg *munbot.Config, filter, args string) {
	log.Debug("edit...")
	if err := cfg.Update(filter, args); err != nil {
		log.Fatal(err)
	}
	fn := filepath.Join(flags.ConfigDir, flags.ConfigFile)
	blob, err := cfg.Bytes()
	if err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Dir(fn), 0770); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(fn, blob, 0660); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s saved", fn)
}
