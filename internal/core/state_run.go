// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package core

import (
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/munbot/master/log"
)

type failmsg struct {
	name string
	err  error
}

var _ State = &SRun{}

type SRun struct {
	m     Machine
	rt    *Mem
	wg    *sync.WaitGroup
	fail  chan failmsg
	wait  time.Duration
	exit  chan bool
	osint chan os.Signal
}

func newRun(m Machine, rt *Mem) State {
	return &SRun{
		m:     m,
		rt:    rt,
		wg:    new(sync.WaitGroup),
		fail:  make(chan failmsg, 1),
		wait:  300 * time.Millisecond,
		exit:  make(chan bool, 1),
		osint: make(chan os.Signal, 1),
	}
}

func (s *SRun) Init() error {
	return ErrInit
}

func (s *SRun) Configure() error {
	return ErrConfigure
}

func (s *SRun) Start() error {
	log.Print("Start...")
	// start master robot
	log.Print("Start master robot...")
	s.rt.Master.ExitNotify(s.exit)
	s.wg.Add(1)
	go func(wg *sync.WaitGroup, fail chan failmsg) {
		defer wg.Done()
		log.Debug("robot start...")
		if err := s.rt.Master.Start(); err != nil {
			log.Error(err)
			fail <- failmsg{"robot", err}
		}
	}(s.wg, s.fail)
	// start api server
	time.Sleep(s.wait)
	log.Print("Start master api...")
	s.wg.Add(1)
	go func(wg *sync.WaitGroup, fail chan failmsg) {
		defer wg.Done()
		log.Debug("api start...")
		if err := s.rt.Api.Start(); err != nil {
			log.Error(err)
			fail <- failmsg{"api", err}
		}
	}(s.wg, s.fail)
	// start console server
	time.Sleep(s.wait)
	log.Print("Start master console...")
	s.wg.Add(1)
	go func(wg *sync.WaitGroup, fail chan failmsg) {
		defer wg.Done()
		log.Debug("console start...")
		if err := s.rt.Console.Start(); err != nil {
			log.Error(err)
			fail <- failmsg{"console", err}
		}
	}(s.wg, s.fail)
	return nil
}

func (s *SRun) Run() error {
	log.Print("Run...")
	var fail failmsg
	abort := false
	signal.Notify(s.osint, os.Interrupt)
LOOP:
	for {
		select {
		case fail = <-s.fail:
			log.Info("fail...")
			break LOOP
		case <-s.exit:
			log.Info("master exit...")
			abort = true
			break LOOP
		case <-s.osint:
			log.Info("os interrupt...")
			abort = true
			break LOOP
		default:
			time.Sleep(s.wait)
			if !s.rt.Master.Running() {
				log.Info("master robot is not running...")
				select {
				case fail = <-s.fail:
					log.Debug("there was a failure before abort")
				default:
				}
				abort = true
				break LOOP
			}
		}
	}
	// close fail channel just in case...
	// maybe not a good idea? but at least it makes them panic...
	defer close(s.fail)
	if fail.err != nil {
		log.Debugf("FAIL: core %s: %s", fail.name, fail.err)
		abort = true
	}
	if abort {
		log.Debug("ABORT!")
		if err := s.m.Abort(); err != nil {
			return err
		}
	}
	return fail.err
}

func (s *SRun) Stop() error {
	log.Print("Stop...")
	var xerr error
	// stop console
	log.Print("Stop master console...")
	if err := s.rt.Console.Stop(); err != nil {
		xerr = log.Error(err)
	}
	// stop api
	log.Print("Stop master api...")
	if err := s.rt.Api.Stop(); err != nil {
		xerr = log.Error(err)
	}
	// stop robot
	if s.rt.Master.Running() {
		log.Print("Stop master robot...")
		if err := s.rt.Master.Stop(); err != nil {
			xerr = log.Error(err)
		}
	}
	// wait for them...
	log.Debug("wait for them to finish...")
	s.wg.Wait()
	if xerr != nil {
		return xerr
	}
	return s.m.SetState(Halt)
}

func (s *SRun) Halt() error {
	return ErrHalt
}
