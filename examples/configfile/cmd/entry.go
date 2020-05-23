// Copyright © 2020 Hedzr Yeh.

package cmd

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/sirupsen/logrus"
	"gopkg.in/hedzr/errors.v2"
)

func Entry() {
	if err := cmdr.Exec(buildRootCmd(),
		cmdr.WithUnhandledErrorHandler(onUnhandledErrorHandler),
	); err != nil {
		logrus.Fatalf("error: %+v", err)
	}
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
