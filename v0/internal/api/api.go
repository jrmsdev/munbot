// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package api defines and implements master api server interface.
package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"path"
	"time"

	"github.com/gorilla/mux"

	"github.com/munbot/master/v0/config/profile"
	"github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/log"
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
	mux    *mux.Router
	server *http.Server
	ln     net.Listener
	enable bool
	net    string
}

func New() Server {
	a := &Api{mux: mux.NewRouter()}
	a.server = newHTTPServer(a.mux)
	return a
}

func (a *Api) Configure() error {
	a.enable = env.GetBool("MBAPI")
	a.net = env.Get("MBAPI_NET")
	if a.net == "tcp" || a.net == "tcp4" || a.net == "tcp6" {
		addr := env.Get("MBAPI_ADDR")
		port := env.GetUint("MBAPI_PORT")
		a.server.Addr = fmt.Sprintf("%s:%d", addr, port)
	} else if a.net == "unix" {
		prof := profile.New()
		a.server.Addr = prof.GetRundirPath("api.socket")
	} else {
		return fmt.Errorf("api: invalid network %q", a.net)
	}
	return nil
}

func (a *Api) Start() error {
	if a.enable {
		var err error
		a.ln, err = net.Listen(a.net, a.server.Addr)
		if err != nil {
			log.Debugf("listen error: %v", err)
			return err
		}
		log.Printf("Api server http://%s", a.server.Addr)
		if err := a.server.Serve(a.ln); err != http.ErrServerClosed {
			return err
		}
	} else {
		log.Warn("Api server is disabled")
	}
	return nil
}

func (a *Api) Stop() error {
	if a.enable && a.ln != nil {
		log.Debugf("server shutdown... timeout in %s", stopTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), stopTimeout)
		defer cancel()
		if err := a.server.Shutdown(ctx); err != nil {
			if err != http.ErrServerClosed {
				log.Debugf("shutdown error: %v", err)
				return err
			}
		}
	} else {
		log.Debugf("avoid stop... enable:%v ln:%v", a.enable, a.ln == nil)
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
