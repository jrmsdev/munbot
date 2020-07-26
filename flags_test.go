// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"flag"
	"testing"

	"github.com/munbot/master/config"
	"github.com/munbot/master/testing/assert"
)

func newTestFS() *flag.FlagSet {
	return flag.NewFlagSet("testing", flag.PanicOnError)
}

func TestFlagsSet(t *testing.T) {
	assert := assert.New(t)
	f := new(Flags)
	f.Set(newTestFS())
	assert.Equal("", f.ApiAddr, "default api.addr")
	assert.Equal(0, f.ApiPort, "default api.port")
}

func TestFlagsParse(t *testing.T) {
	assert := assert.New(t)
	f := new(Flags)
	f.Set(newTestFS())
	c := config.New()
	f.Parse(c)
	assert.Equal("ECFGMISS:default.netaddr", f.ApiAddr, "default api.addr")
	assert.Equal(int(0), f.ApiPort, "default api.port")
}

func TestFlagsDefaults(t *testing.T) {
	assert := assert.New(t)
	f := new(Flags)
	f.Set(newTestFS())
	c := config.New()
	c.SetDefaults(config.Defaults)
	f.Parse(c)
	assert.Equal("0.0.0.0", f.ApiAddr, "default api.addr")
	assert.Equal(int(6492), f.ApiPort, "default api.port")
}
