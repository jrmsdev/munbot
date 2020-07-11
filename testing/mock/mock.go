// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package mock

import (
	"github.com/stretchr/testify/mock"

	"github.com/munbot/master/vfs"
)

type Mock mock.Mock

type Filesystem struct {
	*vfs.MockFilesystem
}

func NewFilesystem(files ...string) *Filesystem {
	return &Filesystem{vfs.NewMockFilesystem(files...)}
}

func SetFilesystem(fs vfs.Filesystem) {
	vfs.SetFilesystem(fs)
}

func SetDefaultFilesystem() {
	vfs.SetFilesystem(vfs.DefaultFilesystem)
}
