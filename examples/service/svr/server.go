/*
 * Copyright © 2019 Hedzr Yeh.
 */

package svr

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/cmdr-examples/examples/service/dex"
	"github.com/kardianos/service"
	"golang.org/x/crypto/acme/autocert"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// NewDaemon creates an `daemon.Daemon` object
func NewDaemon() dex.Daemon {
	return NewDaemonWithConfig(&service.Config{
		Name:        "my-daemon",
		DisplayName: "My Daemon",
		Description: "My Daemon/Service here",
	})
}

func NewDaemonWithConfig(config *service.Config) dex.Daemon {
	d := &daemonImpl{
		// exit:   make(chan struct{}),
		config: config,
		Type:   typeIris,
	}
	return d
}

type daemonImpl struct {
	config *service.Config
	// service service.Service
	// logger  service.Logger
	// cmd     *exec.Cmd
	// exit    chan struct{}

	appTag      string
	certManager *autocert.Manager
	Type        muxType
	mux         *http.ServeMux
	routerImpl  routerMux
	// router      *gin.Engine
	// irisApp     *iris.Application
}

func (d *daemonImpl) Config() (config *service.Config) {
	return d.config
}

func (d *daemonImpl) OnRun(prog *dex.Program, stopCh, doneCh chan struct{}, listener net.Listener) (err error) {
	serverType := cmdr.GetStringR("server.start.Server-Type")

	prog.Logger.Infof("demo daemon OnRun (Server-Type = %q), pid = %v, ppid = %v", serverType, os.Getpid(), os.Getppid())

	if serverType == "h2-server" {
		err = d.onRunHttp2Server(prog, stopCh, doneCh, listener)
		if err == nil {
			err = d.enterLoop(prog, stopCh, doneCh, listener)
		}
		return
	}

	worker(prog, stopCh, doneCh)
	return
}

func worker(prog *dex.Program, stopCh, doneCh chan struct{}) {
	fullExec, errx := exec.LookPath("git")
	if errx != nil {
		prog.Logger.Errorf("Failed to find executable %q: %v", "git --version", errx)
	}

	var args []string = []string{"--version"}
	var env []string

	f1, err := os.OpenFile("/tmp/1.err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		prog.Logger.Warningf("Failed to open std err %q: %v", f1, err)
		return
	}
	defer f1.Close()

	f2, err2 := os.OpenFile("/tmp/1.out", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err2 != nil {
		// logger.Warningf("Failed to open std out %q: %v", p.Stdout, err)
		prog.Logger.Errorf("Failed to open std out %q: %v\n", f2, err)
		return
	}
	defer f2.Close()

	ticker := time.NewTicker(5 * time.Second)
	defer func() {
		ticker.Stop()
		if doneCh != nil && prog.InvokedInDaemon {
			doneCh <- struct{}{}
		}
	}()

LOOP:
	for {
		// time.Sleep(3 * time.Second) // this is work to be done by worker.
		select {
		case <-stopCh:
			break LOOP
		case tc := <-ticker.C:
			cmd := exec.Command(fullExec, args...)
			cmd.Dir = "/tmp"
			cmd.Env = append(os.Environ(), env...)
			cmd.Stdout = f2
			cmd.Stderr = f1

			pwd, _ := os.Getwd()
			prog.Logger.Infof("demo running at %d [dir: %q], inDaemon: %v, tick: %v, OS=%v\n", os.Getpid(), pwd, prog.InvokedInDaemon, tc, runtime.GOOS)
			err = cmd.Run()
			cmd.Wait()
			if !prog.InvokedInDaemon {
				return
			}
		}
	}
}

func (*daemonImpl) OnStop(prog *dex.Program) (err error) {
	prog.Logger.Infof("demo daemon OnStop")
	return
}

func (*daemonImpl) OnReload(prog *dex.Program) {
	prog.Logger.Infof("demo daemon OnReload")
}

func (*daemonImpl) OnStatus(prog *dex.Program, p *os.Process) (err error) {
	fmt.Printf("%v v%v\n", prog.Command.GetRoot().AppName, prog.Command.GetRoot().Version)
	// fmt.Printf("PID=%v\nLOG=%v\n", cxt.PidFileName, cxt.LogFileName)
	return
}

func (*daemonImpl) OnInstall(prog *dex.Program) (err error) {
	prog.Logger.Infof("demo daemon OnInstall")
	return
	// panic("implement me")
}

func (*daemonImpl) OnUninstall(prog *dex.Program) (err error) {
	prog.Logger.Infof("demo daemon OnUninstall")
	return
	// panic("implement me")
}
