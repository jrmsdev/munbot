// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/munbot/master/config"
	"github.com/munbot/master/testing/assert"
	"github.com/munbot/master/testing/require"
)

func newCfg() *config.Munbot {
	return &config.Munbot{Master: &config.Master{Name: "testing", Enable: true}}
}

func TestParse(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	c := newCfg()
	m, err := Parse(c, "")
	require.NoError(err, "parse error")
	assert.Equal("testing", m["master.name"], "paser master.name")
	assert.Equal("true", m["master.enable"], "paser master.enable")
}

func TestParseFilter(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	c := newCfg()

	m, err := Parse(c, "master")
	require.NoError(err, "parse error")
	assert.Equal("testing", m["master.name"], "parser master.name")
	assert.Equal("true", m["master.enable"], "parser master.enable")

	m, err = Parse(c, "master.name")
	require.NoError(err, "parse error")
	assert.Equal("testing", m["master.name"], "parser master.name")
	assert.Equal("", m["master.enable"], "parser master.enable")
}

func mockMarshal(v interface{}) ([]byte, error) {
	return nil, errors.New("mock marshal error")
}

func TestMarshalError(t *testing.T) {
	assert := assert.New(t)
	jsonMarshal = mockMarshal
	defer func() {
		jsonMarshal = json.Marshal
	}()
	c := newCfg()
	_, err := Parse(c, "")
	assert.EqualError(err, "mock marshal error", "marshal error")
}

func mockUnmarshal(b []byte, v interface{}) error {
	return errors.New("mock unmarshal error")
}

func TestUnmarshalError(t *testing.T) {
	assert := assert.New(t)
	jsonUnmarshal = mockUnmarshal
	defer func() {
		jsonUnmarshal = json.Unmarshal
	}()
	c := newCfg()
	_, err := Parse(c, "")
	assert.EqualError(err, "mock unmarshal error", "unmarshal error")
}
