// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env_test

import (
	"testing"

	"github.com/munbot/master/env"
	"github.com/munbot/master/testing/assert"
)

func TestEnvFile(t *testing.T) {
	check := assert.New(t)
	check.Equal("__UNSET__", env.Get("MBTEST_UNSET_ENV_VAR"))
	check.Equal("test", env.Get("MBENV"), "MBENV")
	check.Equal("test.env", env.Get("MBTEST_ENVFILE"), "MBTEST_ENVFILE")
	check.Equal("testing", env.Get("MB_PROFILE"), "MB_PROFILE")
	check.Equal("etc", env.Get("MB_CONFIG"), "MB_CONFIG")
}
