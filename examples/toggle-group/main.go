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

	tg(root)

	return
}

func tg(root cmdr.OptCmd) {
	// toggle-group

	c := cmdr.NewSubCmd().Titles("toggle-group", "tg").
		Description("soundex test").
		Group("Test").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			selectedMuxType := cmdr.GetStringR("toggle-group.mux-type")
			fmt.Printf("Flag 'echo' = %v\n", cmdr.GetBoolR("toggle-group.echo"))
			fmt.Printf("Flag 'gin' = %v\n", cmdr.GetBoolR("toggle-group.gin"))
			fmt.Printf("Flag 'gorilla' = %v\n", cmdr.GetBoolR("toggle-group.gorilla"))
			fmt.Printf("Flag 'iris' = %v\n", cmdr.GetBoolR("toggle-group.iris"))
			fmt.Printf("Flag 'std' = %v\n", cmdr.GetBoolR("toggle-group.std"))
			fmt.Printf("Toggle Group 'mux-type' = %v\n", selectedMuxType)
			return
		}).
		AttachTo(root)

	cmdr.NewBool(false).Titles("echo", "echo").Description("using 'echo' mux").ToggleGroup("mux-type").Group("Mux").AttachTo(c)
	cmdr.NewBool(false).Titles("gin", "gin").Description("using 'gin' mux").ToggleGroup("mux-type").Group("Mux").AttachTo(c)
	cmdr.NewBool(false).Titles("gorilla", "gorilla").Description("using 'gorilla' mux").ToggleGroup("mux-type").Group("Mux").AttachTo(c)
	cmdr.NewBool(true).Titles("iris", "iris").Description("using 'iris' mux").ToggleGroup("mux-type").Group("Mux").AttachTo(c)
	cmdr.NewBool(false).Titles("std", "std").Description("using standardlib http mux mux").ToggleGroup("mux-type").Group("Mux").AttachTo(c)
}

const (
	appName   = "toggle-group"
	copyright = "toggle-group is an effective devops tool"
	desc      = "toggle-group is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "toggle-group is an effective devops tool. It make an demo application for `cmdr`."
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
