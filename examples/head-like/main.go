// Copyright © 2020 Hedzr Yeh.

package main

import (
	"fmt"
	"github.com/hedzr/cmdr"
	cmdr_examples "github.com/hedzr/cmdr-examples"
)

func main() {
	Entry()
}

func Entry() {
	if err := cmdr.Exec(buildRootCmd()); err != nil {
		fmt.Printf("error: %+v\n", err)
	}
}

func buildRootCmd() (rootCmd *cmdr.RootCommand) {
	root := cmdr.Root(appName, cmdr_examples.Version).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	headLike(root)

	return
}

func headLike(root cmdr.OptCmd) {
	// head-like

	cmdr.NewInt(1).
		Titles("lines", "n").
		Description("how many lines to be processed").
		HeadLike(true, 0, 1000).
		AttachTo(root)
	
	root.Action(func(cmd *cmdr.Command, args []string) (err error) {
		fmt.Printf("Got --lines: %v\n", cmdr.GetIntR("lines"))
		return
	})
}

const (
	appName   = "head-like"
	copyright = "head-like is an effective devops tool"
	desc      = "head-like is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "head-like is an effective devops tool. It make an demo application for `cmdr`."
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
