// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package api defines and implements master api server interface.
package api

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/gorilla/mux"

	"github.com/munbot/master/log"
)

var serverTimeout time.Duration = 15 * time.Second
var stopTimeout time.Duration = 30 * time.Second

func newHTTPServer(h http.Handler) *http.Server {
	return &http.Server{
		Handler:      h,
		WriteTimeout: serverTimeout,
		ReadTimeout:  serverTimeout,
	}
}

var _ Server = &Api{}

type Api struct {
	enable bool
	mux    *mux.Router
	server *http.Server
}

func New() Server {
	a := &Api{mux: mux.NewRouter()}
	a.server = newHTTPServer(a.mux)
	return a
}

func (a *Api) Configure(c *ServerConfig) error {
	a.enable = c.Enable
	a.server.Addr = fmt.Sprintf("%s:%d", c.Addr, c.Port)
	return nil
}

func (a *Api) Start() error {
	if a.enable {
		log.Printf("Api server http://%s/", a.server.Addr)
		if err := a.server.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
	} else {
		log.Warn("Api server is disabled")
	}
	return nil
}

func (a *Api) Stop() error {
	if a.enable {
		log.Debugf("server shutdown... timeout in %s", stopTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), stopTimeout)
		defer cancel()
		if err := a.server.Shutdown(ctx); err != http.ErrServerClosed {
			return err
		}
	}
	return nil
}

func (a *Api) cleanPath(p string) string {
	p = path.Clean(p)
	if p == "." {
		return "/"
	}
	return p
}

func (a *Api) Mount(prefix string, handler http.Handler) {
	prefix = a.cleanPath(prefix)
	if prefix == "/" {
		a.mux.PathPrefix(prefix).Handler(handler)
	} else {
		a.mux.PathPrefix(prefix).Handler(http.StripPrefix(prefix, handler))
	}
}
