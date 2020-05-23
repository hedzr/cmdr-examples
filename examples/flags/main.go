// Copyright © 2020 Hedzr Yeh.

package main

import (
	"github.com/hedzr/cmdr"
	cmdr_examples "github.com/hedzr/cmdr-examples"
	"github.com/hedzr/cmdr-examples/examples/flags/cmd"
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

	cmd.AddTags(root)
	cmd.AddFlags(root)

	return
}

const (
	appName   = "flags"
	copyright = "flags is an effective devops tool"
	desc      = "flags is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "flags is an effective devops tool. It make an demo application for `cmdr`."
	examples  = `
$ {{.AppName}} gen shell [--bash|--zsh|--auto]
  generate bash/shell completion scripts
$ {{.AppName}} gen man
  generate linux man page 1
$ {{.AppName}} --help
  show help screen.
`
)
