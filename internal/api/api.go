// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	"os"
	"path/filepath"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"

	"github.com/jrmsdev/munbot/config"
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
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

func Start(m *gobot.Master, cfg *config.Api) {
	a := api.NewAPI(m)
	if flags.Debug {
		a.Debug()
	}

	a.Host = cfg.Host.String()
	a.Port = cfg.Port.String()
	a.Cert, a.Key = sslFiles(cfg)

	//~ a.AddHandler(api.BasicAuth("munbot", "tobnum"))
	a.Start()
}

func sslFiles(cfg *config.Api) (string, string) {
	cert := filepath.Join(flags.ConfigDir, cfg.Cert.String())
	key := filepath.Join(flags.ConfigDir, cfg.Key.String())
	ok := true
	_, err := os.Stat(cert)
	if err != nil {
		ok = false
		log.Error(err)
	}
	_, err = os.Stat(key)
	if err != nil {
		ok = false
		log.Error(err)
	}
	if ok {
		return cert, key
	}
	return "", ""
}
