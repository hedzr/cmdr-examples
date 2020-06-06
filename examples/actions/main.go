// Copyright © 2020 Hedzr Yeh.

package main

import (
	"fmt"
	"github.com/hedzr/cmdr"
	cmdr_examples "github.com/hedzr/cmdr-examples"
	"github.com/hedzr/cmdr-examples/internal/trace"
	"gopkg.in/hedzr/errors.v2"
)

func main() {
	Entry()
}

func Entry() {
	if err := cmdr.Exec(buildRootCmd(), cmdr.WithUnhandledErrorHandler(onUnhandledErrorHandler)); err != nil {
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
	root := cmdr.Root(appName, cmdr_examples.Version).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	cmdr.NewBool(false).
		Titles("enable-ueh", "ueh").
		Description("Enables the unhandled exception handler?").
		AttachTo(root)

	soundex(root)
	panicTest(root)

	return
}

func prd(key string, val interface{}, format string, params ...interface{}) {
	fmt.Printf("         [--%v] %v, %v\n", key, val, fmt.Sprintf(format, params...))
}

func soundex(root cmdr.OptCmd) {
	// soundex

	parent := root.NewSubCommand("soundex", "snd", "sndx", "sound").
		Description("soundex test").
		Group("Test").
		TailPlaceholder("[text1, text2, ...]").
		PreAction(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Printf("[PRE] DebugMode=%v, TraceMode=%v. InDebugging/IsDebuggerAttached=%v\n", cmdr.GetDebugMode(), trace.IsEnabled(), cmdr.InDebugging())
			for ix, s := range args {
				fmt.Printf("[PRE] %5d. %s\n", ix, s)
			}

			fmt.Printf("[PRE] Debug=%v, Trace=%v\n", cmdr.GetBoolR("debug"), cmdr.GetBoolR("trace"))

			// return nil to be continue, 
			// return cmdr.ErrShouldBeStopException to stop the following actions without error
			// return other errors for application purpose
			return
		}).
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			for ix, s := range args {
				// fmt.Printf("[ACTION] %5d. %s\n", ix, s)
				fmt.Printf("[ACTION] %5d. %s => %s\n", ix, s, cmdr.Soundex(s))
			}

			prd("bool", cmdr.GetBoolR("flags.bool"), "")
			prd("int", cmdr.GetIntR("flags.int"), "")
			prd("int64", cmdr.GetInt64R("flags.int64"), "")
			prd("uint", cmdr.GetUintR("flags.uint"), "")
			prd("uint64", cmdr.GetUint64R("flags.uint64"), "")
			prd("float32", cmdr.GetFloat32R("flags.float32"), "")
			prd("float64", cmdr.GetFloat64R("flags.float64"), "")
			prd("complex64", cmdr.GetComplex64R("flags.complex64"), "")
			prd("complex128", cmdr.GetComplex128R("flags.complex128"), "")

			return
		}).
		PostAction(func(cmd *cmdr.Command, args []string) {
			for ix, s := range args {
				fmt.Printf("[POST] %5d. %s\n", ix, s)
			}
		})

	cmdr.NewBool(false).
		Titles("bool", "b").
		Description("A bool flag", "").
		Group("").
		EnvKeys("").
		AttachTo(parent)

	cmdr.NewInt(1).
		Titles("int", "i").
		Description("A int flag", "").
		Group("1000.Integer").
		EnvKeys("").
		AttachTo(parent)
	cmdr.NewInt64(2).
		Titles("int64", "i64").
		Description("A int64 flag", "").
		Group("1000.Integer").
		EnvKeys("").
		AttachTo(parent)
	cmdr.NewUint(3).
		Titles("uint", "u").
		Description("A uint flag", "").
		Group("1000.Integer").
		EnvKeys("").
		AttachTo(parent)
	cmdr.NewUint64(4).
		Titles("uint64", "u64").
		Description("A uint64 flag", "").
		Group("1000.Integer").
		EnvKeys("").
		AttachTo(parent)

	cmdr.NewFloat32(2.71828).
		Titles("float32", "f", "float").
		Description("A float32 flag with 'e' value", "").
		Group("2000.Float").
		EnvKeys("E", "E2").
		AttachTo(parent)
	cmdr.NewFloat64(3.14159265358979323846264338327950288419716939937510582097494459230781640628620899).
		Titles("float64", "f64").
		Description("A float64 flag with a `PI` value", "").
		Group("2000.Float").
		EnvKeys("PI").
		AttachTo(parent)
	cmdr.NewComplex64(3.14+9i).
		Titles("complex64", "c64").
		Description("A complex64 flag", "").
		Group("2010.Complex").
		EnvKeys("").
		AttachTo(parent)
	cmdr.NewComplex128(3.14+9i).
		Titles("complex128", "c128").
		Description("A complex128 flag", "").
		Group("2010.Complex").
		EnvKeys("").
		AttachTo(parent)

}

func panicTest(root cmdr.OptCmd) {
	// panic test

	pa := root.NewSubCommand("panic-test", "pa").
		Description("test panic inside cmdr actions", "").
		Group("Test")

	val := 9
	zeroVal := zero

	pa.NewSubCommand("division-by-zero", "dz").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Println(val / zeroVal)
			return
		})

	pa.NewSubCommand("panic", "pa").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			panic(9)
			return
		})
}

const (
	appName   = "simple"
	copyright = "simple is an effective devops tool"
	desc      = "simple is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "simple is an effective devops tool. It make an demo application for `cmdr`."
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
