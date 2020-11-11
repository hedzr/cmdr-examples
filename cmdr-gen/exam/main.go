// Copyright © 2020 Hedzr Yeh.

//go:generate cmdr-gen app -n cli-exam

//go:generate cmdr-gen flag -d '
// gender: [male|female]
// '

package main

//import "github.com/hedzr/cmdr-examples/cmdr-gen/exam/cli"
//
//func main() {
//	cli.Entry()
//}

type Status int

// go :generate myenumstr -type Status,Color
const (
	Offline Status = iota
	Online
	Disable
	Deleted
)

type Color int

const (
	Write Color = iota
	Red
	Blue
)
