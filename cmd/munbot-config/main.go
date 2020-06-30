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
)

func main() {
	parser := flags.Init("munbot-config")
	parser.BoolVar(&listAll, "a", false, "list all config options including default values")
	flags.Parse(os.Args[1:])
	log.Debug("start")
	cfg := munbot.Configure()
	cfg.Dump(os.Stdout, listAll, parser.Arg(0))
	log.Debug("end")
}

func debug() {
	flags.Init("munbot-config")
	flags.Parse(os.Args[1:])

	log.Debug("start")
	cfg := munbot.Configure()

	//~ log.Print("dump1")
	//~ cfg.Dump(os.Stdout, false)
	//~ log.Printf("master.name=%s", cfg.Master.Name)

	//~ log.Print("write1")
	//~ log.Printf("%#v", cfg)
	log.Printf("write1 error: %v", cfg.Write(os.Stdout))

	log.Printf("update error status %v", cfg.Update("master", "name", "saskia"))

	//~ log.Print("dump2")
	//~ cfg.Dump(os.Stdout, false)
	//~ log.Printf("master.name=%s", cfg.Master.Name)

	//~ log.Print("write2")
	//~ log.Printf("%#v", cfg)
	log.Printf("write2 error: %v", cfg.Write(os.Stdout))

	log.Debug("end")
}
