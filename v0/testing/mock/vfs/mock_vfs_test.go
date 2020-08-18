// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package vfs

import (
	"testing"

	"github.com/munbot/master/v0/vfs"

	"github.com/munbot/master/v0/testing/require"
)

func TestAssertions(t *testing.T) {
	check := require.New(t)
	fs := NewMockFilesystem()
	check.IsType(new(vfs.MockFilesystem), fs)
	SetFilesystem(fs)
	SetDefaultFilesystem()
}
