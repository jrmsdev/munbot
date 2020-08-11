// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

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
	f := NewFlags()
	f.Set(newTestFS())
	assert.Equal("", f.apiAddr, "default api.addr")
	assert.Equal(uint(0), f.apiPort, "default api.port")
}

func TestFlagsDefaults(t *testing.T) {
	assert := assert.New(t)
	f := NewFlags()
	f.Set(newTestFS())
	f.Parse()
	assert.Equal("", f.apiAddr, "default api.addr")
	assert.Equal(uint(0), f.apiPort, "default api.port")
	assert.Equal("", f.consoleAddr, "default console.addr")
	assert.Equal(uint(0), f.consolePort, "default console.port")
}
