// Copyright © 2020 Hedzr Yeh.

package main

import (
	"fmt"
	"github.com/hedzr/cmdr"
	cmdrexamples "github.com/hedzr/cmdr-examples"
	"github.com/hedzr/cmdr/tool"
	"github.com/hedzr/log"
)

func main() {
	if err := cmdr.Exec(buildRootCmd(),
		cmdr.WithLogx(log.NewStdLogger()),
	); err != nil {
		cmdr.Logger.Printf("error: %+v\n", err)
	}
}

func buildRootCmd() (rootCmd *cmdr.RootCommand) {
	root := cmdr.
		Root(appName, cmdrexamples.Version).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()
	soundex(root)
	return
}

func soundex(root cmdr.OptCmd) {
	cmdr.NewSubCmd().Titles("soundex", "snd", "sndx", "sound").
		Description("soundex test").
		Group("Test").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			for ix, s := range args {
				fmt.Printf("%5d. %s => %s\n", ix, s, tool.Soundex(s))
			}
			return
		}).AttachTo(root)
}

const (
	appName   = "getting-start"
	copyright = "getting-start is an effective devops tool"
	desc      = "getting-start is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "getting-start is an effective devops tool. It make an demo application for `cmdr`."
	examples  = `
$ {{.AppName}} gen shell [--bash|--zsh|--auto]
  generate bash/shell completion scripts
$ {{.AppName}} gen man
  generate linux man page 1
$ {{.AppName}} --help
  show help screen.
`
)
