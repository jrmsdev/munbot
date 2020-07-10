// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"path/filepath"
	"reflect"
	"testing"
)

var tdir string = filepath.FromSlash("./testdata")
var defcfg *Munbot = &Munbot{
	Master: &Master{
		Enable: true,
		Name: "munbot",
	},
}

func TestDefaults(t *testing.T) {
	c := New("empty")
	c.SetDefaults()
	if !reflect.DeepEqual(c.Munbot, defcfg) {
		t.Fatalf("default config: '%#v' - expect: '%#v'", c.Munbot, defcfg)
	}
}

func TestRead(t *testing.T) {
	c := New("config.json", tdir)
	c.Read()
}
