// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

var (
	listAll bool
	cmdNew bool
)

func main() {
	fs := flags.Init("munbot-config")
	fs.BoolVar(&listAll, "a", false, "list all options including default values")
	flags.Parse(os.Args[1:])

	log.Debug("start")
	cfg := munbot.Configure()

	filter := fs.Arg(0)
	args := fs.Arg(1)
	dumpOrEdit(os.Stdout, cfg, filter, args)

	log.Debug("end")
}

func dumpOrEdit(out io.Writer, cfg *munbot.Config, filter, args string) {
	if args != "" {
		edit(cfg, filter, args)
	} else {
		log.Debug("dump...")
		cfg.Dump(os.Stdout, listAll, filter)
	}
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
	if err := ioutil.WriteFile(fn, blob, 0660); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s saved", fn)
}
