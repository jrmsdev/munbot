// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package vfs

import (
	"io"
	"os"
)

type File interface {
	io.ReadWriteCloser
	io.StringWriter
}

type Filesystem interface {
	OpenFile(string, int, os.FileMode) (File, error)
	Stat(string) (os.FileInfo, error)
}

var fs Filesystem
var DefaultFilesystem Filesystem

func init() {
	DefaultFilesystem = new(NativeFilesystem)
	fs = DefaultFilesystem
}

func SetFilesystem(newfs Filesystem) {
	fs = nil
	fs = newfs
}

func OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	return fs.OpenFile(name, flag, perm)
}

func Stat(name string) (os.FileInfo, error) {
	return fs.Stat(name)
}

func Open(name string) (File, error) {
	return fs.OpenFile(name, os.O_RDONLY, 0)
}
