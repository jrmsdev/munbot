// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package mock is just a handy shorcut (?) to import
// github.com/stretchr/testify/mock functionality.
package mock

import (
	"github.com/stretchr/testify/mock"

	"github.com/munbot/master/vfs"
)

type Mock mock.Mock

// Filesystes is a handy shortcut for vfs.MockFilesystem.
type Filesystem struct {
	*vfs.MockFilesystem
}

// NewFilesystem returns a new mock filesystem with the supplied filenames
// created in the root tree, if any supplied.
func NewFilesystem(files ...string) *Filesystem {
	return &Filesystem{vfs.NewMockFilesystem(files...)}
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
