// Copyright © 2020 Hedzr Yeh.

package cmd

import (
	"github.com/hedzr/cmdr"
)

func AddTags(root cmdr.OptCmd) {
	// tags sub-commands

	msTagsCmd := root.NewSubCommand("tags", "t").
		Description("tags operations of a micro-service", "").
		Group("")

	attachConsulConnectFlags(msTagsCmd)

	// ms tags ls

	msTagsCmd.NewSubCommand("list", "ls", "l", "lst", "dir").
		Description("list tags").
		Group("2333.List").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		})

	// ms tags add

	tagsAdd := msTagsCmd.NewSubCommand("add", "a", "new", "create").
		Description("add tags").
		Deprecated("0.2.1").
		Group("").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		})

	tagsAdd.NewFlagV([]string{}, "list", "ls", "l", "lst", "dir").
		Description("a comma list to be added").
		Group("").
		Placeholder("LIST")

	c1 := tagsAdd.NewSubCommand("check", "c", "chk").
		Description("[sub] check").
		Group("").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		})

	c2 := c1.NewSubCommand("check-point", "pt", "chk-pt").
		Description("[sub][sub] checkpoint").
		Group("").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		})

	c2.NewFlagV([]string{}, "add", "a", "add-list").
		Description("a comma list to be added.").
		Placeholder("LIST").
		Group("List")
	c2.NewFlagV([]string{}, "remove", "r", "rm-list", "rm", "del", "delete").
		Description("a comma list to be removed.", ``).
		Placeholder("LIST").
		Group("List")

	c3 := c1.NewSubCommand("check-in", "in", "chk-in").
		Description("[sub][sub] check-in").
		Group("")

	c3.NewFlag(cmdr.OptFlagTypeString).
		Titles("name", "n").
		Description("a string to be added.").
		DefaultValue("", "")

	c3.NewSubCommand("demo-1", "d1").
		Description("[sub][sub] check-in sub").
		Group("")

	c3.NewSubCommand("demo-2", "d2").
		Description("[sub][sub] check-in sub").
		Group("")

	c3.NewSubCommand("demo-3", "d3").
		Description("[sub][sub] check-in sub").
		Group("")

	c1.NewSubCommand("check-out", "out", "chk-out").
		Description("[sub][sub] check-out").
		Group("").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		})

	// ms tags rm

	tagsRm := msTagsCmd.NewSubCommand("rm", "r", "remove", "delete", "del", "erase").
		Description("remove tags").
		Group("").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		})

	tagsRm.NewFlagV([]string{}, "list", "ls", "l", "lst", "dir").
		Description("a comma list to be added").
		Group("").
		Placeholder("LIST")

	// ms tags modify

	msTagsModifyCmd := msTagsCmd.NewSubCommand("modify", "m", "mod", "modi", "update", "change").
		Description("modify tags of a service.").
		Action(msTagsModify)

	attachModifyFlags(msTagsModifyCmd)

	msTagsModifyCmd.NewFlagV([]string{}, "add", "a", "add-list").
		Description("a comma list to be added.").
		Placeholder("LIST").
		Group("List")
	msTagsModifyCmd.NewFlagV([]string{}, "remove", "r", "rm-list", "rm", "del", "delete").
		Description("a comma list to be removed.").
		Placeholder("LIST").
		Group("List")

	// ms tags toggle

	tagsTog := msTagsCmd.NewSubCommand("toggle", "t", "tog", "switch").
		Description("toggle tags").
		Group("").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		})

	attachModifyFlags(tagsTog)

	tagsTog.NewFlagV([]string{}, "set", "s").
		Description("a comma list to be set").
		Group("").
		Placeholder("LIST")

	tagsTog.NewFlagV([]string{}, "unset", "un").
		Description("a comma list to be unset").
		Group("").
		Placeholder("LIST")

	tagsTog.NewFlagV("", "address", "a", "addr").
		Description("the address of the service (by id or name)").
		Placeholder("HOST:PORT")
}

func attachModifyFlags(cmd cmdr.OptCmd) {
	cmd.NewFlagV("=", "delim", "d").
		Description("delimiter char in `non-plain` mode.").
		Placeholder("")

	cmd.NewFlagV(false, "clear", "c").
		Description("clear all tags.").
		Placeholder("").
		Group("Operate")

	cmd.NewFlagV(false, "string", "g", "string-mode").
		Description("In 'String Mode', default will be disabled: default, a tag string will be split by comma(,), and treated as a string list.").
		Placeholder("").
		Group("Mode")

	cmd.NewFlagV(false, "meta", "m", "meta-mode").
		Description("In 'Meta Mode', service 'NodeMeta' field will be updated instead of 'Tags'. (--plain assumed false).").
		Placeholder("").
		Group("Mode")

	cmd.NewFlagV(false, "both", "2", "both-mode").
		Description("In 'Both Mode', both of 'NodeMeta' and 'Tags' field will be updated.").
		Placeholder("").
		Group("Mode")

	cmd.NewFlagV(false, "plain", "p", "plain-mode").
		Description("In 'Plain Mode', a tag be NOT treated as `key=value` or `key:value`, and modify with the `key`.").
		Placeholder("").
		Group("Mode")

	cmd.NewFlagV(true, "tag", "t", "tag-mode").
		Description("In 'Tag Mode', a tag be treated as `key=value` or `key:value`, and modify with the `key`.").
		Placeholder("").
		Group("Mode")

}

func attachConsulConnectFlags(cmd cmdr.OptCmd) {

	cmd.NewFlagV("localhost", "addr", "a").
		Description("Consul ip/host and port: HOST[:PORT] (No leading 'http(s)://')", ``).
		Placeholder("HOST[:PORT]").
		Group("Consul")
	cmd.NewFlagV(8500, "port", "p").
		Description("Consul port", ``).
		Placeholder("PORT").
		Group("Consul")
	cmd.NewFlagV(true, "insecure", "K").
		Description("Skip TLS host verification", ``).
		Placeholder("").
		Group("Consul")
	cmd.NewFlagV("/", "prefix", "px").
		Description("Root key prefix", ``).
		Placeholder("ROOT").
		Group("Consul")
	cmd.NewFlagV("", "cacert").
		Description("Consul Client CA cert)", ``).
		Placeholder("FILE").
		Group("Consul")
	cmd.NewFlagV("", "cert").
		Description("Consul Client cert", ``).
		Placeholder("FILE").
		Group("Consul")
	cmd.NewFlagV("http", "scheme").
		Description("Consul connection protocol", ``).
		Placeholder("SCHEME").
		Group("Consul")
	cmd.NewFlagV("", "username", "u", "user", "usr", "uid").
		Description("HTTP Basic auth user", ``).
		Placeholder("USERNAME").
		Group("Consul")
	cmd.NewFlagV("", "password", "pw", "passwd", "pass", "pwd").
		Description("HTTP Basic auth `password`", ``).
		// Placeholder("PASSWORD").
		Group("Consul").
		ExternalTool(cmdr.ExternalToolPasswordInput)

}
