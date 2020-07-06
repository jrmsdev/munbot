// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/cmd/internal/flags"
	"github.com/jrmsdev/munbot/log"
	"github.com/jrmsdev/munbot/version"

	cf "github.com/jrmsdev/munbot/config/flags"
)

func main() {
	fs := flags.Init("munbot")
	fs.BoolVar(&cf.DebugApi, "debug.api", false, "enable api debug")
	fs.StringVar(&cf.ApiAddr, "api.addr", cf.ApiAddr, "api network `address`")
	fs.IntVar(&cf.ApiPort, "api.port", cf.ApiPort, "api tcp port `number`")
	fs.StringVar(&cf.ApiCert, "api.cert", cf.ApiCert, "api tls cert file `relpath`")
	fs.StringVar(&cf.ApiKey, "api.key", cf.ApiKey, "api tls key file `relpath`")
	flags.Parse(os.Args[1:])

	log.Printf("munbot version %s", version.String())
	if err := munbot.Configure(fs); err != nil {
		log.Fatal(err)
	}
	master := munbot.New()
	master.Main(munbot.Config.Master)
}
