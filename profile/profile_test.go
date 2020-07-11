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

func TestSetDefaults(t *testing.T) {
	assert := assert.New(t)
	homeDir = ""
	ConfigDir = ""
	setDefaults()
	home := filepath.FromSlash("./.munbot")
	assert.Equal(home, homeDir, "home dir")
	assert.Equal(filepath.Join(home, "config"), ConfigDir, "config dir")
	assert.Equal("munbot", Name, "profile name")
	assert.Equal("config.json", ConfigFilename, "config filename")
	assert.Equal(filepath.FromSlash("/usr/local/etc/munbot"),
		ConfigSysDir, "config sys dir")
	assert.Equal(filepath.FromSlash("/etc/munbot"),
		ConfigDistDir, "config dist dir")
}
