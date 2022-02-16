// Copyright © 2020 Hedzr Yeh.

package cmd

import (
	"fmt"
	"github.com/hedzr/cmdr"
	cmdr_examples "github.com/hedzr/cmdr-examples"
)

func buildRootCmd() (rootCmd *cmdr.RootCommand) {

	// var cmd *Command

	// cmdr.Root("aa", "1.0.1").
	// 	Header("sds").
	// 	NewSubCommand().
	// 	Titles("ms", "microservice").
	// 	Description("", "").
	// 	Group("").
	// 	Action(func(cmd *cmdr.Command, args []string) (err error) {
	// 		return
	// 	})

	// root

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
		// Header("fluent - test for cmdr - no version - hedzr").
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	// root.NewSubCommand("", "go112113").
	// 	Description("test build tags for go1.13 or later and go1.12 and below", "").
	// 	Group("").Action(func(cmd *cmdr.Command, args []string) (err error) {
	// 	// go112113.Fate()
	// 	return nil
	// })

	cmdrMoreCommandsForTest(root)
	kvCommand(root)
	msCommand(root)

	return
}

const (
	appName   = "my-service"
	copyright = "my-service is an effective devops tool"
	desc      = "my-service is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "my-service is an effective devops tool. It make an demo application for `cmdr`."
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
