// Copyright © 2020 Hedzr Yeh.

package main

import (
	"fmt"
	"github.com/hedzr/cmdr"
	cmdrexamples "github.com/hedzr/cmdr-examples"
	"gopkg.in/hedzr/errors.v2"
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
	root := cmdr.Root(appName, cmdrexamples.Version).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	cmdrPanic(root)
	cmdrKbPrint(root)

	return
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

func cmdrKbPrint(root cmdr.OptCmd) {

	// kb-print

	kb := cmdr.NewSubCmd().Titles("kb-print", "kb").
		Description("kilobytes test", "test kibibytes' input,\nverbose long descriptions here.").
		Group("Test").
		Examples(`
$ {{.AppName}} kb --size 5kb
  5kb = 5,120 bytes
$ {{.AppName}} kb --size 8T
  8TB = 8,796,093,022,208 bytes
$ {{.AppName}} kb --size 1g
  1GB = 1,073,741,824 bytes
		`).
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Printf("Got size: %v (literal: %v)\n\n", cmdr.GetKibibytesR("kb-print.size"), cmdr.GetStringR("kb-print.size"))
			return
		}).
		AttachTo(root)

	cmdr.NewString("1k").Titles("size", "s").
		Description("max message size. Valid formats: 2k, 2kb, 2kB, 2KB. Suffixes: k, m, g, t, p, e.", "").
		Group("").
		AttachTo(kb)

}

const (
	appName   = "kilo-bytes"
	copyright = "kilo-bytes is an effective devops tool"
	desc      = "kilo-bytes is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "kilo-bytes is an effective devops tool. It make an demo application for `cmdr`."
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
