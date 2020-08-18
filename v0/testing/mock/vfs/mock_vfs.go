// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package vfs wraps the mock fs from vfs package.
package vfs

import (
	"github.com/munbot/master/v0/vfs"
)

// MockFilesystem wraps vfs.MockFilesystem.
type MockFilesystem struct {
	*vfs.MockFilesystem
}

// NewFilesystem returns a new mock filesystem with the supplied filenames
// created in the root tree, if any...
func NewMockFilesystem(files ...string) *MockFilesystem {
	return &MockFilesystem{vfs.NewMockFilesystem(files...)}
}

// SetFilesystem sets the current filesystem manager to the provided one.
func SetFilesystem(fs vfs.Filesystem) {
	vfs.SetFilesystem(fs)
}

// SetDefaultFilesystem sets the current fs manager to vfs.DefaultFilesystem,
// which should be vfs.NativeFilesystem.
func SetDefaultFilesystem() {
	vfs.SetFilesystem(vfs.DefaultFilesystem)
}
