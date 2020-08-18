// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"time"

	"gobot.io/x/gobot"
)

type Error struct {
	Msg string `json:"error"`
}

func (m *Robot) addCommands(mbot *gobot.Master) {
	mbot.AddCommand("status", m.status)
	mbot.AddCommand("exit", m.exit)
}

type Status struct {
	Born   string `json:"born"`
	Uptime string `json:"uptime"`
	State  string `json:"state"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Die    string `json:"die,omitempty"`
}

func (m *Robot) newStatus() *Status {
	status := "ok"
	err := ""
	if m.err != nil {
		status = "error"
		err = m.err.Error()
	}
	return &Status{
		Born:   m.born.String(),
		Uptime: m.Uptime().String(),
		State:  m.state,
		Status: status,
		Error:  err,
		Die:    "",
	}
}

func (m *Robot) status(args map[string]interface{}) interface{} {
	return m.newStatus()
}

func (m *Robot) exit(args map[string]interface{}) interface{} {
	if m.exitc == nil {
		return Error{"nothing to do here"}
	}
	s := m.newStatus()
	s.Die = time.Now().String()
	s.Status = "exit"
	m.exitc <- true
	return s
}
