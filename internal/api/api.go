// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	//~ "fmt"
	//~ "html"
	//~ "net/http"
	//~ "time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"

	//~ "github.com/jrmsdev/munbot"
)

func Start(m *gobot.Master) {
	a := api.NewAPI(m)
	//~ a.AddHandler(api.BasicAuth("munbot", "tobnum"))
	a.Debug()
	a.Start()
}
