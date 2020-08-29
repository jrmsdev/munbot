// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package master implements the master robot.
package master

import (
	"os"
	"os/signal"
	"sync"

	"gobot.io/x/gobot"

	"github.com/munbot/master/v0/adaptor"
	"github.com/munbot/master/v0/driver/sshd"
	"github.com/munbot/master/v0/log"
	"github.com/munbot/master/v0/robot"
	"github.com/munbot/master/v0/utils/net"
)

// Robot works around a gobot.Master.
type Robot struct {
	*gobot.Master
	robots *gobot.Robots
	rw     *sync.RWMutex
	done   chan bool
	bot    *robot.Munbot
}

// New creates a new master robot.
func New() *Robot {
	return &Robot{
		Master: gobot.NewMaster(),
		robots: new(gobot.Robots),
		rw:     new(sync.RWMutex),
		done:   make(chan bool, 0),
	}
}

// AddRobot adds a new robot.
func (m *Robot) AddRobot(r *gobot.Robot) *gobot.Robot {
	m.rw.Lock()
	defer m.rw.Unlock()
	r.AutoRun = false
	if m.robots == nil {
		m.Master.AddRobot(r)
	} else {
		*m.robots = append(*m.robots, r)
	}
	return r
}

// Start starts the core runtime and then the gobot master robot.
func (m *Robot) Start() error {
	log.Print("Start master robot.")
	m.rw.Lock()
	m.Master.AutoRun = false
	if err := m.Master.Start(); err != nil {
		m.rw.Unlock()
		return log.Error(err)
	}
	for _, r := range *m.robots {
		r.AutoRun = false
		m.Master.AddRobot(r)
	}
	m.robots = nil
	var err error
	autorun := false
	start := func(r *gobot.Robot) {
		log.Debugf("start %s robot", r.Name)
		if x := r.Start(autorun); x != nil {
			err = log.Error(x)
		}
	}
	m.Robots().Each(start)
	m.rw.Unlock()
	if err != nil {
		return err
	}
	log.Debug("wait for stop signal")
	<-m.done
	return nil
}

// Stop stops the gobot master robot and then the core runtime.
func (m *Robot) Stop() error {
	log.Print("Stop master robot.")
	var err error
	if m.Running() {
		log.Debug("stop master robot")
		if x := m.Master.Stop(); x != nil {
			err = log.Error(x)
		}
	} else {
		log.Debug("master robot not running!")
	}
	stop := func(r *gobot.Robot) {
		if r.Running() {
			log.Debugf("stop %s robot", r.Name)
			if x := r.Stop(); x != nil {
				log.Errorf("Stop %s robot: %v", r.Name, x)
				err = x
			}
		}
	}
	m.Master.Robots().Each(stop)
	m.done <- true
	return err
}

// Run runs the robot's main loop.
func (m *Robot) Run() error {
	log.Debug("run...")
	c := make(chan os.Signal, 0)
	log.Debug("trap os interrupt signal")
	signal.Notify(c, os.Interrupt)
	go func(c <-chan os.Signal) {
		select {
		case <-c:
			log.Info("OS interrupt!")
			m.Stop()
			return
		case <-m.done:
			log.Debug("trap os interrupt done..")
			return
		}
	}(c)
	log.Debug("add core munbot")
	m.bot = robot.New(adaptor.New(m.Master))
	m.AddRobot(m.bot.Gobot())
	log.Debug("start...")
	if err := m.Start(); err != nil {
		log.Debugf("start error: %v", err)
		log.Error("Abort start!")
		return m.Stop()
	}
	log.Debug("exit")
	return nil
}

// Addr returns the ssh server network address information.
func (m *Robot) Addr() *net.Addr {
	if m.bot != nil {
		d := m.bot.Device("munbot.sshd").(*sshd.Driver)
		if d != nil {
			return d.Addr()
		}
	}
	return nil
}
