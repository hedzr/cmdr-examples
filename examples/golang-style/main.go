// Copyright © 2020 Hedzr Yeh.

package main

import (
	"fmt"
	"github.com/hedzr/cmdr"
)

func main() {
	if err := cmdr.Exec(buildRootCmd(),
	); err != nil {
		fmt.Printf("error: %+v\n", err)
	}
}

func buildRootCmd() (rootCmd *cmdr.RootCommand) {
	root := cmdr.Root(appName, version).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	cmdr.NewInt(28).
		Titles("age", "age").
		Description("Input Your Age").
		AttachTo(root)

	cmdr.NewInt(1234).
		Titles("flagname", "flagname").
		Description("Just for demo").
		AttachTo(root)

	cmdr.NewString("male").
		Titles("gender", "gender").
		Description("Input Your Gender").
		AttachTo(root)

	cmdr.NewString("nick").
		Titles("name", "name").
		Description("Input Your Name").
		AttachTo(root)

	return
}

const (
	version   = "1.0.0"
	appName   = "golang-style"
	copyright = "golang-style is an effective devops tool"
	desc      = "golang-style is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "golang-style is an effective devops tool. It make an demo application for `cmdr`."
	examples  = `
$ {{.AppName}} gen shell [--bash|--zsh|--auto]
  generate bash/shell completion scripts
$ {{.AppName}} gen man
  generate linux man page 1
$ {{.AppName}} --help
  show help screen.
`
)
