// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	"crypto/tls"
	"os"
	"path/filepath"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"

	"github.com/munbot/master/config"
	"github.com/munbot/master/config/flags"
	"github.com/munbot/master/log"
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
	a.ctl.Host = flags.ApiAddr
	a.ctl.Port = config.Itoa(flags.ApiPort)
	a.ctl.Cert, a.ctl.Key = tlsFiles()

	protocol := "https"
	if a.ctl.Cert == "" || a.ctl.Key == "" {
		log.Warn("api ssl check failed, forcing http on localhost...")
		a.ctl.Cert = ""
		a.ctl.Key = ""
		a.ctl.Host = "localhost"
		protocol = "http"
	} else {
		tlsCheck(a.ctl.Cert, a.ctl.Key)
	}

	//~ a.ctl.AddHandler(api.BasicAuth("munbot", "tobnum"))

	log.Printf("Start api %s://%s:%s/", protocol, a.ctl.Host, a.ctl.Port)
	a.ctl.Start()
}

func tlsFiles() (string, string) {
	cert := filepath.Join(flags.ConfigDir, flags.ApiCert)
	key := filepath.Join(flags.ConfigDir, flags.ApiKey)
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
	// TODO: check key file permissions are not too open?
	//       refuse to start if not 600 at least or even 400
	if ok {
		log.Debugf("cert file %s", cert)
		log.Debugf("key file %s", key)
		return cert, key
	}
	return "", ""
}

func tlsCheck(cert, key string) {
	log.Debug("check tls load x509 key pair")
	if _, err := tls.LoadX509KeyPair(cert, key); err != nil {
		log.Fatal(err)
	}
	log.Printf("TLS x509 cert %s", cert)
	log.Printf("TLS x509 key %s", key)
}
