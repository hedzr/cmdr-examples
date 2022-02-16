package cmd

import "github.com/hedzr/cmdr"

func kvCommand(root cmdr.OptCmd) {

	// kv

	kvCmd := cmdr.NewSubCmd().Titles("kvstore", "kv").
		Description("consul kv store operations...", ``).
		AttachTo(root)

	attachConsulConnectFlags(kvCmd)

	kvBackupCmd := cmdr.NewSubCmd().Titles("backup", "b", "bf", "bkp").
		Description("Dump Consul's KV database to a JSON/YAML file", ``).
		Action(kvBackup).
		AttachTo(kvCmd)
	cmdr.NewString("consul-backup.json").Titles("output", "o").
		Description("Write output to a file (*.json / *.yml)", ``).
		Placeholder("FILE").
		CompletionActionStr(`*.(json|yml|yaml)`). //  \*.\(ps\|eps\)
		// ':postscript file:_files -g \*.\(ps\|eps\)'
		AttachTo(kvBackupCmd)

	kvRestoreCmd := cmdr.NewSubCmd().Titles("restore", "r").
		Description("restore to Consul's KV store, from a a JSON/YAML backup file", ``).
		Action(kvRestore).
		AttachTo(kvCmd)
	cmdr.NewString("consul-backup.json").Titles("input", "i").
		Description("Read the input file (*.json / *.yml)", ``).
		Placeholder("FILE").
		AttachTo(kvRestoreCmd)

}

func msCommand(root cmdr.OptCmd) {

	// ms

	msCmd := cmdr.NewSubCmd().Titles("micro-service", "ms", "microservice").
		Description("micro-service operations...", "").
		Group("").
		AttachTo(root)

	cmdr.NewBool().Titles("money", "mm").
		Description("A placeholder flag - money.", "").
		Group("").
		Placeholder("").
		AttachTo(msCmd)

	cmdr.NewString("").Titles("name", "n").
		Description("name of the service", ``).
		Placeholder("NAME").
		AttachTo(msCmd)
	cmdr.NewString().Titles("id", "i", "ID").
		Description("unique id of the service", ``).
		Placeholder("ID").
		AttachTo(msCmd)
	cmdr.NewBool().Titles("all", "a").
		Description("all services", ``).
		Placeholder("").
		AttachTo(msCmd)

	cmdr.NewInt(3).Titles("retry", "t").
		Description("retry times for ms cmd", "").
		Group("").
		Placeholder("RETRY").
		AttachTo(msCmd)

	// ms ls

	cmdr.NewSubCmd().Titles("list", "ls", "l", "lst", "dir").
		Description("list tags for ms cmd", "").
		Group("2333.List").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		}).
		AttachTo(msCmd)

	// ms tags

	msTagsCmd := cmdr.NewSubCmd().Titles("tags", "t").
		Description("tags operations of a micro-service", "").
		Group("").
		AttachTo(msCmd)

	// cTags.NewFlag(cmdr.OptFlagTypeString).
	// 	Titles("n", "name").
	// 	Description("name of the service", "").
	// 	Group("").
	// 	DefaultValue("", "NAME")
	//
	// cTags.NewFlag(cmdr.OptFlagTypeString).
	// 	Titles("i", "id").
	// 	Description("unique id of the service", "").
	// 	Group("").
	// 	DefaultValue("", "ID")
	//
	// cTags.NewFlag(cmdr.OptFlagTypeString).
	// 	Titles("a", "addr").
	// 	Description("", "").
	// 	Group("").
	// 	DefaultValue("consul.ops.local", "ADDR")

	attachConsulConnectFlags(msTagsCmd)

	// ms tags ls

	cmdr.NewSubCmd().Titles("list", "ls", "l", "lst", "dir").
		Description("list tags for ms tags cmd").
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
		//Action(func(cmd *cmdr.Command, args []string) (err error) {
		//	return
		//}).
		AttachTo(msTagsCmd)

	cmdr.NewStringSlice().Titles("list", "ls", "l", "lst", "dir").
		Description("tags add: a comma list to be added").
		Group("").
		Placeholder("LIST").
		AttachTo(tagsAdd)

	c1 := cmdr.NewSubCmd().Titles("check", "c", "chk").
		Description("[sub] check").
		Group("").
		//Action(func(cmd *cmdr.Command, args []string) (err error) {
		//	return
		//}).
		AttachTo(tagsAdd)

	c2 := cmdr.NewSubCmd().Titles("check-point", "pt", "chk-pt").
		Description("[sub][sub] checkpoint").
		Group("").
		//Action(func(cmd *cmdr.Command, args []string) (err error) {
		//	return
		//}).
		AttachTo(c1)

	cmdr.NewStringSlice().Titles("add", "a", "add-list").
		Description("checkpoint: a comma list to be added.").
		Placeholder("LIST").
		Group("List").
		AttachTo(c2)
	cmdr.NewStringSlice().Titles("remove", "r", "rm-list", "rm", "del", "delete").
		Description("checkpoint: a comma list to be removed.", ``).
		Placeholder("LIST").
		Group("List").
		AttachTo(c2)

	c3 := cmdr.NewSubCmd().Titles("check-in", "in", "chk-in").
		Description("[sub][sub] check-in").
		Group("").
		AttachTo(c1)

	cmdr.NewString().
		Titles("n", "name").
		Description("check-in name: a string to be added.").
		DefaultValue("", "").
		AttachTo(c3)

	cmdr.NewSubCmd().Titles("demo-1", "d1").
		Description("[sub][sub] check-in sub, d1").
		Group("").
		AttachTo(c3)

	cmdr.NewSubCmd().Titles("demo-2", "d2").
		Description("[sub][sub] check-in sub, d2").
		Group("").
		AttachTo(c3)

	cmdr.NewSubCmd().Titles("demo-3", "d3").
		Description("[sub][sub] check-in sub, d3").
		Group("").
		AttachTo(c3)

	cmdr.NewSubCmd().Titles("check-out", "out", "chk-out").
		Description("[sub][sub] check-out").
		Group("").
		//Action(func(cmd *cmdr.Command, args []string) (err error) {
		//	return
		//}).
		AttachTo(c3)

	// ms tags rm

	tagsRm := cmdr.NewSubCmd().Titles("rm", "r", "remove", "delete", "del", "erase").
		Description("remove tags").
		Group("").
		//Action(func(cmd *cmdr.Command, args []string) (err error) {
		//	return
		//}).
		AttachTo(msTagsCmd)

	cmdr.NewStringSlice().Titles("list", "ls", "l", "lst", "dir").
		Description("tags rm: a comma list to be added").
		Group("").
		Placeholder("LIST").
		AttachTo(tagsRm)

	// ms tags modify

	msTagsModifyCmd := cmdr.NewSubCmd().Titles("modify", "m", "mod", "modi", "update", "change").
		Description("modify tags of a service.").
		Action(msTagsModify).AttachTo(msTagsCmd)

	attachModifyFlags(msTagsModifyCmd)

	cmdr.NewStringSlice().Titles("add", "a", "add-list").
		Description("tags modify: a comma list to be added.").
		Placeholder("LIST").
		Group("List").
		AttachTo(msTagsModifyCmd)
	cmdr.NewStringSlice().Titles("remove", "r", "rm-list", "rm", "del", "delete").
		Description("tags modify: a comma list to be removed.").
		Placeholder("LIST").
		Group("List").
		AttachTo(msTagsModifyCmd)

	// ms tags toggle

	tagsTog := cmdr.NewSubCmd().Titles("toggle", "t", "tog", "switch").
		Description("toggle tags").
		Group("").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		}).AttachTo(msTagsCmd)

	attachModifyFlags(tagsTog)

	cmdr.NewStringSlice().Titles("set", "s").
		Description("tags toggle: a comma list to be set").
		Group("").
		Placeholder("LIST").
		AttachTo(tagsTog)

	cmdr.NewStringSlice().Titles("unset", "un").
		Description("tags toggle: a comma list to be unset").
		Group("").
		Placeholder("LIST").
		AttachTo(tagsTog)

	cmdr.NewString().Titles("address", "a", "addr").
		Description("tags toggle: the address of the service (by id or name)").
		Placeholder("HOST:PORT").
		AttachTo(tagsTog)

}
