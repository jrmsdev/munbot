// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package vfs

import (
	"os"
)

type NativeFilesystem struct{}

func (fs *NativeFilesystem) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	return os.OpenFile(name, flag, perm)
}

func (fs *NativeFilesystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
