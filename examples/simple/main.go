// Copyright © 2020 Hedzr Yeh.

package main

import (
	"fmt"
	"github.com/hedzr/cmdr"
	cmdr_examples "github.com/hedzr/cmdr-examples"
	"github.com/hedzr/cmdr/tool"
	"gopkg.in/hedzr/errors.v3"
)

func main() {
	Entry()
}

func Entry() {
	if err := cmdr.Exec(buildRootCmd(),
		cmdr.WithUnhandledErrorHandler(onUnhandledErrorHandler),
	); err != nil {
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
		EnvKeys("ENABLE_UEH").
		Description("Enables the unhandled exception handler?").
		OnSet(func(keyPath string, value interface{}) {
			if value == true {
				panic("unexpected value: true")
			}
		}).
		AttachTo(root)

	soundex(root)
	panicTest(root)

	return
}

func soundex(root cmdr.OptCmd) {
	// soundex

	cmdr.NewSubCmd().Titles("soundex", "snd", "sndx", "sound").
		Description("soundex test").
		Group("Test").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			for ix, s := range args {
				fmt.Printf("%5d. %s => %s\n", ix, s, tool.Soundex(s))
			}
			return
		}).
		AttachTo(root)
}

func panicTest(root cmdr.OptCmd) {
	// panic test

	pa := cmdr.NewSubCmd().Titles("panic-test", "pa").
		Description("test panic inside cmdr actions", "").
		Group("Test").
		AttachTo(root)

	val := 9
	zeroVal := zero

	cmdr.NewSubCmd().Titles("division-by-zero", "dz").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Println(val / zeroVal)
			return
		}).AttachTo(pa)

	cmdr.NewSubCmd().Titles("panic", "pa").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			panic(9)
			return
		}).AttachTo(pa)
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
