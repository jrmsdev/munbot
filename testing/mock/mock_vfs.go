// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package mock

import (
	"gobot.io/x/gobot/sysfs"
)

type MockFilesystem struct {
	*sysfs.MockFilesystem
}

func NewFilesystem(files []string) *MockFilesystem {
	return &MockFilesystem{sysfs.NewMockFilesystem(files)}
}
