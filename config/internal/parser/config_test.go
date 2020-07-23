// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"testing"

	"github.com/munbot/master/config/value"
	"github.com/munbot/master/testing/assert"
	"github.com/munbot/master/testing/require"
)

var tcfg = []byte(`{"master":{"name":"testing"}}`)

type testCfg struct {
	t       *testing.T
	test    *Config
	assert  *assert.Assertions
	require *require.Assertions
}

func newTestCfg(t *testing.T) *testCfg {
	return &testCfg{t, New(), assert.New(t), require.New(t)}
}

func (c *testCfg) setDefaults() {
	tdef := value.DB{
		"master": value.Map{
			"name": "munbot",
		},
	}
	c.test.SetDefaults(tdef)
}

func (c *testCfg) loadCfg(s string) {
	err := c.test.Load([]byte(s))
	c.require.NoError(err, "load cfg error")
}

func (c *testCfg) loadTestCfg() {
	err := c.test.Load(tcfg)
	c.require.NoError(err, "load test cfg error")
}

func TestNew(t *testing.T) {
	c := newTestCfg(t)
	c.setDefaults()

	blob, err := c.test.Dump()
	c.require.NoError(err, "dump error")
	c.assert.Equal(blob, []byte(`{"master":{"name":"munbot"}}`), "dump blob")

	c.assert.Equal("munbot", c.test.Get("master", "name"), "master.name value")
	c.assert.Equal("ECFGMISS:master.missing", c.test.Get("master", "missing"), "get missing value")

	c.loadTestCfg()
	blob, err = c.test.Dump()
	c.require.NoError(err, "dump error")
	c.assert.Equal(blob, tcfg, "dump blob")
}

var evalCfg = `{
	"master": {
		"name":"testing"
	},
	"test": {
		"master_name": "${master.name}",
		"alias": "${test.master_name}",
		"loop": "${test.loop}",
		"loop2": "${test.loop}",
		"loop3": "${test.loop4}",
		"loop4": "${test.loop3}",
		"loop5": "${test.loop6}",
		"loop6": "${test.loop7}",
		"loop7": "${test.loop5}"
	}
}`

func TestEval(t *testing.T) {
	c := newTestCfg(t)
	c.loadCfg(evalCfg)
	c.require.Equal("testing", c.test.Get("master", "name"), "master.name")
	c.assert.Equal("testing", c.test.Get("test", "master_name"), "test master_name")
	c.assert.Equal("testing", c.test.Get("test", "alias"), "test alias")
	c.assert.Equal("ECFGLOOP:test.loop", c.test.Get("test", "loop"), "test loop")
	c.assert.Equal("ECFGLOOP:test.loop", c.test.Get("test", "loop2"), "test loop2")
	c.assert.Equal("ECFGLOOP:test.loop4", c.test.Get("test", "loop3"), "test loop3")
	c.assert.Equal("ECFGLOOP:test.loop3", c.test.Get("test", "loop4"), "test loop4")
	c.assert.Equal("ECFGLOOP:test.loop6", c.test.Get("test", "loop5"), "test loop5")
	c.assert.Equal("ECFGLOOP:test.loop7", c.test.Get("test", "loop6"), "test loop6")
	c.assert.Equal("ECFGLOOP:test.loop5", c.test.Get("test", "loop7"), "test loop7")
}

func TestCopy(t *testing.T) {
	c := newTestCfg(t)
	c.setDefaults()
	ct := c.test
	ct2 := ct.Copy()
	c.assert.Equal("munbot", ct.Get("master", "name"), "config val")
	c.assert.Equal("munbot", ct2.Get("master", "name"), "copy val")

	ct.db["master"]["name"] = "ct"
	ct2.db["master"]["name"] = "ct2"
	c.assert.Equal("ct", ct.Get("master", "name"), "config val")
	c.assert.Equal("ct2", ct2.Get("master", "name"), "copy val")
}
