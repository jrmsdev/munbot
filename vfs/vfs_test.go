// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package vfs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/munbot/master/testing/assert"
	"github.com/munbot/master/testing/require"
	"github.com/munbot/master/testing/suite"
)

var defFS Filesystem
var testFS Filesystem

func init() {
	defFS = fs
	testFS = NewMockFilesystem("stat.txt")
}

func TestDefaultFS(t *testing.T) {
	switch typ := defFS.(type) {
	case *NativeFilesystem:
	default:
		t.Fatalf("wrong default filesystem: %T", typ)
	}
}

type Suite struct {
	*suite.Suite
}

func (s *Suite) SetupTest() {
	SetFilesystem(testFS)
}

func (s *Suite) TearDownTest() {
	SetFilesystem(defFS)
}

func (s *Suite) TestStat() {
	require := require.New(s.T())
	assert := assert.New(s.T())
	i, err := Stat("stat.txt")
	require.NoError(err, "stat error")
	assert.Equal(int64(0), i.Size(), "file size")
}

func (s *Suite) TestOpen() {
	require := require.New(s.T())
	_, err := Open("stat.txt")
	require.NoError(err, "open error")
}

func (s *Suite) TestOpenError() {
	require := require.New(s.T())
	_, err := Open("stat.err")
	require.Error(err, "open error")
}

func (s *Suite) TestOpenFile() {
	require := require.New(s.T())
	_, err := OpenFile("stat.txt", os.O_RDONLY, 0)
	require.NoError(err, "open error")
}

func (s *Suite) TestOpenFileError() {
	require := require.New(s.T())
	_, err := OpenFile("stat.err", os.O_RDONLY, 0)
	require.Error(err, "open error")
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{suite.New()})
}

func TestNativeOpenFile(t *testing.T) {
	require := require.New(t)
	fs := new(NativeFilesystem)
	fn := filepath.FromSlash("./testdata/native.test")
	fh, err := fs.OpenFile(fn, os.O_RDONLY, 0)
	require.NoError(err, "open file")
	require.Implements((*File)(nil), fh, "file interface")
}

func TestNativeOpenFileError(t *testing.T) {
	require := require.New(t)
	fs := new(NativeFilesystem)
	fn := filepath.FromSlash("./testdata/nofile.test")
	_, err := fs.OpenFile(fn, os.O_RDONLY, 0)
	require.Error(err, "open file error")
	require.True(os.IsNotExist(err), "open file not found error type")
}

func TestNativeStat(t *testing.T) {
	require := require.New(t)
	fs := new(NativeFilesystem)
	fn := filepath.FromSlash("./testdata/native.test")
	fi, err := fs.Stat(fn)
	require.NoError(err, "fs stat")
	require.Equal(int64(len("testing\n")), fi.Size(), "stat file size")
}

func TestNativeStatError(t *testing.T) {
	require := require.New(t)
	fs := new(NativeFilesystem)
	fn := filepath.FromSlash("./testdata/nofile.test")
	_, err := fs.Stat(fn)
	require.Error(err, "stat error")
	require.True(os.IsNotExist(err), "stat file not found error type")
}
