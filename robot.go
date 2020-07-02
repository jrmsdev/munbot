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
	conn gobot.Connection
	dev gobot.Device
}

//~ conn := adaptor.New()
//~ dev := driver.New(conn)
//~ work := func() {
//~ 	dev.On(dev.Event(driver.Hello), func(data interface{}) {
//~ 		fmt.Println(data)
//~ 	})
//~ 	gobot.Every(3000*time.Millisecond, func() {
//~ 		fmt.Println(dev.Ping())
//~ 	})
//~ }
//~ robot := gobot.NewRobot(
//~ 	"munbot",
//~ 	[]gobot.Connection{conn},
//~ 	[]gobot.Device{dev},
//~ 	work,
//~ )

func NewRobot(cfg *config.Robot) *Robot {
	name := cfg.Name.String()
	bot := gobot.NewRobot(name)
	bot.AutoRun = cfg.AutoRun.Value()
	conn := NewAdaptor(name)
	bot.AddConnection(conn)
	dev := NewDriver(conn)
	bot.AddDevice(dev)
	r := &Robot{bot, name, conn, dev}
	bot.Work = r.Work
	return r
}

func (r *Robot) Work() {
	log.Printf("Robot %s work", r.name)
}
