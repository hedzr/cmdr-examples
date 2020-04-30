/*
 * Copyright © 2019 Hedzr Yeh.
 */

package dex

import (
	"github.com/hedzr/cmdr"
	"net"
	"os"
)

// Daemon interface should be implemented when you are using `daemon.Enable()`.
type Daemon interface {
	// Config() (config *service.Config)
	
	OnRun(prog *Program, stopCh, doneCh chan struct{}, listener net.Listener) (err error)
	OnStop(prog *Program) (err error)
	OnReload(prog *Program)
	OnStatus(prog *Program, p *os.Process) (err error)
	OnInstall(prog *Program) (err error)
	OnUninstall(prog *Program) (err error)

	// OnReadConfigFromCommandLine(root *cmdr.RootCommand)
	OnPrepare(prog *Program, root *cmdr.RootCommand) (err error)
}

// HotReloadable enables hot-restart/hot-reload feature
type HotReloadable interface {
	OnHotReload(prog *Program) (err error)
}

var daemonImpl Daemon
