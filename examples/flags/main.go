// Copyright © 2020 Hedzr Yeh.

package main

import (
	"github.com/hedzr/cmdr"
	cmdr_examples "github.com/hedzr/cmdr-examples"
	"github.com/hedzr/cmdr-examples/examples/flags/cmd"
	"github.com/hedzr/log"
	"github.com/hedzr/logex/build"
)

func main() {
	Entry()
}

func Entry() {
	// logConfig := log.NewLoggerConfigWith(true, "logrus", "debug")
	if err := cmdr.Exec(buildRootCmd(),
		cmdr.WithLogx(build.New(cmdr.NewLoggerConfigWith(true, "logrus", "debug"))),
		cmdr.WithOptionMergeModifying(func(keyPath string, value, oldVal interface{}) {
			cmdr.Logger.Debugf("-> -> onOptionMergeModifying: %q - %v -> %v", keyPath, oldVal, value)
		}),
		cmdr.WithOptionModifying(func(keyPath string, value, oldVal interface{}) {
			cmdr.Logger.Debugf("-> -> onOptionModifying: %q - %v -> %v", keyPath, oldVal, value)
		}),
	); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func buildRootCmd() (rootCmd *cmdr.RootCommand) {
	root := cmdr.Root(appName, cmdr_examples.Version).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	cmdr.NewBool().
		Titles("test-bool", "tb").
		Description("test-bool flag").
		OnSet(func(keyPath string, value interface{}) {
			cmdr.Logger.Debugf("-> -> onSet: %q <- %v", "care", value)
		}).
		AttachTo(root)
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
