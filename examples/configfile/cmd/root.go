// Copyright © 2020 Hedzr Yeh.

package cmd

import (
	"github.com/hedzr/cmdr"
	cmdr_examples "github.com/hedzr/cmdr-examples"
)

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

	addFlags(root)

	return
}

const (
	appName   = "config-file"
	copyright = "config-file is an effective devops tool"
	desc      = "config-file is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "config-file is an effective devops tool. It make an demo application for `cmdr`."
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
