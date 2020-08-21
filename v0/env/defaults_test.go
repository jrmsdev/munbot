// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env_test

import (
	"testing"

	"github.com/munbot/master/v0/env"
	"github.com/munbot/master/v0/testing/assert"
)

func TestDefaults(t *testing.T) {
	check := assert.New(t)

	check.Equal("life", env.MBENV, "env.MBENV")
	check.Equal("", env.MBENV_CONFIG, "env.MBENV_CONFIG")

	check.Equal("master", env.Defaults["MUNBOT"], "MUNBOT")

	check.Equal("false", env.Defaults["MB_DEBUG"], "MB_DEBUG")
	check.Equal("default", env.Defaults["MB_PROFILE"], "MB_PROFILE")

	check.Equal("verbose", env.Defaults["MB_LOG"], "MB_LOG")
	check.Equal("auto", env.Defaults["MB_LOG_COLORS"], "MB_LOG_COLORS")
	check.Equal("", env.Defaults["MB_LOG_DEBUG"], "MB_LOG_DEBUG")

	check.Equal("", env.Defaults["MB_HOME"], "MB_HOME")
	check.Equal("", env.Defaults["MB_CONFIG"], "MB_CONFIG")
	check.Equal("", env.Defaults["MB_RUN"], "MB_RUN")

	check.Equal("true", env.Defaults["MBAPI"], "MBAPI")
	check.Equal("false", env.Defaults["MBAPI_DEBUG"], "MBAPI_DEBUG")
	check.Equal("tcp", env.Defaults["MBAPI_NET"], "MBAPI_NET")
	check.Equal("127.0.0.1", env.Defaults["MBAPI_ADDR"], "MBAPI_ADDR")
	check.Equal("6490", env.Defaults["MBAPI_PORT"], "MBAPI_PORT")
	check.Equal("/", env.Defaults["MBAPI_PATH"], "MBAPI_PATH")

	check.Equal("true", env.Defaults["MBAUTH"], "MBAUTH")

	check.Equal("true", env.Defaults["MBCONSOLE"], "MBCONSOLE")
	check.Equal("0.0.0.0", env.Defaults["MBCONSOLE_ADDR"], "MBCONSOLE_ADDR")
	check.Equal("6492", env.Defaults["MBCONSOLE_PORT"], "MBCONSOLE_PORT")
}
