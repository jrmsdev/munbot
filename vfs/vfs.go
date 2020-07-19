// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package vfs implements the utilities to interact with the filesystem.
package vfs

import (
	"io"
	"os"
)

// File handler interface.
type File interface {
	io.ReadWriteCloser
	io.StringWriter
}

// Filesystem handler interface.
type Filesystem interface {
	OpenFile(string, int, os.FileMode) (File, error)
	Stat(string) (os.FileInfo, error)
}

var fs Filesystem

// DefaultFilesystem is set as NativeFilesystem at init time.
var DefaultFilesystem Filesystem

func init() {
	DefaultFilesystem = new(NativeFilesystem)
	fs = DefaultFilesystem
}

// SetFilesystem sets the fs manager to the provided one.
func SetFilesystem(newfs Filesystem) {
	fs = nil
	fs = newfs
}

// OpenFile calls current fs manager OpenFile method.
func OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	return fs.OpenFile(name, flag, perm)
}

// Stat calls current fs manager Stat method.
func Stat(name string) (os.FileInfo, error) {
	return fs.Stat(name)
}

// Open opens the named file as read only.
func Open(name string) (File, error) {
	return fs.OpenFile(name, os.O_RDONLY, 0)
}

// Create opens the named file with read and write access, creates it if it does
// not exists already and truncates its content if it exists. Permissions before
// umask are set as 0660.
func Create(name string) (File, error) {
	return fs.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0660)
}
