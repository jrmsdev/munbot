// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package vfs

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
)

type MockFile struct {
	*bytes.Buffer
	fs *MockFilesystem
	closed bool
}

func (f MockFile) Close() error {
	f.closed = true
	f.Buffer.Reset()
	return nil
}

func (f MockFile) Read(b []byte) (int, error) {
	if f.fs.WithReadError {
		return 0, errors.New("mock read error")
	}
	if f.closed {
		return 0, errors.New("mock file is closed")
	}
	return f.Buffer.Read(b)
}

func (f MockFile) Write(b []byte) (int, error) {
	if f.fs.WithWriteError {
		return 0, errors.New("mock write error")
	}
	if f.closed {
		return 0, errors.New("mock file is closed")
	}
	return f.Buffer.Write(b)
}

func (f MockFile) WriteString(s string) (int, error) {
	if f.fs.WithWriteError {
		return 0, errors.New("mock write error")
	}
	if f.closed {
		return 0, errors.New("mock file is closed")
	}
	return f.Buffer.WriteString(s)
}

type MockFilesystem struct {
	root map[string]File
	WithOpenError bool
	WithReadError bool
	WithWriteError bool
}

func NewMockFilesystem(files ...string) *MockFilesystem {
	fs := &MockFilesystem{}
	fs.root = make(map[string]File)
	for i := range files {
		fn := files[i]
		fs.Add(fn)
	}
	return fs
}

func (fs *MockFilesystem) Add(filename string) *MockFile {
	_, found := fs.root[filename]
	if found {
		fs.root[filename] = nil
	}
	fs.root[filename] = &MockFile{new(bytes.Buffer), fs, false}
	return fs.root[filename].(*MockFile)
}

func (fs *MockFilesystem) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	if fs.root == nil {
		return nil, fs.notfound(name)
	}
	fh, found := fs.root[name]
	if !found {
		return nil, fs.notfound(name)
	}
	if fs.WithOpenError {
		return nil, errors.New("mock open error")
	}
	return fh, nil
}

func (fs *MockFilesystem) Stat(name string) (os.FileInfo, error) {
	if fs.root == nil {
		return nil, fs.notfound(name)
	}
	_, found := fs.root[name]
	if !found {
		return nil, fs.notfound(name)
	}
	fh, err := ioutil.TempFile("", name)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	defer os.Remove(fh.Name())
	s, serr := fh.Stat()
	if serr != nil {
		return nil, err
	}
	return s, nil
}

func (fs *MockFilesystem) notfound(name string) error {
	_, err := os.Stat(name + ".mock-notfound")
	return err
}
