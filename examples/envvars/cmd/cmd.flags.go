// Copyright © 2020 Hedzr Yeh.

package cmd

import (
	"fmt"
	"github.com/hedzr/cmdr"
)

func prd(key string, val interface{}, format string, params ...interface{}) {
	fmt.Printf("[--%v] %v, %v\n", key, val, fmt.Sprintf(format, params...))
}

func AddFlags(root cmdr.OptCmd) {
	// tags sub-commands

	parent := cmdr.NewSubCmd().Titles("flags", "f").
		Description("flags demo", "").
		Group("").
		Action(flagsAction).
		AttachTo(root)

	cmdr.NewBool(false).
		Titles("bool", "b").
		Description("A bool flag", "").
		Group("").
		AttachTo(parent)

	cmdr.NewInt(1).
		Titles("int", "i").
		Description("A int flag", "").
		Group("1000.Integer").
		AttachTo(parent)
	cmdr.NewInt64(2).
		Titles("int64", "i64").
		Description("A int64 flag", "").
		Group("1000.Integer").
		AttachTo(parent)
	cmdr.NewUint(3).
		Titles("uint", "u").
		Description("A uint flag", "").
		Group("1000.Integer").
		AttachTo(parent)
	cmdr.NewUint64(4).
		Titles("uint64", "u64").
		Description("A uint64 flag", "").
		Group("1000.Integer").
		AttachTo(parent)

	cmdr.NewFloat32(2.71828).
		Titles("float32", "f", "float").
		Description("A float32 flag with 'e' value", "").
		Group("2000.Float").
		AttachTo(parent)
	cmdr.NewFloat64(3.14159265358979323846264338327950288419716939937510582097494459230781640628620899).
		Titles("float64", "f64").
		Description("A float64 flag with a `PI` value", "").
		Group("2000.Float").
		AttachTo(parent)
	cmdr.NewComplex64(3.14+9i).
		Titles("complex64", "c64").
		Description("A complex64 flag", "").
		Group("2010.Complex").
		AttachTo(parent)
	cmdr.NewComplex64(3.14+9i).
		Titles("complex128", "c128").
		Description("A complex128 flag", "").
		Group("2010.Complex").
		AttachTo(parent)

	// a set of booleans

	cmdr.NewBool(false).
		Titles("single", "s").
		Description("A bool flag: single", "").
		Group("Boolean").
		EnvKeys("").
		AttachTo(parent)

	cmdr.NewBool(false).
		Titles("double", "d").
		Description("A bool flag: double", "").
		Group("Boolean").
		EnvKeys("").
		AttachTo(parent)

	cmdr.NewBool(false).
		Titles("norway", "n", "nw").
		Description("A bool flag: norway", "").
		Group("Boolean").
		EnvKeys("").
		AttachTo(parent)

	cmdr.NewBool(false).
		Titles("mongo", "m").
		Description("A bool flag: mongo", "").
		Group("Boolean").
		EnvKeys("").
		AttachTo(parent)

}

func flagsAction(cmd *cmdr.Command, args []string) (err error) {
	prd("bool", cmdr.GetBoolR("flags.bool"), "")
	prd("int", cmdr.GetIntR("flags.int"), "")
	prd("int64", cmdr.GetInt64R("flags.int64"), "")
	prd("uint", cmdr.GetUintR("flags.uint"), "")
	prd("uint64", cmdr.GetUint64R("flags.uint64"), "")
	prd("float32", cmdr.GetFloat32R("flags.float32"), "")
	prd("float64", cmdr.GetFloat64R("flags.float64"), "")
	prd("complex64", cmdr.GetComplex64R("flags.complex64"), "")
	prd("complex128", cmdr.GetComplex128R("flags.complex128"), "")

	prd("single", cmdr.GetBoolR("flags.single"), "")
	prd("double", cmdr.GetBoolR("flags.double"), "")
	prd("norway", cmdr.GetBoolR("flags.norway"), "")
	prd("mongo", cmdr.GetBoolR("flags.mongo"), "")

	return
}
