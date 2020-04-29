// Copyright © 2020 Hedzr Yeh.

package dex

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/kardianos/service"
	"github.com/sirupsen/logrus"
	"log"
	"runtime"
)

// WithDaemon enables daemon plugin:
// - add daemon commands and sub-commands: start/run, stop, restart/reload, status, install/uninstall
// - pidfile
// -
func WithDaemon(daemonImplObject Daemon,
	modifier func(daemonServerCommand *cmdr.Command) *cmdr.Command,
	preAction func(cmd *cmdr.Command, args []string) (err error),
	postAction func(cmd *cmdr.Command, args []string),
	opts ...Opt,
) cmdr.ExecOption {

	pd = &Program{
		Config: &service.Config{
			Name:        "the-daemon",
			DisplayName: "The Daemon",
			Description: "The Daemon/Service here",
		},
		daemon:  daemonImplObject,
		Service: nil,
		Logger:  nil,
		Command: nil,
		// exit:    make(chan struct{}),
		// done:    make(chan struct{}),
	}

	pd.Config = daemonImplObject.Config()
	if len(pd.Config.Arguments) == 0 {
		pd.Config.Arguments = []string{
			"server", "run", "--in-daemon",
		}
	}

	for _, opt := range opts {
		opt()
	}

	return func(w *cmdr.ExecWorker) {
		w.AddOnBeforeXrefBuilding(func(root *cmdr.RootCommand, args []string) {

			if modifier != nil {
				root.SubCommands = append(root.SubCommands, modifier(DaemonServerCommand))
			} else {
				root.SubCommands = append(root.SubCommands, DaemonServerCommand)
			}

			// prefix = strings.Join(append(cmdr.RxxtPrefix, "server"), ".")
			// prefix = "server"

			attachPreAction(root, preAction)
			attachPostAction(root, postAction)

			if err := prepare(daemonImplObject, root); err != nil {
				logrus.Fatal(err)
			}

		})
	}
}

// Opt is functional option type
type Opt func()

func WithServiceConfig(config *service.Config) Opt {
	return func() {
		pd.Config = config
	}
}

// // WithOnGetListener returns tcp/http listener for daemon hot-restarting
// func WithOnGetListener(fn func() net.Listener) Opt {
// 	return func() {
// 		impl.SetOnGetListener(fn)
// 	}
// }

func attachPostAction(root *cmdr.RootCommand, postAction func(cmd *cmdr.Command, args []string)) {
	if root.PostAction != nil {
		savedPostAction := root.PostAction
		root.PostAction = func(cmd *cmdr.Command, args []string) {
			if postAction != nil {
				postAction(cmd, args)
			}
			pidfile.Destroy()
			savedPostAction(cmd, args)
			return
		}
	} else {
		root.PostAction = func(cmd *cmdr.Command, args []string) {
			if postAction != nil {
				postAction(cmd, args)
			}
			pidfile.Destroy()
			return
		}
	}
}

func attachPreAction(root *cmdr.RootCommand, preAction func(cmd *cmdr.Command, args []string) (err error)) {
	if root.PreAction != nil {
		savedPreAction := root.PreAction
		root.PreAction = func(cmd *cmdr.Command, args []string) (err error) {
			pidfile.Create(cmd)
			logger.Setup(cmd)
			if err = savedPreAction(cmd, args); err != nil {
				return
			}
			if preAction != nil {
				err = preAction(cmd, args)
			}
			return
		}
	} else {
		root.PreAction = func(cmd *cmdr.Command, args []string) (err error) {
			pidfile.Create(cmd)
			logger.Setup(cmd)
			if preAction != nil {
				err = preAction(cmd, args)
			}
			return
		}
	}
}

func prepare(daemonImplObject Daemon, root *cmdr.RootCommand) (err error) {

	pd.Service, err = service.New(pd, pd.Config)
	if err != nil {
		return
	}

	errs := make(chan error, 5)
	pd.Logger, err = pd.Service.Logger(errs)
	if err != nil {
		return
	}

	// pd.daemon.OnReadConfigFromCommandLine(root)

	if err = pd.daemon.OnPrepare(pd, root); err != nil {
		return
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	pd.Logger.Info("daemonex prepared.")
	return
}

func daemonStart(cmd *cmdr.Command, args []string) (err error) {
	pd.Command, pd.Args = cmd, args
	pd.InvokedInDaemon = cmdr.GetBoolRP("server.start", "in-daemon")
	foreground := cmdr.GetBoolRP("server.start", "foreground")
	pd.Logger.Infof("daemonStart: foreground: %v, in-daemon: %v, hit: %v", foreground, pd.InvokedInDaemon, cmd.GetHitStr())
	// ctx := impl.GetContext(Command, Args, daemonImpl, onHotReloading)
	if foreground || cmd.GetHitStr() == "run" {
		err = run(cmd, args)
	} else {
		err = runAsDaemon(cmd, args)
	}
	pd.Logger.Infof("daemonStart END: err: %v", err)
	if pd.InvokedInDaemon || runtime.GOOS == "windows" {
		err = pd.Service.Run()
	}
	return
}

// func onHotReloading(ctx *impl.Context) (err error) {
// 	// if hr, ok := ctx.DaemonImpl.(HotReloadable); ok {
// 	// 	err = hr.OnHotReload(ctx)
// 	// }
// 	return
// }

func runAsDaemon(cmd *cmdr.Command, args []string) (err error) {
	err = service.Control(pd.Service, "start")
	if err != nil {
		pd.Logger.Errorf("Valid actions: %q\n", service.ControlAction)
		return // log.Fatal(err)
	}
	return
}

func run(cmd *cmdr.Command, args []string) (err error) {
	if runtime.GOOS != "windows" {
		pd.run()
		err = pd.err
	}
	
	// defer func() {
	// 	if service.Interactive() {
	// 		err = pd.Stop(pd.service)
	// 	} else {
	// 		err = pd.service.Stop()
	// 	}
	// }()
	// err = pd.daemon.OnRun(Command, Args, nil, nil, nil)
	return
}

func daemonStop(cmd *cmdr.Command, args []string) (err error) {
	pd.Command, pd.Args = cmd, args
	err = service.Control(pd.Service, "stop")
	if err != nil {
		pd.Logger.Errorf("Valid actions: %q\n", service.ControlAction)
		return // log.Fatal(err)
	}
	return
}

func daemonRestart(cmd *cmdr.Command, args []string) (err error) {
	pd.Command, pd.Args = cmd, args
	err = service.Control(pd.Service, "restart")
	if err != nil {
		pd.Logger.Errorf("Valid actions: %q\n", service.ControlAction)
		return // log.Fatal(err)
	}

	// getContext(Command, Args)
	//
	// p, err := daemonCtx.Search()
	// if err != nil {
	// 	fmt.Printf("%v is stopped.\n", Command.GetRoot().AppName)
	// } else {
	// 	if err = p.Signal(syscall.SIGHUP); err != nil {
	// 		return
	// 	}
	// }

	// ctx := impl.GetContext(Command, Args, daemonImpl, onHotReloading)
	// impl.Reload(Command.GetRoot().AppName, ctx)
	return
}

func daemonHotRestart(cmd *cmdr.Command, args []string) (err error) {
	pd.Command, pd.Args = cmd, args
	// ctx := impl.GetContext(Command, Args, daemonImpl, onHotReloading)
	// impl.HotReload(Command.GetRoot().AppName, ctx)
	return
}

func daemonStatus(cmd *cmdr.Command, args []string) (err error) {
	pd.Command, pd.Args = cmd, args
	pd.Logger.Infof("Args: %v", args)
	// err = service.Control(pd.service, "status")
	// if err != nil {
	// 	logrus.Errorf("Valid actions: %q\n", service.ControlAction)
	// 	return // log.Fatal(err)
	// }
	var st service.Status
	st, err = pd.Service.Status()
	var sst string
	switch st {
	case service.StatusStopped:
		sst = "Stopped"
	case service.StatusRunning:
		sst = "Running"
	default:
		if err == service.ErrNotInstalled {
			sst = "Not Installed"
		} else {
			sst = "Unknown"
		}
	}
	fmt.Printf("Status: %v\n", sst)
	if pd.daemon != nil {
		err = pd.daemon.OnStatus(pd, nil)
	}

	// getContext(Command, Args)
	//
	// p, err := daemonCtx.Search()
	// if err != nil {
	// 	fmt.Printf("%v is stopped.\n", Command.GetRoot().AppName)
	// } else {
	// 	fmt.Printf("%v is running as %v.\n", Command.GetRoot().AppName, p.Pid)
	// }
	//
	// if daemonImpl != nil {
	// 	err = daemonImpl.OnStatus(&Context{Context: *daemonCtx}, Command, p)
	// }

	// ctx := impl.GetContext(Command, Args, daemonImpl, onHotReloading)
	// present, process := impl.FindDaemonProcess(ctx)
	// if present && daemonImpl != nil {
	// 	err = daemonImpl.OnStatus(ctx, Command, process)
	// }
	return
}

func daemonInstall(cmd *cmdr.Command, args []string) (err error) {
	pd.Command, pd.Args = cmd, args
	err = service.Control(pd.Service, "install")
	if err != nil {
		pd.Logger.Errorf("Valid actions: %q\n", service.ControlAction)
		return // log.Fatal(err)
	}

	// ctx := impl.GetContext(Command, Args, daemonImpl, onHotReloading)
	// 
	// err = runInstaller(Command, Args)
	// if err != nil {
	// 	return
	// }
	// if daemonImpl != nil {
	// 	err = daemonImpl.OnInstall(ctx /*&Context{Context: *daemonCtx}*/, Command, Args)
	// }
	return
}

func daemonUninstall(cmd *cmdr.Command, args []string) (err error) {
	pd.Command, pd.Args = cmd, args
	err = service.Control(pd.Service, "uninstall")
	if err != nil {
		pd.Logger.Errorf("Valid actions: %q\n", service.ControlAction)
		return // log.Fatal(err)
	}

	// ctx := impl.GetContext(Command, Args, daemonImpl, onHotReloading)
	// 
	// err = runUninstaller(Command, Args)
	// if err != nil {
	// 	return
	// }
	// if daemonImpl != nil {
	// 	err = daemonImpl.OnUninstall(ctx /*&Context{Context: *daemonCtx}*/, Command, Args)
	// }
	return
}

//
//
//
