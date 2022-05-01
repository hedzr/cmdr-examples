// Copyright © 2020 Hedzr Yeh.

package gen

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/cmdr/tool"
	"github.com/hedzr/log/dir"
	"gopkg.in/hedzr/errors.v3"
	"log"
	"os"
)

func Entry() {
	if err := cmdr.Exec(buildRootCmd(),
		cmdr.WithIgnoreWrongEnumValue(true),
		cmdr.WithUnhandledErrorHandler(onUnhandledErrorHandler),
	); err != nil {
		fmt.Printf("error: %+v\n", err)
	}
}

func onUnhandledErrorHandler(err interface{}) {
	if cmdr.GetBoolR("enable-ueh") {
		dumpStacks()
		return
	}

	panic(err)
}

func dumpStacks() {
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===\n", errors.DumpStacksAsString(true))
}

func buildRootCmd() (rootCmd *cmdr.RootCommand) {
	root := cmdr.Root(appName, Version).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	gen(root)

	mx(root)

	return
}

func gen(root cmdr.OptCmd) {
	genNewApp(root)
	genNewCommand(root)
	genNewFlag(root)
}

func genNewApp(root cmdr.OptCmd) {
	var cc cmdr.OptCmd

	cc = cmdr.NewSubCmd().Titles("application", "a", "app").
		Description("create an app", "test new features,\nverbose long descriptions here.").
		Group("Test").
		Action(genApp).
		AttachTo(root)

	cmdr.NewString("example").
		Titles("module-name", "mn").
		Description("the module name of your application").
		EnvKeys("GOMODNAME").
		AttachTo(cc)
	cmdr.NewString("example").
		Titles("name", "n", "appname", "app-name").
		Description("the name of your application").
		EnvKeys("GOAPPNAME").
		AttachTo(cc)
	cmdr.NewString("github.com/yourname/example").
		Titles("package", "p", "pkg").
		Description("the package name of your application").
		EnvKeys("GOPACKAGE").
		AttachTo(cc)
	cmdr.NewString().
		Titles("processing-filename", "file").
		Description("the processing filename for go generate").
		EnvKeys("GOFILE").
		AttachTo(cc)
	cmdr.NewString().
		Titles("processing-file-lineno", "lineno").
		Description("the processing file line number for go generate").
		EnvKeys("GOLINE").
		AttachTo(cc)
}

func genNewCommand(root cmdr.OptCmd) {

	var cc cmdr.OptCmd

	cc = cmdr.NewSubCmd().Titles("flag", "f").
		Description("create a flag from YAML definition", "test new features,\nverbose long descriptions here.").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		}).
		AttachTo(root)

	cmdr.NewString().
		Titles("define", "d").
		Description("the YAML definition").
		AttachTo(cc)

}
func genNewFlag(root cmdr.OptCmd) {

	var cc cmdr.OptCmd

	cc = cmdr.NewSubCmd().Titles("command", "c", "cmd").
		Description("create a command from YAML definition", "test new features,\nverbose long descriptions here.").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		}).
		AttachTo(root)

	cmdr.NewString().
		Titles("define", "d").
		Description("the YAML definition").
		AttachTo(cc)

}

func mx(root cmdr.OptCmd) {
	// mx-test

	mx := cmdr.NewSubCmd().Titles("mx-test", "mx").
		Description("test new features", "test new features,\nverbose long descriptions here.").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			// cmdr.Set("test.1", 8)
			cmdr.Set("test.deep.branch.1", "test")
			z := cmdr.GetString("app.test.deep.branch.1")
			fmt.Printf("*** Got app.test.deep.branch.1: %s\n", z)
			if z != "test" {
				log.Fatalf("err, expect 'test', but got '%v'", z)
			}

			cmdr.DeleteKey("app.test.deep.branch.1")
			if cmdr.HasKey("app.test.deep.branch.1") {
				log.Fatalf("FAILED, expect key not found, but found: %v", cmdr.Get("app.test.deep.branch.1"))
			}
			fmt.Printf("*** Got app.test.deep.branch.1 (after deleted): %s\n", cmdr.GetString("app.test.deep.branch.1"))

			fmt.Printf("*** Got pp: %s\n", cmdr.GetString("app.mx-test.password"))
			fmt.Printf("*** Got msg: %s\n", cmdr.GetString("app.mx-test.message"))
			fmt.Printf("*** Got fruit (toggle group): %v\n", cmdr.GetString("app.mx-test.fruit"))
			fmt.Printf("*** Got head (head-like): %v\n", cmdr.GetInt("app.mx-test.head"))
			fmt.Println()
			fmt.Printf("*** test text: %s\n", cmdr.GetStringR("mx-test.test"))
			fmt.Println()
			fmt.Printf("> InTesting: args[0]=%v \n", tool.SavedOsArgs[0])
			fmt.Println()
			fmt.Printf("> Used config file: %v\n", cmdr.GetUsedConfigFile())
			fmt.Printf("> Used config files: %v\n", cmdr.GetUsingConfigFiles())
			fmt.Printf("> Used config sub-dir: %v\n", cmdr.GetUsedConfigSubDir())

			fmt.Printf("> STDIN MODE: %v \n", cmdr.GetBoolR("mx-test.stdin"))
			fmt.Println()

			// logrus.Debug("debug")
			// logrus.Info("debug")
			// logrus.Warning("debug")
			// logrus.WithField(logex.SKIP, 1).Warningf("dsdsdsds")

			if cmdr.GetBoolR("mx-test.stdin") {
				fmt.Println("> Type your contents here, press Ctrl-D to end it:")
				var data []byte
				data, err = dir.ReadAll(os.Stdin)
				if err != nil {
					log.Printf("error: %+v", err)
					return
				}
				fmt.Println("> The input contents are:")
				fmt.Print(string(data))
				fmt.Println()
			}
			return
		}).
		AttachTo(root)

	cmdr.NewString().Titles("test", "t").
		Description("the test text.", "").
		EnvKeys("COOL", "TEST").
		Group("").
		AttachTo(mx)
	cmdr.NewString().Titles("password", "pp").
		Description("the password requesting.", "").
		Group("").
		Placeholder("PASSWORD").
		ExternalTool(cmdr.ExternalToolPasswordInput).
		AttachTo(mx)
	cmdr.NewString().Titles("message", "m", "msg").
		Description("the message requesting.", "").
		Group("").
		Placeholder("MESG").
		ExternalTool(cmdr.ExternalToolEditor).
		AttachTo(mx)
	cmdr.NewString().Titles("fruit", "fr").
		Description("the message.", "").
		Group("").
		Placeholder("FRUIT").
		ValidArgs("apple", "banana", "orange").
		AttachTo(mx)
	cmdr.NewInt(1).Titles("head", "hd").
		Description("the head lines.", "").
		Group("").
		Placeholder("LINES").
		HeadLike(true, 1, 3000).
		AttachTo(mx)
	cmdr.NewBool().Titles("stdin", "c").
		Description("read file content from stdin.", "").
		Group("").
		AttachTo(mx)
}

const (
	Version   = "1.0.0"
	appName   = "cmdr-gen"
	copyright = "cmdr-gen is an effective devops tool"
	desc      = "cmdr-gen is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "cmdr-gen is an effective devops tool. It make an demo application for `cmdr`."
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
