// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"gobot.io/x/gobot"

	"github.com/jrmsdev/munbot/config"
	"github.com/jrmsdev/munbot/log"
)

type Robot struct {
	*gobot.Robot
	name string
}

func NewRobot(cfg *config.Robot) *Robot {
	name := cfg.Name.String()
	bot := gobot.NewRobot(name)
	bot.AutoRun = cfg.AutoRun.Value()
	r := &Robot{bot, name}
	bot.Work = r.Work
	return r
}

func (r *Robot) Work() {
	log.Debugf("%s work", r.name)
}
