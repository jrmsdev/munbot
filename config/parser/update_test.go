// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"testing"
)

func TestUpdate(t *testing.T) {
	c := New()
	c.SetDefaults(tdef)
	s := c.Section("master")
	blob, err := c.Dump()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", blob)
	t.Logf("master.name: %s", s.Get("name"))
	t.Log(Update(c, "master.name", "testing"))
	blob, err = c.Dump()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", blob)
	s = c.Section("master")
	t.Logf("master.name: %s", s.Get("name"))
}
