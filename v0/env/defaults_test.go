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
	check.Equal("master", env.Init["MUNBOT"], "MUNBOT")

	check.Equal("false", env.Init["MB_DEBUG"], "MB_DEBUG")
	check.Equal("default", env.Init["MB_PROFILE"], "MB_PROFILE")

	check.Equal("verbose", env.Init["MB_LOG"], "MB_LOG")
	check.Equal("auto", env.Init["MB_LOG_COLORS"], "MB_LOG_COLORS")
	check.Equal("", env.Init["MB_LOG_DEBUG"], "MB_LOG_DEBUG")

	check.Equal("", env.Init["MB_HOME"], "MB_HOME")
	check.Equal("", env.Init["MB_CONFIG"], "MB_CONFIG")
	check.Equal("", env.Init["MB_RUN"], "MB_RUN")

	check.Equal("true", env.Init["MBAPI"], "MBAPI")
	check.Equal("false", env.Init["MBAPI_DEBUG"], "MBAPI_DEBUG")
	check.Equal("tcp", env.Init["MBAPI_NET"], "MBAPI_NET")
	check.Equal("127.0.0.1", env.Init["MBAPI_ADDR"], "MBAPI_ADDR")
	check.Equal("6490", env.Init["MBAPI_PORT"], "MBAPI_PORT")
	check.Equal("/", env.Init["MBAPI_PATH"], "MBAPI_PATH")

	check.Equal("true", env.Init["MBAUTH"], "MBAUTH")

	check.Equal("true", env.Init["MBCONSOLE"], "MBCONSOLE")
	check.Equal("0.0.0.0", env.Init["MBCONSOLE_ADDR"], "MBCONSOLE_ADDR")
	check.Equal("6492", env.Init["MBCONSOLE_PORT"], "MBCONSOLE_PORT")
}
