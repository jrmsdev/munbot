// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package vfs

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/munbot/master/testing/assert"
	"github.com/munbot/master/testing/require"
	"github.com/munbot/master/testing/suite"
)

func TestMockSuite(t *testing.T) {
	suite.Run(t, &MockSuite{Suite: suite.New()})
}

type MockSuite struct {
	*suite.Suite
	fs      *MockFilesystem
	assert  *assert.Assertions
	require *require.Assertions
}

func (s *MockSuite) SetupTest() {
	s.assert = assert.New(s.T())
	s.require = require.New(s.T())
	s.fs = NewMockFilesystem()
}

func (s *MockSuite) TearDownTest() {
	s.assert = nil
	s.require = nil
	s.fs = nil
}

func (s *MockSuite) TestClose() {
	fh := s.fs.Add("testing.txt")
	s.assert.False(fh.closed)
	err := fh.Close()
	s.require.NoError(err, "close error")
	s.assert.True(fh.closed)
}

func (s *MockSuite) TestRead() {
	fh := s.fs.Add("testing.txt")
	defer fh.Close()
	fh.WriteString("testing")
	blob, err := ioutil.ReadAll(fh)
	s.require.NoError(err, "read error")
	s.assert.Equal([]byte("testing"), blob, "file content")
}

func (s *MockSuite) TestReadWithError() {
	var b []byte
	fh := s.fs.Add("testing.txt")
	s.fs.WithReadError = true
	_, err := fh.Read(b)
	s.require.EqualError(err, "mock read error", "read error")
}

func (s *MockSuite) TestReadClosedError() {
	var b []byte
	fh := s.fs.Add("testing.txt")
	fh.Close()
	_, err := fh.Read(b)
	s.require.EqualError(err, "mock file is closed", "read closed")
}

func (s *MockSuite) TestWrite() {
	fh := s.fs.Add("testing.txt")
	defer fh.Close()
	n, err := fh.Write([]byte("testing"))
	s.require.NoError(err, "write error")
	s.assert.Equal(len("testing"), n, "file written bytes")
}

func (s *MockSuite) TestWriteWithError() {
	s.fs.WithWriteError = true
	fh := s.fs.Add("testing.txt")
	_, err := fh.Write([]byte("testing"))
	s.require.EqualError(err, "mock write error", "write error")
}

func (s *MockSuite) TestWriteClosedError() {
	fh := s.fs.Add("testing.txt")
	fh.Close()
	_, err := fh.Write([]byte("testing"))
	s.require.EqualError(err, "mock file is closed", "write closed")
}

func (s *MockSuite) TestWriteString() {
	fh := s.fs.Add("testing.txt")
	defer fh.Close()
	n, err := fh.WriteString("testing")
	s.require.NoError(err, "write string error")
	s.assert.Equal(len("testing"), n, "file written bytes")
}

func (s *MockSuite) TestWriteStringWithError() {
	s.fs.WithWriteError = true
	fh := s.fs.Add("testing.txt")
	_, err := fh.WriteString("testing")
	s.require.EqualError(err, "mock write error", "write string error")
}

func (s *MockSuite) TestWriteStringClosedError() {
	fh := s.fs.Add("testing.txt")
	fh.Close()
	_, err := fh.WriteString("testing")
	s.require.EqualError(err, "mock file is closed", "write string closed")
}

func (s *MockSuite) TestAddTwice() {
	s.require.Equal(len(s.fs.root), 0, "fs root init size")
	s.fs.Add("testing.txt")
	s.require.Equal(len(s.fs.root), 1, "fs root add size")
	s.fs.Add("testing.txt")
	s.require.Equal(len(s.fs.root), 1, "fs root add dup size")
}

func (s *MockSuite) TestOpenFile() {
	s.fs.Add("testing.txt")
	_, err := s.fs.OpenFile("testing.txt", 0, 0)
	s.require.NoError(err, "open file")
}

func (s *MockSuite) TestOpenFileNotFound() {
	_, err := s.fs.OpenFile("testing.txt", 0, 0)
	s.require.Error(err, "open file not found")
	s.require.True(os.IsNotExist(err), "open file not found error type")
}

func (s *MockSuite) TestOpenFileWithError() {
	s.fs.WithOpenError = true
	s.fs.Add("testing.txt")
	_, err := s.fs.OpenFile("testing.txt", 0, 0)
	s.require.EqualError(err, "mock open error", "open file error")
}

func (s *MockSuite) TestStat() {
	s.fs.Add("testing.txt")
	i, err := s.fs.Stat("testing.txt")
	s.require.NoError(err, "fs stat")
	s.assert.Equal(i.Size(), int64(0), "fs stat size")
}

func (s *MockSuite) TestStatNotFound() {
	_, err := s.fs.Stat("testing.txt")
	s.require.Error(err, "fs stat not found")
	s.require.True(os.IsNotExist(err), "stat not found error type")
}

func mockTempFileError() tempFileFunc {
	return func(dir, pat string) (*os.File, error) {
		return nil, errors.New("mock tempfile error")
	}
}

func (s *MockSuite) TestStatTempFileError() {
	s.fs.tempfile = mockTempFileError()
	s.fs.Add("testing.txt")
	_, err := s.fs.Stat("testing.txt")
	s.require.EqualError(err, "mock tempfile error", "fs stat tempfile error")
}
