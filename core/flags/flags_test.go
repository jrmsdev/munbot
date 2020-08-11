// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"testing"

	"github.com/munbot/master/testing/assert"
)

func newTestFS() *flag.FlagSet {
	return flag.NewFlagSet("testing", flag.PanicOnError)
}

func TestFlagsSet(t *testing.T) {
	assert := assert.New(t)
	f := New()
	f.Set(newTestFS())
	assert.Equal("", f.ApiAddr, "default api.addr")
	assert.Equal(uint(0), f.ApiPort, "default api.port")
}

func TestFlagsDefaults(t *testing.T) {
	assert := assert.New(t)
	f := New()
	f.Set(newTestFS())
	f.Parse()
	assert.Equal("", f.ApiAddr, "default api.addr")
	assert.Equal(uint(0), f.ApiPort, "default api.port")
	assert.Equal("", f.ConsoleAddr, "default console.addr")
	assert.Equal(uint(0), f.ConsolePort, "default console.port")
}
