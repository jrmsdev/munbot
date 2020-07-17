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
			"enable": "true",
		},
	}}
}

func TestParse(t *testing.T) {
	assert := assert.New(t)
	c := newCfg()
	m := Parse(c, "")
	assert.Equal("testing", m["master.name"], "paser master.name")
	assert.Equal("true", m["master.enable"], "paser master.enable")
}

func TestParseFilter(t *testing.T) {
	assert := assert.New(t)
	c := newCfg()

	m := Parse(c, "master")
	assert.Equal("testing", m["master.name"], "parser master.name")
	assert.Equal("true", m["master.enable"], "parser master.enable")

	m = Parse(c, "master.name")
	assert.Equal("testing", m["master.name"], "parser master.name")
	assert.Equal("", m["master.enable"], "parser master.enable")
}
