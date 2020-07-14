// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"testing"

	"github.com/munbot/master/cmd/mbcfg/internal/config"
)

func TestUpdate(t *testing.T) {
	c := config.New()
	config.SetDefaults(c)
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
	//~ t.Logf("%s", Update(c, "master.dot.opt", "testing"))
	//~ t.Logf("%s", Update(c, "master.api.addr", "testing"))
	//~ t.Logf("%s", Update(c, "master.api.dot.opt", "testing"))
}
