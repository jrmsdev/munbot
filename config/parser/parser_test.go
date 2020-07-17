// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"testing"

	"github.com/munbot/master/testing/assert"
)

func newCfg() *Config {
	return &Config{db: DB{
		"master": Map{
			"name": "testing",
		},
	}}
}

func TestParse(t *testing.T) {
	assert := assert.New(t)
	c := newCfg()
	m := Parse(c, "")
	assert.Equal("testing", m["master.name"], "paser master.name")
}

func TestParseFilter(t *testing.T) {
	assert := assert.New(t)
	c := newCfg()

	m := Parse(c, "master")
	assert.Equal("testing", m["master.name"], "parser master.name")

	m = Parse(c, "master.name")
	assert.Equal("testing", m["master.name"], "parser master.name")
	assert.Equal("", m["master.enable"], "parser master.enable")
}
