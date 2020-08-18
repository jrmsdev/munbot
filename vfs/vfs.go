// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package vfs implements the utilities to interact with the filesystem.
package vfs

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var dirPerm os.FileMode = 0770
var filePerm os.FileMode = 0660

// File handler interface.
type File interface {
	io.ReadWriteCloser
	io.StringWriter
}

// Filesystem handler interface.
type Filesystem interface {
	OpenFile(string, int, os.FileMode) (File, error)
	Stat(string) (os.FileInfo, error)
	Mkdir(string) error
	MkdirAll(string) error
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

// StatHash returns a hash string from stat information of named file.
func StatHash(name string) (string, error) {
	st, err := Stat(name)
	if err != nil {
		return "", err
	}
	return hash(fileInfoString(st)), nil
}

func fileInfoString(st os.FileInfo) string {
	return fmt.Sprintf("Name:%s:Size:%d:Mode:%o:Time:%s:Dir:%v",
		st.Name(),
		st.Size(),
		st.Mode(),
		st.ModTime(),
		st.IsDir(),
	)
}

func hash(s string) string {
	buf := new(bytes.Buffer)
	buf.WriteString(s)
	h := sha256.Sum256(buf.Bytes())
	return fmt.Sprintf("%x", h)
}

// Mkdir creates a new directory on current filesystem.
func Mkdir(path string) error {
	return fs.Mkdir(path)
}

// MkdirAll creates a new directory on current filesystem.
func MkdirAll(path string) error {
	return fs.MkdirAll(path)
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

// Exist checks if the named file exists on current filesystem.
func Exist(name string) bool {
	_, err := fs.Stat(name)
	if err == nil {
		return true
	}
	return false
}

// ReadFile calls vfs.Open(name) and then ioutil.ReadAll.
func ReadFile(name string) ([]byte, error) {
	fh, err := Open(name)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	return ioutil.ReadAll(fh)
}
