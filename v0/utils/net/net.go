// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package net provides some net related utils.
package net

import (
	"net"
	"net/url"
)

type Addr struct {
	net.Addr
	addr string
	net  string
	uri  *url.URL
}

func NewAddr(scheme, network, addr string) *Addr {
	return &Addr{addr: addr, net: network, uri: &url.URL{Scheme: scheme, Host: addr}}
}

func (a *Addr) String() string {
	return a.addr
}

func (a *Addr) Network() string {
	return a.net
}

func (a *Addr) Hostname() string {
	return a.uri.Hostname()
}

func (a *Addr) Port() string {
	return a.uri.Port()
}
