// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env_test

import (
	"path/filepath"
	"testing"

	"github.com/munbot/master/env"
	"github.com/munbot/master/testing/assert"
)

func TestEnvFile(t *testing.T) {
	check := assert.New(t)
	check.Equal("test", env.Get("MBENV"), "MBENV")
	cfgdir, err := filepath.Abs(filepath.FromSlash("."))
	check.NoError(err)
	check.Equal(cfgdir, env.Get("MBENV_CONFIG"), "MBENV_CONFIG")
	check.Equal("test.env", env.Get("MBTEST_ENVFILE"), "MBTEST_ENVFILE")

	check.Equal("testing", env.Get("MB_PROFILE"), "MB_PROFILE")
	check.Equal("debug", env.Get("MB_LOG"), "MB_LOG")
	check.Equal("home", env.Get("MB_HOME"), "MB_HOME")
	check.Equal("etc", env.Get("MB_CONFIG"), "MB_CONFIG")
	check.Equal("run", env.Get("MB_RUN"), "MB_RUN")

	check.Equal("true", env.Get("MBAPI"), "MBAPI")
	check.Equal("tcp", env.Get("MBAPI_NET"), "MBAPI_NET")
	check.Equal("127.0.0.1", env.Get("MBAPI_ADDR"), "MBAPI_ADDR")

	check.Equal("true", env.Get("MBAUTH"), "MBAUTH")

	check.Equal("true", env.Get("MBCONSOLE"), "MBCONSOLE")
	check.Equal("127.0.0.1", env.Get("MBCONSOLE_ADDR"), "MBCONSOLE_ADDR")
}
