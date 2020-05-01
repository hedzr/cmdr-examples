// Copyright © 2020 Hedzr Yeh.

package cmd

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/cmdr-addons/pkg/plugins/dex"
	"github.com/hedzr/cmdr-examples/examples/service/svr"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

func Entry() {
	// logrus.SetLevel(logrus.DebugLevel)
	// logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	// defer func() {
	// 	fmt.Println("defer caller")
	// 	if err := recover(); err != nil {
	// 		fmt.Printf("recover success. error: %v", err)
	// 	}
	// }()

	if err := cmdr.Exec(buildRootCmd(),
		// To disable internal commands and flags, uncomment the following codes
		// cmdr.WithBuiltinCommands(false, false, false, false, false),

		dex.WithDaemon(svr.NewDaemon(), dex.WithCommandsModifier(modifier), dex.WithLoggerForward(true)),
		// server.WithCmdrDaemonSupport(),
		// server.WithCmdrHook(),

		// integrate with logex library
		cmdr.WithLogex(cmdr.DebugLevel),
		cmdr.WithLogexPrefix("logger"),

		cmdr.WithHelpTabStop(51),

		cmdr.WithWatchMainConfigFileToo(true),
		// cmdr.WithNoWatchConfigFiles(false),
		cmdr.WithOptionMergeModifying(func(keyPath string, value, oldVal interface{}) {
			logrus.Debugf("%%-> -> %q: %v -> %v", keyPath, oldVal, value)
			if strings.HasSuffix(keyPath, ".mqtt.server.stats.enabled") {
				// mqttlib.FindServer().EnableSysStats(!vxconf.ToBool(value))
			}
			if strings.HasSuffix(keyPath, ".mqtt.server.stats.log.enabled") {
				// mqttlib.FindServer().EnableSysStatsLog(!vxconf.ToBool(value))
			}
		}),
		cmdr.WithOptionModifying(func(keyPath string, value, oldVal interface{}) {
			logrus.Infof("%%-> -> %q: %v -> %v", keyPath, oldVal, value)
		}),

		// sample.WithSampleCmdrOption(),
		// trace.WithTraceEnable(true),

		cmdr.WithUnknownOptionHandler(onUnknownOptionHandler),
		cmdr.WithUnhandledErrorHandler(onUnhandledErrorHandler),

		optAddTraceOption,
		optAddServerExtOption,

		cmdr.WithOnSwitchCharHit(onSwitchCharHit),
		cmdr.WithOnPassThruCharHit(onPassThruCharHit),
	); err != nil {
		logrus.Fatalf("error: %+v", err)
	}
}

func onSwitchCharHit(parsed *cmdr.Command, switchChar string, args []string) (err error) {
	if parsed != nil {
		fmt.Printf("the last parsed command is %q - %q\n", parsed.GetTitleNames(), parsed.Description)
	}
	fmt.Printf("SwitchChar FOUND: %v\nremains: %v\n\n", switchChar, args)
	return // cmdr.ErrShouldBeStopException
}

func onPassThruCharHit(parsed *cmdr.Command, switchChar string, args []string) (err error) {
	if parsed != nil {
		fmt.Printf("the last parsed command is %q - %q\n", parsed.GetTitleNames(), parsed.Description)
	}
	fmt.Printf("PassThrough flag FOUND: %v\nremains: %v\n\n", switchChar, args)
	return // ErrShouldBeStopException
}

func onUnknownOptionHandler(isFlag bool, title string, cmd *cmdr.Command, args []string) (fallbackToDefaultDetector bool) {
	return true
}

func onUnhandledErrorHandler(err interface{}) {
	// debug.PrintStack()
	// pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	if e, ok := err.(error); ok {
		logrus.Errorf("%+v", e)
	} else {
		logrus.Errorf("%+v", err)
		dumpStacks()
	}
}

func dumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, false)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===\n", buf)
	// fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===\n", errors.DumpStacksAsString(true))
}

var optAddTraceOption, optAddServerExtOption cmdr.ExecOption

func init() {
	// attaches `--trace` to root command
	optAddTraceOption = cmdr.WithXrefBuildingHooks(func(root *cmdr.RootCommand, args []string) {
		cmdr.NewBool(false).
			Titles("tr", "trace").
			Description("enable trace mode for tcp/mqtt send/recv data dump", "").
			AttachToRoot(root)
	}, nil)

	// the following statements show you how to attach an option to a sub-command
	optAddServerExtOption = cmdr.WithXrefBuildingHooks(func(root *cmdr.RootCommand, args []string) {
		serverCmd := cmdr.FindSubCommandRecursive("server", nil)
		serverStartCmd := cmdr.FindSubCommand("start", serverCmd)
		cmdr.NewInt(5100).
			Titles("vnc", "vnc-server").
			Description("start as a vnc server (just a demo)", "").
			Placeholder("PORT").
			AttachTo(cmdr.NewCmdFrom(serverStartCmd))
	}, nil)
}
