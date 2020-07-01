// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"time"

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
	sslCheck(a.Host, a.Port, a.Cert, a.Key)

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
		log.Debugf("cert file %s", cert)
		log.Debugf("key file %s", key)
		return cert, key
	}
	return "", ""
}

func sslCheck(host, port, cert, key string) {
	log.Debug("ssl check")
	if cert == "" && key == "" {
		return
	}
	s := &http.Server{Addr: host+":"+port}
	go func() {
		err := s.ListenAndServeTLS(cert, key)
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	time.Sleep(time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
