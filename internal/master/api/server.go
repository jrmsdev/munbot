// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	"crypto/tls"
	"net/http"

	"gobot.io/x/gobot/api"
)

func newServer(h *api.API) *http.Server {
	return &http.Server{
		Addr: h.Host + ":" + h.Port,
		Handler: h,
		TLSConfig: &tls.Config{
			ServerName: h.Host,
			ClientAuth: tls.RequestClientCert,
		},
	}
}
