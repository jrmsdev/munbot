// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package adaptor

import (
	"testing"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/gobottest"
)

var _ gobot.Adaptor = (*Adaptor)(nil)

func TestAdaptor(t *testing.T) {
	a := New()
	gobottest.Assert(t, a.Name(), "munbot")
	gobottest.Assert(t, len(a.Connect()), 0)
	gobottest.Assert(t, a.Ping(), "pong")
	gobottest.Assert(t, len(a.Connect()), 0)
	gobottest.Assert(t, len(a.Finalize()), 0)
}
