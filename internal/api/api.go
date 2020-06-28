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

	"github.com/jrmsdev/munbot/flags"
)

//~ func main() {
//~ master := gobot.NewMaster()

//~ a := api.NewAPI(master)
//a.AddHandler(api.BasicAuth("munbot", "lalala"))
//~ a.Debug()

//~ a.AddHandler(func(w http.ResponseWriter, r *http.Request) {
//~ fmt.Fprintf(w, "Hello, %q \n", html.EscapeString(r.URL.Path))
//~ })
//~ a.Start()

func Start(m *gobot.Master) {
	a := api.NewAPI(m)
	if flags.Debug {
		a.Debug()
	}
	//~ a.AddHandler(api.BasicAuth("munbot", "tobnum"))
	a.Start()
}
