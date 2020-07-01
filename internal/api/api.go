// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	"crypto/tls"
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

	protocol := "https"
	sslCheck(a.Host, a.Port, a.Cert, a.Key)
	if a.Cert == "" || a.Key == "" {
		log.Warn("api ssl check failed, forcing http on localhost...")
		a.Host = "localhost"
		protocol = "http"
	}

	//~ a.AddHandler(api.BasicAuth("munbot", "tobnum"))
	h := a.Host
	if h == "" {
		h = "0.0.0.0"
	}
	log.Printf("start api %s://%s:%s/", protocol, h, a.Port)
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
	if _, err := tls.LoadX509KeyPair(cert, key); err != nil {
		log.Fatal(err)
	}
}
