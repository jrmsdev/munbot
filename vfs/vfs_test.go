// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package vfs

import (
	"os"
	"testing"

	"gobot.io/x/gobot/sysfs"
)

var defFS Filesystem
var testFS Filesystem

func init() {
	defFS = fs
	testFS = sysfs.NewMockFilesystem([]string{"stat.txt"})
	SetFilesystem(testFS)
}

func TestDefaultFS(t *testing.T) {
	switch typ := defFS.(type) {
	case *sysfs.NativeFilesystem:
	default:
		t.Fatalf("wrong default filesystem: %T", typ)
	}
}

func TestStat(t *testing.T) {
	i, err := Stat("stat.txt")
	if err != nil {
		t.Fatalf("stat error: %v", err)
	}
	if i.Size() != 0 {
		t.Errorf("stat file size: '%d' - expected: 0", i.Size())
	}
}

func TestOpen(t *testing.T) {
	fh, err := Open("stat.txt")
	if err != nil {
		t.Fatalf("open error: %v", err)
	}
	if fh.Name() != "stat.txt" {
		t.Errorf("filename: '%s' - expected: stat.txt", fh.Name())
	}
}

func TestOpenError(t *testing.T) {
	_, err := Open("stat.err")
	if err == nil {
		t.Fatalf("open did not fail")
	}
}

func TestOpenFile(t *testing.T) {
	fh, err := OpenFile("stat.txt", os.O_RDONLY, 0)
	if err != nil {
		t.Fatalf("open error: %v", err)
	}
	if fh.Name() != "stat.txt" {
		t.Errorf("filename: '%s' - expected: stat.txt", fh.Name())
	}
}

func TestOpenFileError(t *testing.T) {
	_, err := OpenFile("stat.err", os.O_RDONLY, 0)
	if err == nil {
		t.Fatalf("open did not fail")
	}
}
