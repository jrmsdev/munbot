// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package driver

import (
	"testing"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/gobottest"

	"github.com/jrmsdev/munbot/adaptor"
)

var _ gobot.Driver = (*Driver)(nil)

func TestDriver(t *testing.T) {
	d := New(adaptor.New())

	gobottest.Assert(t, d.Name(), "munbot")
	gobottest.Assert(t, d.Connection().Name(), "munbot")

	ret := d.Command(Hello)(nil)
	gobottest.Assert(t, ret.(string), "hello from Munbot!")

	gobottest.Assert(t, d.Ping(), "pong")
	gobottest.Assert(t, len(d.Start()), 0)

	//~ time.Sleep(d.interval)

	sem := make(chan bool, 0)

	d.On(d.Event(Hello), func(data interface{}) {
		sem <- true
	})
	select {
	case <-sem:
	case <-time.After(600 * time.Millisecond):
		t.Errorf("Hello Event was not published")
	}

	gobottest.Assert(t, len(d.Halt()), 0)

	d.On(d.Event(Hello), func(data interface{}) {
		sem <- true
	})
	select {
	case <-sem:
		t.Errorf("Hello Event should not publish after Halt")
	case <-time.After(600 * time.Millisecond):
	}
}
