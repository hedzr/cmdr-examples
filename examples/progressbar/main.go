// Copyright © 2020 Hedzr Yeh.

package main

import (
	"fmt"
	"github.com/hedzr/cmdr"
	cmdrexamples "github.com/hedzr/cmdr-examples"
	"github.com/hedzr/cmdr-examples/examples/progressbar/pbar"
	"github.com/hedzr/cmdr/tool"
	"math/rand"
	"strings"
	"time"

	"github.com/superhawk610/bar"
	"github.com/ttacon/chalk"
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

	pBar(root)
	cmdrPanic(root)

	return
}

func pBar(root cmdr.OptCmd) {
	// progressBar

	cmdr.NewSubCmd().Titles("progress-bar", "p", "pb").
		Description("progress bar test").
		Group("Test").
		TailPlaceholder("[text1, text2, ...]").
		Action(pBar2).
		AttachTo(root)
}

func pBar1(cmd *cmdr.Command, args []string) (err error) {
	for ix, s := range args {
		fmt.Printf("%5d. %s => %s\n", ix, s, tool.Soundex(s))
	}
	fmt.Println(strings.Repeat("\n", rand.Intn(7)))

	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer func() {
			ticker.Stop()
		}()
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Println(strings.Repeat("rand string\n", rand.Intn(2)))
			}
		}
	}()

	pbar.Run(done)
	return
}

func pBar2(cmd *cmdr.Command, args []string) (err error) {

	for ix, s := range args {
		fmt.Printf("%5d. %s => %s\n", ix, s, tool.Soundex(s))
	}

	fmt.Println(strings.Repeat("\n", rand.Intn(7)))

	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer func() {
			ticker.Stop()
		}()
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				// fmt.Println(strings.Repeat("rand string\n", rand.Intn(2)))
			}
		}
	}()

	n := 20
	b := bar.NewWithOpts(
		bar.WithDimensions(n, 30),
		bar.WithFormat(
			fmt.Sprintf(
				"   %sloading...%s :percent :bar %s:rate ops/s%s ",
				chalk.Blue,
				chalk.Reset,
				chalk.Green,
				chalk.Reset,
			),
		),
	)

	fmt.Println()
	fmt.Println()

	for i := 0; i < n; i++ {
		b.Tick()
		time.Sleep(500 * time.Millisecond)
	}

	b.Done()

	fmt.Println()
	fmt.Println()

	return
}

func cmdrPanic(root cmdr.OptCmd) {
	// panic test

	pa := cmdr.NewSubCmd().
		Titles("panic-test", "pa").
		Description("test panic inside cmdr actions", "").
		Group("Test").
		AttachTo(root)

	val := 9
	zeroVal := zero

	cmdr.NewSubCmd().
		Titles("division-by-zero", "dz").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Println(val / zeroVal)
			return
		}).
		AttachTo(pa)

	cmdr.NewSubCmd().
		Titles("panic", "pa").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			panic(9)
			return
		}).
		AttachTo(pa)

}

const (
	appName   = "progress-bar"
	copyright = "progress-bar is an effective devops tool"
	desc      = "progress-bar is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "progress-bar is an effective devops tool. It make an demo application for `cmdr`."
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
