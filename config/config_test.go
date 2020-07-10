// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"path/filepath"
	"testing"
)

var tdir = filepath.FromSlash("./testdata")

func TestRead(t *testing.T) {
	c := New("config.json", tdir)
	c.Read()
}
