// Copyright © 2020 Hedzr Yeh.

package cmd

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/log"
	"github.com/hedzr/logex/build"
	"gopkg.in/hedzr/errors.v2"
)

func Entry() {
	logConfig := log.NewLoggerConfig()
	logConfig.Level = "debug"
	logConfig.Backend = "logrus"
	
	if err := cmdr.Exec(buildRootCmd(),
		cmdr.WithLogx(build.New(logConfig)),
		cmdr.WithUnhandledErrorHandler(onUnhandledErrorHandler),
	); err != nil {
		cmdr.Logger.Fatalf("error: %+v", err)
	}
	
	cmdr.Logger.Debugf("hello")
}

func onUnhandledErrorHandler(err interface{}) {
	if cmdr.GetBoolR("enable-ueh") {
		dumpStacks()
		return
	}

	panic(err)
}

func dumpStacks() {
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===\n", errors.DumpStacksAsString(true))
}
