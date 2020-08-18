// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package vfs

import (
	"os"
)

// NativeFilesystem it's just a wrapper for os package functions.
type NativeFilesystem struct{}

// OpenFile calls os.OpenFile.
func (fs *NativeFilesystem) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	return os.OpenFile(name, flag, perm)
}

// Stat calls os.Stat.
func (fs *NativeFilesystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

// Mkdir calls os.Mkdir.
func (fs *NativeFilesystem) Mkdir(path string, perm os.FileMode) error {
	return os.Mkdir(path, perm)
}

// MkdirAll calls os.MkdirAll.
func (fs *NativeFilesystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}
