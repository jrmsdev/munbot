// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"testing"

	"github.com/munbot/master/testing/assert"
	"github.com/munbot/master/testing/require"
)

var tdef DB = DB{
	"master": Map{
		"name": "munbot",
	},
}

var tcfg = []byte(`{"master":{"name":"testing"}}`)

type testCfg struct {
	t *testing.T
	test *Config
	assert *assert.Assertions
	require *require.Assertions
}

func newTestCfg(t *testing.T) *testCfg {
	return &testCfg{t, New(), assert.New(t), require.New(t)}
}

func (c *testCfg) setDefaults() {
	c.test.SetDefaults(tdef)
}

func TestNew(t *testing.T) {
	c := newTestCfg(t)
	c.setDefaults()

	blob, err := c.test.Dump()
	c.require.NoError(err, "dump error")
	c.assert.Equal(blob, []byte(`{"master":{"name":"munbot"}}`), "dump blob")

	s := c.test.Section("master")
	c.require.Equal(s.Name(), "master", "section name")
	c.assert.Equal(s.Get("name"), "munbot", "master.name value")
	c.assert.Equal(s.Get("missing"), "ECFGMISS:master.missing", "get missing value")

	s = c.test.Section("missing")
	c.require.Equal(s.Name(), "ECFGSECT:missing", "missing section name")

	err = c.test.Load(tcfg)
	c.require.NoError(err, "load error")

	blob, err = c.test.Dump()
	c.require.NoError(err, "dump error")
	c.assert.Equal(blob, tcfg, "dump blob")
}
