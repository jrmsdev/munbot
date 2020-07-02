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

type Api struct {
	ctl *api.API
}

func New(m *gobot.Master) *Api {
	return &Api{api.NewAPI(m)}
}

//~ func main() {
//~ master := gobot.NewMaster()

//~ a := api.NewAPI(master)
//a.AddHandler(api.BasicAuth("munbot", "lalala"))
//~ a.Debug()

//~ a.AddHandler(func(w http.ResponseWriter, r *http.Request) {
//~ fmt.Fprintf(w, "Hello, %q \n", html.EscapeString(r.URL.Path))
//~ })
//~ a.Start()

func (a *Api) Start(cfg *config.Api) {
	log.Debug("start")
	if flags.DebugApi {
		a.ctl.Debug()
	}
	a.ctl.Host = cfg.Addr.String()
	a.ctl.Port = cfg.Port.String()
	a.ctl.Cert, a.ctl.Key = tlsFiles(cfg)

	protocol := "https"
	tlsCheck(a.ctl.Host, a.ctl.Port, a.ctl.Cert, a.ctl.Key)
	if a.ctl.Cert == "" || a.ctl.Key == "" {
		log.Warn("api ssl check failed, forcing http on localhost...")
		a.ctl.Host = "localhost"
		protocol = "http"
	}

	//~ a.ctl.AddHandler(api.BasicAuth("munbot", "tobnum"))

	h := a.ctl.Host
	if h == "" {
		h = "0.0.0.0"
	}
	log.Printf("Start api %s://%s:%s/", protocol, h, a.ctl.Port)
	a.ctl.Start()
}

func tlsFiles(cfg *config.Api) (string, string) {
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

func tlsCheck(host, port, cert, key string) {
	log.Debug("check tls load x509 key pair")
	if cert == "" && key == "" {
		return
	}
	if _, err := tls.LoadX509KeyPair(cert, key); err != nil {
		log.Fatal(err)
	}
	log.Printf("TLS x509 cert %s", cert)
	log.Printf("TLS x509 key %s", key)
}
