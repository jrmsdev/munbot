// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"github.com/jrmsdev/munbot/config"
	"github.com/jrmsdev/munbot/log"
)

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

func (m *Master) Main(cfg *config.Master) {
	log.Debug(cfg.Name)

	if cfg.Api.Enable.IsTrue() {
		m.api.Start(cfg.Api)
	} else {
		log.Warn("master api is disabled")
	}

	if cfg.Robot.Enable.IsTrue() {
		log.Printf("Add robot %s", cfg.Robot.Name)
		bot := NewRobot(cfg.Robot)
		m.AddRobot(bot.Robot)
	} else {
		log.Warn("master robot is disabled")
	}

	//~ conn := adaptor.New()
	//~ dev := driver.New(conn)
	//~ work := func() {
	//~ dev.On(dev.Event(driver.Hello), func(data interface{}) {
	//~ fmt.Println(data)
	//~ })
	//~ gobot.Every(3000*time.Millisecond, func() {
	//~ fmt.Println(dev.Ping())
	//~ })
	//~ }
	//~ robot := gobot.NewRobot(
	//~ "munbot",
	//~ []gobot.Connection{conn},
	//~ []gobot.Device{dev},
	//~ work,
	//~ )

	//~ m.AddRobot(robot)
	m.Start()
}
