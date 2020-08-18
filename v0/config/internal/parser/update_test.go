// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"testing"
)

func TestUpdate(t *testing.T) {
	c := newTestCfg(t)
	c.setDefaults()
	c.require.Equal("munbot", c.test.Get("master", "name"), "master name default")
	err := Update(c.test, "master.name", "testing")
	c.require.NoError(err, "update master.name error")
	c.assert.Equal("testing", c.test.Get("master", "name"), "master name update")
	err = Update(c.test, "nosect.opt", "val")
	c.assert.EqualError(err, "update invalid section: nosect", "update error")
	err = Update(c.test, "master.noopt", "val")
	c.assert.EqualError(err, "update invalid option: master.noopt", "update error")
}

func TestSet(t *testing.T) {
	c := newTestCfg(t)
	err := Set(c.test, "test.opt", "testing")
	c.require.NoError(err, "set error")
	c.assert.Equal("testing", c.test.Get("test", "opt"), "test opt")
	err = Set(c.test, "test.opt", "dup")
	c.require.Error(err, "set dup error")
	c.assert.Equal("testing", c.test.Get("test", "opt"), "test opt")
}

func TestSetEmptyOption(t *testing.T) {
	c := newTestCfg(t)
	err := Set(c.test, "test", "opt")
	c.assert.EqualError(err, "set invalid format: test opt", "set error")
}

func TestUnset(t *testing.T) {
	c := newTestCfg(t)
	err := Unset(c.test, "test.opt")
	c.require.NoError(err, "unset error")
	c.require.Equal("ECFGMISS:test.opt", c.test.Get("test", "opt"), "test opt")

	err = Set(c.test, "test.opt", "testing")
	c.require.NoError(err, "set error")
	c.require.Equal("testing", c.test.Get("test", "opt"), "test opt")

	err = Unset(c.test, "test.opt")
	c.require.NoError(err, "unset error")
	c.require.Equal("ECFGMISS:test.opt", c.test.Get("test", "opt"), "test opt")
}
