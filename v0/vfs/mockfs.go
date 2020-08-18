// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package vfs

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
)

// MockFile implements File interface for testing purposes mainly. It's mainly
// a wrapper to bytes.Buffer.
type MockFile struct {
	*bytes.Buffer
	fs     *MockFilesystem
	closed bool
	isdir  bool
}

// Close sets the file as closed so it will raise an error if future read/write
// are tried. It also resets the underlying buffer.
func (f *MockFile) Close() error {
	f.closed = true
	f.Buffer.Reset()
	return nil
}

// Read reads from the buffer. If the filesystem was set to return a read error,
// a mock error will be returned.
func (f *MockFile) Read(b []byte) (int, error) {
	if f.fs.WithReadError {
		return 0, errors.New("mock read error")
	}
	if f.closed {
		return 0, errors.New("mock file is closed")
	}
	return f.Buffer.Read(b)
}

// Write writes to the buffer. If the filesystem was set to return a write error,
// a mock error will be returned.
func (f *MockFile) Write(b []byte) (int, error) {
	if f.fs.WithWriteError {
		return 0, errors.New("mock write error")
	}
	if f.closed {
		return 0, errors.New("mock file is closed")
	}
	return f.Buffer.Write(b)
}

// WriteString writes an string to the buffer. If the filesystem was set to
// return a write error, a mock error will be returned.
func (f *MockFile) WriteString(s string) (int, error) {
	if f.fs.WithWriteError {
		return 0, errors.New("mock write error")
	}
	if f.closed {
		return 0, errors.New("mock file is closed")
	}
	return f.Buffer.WriteString(s)
}

type tempFileFunc func(string, string) (*os.File, error)

// MockFilesystem implements Filesystem interface for testing purposes mainly.
// It can be set to return mock errors on open time (WithOpenError), reads
// (WithReadError) and/or at writes (WithWriteError).
type MockFilesystem struct {
	root           map[string]File
	stat           map[string]os.FileInfo
	tempfile       tempFileFunc
	WithOpenError  bool
	WithReadError  bool
	WithWriteError bool
}

// NewMockFilesystem returns a new mock filesystem with the supplied filenames
// created in the root tree (empty though), if any supplied.
func NewMockFilesystem(files ...string) *MockFilesystem {
	fs := &MockFilesystem{}
	fs.root = make(map[string]File)
	fs.stat = make(map[string]os.FileInfo)
	fs.tempfile = ioutil.TempFile
	for i := range files {
		fn := files[i]
		fs.Add(fn)
	}
	return fs
}

// Mkdir creates a mocking dir path.
func (fs *MockFilesystem) Mkdir(path string) error {
	fs.root[path] = &MockFile{nil, fs, false, true}
	return nil
}

// MkdirAll creates a mocking dir path.
func (fs *MockFilesystem) MkdirAll(path string) error {
	fs.root[path] = &MockFile{nil, fs, false, true}
	return nil
}

// Add adds a new file to the root tree (if it already exists it is silently
// overriden). It returns the new file handler.
func (fs *MockFilesystem) Add(filename string) *MockFile {
	_, found := fs.root[filename]
	if found {
		fs.root[filename] = nil
	}
	fs.root[filename] = &MockFile{new(bytes.Buffer), fs, false, false}
	return fs.root[filename].(*MockFile)
}

// OpenFile opens a file that MUST exist in the root tree. Even if empty.
// Otherwise a proper "file not found" mocked *os.PathError is returned.
// If WithOpenError is set a mock error is returned, even if the filename is not
// in the root tree.
func (fs *MockFilesystem) OpenFile(name string, flag int) (File, error) {
	if fs.WithOpenError {
		return nil, errors.New("mock open error")
	}
	fh, found := fs.root[name]
	if !found {
		return nil, fs.notfound(name)
	}
	return fh, nil
}

// Stat creates a "real" tempfile and returns its stats. The filename returned
// in the stats is NOT the same a the provided name.
func (fs *MockFilesystem) Stat(name string) (os.FileInfo, error) {
	_, found := fs.root[name]
	if !found {
		return nil, fs.notfound(name)
	}
	i, ok := fs.stat[name]
	if ok {
		return i, nil
	}
	fh, err := fs.tempfile("", name)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	defer os.Remove(fh.Name())
	fs.stat[name], err = fh.Stat()
	if err != nil {
		return nil, err
	}
	return fs.stat[name], nil
}

// returns a "real" file not found error.
func (fs *MockFilesystem) notfound(name string) error {
	_, err := os.Stat(name + ".mock-notfound")
	return err
}
