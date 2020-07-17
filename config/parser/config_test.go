// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"testing"
)

var tcfg = []byte(`{"master":{"enable":"false","name":"testing"}}`)

func TestNew(t *testing.T) {
	c := New()
	t.Logf("%#v", c)
	SetDefaults(c)
	t.Logf("%#v", c)
	blob, err := c.Dump()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", blob)
	s := c.Section("master")
	t.Log(s)
	t.Logf("master.name: %s", s.Get("name"))
	t.Logf("master.enable: %v", s.GetBool("enable"))
	t.Logf("missing bool: %v", s.GetBool("missing"))
	s = c.Section("master.api")
	t.Log(s)
	t.Logf("master.api.enable: %v", s.GetBool("enable"))

	err = c.Load(tcfg)
	if err != nil {
		t.Fatal(err)
	}
	blob, err = c.Dump()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", blob)
}
