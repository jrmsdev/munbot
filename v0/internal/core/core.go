// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package core implements the runtime utils.
package core

import (
	"errors"

	"github.com/munbot/master/v0/log"
	"github.com/munbot/master/v0/utils/lock"
)

var ErrLock error = errors.New("core: lock failed")
var UnlockPanic error = errors.New("core: unlock failed")

var k *lock.Locker

func init() {
	k = lock.New()
}

func Lock() error {
	if !k.TryLock(nil) {
		return log.Error(ErrLock)
	}
	return nil
}

func Unlock() {
	var err error
	defer func() {
		if r := recover(); r != nil {
			msg, ok := r.(string)
			if !ok {
				panic(r)
			}
			if msg == "Unlock() failed" {
				err = UnlockPanic
			}
		}
	}()
	k.Unlock()
	if err != nil {
		panic(err)
	}
}
