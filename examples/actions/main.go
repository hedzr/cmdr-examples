// Copyright © 2020 Hedzr Yeh.

package main

import (
	"fmt"
	"github.com/hedzr/cmdr"
	cmdr_examples "github.com/hedzr/cmdr-examples"
	"github.com/hedzr/cmdr-examples/internal/trace"
	"gopkg.in/hedzr/errors.v2"
)

func main() {
	Entry()
}

func Entry() {
	if err := cmdr.Exec(buildRootCmd(), cmdr.WithUnhandledErrorHandler(onUnhandledErrorHandler)); err != nil {
		fmt.Printf("error: %+v\n", err)
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

func buildRootCmd() (rootCmd *cmdr.RootCommand) {
	root := cmdr.Root(appName, cmdr_examples.Version).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	cmdr.NewBool(false).
		Titles("enable-ueh", "ueh").
		Description("Enables the unhandled exception handler?").
		AttachTo(root)

	soundex(root)
	panicTest(root)

	return
}

func soundex(root cmdr.OptCmd) {
	// soundex

	root.NewSubCommand("soundex", "snd", "sndx", "sound").
		Description("soundex test").
		Group("Test").
		TailPlaceholder("[text1, text2, ...]").
		PreAction(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Printf("[PRE] DebugMode=%v, TraceMode=%v. InDebugging/IsDebuggerAttached=%v\n", cmdr.GetDebugMode(), trace.IsEnabled(), cmdr.InDebugging())
			for ix, s := range args {
				fmt.Printf("[PRE] %5d. %s\n", ix, s)
			}

			fmt.Printf("[PRE] Debug=%v, Trace=%v\n", cmdr.GetBoolR("debug"), cmdr.GetBoolR("trace"))

			// return nil to be continue, 
			// return cmdr.ErrShouldBeStopException to stop the following actions without error
			// return other errors for application purpose
			return
		}).
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			for ix, s := range args {
				fmt.Printf("[ACTION] %5d. %s\n", ix, s)
				fmt.Printf("%5d. %s => %s\n", ix, s, cmdr.Soundex(s))
			}
			return
		}).
		PostAction(func(cmd *cmdr.Command, args []string) {
			for ix, s := range args {
				fmt.Printf("[POST] %5d. %s\n", ix, s)
			}
		})
}

func panicTest(root cmdr.OptCmd) {
	// panic test

	pa := root.NewSubCommand("panic-test", "pa").
		Description("test panic inside cmdr actions", "").
		Group("Test")

	val := 9
	zeroVal := zero

	pa.NewSubCommand("division-by-zero", "dz").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Println(val / zeroVal)
			return
		})

	pa.NewSubCommand("panic", "pa").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			panic(9)
			return
		})
}

const (
	appName   = "simple"
	copyright = "simple is an effective devops tool"
	desc      = "simple is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "simple is an effective devops tool. It make an demo application for `cmdr`."
	examples  = `
$ {{.AppName}} gen shell [--bash|--zsh|--auto]
  generate bash/shell completion scripts
$ {{.AppName}} gen man
  generate linux man page 1
$ {{.AppName}} --help
  show help screen.
`
	overview = ``

	zero = 0
)
