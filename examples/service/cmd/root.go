// Copyright © 2020 Hedzr Yeh.

package cmd

import (
	"fmt"
	"github.com/hedzr/cmdr"
	cmdr_examples "github.com/hedzr/cmdr-examples"
	"github.com/hedzr/cmdr/tool"
	"io/ioutil"
	"log"
	"os"
)

func buildRootCmd() (rootCmd *cmdr.RootCommand) {

	// var cmd *Command

	// cmdr.Root("aa", "1.0.1").
	// 	Header("sds").
	// 	NewSubCommand().
	// 	Titles("ms", "microservice").
	// 	Description("", "").
	// 	Group("").
	// 	Action(func(cmd *cmdr.Command, args []string) (err error) {
	// 		return
	// 	})

	// root

	root := cmdr.Root(appName, cmdr_examples.Version).
		AddGlobalPreAction(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Println("# global pre-action 1")
			return
		}).
		AddGlobalPreAction(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Println("# global pre-action 2")
			return
		}).
		AddGlobalPostAction(func(cmd *cmdr.Command, args []string) {
			fmt.Println("# global post-action 1")
		}).
		AddGlobalPostAction(func(cmd *cmdr.Command, args []string) {
			fmt.Println("# global post-action 2")
		}).
		// Header("fluent - test for cmdr - no version - hedzr").
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	// root.NewSubCommand("", "go112113").
	// 	Description("test build tags for go1.13 or later and go1.12 and below", "").
	// 	Group("").Action(func(cmd *cmdr.Command, args []string) (err error) {
	// 	// go112113.Fate()
	// 	return nil
	// })

	panicTest(root)
	kbPrint(root)
	soundex(root)
	mx(root)
	kv(root)
	ms(root)

	return
}

func soundex(root cmdr.OptCmd) {
	// soundex

	root.NewSubCommand("soundex", "snd", "sndx", "sound").
		Description("soundex test").
		Group("Test").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			for ix, s := range args {
				fmt.Printf("%5d. %s => %s\n", ix, s, tool.Soundex(s))
			}
			return
		})
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

func kbPrint(root cmdr.OptCmd) {
	// kb-print

	kb := root.NewSubCommand("kb-print", "kb").
		Description("kilobytes test", "test kibibytes' input,\nverbose long descriptions here.").
		Group("Test").
		Examples(`
$ {{.AppName}} kb --size 5kb
  5kb = 5,120 bytes
$ {{.AppName}} kb --size 8T
  8TB = 8,796,093,022,208 bytes
$ {{.AppName}} kb --size 1g
  1GB = 1,073,741,824 bytes
		`).
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Printf("Got size: %v (literal: %v)\n\n", cmdr.GetKibibytesR("kb-print.size"), cmdr.GetStringR("kb-print.size"))
			return
		})

	kb.NewFlagV("1k", "size", "s").
		Description("max message size. Valid formats: 2k, 2kb, 2kB, 2KB. Suffixes: k, m, g, t, p, e.", "").
		Group("")

	// xy-print

	root.NewSubCommand("xy-print", "xy").
		Description("test terminal control sequences", "test terminal control sequences,\nverbose long descriptions here.").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			//
			// https://en.wikipedia.org/wiki/ANSI_escape_code
			// https://zh.wikipedia.org/wiki/ANSI%E8%BD%AC%E4%B9%89%E5%BA%8F%E5%88%97
			// https://en.wikipedia.org/wiki/POSIX_terminal_interface
			//

			fmt.Println("\x1b[2J") // clear screen

			for i, s := range args {
				fmt.Printf("\x1b[s\x1b[%d;%dH%s\x1b[u", 15+i, 30, s)
			}

			return
		})
}

func mx(root cmdr.OptCmd) {
	// mx-test

	mx := root.NewSubCommand("mx-test", "mx").
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

			cmdr.Logger.Debugf("debug")
			cmdr.Logger.Infof("info")
			cmdr.Logger.Warnf("warning")
			// cmdr.Logger.WithField(logex.SKIP, 1).Warningf("dsdsdsds")

			if cmdr.GetBoolR("mx-test.stdin") {
				fmt.Println("> Type your contents here, press Ctrl-D to end it:")
				var data []byte
				data, err = ioutil.ReadAll(os.Stdin)
				if err != nil {
					log.Printf("error: %+v", err)
					return
				}
				fmt.Println("> The input contents are:")
				fmt.Print(string(data))
				fmt.Println()
			}
			return
		})
	mx.NewFlagV("", "test", "t").
		Description("the test text.", "").
		EnvKeys("COOLT", "TEST").
		Group("")
	mx.NewFlagV("", "password", "pp").
		Description("the password requesting.", "").
		Group("").
		Placeholder("PASSWORD").
		ExternalTool(cmdr.ExternalToolPasswordInput)
	mx.NewFlagV("", "message", "m", "msg").
		Description("the message requesting.", "").
		Group("").
		Placeholder("MESG").
		ExternalTool(cmdr.ExternalToolEditor)
	mx.NewFlagV("", "fruit", "fr").
		Description("the message.", "").
		Group("").
		Placeholder("FRUIT").
		ValidArgs("apple", "banana", "orange")
	mx.NewFlagV(1, "head", "hd").
		Description("the head lines.", "").
		Group("").
		Placeholder("LINES").
		HeadLike(true, 1, 3000)
	mx.NewFlagV(false, "stdin", "c").
		Description("read file content from stdin.", "").
		Group("")
}

func kv(root cmdr.OptCmd) {
	// kv

	kvCmd := root.NewSubCommand("kvstore", "kv").
		Description("consul kv store operations...", ``)

	attachConsulConnectFlags(kvCmd)

	kvBackupCmd := kvCmd.NewSubCommand("backup", "b", "bf", "bkp").
		Description("Dump Consul's KV database to a JSON/YAML file", ``).
		Action(kvBackup)
	kvBackupCmd.NewFlagV("consul-backup.json", "output", "o").
		Description("Write output to a file (*.json / *.yml)", ``).
		Placeholder("FILE")

	kvRestoreCmd := kvCmd.NewSubCommand("restore", "r").
		Description("restore to Consul's KV store, from a a JSON/YAML backup file", ``).
		Action(kvRestore)
	kvRestoreCmd.NewFlagV("consul-backup.json", "input", "i").
		Description("Read the input file (*.json / *.yml)", ``).
		Placeholder("FILE")
}

func ms(root cmdr.OptCmd) {
	// ms

	msCmd := root.NewSubCommand("micro-service", "ms", "microservice").
		Description("micro-service operations...", "").
		Group("")

	msCmd.NewFlagV(false, "money", "mm").
		Description("A placeholder flag.", "").
		Group("").
		Placeholder("")

	msCmd.NewFlagV("", "name", "n").
		Description("name of the service", ``).
		Placeholder("NAME")
	msCmd.NewFlagV("", "id", "i", "ID").
		Description("unique id of the service", ``).
		Placeholder("ID")
	msCmd.NewFlagV(false, "all", "a").
		Description("all services", ``).
		Placeholder("")

	msCmd.NewFlagV(3, "retry", "t").
		Description("", "").
		Group("").
		Placeholder("RETRY")

	// ms ls

	msCmd.NewSubCommand("list", "ls", "l", "lst", "dir").
		Description("list tags", "").
		Group("2333.List").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			return
		})

	tags(msCmd)
}

func tags(msCmd cmdr.OptCmd) {
	// ms tags

	msTagsCmd := msCmd.NewSubCommand("tags", "t").
		Description("tags operations of a micro-service", "").
		Group("")

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
		Titles("n", "name").
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
		Description("delimitor char in `non-plain` mode.").
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
		Description("HTTP Basic auth password", ``).
		Placeholder("PASSWORD").
		Group("Consul").
		ExternalTool(cmdr.ExternalToolPasswordInput)

}

func kvBackup(cmd *cmdr.Command, args []string) (err error) {
	// err = consul.Backup()
	return
}

func kvRestore(cmd *cmdr.Command, args []string) (err error) {
	// err = consul.Restore()
	return
}

func msList(cmd *cmdr.Command, args []string) (err error) {
	// err = consul.ServiceList()
	return
}

func msTagsList(cmd *cmdr.Command, args []string) (err error) {
	// err = consul.TagsList()
	return
}

func msTagsAdd(cmd *cmdr.Command, args []string) (err error) {
	// err = consul.Tags()
	return
}

func msTagsRemove(cmd *cmdr.Command, args []string) (err error) {
	// err = consul.Tags()
	return
}

func msTagsModify(cmd *cmdr.Command, args []string) (err error) {
	// err = consul.Tags()
	return
}

func msTagsToggle(cmd *cmdr.Command, args []string) (err error) {
	// err = consul.TagsToggle()
	return
}

const (
	appName   = "my-service"
	copyright = "my-service is an effective devops tool"
	desc      = "my-service is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "my-service is an effective devops tool. It make an demo application for `cmdr`."
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
