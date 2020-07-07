// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package munbot

import (
	"time"

	"gobot.io/x/gobot"

	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
)

type Robot struct {
	*gobot.Robot
	cfg  *config.Robot
	name string
	conn *Adaptor
	dev  *Driver
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
	name := cfg.Name
	bot := gobot.NewRobot(name)
	bot.AutoRun = cfg.AutoRun
	conn := NewAdaptor(name)
	bot.AddConnection(conn)
	dev := NewDriver(conn)
	bot.AddDevice(dev)
	r := &Robot{bot, cfg, name, conn, dev}
	bot.Work = r.Work
	return r
}

func (r *Robot) Work() {
	log.Printf("Robot %s work", r.name)
	//~ TODO: create new config value TimeDuration and use time.ParseDuration
	//~ gobot.Every(cfg.Ping.EverySecond.Value()*time.Second, r.ping)
	gobot.Every(15*time.Second, r.ping)
}

func (r *Robot) ping() {
	log.Debugf("%s ping<->%s", r.name, r.dev.Ping())
}
