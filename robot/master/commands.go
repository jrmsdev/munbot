// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"time"

	"gobot.io/x/gobot"
)

type Error struct {
	Msg string `json:"error,omitempty"`
}

func (m *Robot) addCommands(mbot *gobot.Master) {
	mbot.AddCommand("status", m.status)
	mbot.AddCommand("exit", m.exit)
}

type Status struct {
	Uptime string `json:"uptime"`
	State  string `json:"state"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func (m *Robot) newStatus() *Status {
	status := "ok"
	err := ""
	if m.err != nil {
		status = "error"
		err = m.err.Error()
	}
	return &Status{
		Uptime: time.Since(m.born).String(),
		State:  m.state,
		Status: status,
		Error:  err,
	}
}

func (m *Robot) status(args map[string]interface{}) interface{} {
	return m.newStatus()
}

func (m *Robot) exit(args map[string]interface{}) interface{} {
	if m.exitc == nil {
		return Error{"nothing to do here"}
	}
	m.exitc <- true
	return m.newStatus()
}
