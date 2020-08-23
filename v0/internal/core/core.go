// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package core implements the runtime utils.
package core

import (
	"errors"

	"github.com/munbot/master/v0/internal/api"
	"github.com/munbot/master/v0/internal/ssh"
	"github.com/munbot/master/v0/log"
	"github.com/munbot/master/v0/utils/lock"
)

var LockPanic error = errors.New("core: lock failed")
var UnlockPanic error = errors.New("core: unlock failed")
var NotLockedPanic error = errors.New("core: not locked")

var k *lock.Locker
var locked bool

func init() {
	k = lock.New()
}

func Lock() {
	if !k.TryLock(nil) {
		panic(LockPanic)
	}
	locked = true
}

func Unlock() {
	defer func() {
		if r := recover(); r != nil {
			msg, ok := r.(string)
			if ok && msg == "Unlock() failed" {
				panic(UnlockPanic)
			} else {
				panic(r)
			}
		}
	}()
	k.Unlock()
	locked = false
}

func checkLock() {
	log.Debug("check lock")
	if !locked {
		panic(NotLockedPanic)
	}
}

type ApiServer api.Server

func NewApiServer() ApiServer {
	log.Debug("new api server")
	checkLock()
	return api.NewServer()
}

type SSHServer ssh.Server

func NewSSHServer() SSHServer {
	log.Debug("new ssh server")
	checkLock()
	return ssh.NewServer()
}
