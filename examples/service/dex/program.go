// Copyright © 2020 Hedzr Yeh.

package dex

import (
	"github.com/hedzr/cmdr"
	"github.com/hedzr/cmdr-examples/examples/service/svr/sig"
	"github.com/kardianos/service"
	"os"
)

var pd *Program

type Program struct {
	daemon          Daemon
	Config          *service.Config
	Service         service.Service
	Logger          service.Logger
	Command         *cmdr.Command
	Args            []string
	Env             []string
	InvokedInDaemon bool
	// exit            chan struct{}
	// done            chan struct{}
	err error
	// Command     *exec.Cmd
}

func (p *Program) Start(s service.Service) error {
	p.Logger.Infof("xx.pp.start; Args: %v;", os.Args)

	// Start should not block. Do the actual work async.

	go p.run()
	return p.err
}

func (p *Program) run() {
	p.runIt(p.Command, p.Args)
}

func (p *Program) runIt(cmd *cmdr.Command, args []string) {
	p.Logger.Infof("xx.pp.runIt; Args: %v;", os.Args)

	// logger.Info("Starting ", p.DisplayName)

	// if runtime.GOOS == "windows" {
	//	defer func() {
	//		if service.Interactive() {
	//			p.err = p.Stop(p.service)
	//		} else {
	//			p.err = p.service.Stop()
	//		}
	//	}()
	// }

	if p.InvokedInDaemon {
		// go func() {
		stop, done := sig.GetChs()
		p.err = pd.daemon.OnRun(p, stop, done, nil)
		// }()
	} else {
		stop, done := sig.GetChs()
		p.err = pd.daemon.OnRun(p, stop, done, nil)
	}
	return
}

func (p *Program) Stop(s service.Service) (err error) {
	p.Logger.Infof("xx.pp.stop; Args: %v;", os.Args)

	err = pd.daemon.OnStop(p)

	// Stop should not block. Return with a few seconds.
	// <-time.After(time.Second * 13)

	stop, done := sig.GetChs()
	close(stop)
	
	// logger.Info("Stopping ", p.DisplayName)
	// if p.Command.ProcessState.Exited() == false {
	// 	err = p.Command.Process.Kill()
	// }
	if service.Interactive() {
		os.Exit(0)
	}
	close(done)
	return
}
