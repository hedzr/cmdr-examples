// Copyright © 2020 Hedzr Yeh.

package cmd

import (
	"github.com/hedzr/cmdr"
)

func AddTags(root cmdr.OptCmd) {
	// tags sub-commands

	msTagsCmd := cmdr.NewSubCmd().Titles("tags", "t").
		Description("tags operations of a micro-service", "").
		Group("").
		AttachTo(root)

	attachConsulConnectFlags(msTagsCmd)

	// ms tags ls

	cmdr.NewSubCmd().Titles("list", "ls", "l", "lst", "dir").
		Description("list tags").
		Group("2333.List").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		}).
		AttachTo(msTagsCmd)

	// ms tags add

	tagsAdd := cmdr.NewSubCmd().Titles("add", "a", "new", "create").
		Description("add tags").
		Deprecated("0.2.1").
		Group("").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		}).
		AttachTo(msTagsCmd)

	cmdr.NewStringSlice().Titles("list", "ls", "l", "lst", "dir").
		Description("a comma list to be added").
		Group("").
		Placeholder("LIST").
		AttachTo(tagsAdd)

	c1 := cmdr.NewSubCmd().Titles("check", "c", "chk").
		Description("[sub] check").
		Group("").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		}).
		AttachTo(tagsAdd)

	c2 := cmdr.NewSubCmd().Titles("check-point", "pt", "chk-pt").
		Description("[sub][sub] checkpoint").
		Group("").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		}).
		AttachTo(c1)

	cmdr.NewStringSlice().Titles("add", "a", "add-list").
		Description("a comma list to be added.").
		Placeholder("LIST").
		Group("List").
		AttachTo(c2)
	cmdr.NewStringSlice().Titles("remove", "r", "rm-list", "rm", "del", "delete").
		Description("a comma list to be removed.", ``).
		Placeholder("LIST").
		Group("List").
		AttachTo(c2)

	c3 := cmdr.NewSubCmd().Titles("check-in", "in", "chk-in").
		Description("[sub][sub] check-in").
		Group("").
		AttachTo(c1)

	cmdr.NewString().
		Titles("name", "n").
		Description("a string to be added.").
		DefaultValue("", "").
		AttachTo(c3)

	cmdr.NewSubCmd().Titles("demo-1", "d1").
		Description("[sub][sub] check-in sub").
		Group("").
		AttachTo(c3)

	cmdr.NewSubCmd().Titles("demo-2", "d2").
		Description("[sub][sub] check-in sub").
		Group("").
		AttachTo(c3)

	cmdr.NewSubCmd().Titles("demo-3", "d3").
		Description("[sub][sub] check-in sub").
		Group("").
		AttachTo(c3)

	cmdr.NewSubCmd().Titles("check-out", "out", "chk-out").
		Description("[sub][sub] check-out").
		Group("").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		}).
		AttachTo(c1)

	// ms tags rm

	tagsRm := cmdr.NewSubCmd().Titles("rm", "r", "remove", "delete", "del", "erase").
		Description("remove tags").
		Group("").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		}).
		AttachTo(msTagsCmd)

	cmdr.NewStringSlice().Titles("list", "ls", "l", "lst", "dir").
		Description("a comma list to be added").
		Group("").
		Placeholder("LIST").
		AttachTo(tagsRm)

	// ms tags modify

	msTagsModifyCmd := cmdr.NewSubCmd().Titles("modify", "m", "mod", "modi", "update", "change").
		Description("modify tags of a service.").
		Action(msTagsModify).
		AttachTo(msTagsCmd)

	attachModifyFlags(msTagsModifyCmd)

	cmdr.NewStringSlice().Titles("add", "a", "add-list").
		Description("a comma list to be added.").
		Placeholder("LIST").
		Group("List").
		AttachTo(msTagsModifyCmd)
	cmdr.NewStringSlice().Titles("remove", "r", "rm-list", "rm", "del", "delete").
		Description("a comma list to be removed.").
		Placeholder("LIST").
		Group("List").
		AttachTo(msTagsModifyCmd)

	// ms tags toggle

	tagsTog := cmdr.NewSubCmd().Titles("toggle", "t", "tog", "switch").
		Description("toggle tags").
		Group("").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		}).
		AttachTo(msTagsCmd)

	attachModifyFlags(tagsTog)

	cmdr.NewStringSlice().Titles("set", "s").
		Description("a comma list to be set").
		Group("").
		Placeholder("LIST").
		AttachTo(tagsTog)

	cmdr.NewStringSlice().Titles("unset", "un").
		Description("a comma list to be unset").
		Group("").
		Placeholder("LIST").
		AttachTo(tagsTog)

	cmdr.NewString().Titles("address", "a", "addr").
		Description("the address of the service (by id or name)").
		Placeholder("HOST:PORT").
		AttachTo(tagsTog)
}

func attachModifyFlags(cmd cmdr.OptCmd) {
	cmdr.NewString("=").Titles("delim", "d").
		Description("delimiter char in `non-plain` mode.").
		Placeholder("").
		AttachTo(cmd)

	cmdr.NewBool().Titles("clear", "c").
		Description("clear all tags.").
		Placeholder("").
		Group("Operate").
		AttachTo(cmd)

	cmdr.NewBool().Titles("string", "g", "string-mode").
		Description("In 'String Mode', default will be disabled: default, a tag string will be split by comma(,), and treated as a string list.").
		Placeholder("").
		ToggleGroup("Mode").
		AttachTo(cmd)

	cmdr.NewBool().Titles("meta", "m", "meta-mode").
		Description("In 'Meta Mode', service 'NodeMeta' field will be updated instead of 'Tags'. (--plain assumed false).").
		Placeholder("").
		ToggleGroup("Mode").
		AttachTo(cmd)

	cmdr.NewBool().Titles("both", "2", "both-mode").
		Description("In 'Both Mode', both of 'NodeMeta' and 'Tags' field will be updated.").
		Placeholder("").
		ToggleGroup("Mode").
		AttachTo(cmd)

	cmdr.NewBool().Titles("plain", "p", "plain-mode").
		Description("In 'Plain Mode', a tag be NOT treated as `key=value` or `key:value`, and modify with the `key`.").
		Placeholder("").
		ToggleGroup("Mode").
		AttachTo(cmd)

	cmdr.NewBool(true).Titles("tag", "t", "tag-mode").
		Description("In 'Tag Mode', a tag be treated as `key=value` or `key:value`, and modify with the `key`.").
		Placeholder("").
		ToggleGroup("Mode").
		AttachTo(cmd)

}

func attachConsulConnectFlags(cmd cmdr.OptCmd) {

	cmdr.NewString("localhost").Titles("addr", "a").
		Description("Consul ip/host and port: HOST[:PORT] (No leading 'http(s)://')", ``).
		Placeholder("HOST[:PORT]").
		Group("Consul").
		AttachTo(cmd)
	cmdr.NewInt(8500).Titles("port", "p").
		Description("Consul port", ``).
		Placeholder("PORT").
		Group("Consul").
		AttachTo(cmd)
	cmdr.NewBool().Titles("insecure", "K").
		Description("Skip TLS host verification", ``).
		Placeholder("").
		Group("Consul").
		AttachTo(cmd)
	cmdr.NewString("/").Titles("prefix", "px").
		Description("Root key prefix", ``).
		Placeholder("ROOT").
		Group("Consul").
		AttachTo(cmd)
	cmdr.NewString().Titles("cacert", "").
		Description("Consul Client CA cert)", ``).
		Placeholder("FILE").
		Group("Consul").
		AttachTo(cmd)
	cmdr.NewString().Titles("cert", "").
		Description("Consul Client cert", ``).
		Placeholder("FILE").
		Group("Consul").
		AttachTo(cmd)
	cmdr.NewString("http").Titles("scheme", "").
		Description("Consul connection protocol", ``).
		Placeholder("SCHEME").
		Group("Consul").
		AttachTo(cmd)
	cmdr.NewString().Titles("username", "u", "user", "usr", "uid").
		Description("HTTP Basic auth user", ``).
		Placeholder("USERNAME").
		Group("Consul").
		AttachTo(cmd)
	cmdr.NewString().Titles("password", "pw", "passwd", "pass", "pwd").
		Description("HTTP Basic auth `password`", ``).
		// Placeholder("PASSWORD").
		Group("Consul").
		ExternalTool(cmdr.ExternalToolPasswordInput).
		AttachTo(cmd)

}
