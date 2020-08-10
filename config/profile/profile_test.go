// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package profile

import (
	"path/filepath"
	"testing"

	"github.com/munbot/master/testing/assert"
)

func TestNewDefaults(t *testing.T) {
	check := assert.New(t)
	p := New("testing")
	home := filepath.FromSlash("/var/local/munbot")
	check.Equal("testing", p.Name)
	check.Equal(home, p.Home)
	check.Equal(filepath.Join(home, "config"), p.Config)
	check.Equal("config.json", p.ConfigFile)
}
