// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package master

import (
	"time"

	"gobot.io/x/gobot"
)

func (m *Robot) addCommands(mbot *gobot.Master) {
	mbot.AddCommand("munbot_status", m.status)
}

type Status struct {
	Uptime string `json:"uptime"`
	State  string `json:"state"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func (m *Robot) status(args map[string]interface{}) interface{} {
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
