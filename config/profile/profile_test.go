// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package profile

import (
	"path/filepath"
	"testing"

	"github.com/munbot/master/testing/assert"
	"github.com/munbot/master/testing/require"
)

func TestInitErrors(t *testing.T) {
	require := require.New(t)
	require.NoError(homeDirErr, "os user home dir")
	require.NoError(configDirErr, "os user config dir")
}

func TestNewDefaults(t *testing.T) {
	assert := assert.New(t)
	homeDir = ""
	configDir = ""
	p := New("testing")
	home := filepath.FromSlash("./.munbot")
	assert.Equal(home, homeDir, "home dir")
	assert.Equal(filepath.Join(home, "config"), configDir, "config dir")
	assert.Equal(configDir, p.ConfigDir, "profile config dir")
	assert.Equal("testing", p.Name, "profile name")
	assert.Equal("config.json", p.Config, "config filename")
	assert.Equal(filepath.FromSlash("/usr/local/etc/munbot"),
		p.ConfigSysDir, "config sys dir")
	assert.Equal(filepath.FromSlash("/etc/munbot"),
		p.ConfigDistDir, "config dist dir")
}
