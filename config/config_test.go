// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"path/filepath"
	"testing"

	"github.com/munbot/master/testing/require"
)

var tdir string = filepath.FromSlash("./testdata")
var defcfg *Munbot = &Munbot{
	Master: &Master{
		Enable: true,
		Name:   "munbot",
	},
}

func TestDefaults(t *testing.T) {
	require := require.New(t)
	c := New("empty")
	c.SetDefaults()
	require.Equal(defcfg, c.Munbot, "default config")
}

func TestRead(t *testing.T) {
	c := New("config.json", tdir)
	c.Read()
}
