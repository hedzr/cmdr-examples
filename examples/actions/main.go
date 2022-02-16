// Copyright © 2020 Hedzr Yeh.

package main

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/cmdr-addons/pkg/plugins/trace"
	cmdr_examples "github.com/hedzr/cmdr-examples"
	"github.com/hedzr/cmdr/tool"
	"gopkg.in/hedzr/errors.v2"
)

func main() {
	Entry()
}

func Entry() {
	if err := cmdr.Exec(buildRootCmd(),
		trace.WithTraceEnable(true),
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
		AddGlobalPreAction(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Println("# global pre-action 1")
			return
		}).
		AddGlobalPreAction(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Println("# global pre-action 2")
			return
		}).
		AddGlobalPostAction(func(cmd *cmdr.Command, args []string) {
			fmt.Println("# global post-action 1")
		}).
		AddGlobalPostAction(func(cmd *cmdr.Command, args []string) {
			fmt.Println("# global post-action 2")
		}).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	cmdr.NewBool(false).
		Titles("enable-ueh", "ueh").
		Description("Enables the unhandled exception handler?").
		AttachTo(root)

	cmdrPanic(root)
	cmdrSoundex(root)

	return
}

func prd(key string, val interface{}, format string, params ...interface{}) {
	fmt.Printf("         [--%v] %v, %v\n", key, val, fmt.Sprintf(format, params...))
}

func cmdrPanic(root cmdr.OptCmd) {
	// panic test

	pa := cmdr.NewSubCmd().
		Titles("panic-test", "pa").
		Description("test panic inside cmdr actions", "").
		Group("Test").
		AttachTo(root)

	val := 9
	zeroVal := zero

	cmdr.NewSubCmd().
		Titles("division-by-zero", "dz").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Println(val / zeroVal)
			return
		}).
		AttachTo(pa)

	cmdr.NewSubCmd().
		Titles("panic", "pa").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			panic(9)
			return
		}).
		AttachTo(pa)

}

func cmdrSoundex(root cmdr.OptCmd) {

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

const (
	appName   = "actions"
	copyright = "actions is an effective devops tool"
	desc      = "actions is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "actions is an effective devops tool. It make an demo application for `cmdr`."
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
