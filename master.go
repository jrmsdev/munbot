// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"gobot.io/x/gobot"

	"github.com/munbot/master/config"
	"github.com/munbot/master/internal/master/api"
	"github.com/munbot/master/log"
)

type Master struct {
	*gobot.Master
	api *api.Api
}

func New() *Master {
	m := gobot.NewMaster()
	return &Master{m, api.New(m)}
}

//~ master.AddCommand("custom_gobot_command",
//~ func(params map[string]interface{}) interface{} {
//~ return "This command is attached to the mcp!"
//~ })

//~ hello := master.AddRobot(gobot.NewRobot("hello"))

//~ hello.AddCommand("hi_there", func(params map[string]interface{}) interface{} {
//~ return fmt.Sprintf("This command is attached to the robot %v", hello.Name)
//~ })

//~ master.Start()
//~ }

func (m *Master) Main(cfg *config.Master) error {
	log.Printf("Name %s", cfg.Name)
	setupInfo()
	if cfg.Api.Enable {
		if err := m.api.Start(cfg.Api); err != nil {
			return err
		}
	} else {
		log.Warn("master api is disabled")
	}
	if cfg.Robot.Enable {
		log.Printf("Add robot %s", cfg.Robot.Name)
		bot := NewRobot(cfg.Robot)
		m.AddRobot(bot.Robot)
	} else {
		log.Warn("master robot is disabled")
	}
	return m.Start()
}
