// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package vfs

import (
	"os"
	"gobot.io/x/gobot/sysfs"
)

type File interface {
	sysfs.File
	Name() string
}

type fileHandler struct {
	sysfs.File
	name string
}

func newFile(name string, fh sysfs.File) File {
	return &fileHandler{fh, name}
}

func (f *fileHandler) Name() string {
	return f.name
}

type Filesystem sysfs.Filesystem

var fs Filesystem

func init() {
	fs = new(sysfs.NativeFilesystem)
}

func SetFilesystem(newfs Filesystem) {
	fs = nil
	fs = newfs
}

func OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	fh , err := fs.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return newFile(name, fh), nil
}

func Stat(name string) (os.FileInfo, error) {
	return fs.Stat(name)
}

func Open(name string) (File, error) {
	fh, err := fs.OpenFile(name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	return newFile(name, fh), nil
}
