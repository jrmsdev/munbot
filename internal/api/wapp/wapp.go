// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package wapp implements the master api webapp.
package wapp

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"gobot.io/x/gobot/api"

	"github.com/munbot/master/log"
)

type Config struct {
	Enable bool
	Debug  bool
	Path   string
}

type Api interface {
	Configure(*Config)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type wapp struct {
	mux     *mux.Router
	cfginit bool
	cpppio  *api.API
}

func New(cpppio *api.API) Api {
	r := mux.NewRouter()
	return &wapp{mux: r, cpppio: cpppio}
}

func (a *wapp) cleanPath(p string) string {
	p = path.Clean(p)
	if p == "." {
		return "/"
	}
	return p
}

func (a *wapp) mount(prefix string, h http.Handler) {
	prefix = a.cleanPath(prefix)
	if prefix == "/" {
		a.mux.PathPrefix(prefix).Handler(h)
	} else {
		a.mux.PathPrefix(prefix).Handler(http.StripPrefix(prefix, h))
	}
}

func (a *wapp) Configure(c *Config) {
	if !a.cfginit {
		if c.Enable {
			a.cpppio.AddRobeauxRoutes()
		}
		a.cfginit = true
	}
	if c.Debug {
		a.cpppio.Debug()
	}
	log.Debugf("api path: %s", c.Path)
	a.mount(c.Path, a.cpppio)
}

func (a *wapp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
