// Copyright © 2020 Hedzr Yeh.

package main

import (
	"fmt"
	"github.com/hedzr/cmdr"
	cmdrexamples "github.com/hedzr/cmdr-examples"
	"gopkg.in/AlecAivazis/survey.v1"
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
	root := cmdr.Root(appName, cmdrexamples.Version).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	cmdr.NewBool(false).Titles("ueh", "enable-ueh", "Enables the unhandled exception handler?")

	prompts(root)

	return
}

func prompts(root cmdr.OptCmd) {
	// prompts

	// root.NewSubCommand("prompts", "p", "pb").
	// 	Description("progress bar test").
	// 	Group("Test").
	// 	TailPlaceholder("[text1, text2, ...]").
	// 	Action(prompts1)

	root.Action(prompts1)
}

func prompts1(cmd *cmdr.Command, args []string) (err error) {
	color := ""
	prompt := &survey.Select{
		Message: "Choose a color:",
		Options: []string{"red", "blue", "green"},
	}
	err = survey.AskOne(prompt, &color, survey.Required) // , survey.WithKeepFilter(true))

	var days []string
	promptTitles := &survey.MultiSelect{
		Message: "What days do you prefer:",
		Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
	}
	err = survey.AskOne(promptTitles, &days, survey.Required)

	return
}

const (
	appName   = "interactive-prompt"
	copyright = "interactive-prompt is an effective devops tool"
	desc      = "interactive-prompt is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "interactive-prompt is an effective devops tool. It make an demo application for `cmdr`."
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
