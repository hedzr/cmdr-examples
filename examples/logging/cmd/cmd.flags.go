// Copyright © 2020 Hedzr Yeh.

package cmd

import "github.com/hedzr/cmdr"

func addFlags(root cmdr.OptCmd) {
	// tags sub-commands

	parent := cmdr.NewSubCmd().Titles("flags", "f").
		Description("flags demo", "").
		Group("").AttachTo(root)

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
	cmdr.NewComplex64(3.14+9i).
		Titles("complex128", "c128").
		Description("A complex128 flag", "").
		Group("2010.Complex").
		EnvKeys("").
		AttachTo(parent)

}
