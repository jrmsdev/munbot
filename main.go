// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	//~ "fmt"
	//~ "html"
	//~ "net/http"
	//~ "time"

	//~ "gobot.io/x/gobot"
	//~ "gobot.io/x/gobot/api"

	//~ "github.com/jrmsdev/munbot/adaptor"
	//~ "github.com/jrmsdev/munbot/driver"
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
	log.Debugf("master main %s", cfg.Name)

	if cfg.Api.Enable.IsTrue() {
		m.api.Start(cfg.Api)
	} else {
		log.Debug("master api is disabled")
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
