// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	//~ "html"
	//~ "net/http"
	"time"

	"gobot.io/x/gobot"
	//~ "gobot.io/x/gobot/api"

	"github.com/jrmsdev/munbot/adaptor"
	"github.com/jrmsdev/munbot/driver"
)

//~ func main() {
//~ master := gobot.NewMaster()

//~ a := api.NewAPI(master)
//a.AddHandler(api.BasicAuth("munbot", "lalala"))
//~ a.Debug()

//~ a.AddHandler(func(w http.ResponseWriter, r *http.Request) {
//~ fmt.Fprintf(w, "Hello, %q \n", html.EscapeString(r.URL.Path))
//~ })
//~ a.Start()

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

func main() {
	master := gobot.NewMaster()

	conn := adaptor.New()
	dev := driver.New(conn)

	work := func() {
		dev.On(dev.Event(driver.Hello), func(data interface{}) {
			fmt.Println(data)
		})
		gobot.Every(1200*time.Millisecond, func() {
			fmt.Println(dev.Ping())
		})
	}

	robot := gobot.NewRobot(
		"munbot",
		[]gobot.Connection{conn},
		[]gobot.Device{dev},
		work,
	)

	master.AddRobot(robot)
	master.Start()
}
