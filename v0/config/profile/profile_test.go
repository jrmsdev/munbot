// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package profile

import (
	"testing"

	"github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/testing/assert"
)

func TestNewDefaults(t *testing.T) {
	check := assert.New(t)
	p := New()
	check.Equal("testing", p.Name)
	check.Equal(env.Get("MB_HOME"), p.Home)
	check.Equal(env.Get("MB_CONFIG"), p.Config)
	check.Equal("config.json", p.ConfigFile)
}
