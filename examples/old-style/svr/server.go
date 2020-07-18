/*
 * Copyright © 2019 Hedzr Yeh.
 */

package svr

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/cmdr/plugin/daemon"
	"github.com/hedzr/cmdr/plugin/daemon/impl"
	"log"
	"net"
	"os"
	"time"
)

// NewDaemon creates an `daemon.Daemon` object
func NewDaemon() daemon.Daemon {
	return &daemonImpl{}
}

type daemonImpl struct {
}

func (*daemonImpl) OnRun(cmd *cmdr.Command, args []string, stopCh, doneCh chan struct{}, listener net.Listener) (err error) {
	log.Printf("demo daemon OnRun, pid = %v, ppid = %v", os.Getpid(), os.Getppid())
	go worker(stopCh, doneCh)
	return
}

func worker(stopCh, doneCh chan struct{}) {
LOOP:
	for {
		time.Sleep(3 * time.Second) // this is work to be done by worker.
		select {
		case <-stopCh:
			break LOOP
		default:
			log.Printf("demo running at %d", os.Getpid())
		}
	}
	doneCh <- struct{}{}
}

func (*daemonImpl) OnStop(cmd *cmdr.Command, args []string) (err error) {
	log.Printf("demo daemon OnStop")
	return
}

func (*daemonImpl) OnReload() {
	log.Printf("demo daemon OnReload")
}

func (*daemonImpl) OnStatus(cxt *impl.Context, cmd *cmdr.Command, p *os.Process) (err error) {
	fmt.Printf("%v v%v\n", cmd.GetRoot().AppName, cmd.GetRoot().Version)
	fmt.Printf("PID=%v\nLOG=%v\n", cxt.PidFileName, cxt.LogFileName)
	return
}

func (*daemonImpl) OnInstall(cxt *impl.Context, cmd *cmdr.Command, args []string) (err error) {
	log.Printf("demo daemon OnInstall")
	return
	// panic("implement me")
}

func (*daemonImpl) OnUninstall(cxt *impl.Context, cmd *cmdr.Command, args []string) (err error) {
	log.Printf("demo daemon OnUninstall")
	return
	// panic("implement me")
}
