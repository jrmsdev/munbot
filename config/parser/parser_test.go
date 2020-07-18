// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	c := newTestCfg(t)
	c.loadTestCfg()
	m := Parse(c.test, "")
	c.assert.Equal("testing", m["master.name"], "paser master.name")
}

func TestParseFilter(t *testing.T) {
	c := newTestCfg(t)
	c.loadTestCfg()

	m := Parse(c.test, "master")
	c.assert.Equal("testing", m["master.name"], "parser master.name")

	m = Parse(c.test, "master.name")
	c.assert.Equal("testing", m["master.name"], "parser master.name")
	c.assert.Equal("", m["master.enable"], "parser master.enable")
}
