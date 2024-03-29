// Copyright © 2020 Hedzr Yeh.

package main

import (
	"fmt"
	"github.com/hedzr/cmdr"
	cmdr_examples "github.com/hedzr/cmdr-examples"
	"github.com/hedzr/cmdr/tool"
	"github.com/sirupsen/logrus"
)

func main() {
	Entry()
}

func Entry() {
	if err := cmdr.Exec(buildRootCmd()); err != nil {
		logrus.Fatalf("error: %+v", err)
	}
}

func buildRootCmd() (rootCmd *cmdr.RootCommand) {
	root := cmdr.Root(appName, cmdr_examples.Version).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	soundex(root)
	panicTest(root)
	nested(root)
	grouped(root)

	return
}

func grouped(root cmdr.OptCmd) {
	// grouped sub-commands

	d1 := cmdr.NewSubCmd().Titles("sorted", "sorted").
		Description("[grouped] Tags operations").
		Group("Grouped").
		AttachTo(root)

	cmdr.NewSubCmd().Titles("demo-1", "d1").
		Description("[sub][sub] check-in sub").
		Group("g001.Group 1").
		AttachTo(d1)
	cmdr.NewSubCmd().Titles("demo-2", "d2").
		Description("[sub][sub] check-in sub").
		Group("g001.Group 1").
		AttachTo(d1)

	cmdr.NewSubCmd().Titles("cmd-1", "c1").
		Description("[sub][sub] check-in sub").
		Group("gz99.Group 99").
		AttachTo(d1)
	cmdr.NewSubCmd().Titles("cmd-2", "c2").
		Description("[sub][sub] check-in sub").
		Group("gz99.Group 99").
		AttachTo(d1)
	cmdr.NewSubCmd().Titles("cmd-3", "c3").
		Description("[sub][sub] check-in sub").
		Group("gz99.Group 99").
		AttachTo(d1)

}

func nested(root cmdr.OptCmd) {
	// nested sub-commands

	d1 := cmdr.NewSubCmd().Titles("demo-1", "d1").
		Description("[sub] check-in sub").
		Group("Nested").
		AttachTo(root)
	d2 := cmdr.NewSubCmd().Titles("demo-2", "d2").
		Description("[sub][sub] check-in sub").
		Group("Nested").
		AttachTo(d1)
	cmdr.NewSubCmd().Titles("demo-3", "d3").
		Description("[sub][sub][sub] check-in sub").
		Group("Nested").
		AttachTo(d2)

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
		}).
		AttachTo(pa)

	cmdr.NewSubCmd().Titles("panic", "pa").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			panic(9)
			return
		}).
		AttachTo(pa)
}

const (
	appName   = "subcommands"
	copyright = "subcommands is an effective devops tool"
	desc      = "subcommands is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "subcommands is an effective devops tool. It make an demo application for `cmdr`."
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
