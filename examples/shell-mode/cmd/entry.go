package cmd

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/cmdr-addons/pkg/plugins/shell"
	cmdrexamples "github.com/hedzr/cmdr-examples"
)

func Entry() {
	if err := cmdr.Exec(
		buildRootCmd(),
		shell.WithShellModule(),
	); err != nil {
		fmt.Printf("error: %+v\n", err)
	}
}

func buildRootCmd() (rootCmd *cmdr.RootCommand) {
	root := cmdr.Root(appName, cmdrexamples.Version).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	cmdrMoreCommandsForTest(root)
	kvCommand(root)
	msCommand(root)

	return
}

const (
	appName   = "shell-mode"
	copyright = "shell-mode is an effective devops tool"
	desc      = "shell-mode is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "shell-mode is an effective devops tool. It make an demo application for `cmdr`."
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
