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
	t.Logf("%s", Update(c, "master.name", "testing"))
	t.Logf("%s", Update(c, "master.dot.opt", "testing"))
	t.Logf("%s", Update(c, "master.api.addr", "testing"))
	t.Logf("%s", Update(c, "master.api.dot.opt", "testing"))
}
