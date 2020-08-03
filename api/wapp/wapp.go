// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package wapp implements the master api webapp.
package wapp

import (
	"net/http"

	"github.com/gorilla/mux"
	"gobot.io/x/gobot/api"
)

type Config struct {
	Enable bool
	Debug  bool
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
	r.PathPrefix("/").Handler(cpppio)
	return &wapp{mux: r, cpppio: cpppio}
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
}

func (a *wapp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
