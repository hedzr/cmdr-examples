# kilo-bytes

`cmdr` provides extracting `kilobytes` value from Option Store. It allows human-readable sizes input from command line.

The samples:

```bash
$ go run ./examples/kilo-bytes/ kb --size 5kb
$ go run ./examples/kilo-bytes/ kb --size 8T
$ go run ./examples/kilo-bytes/ kb --size 1g
$ go run ./examples/kilo-bytes/ kb --size 329eb
```

Sometimes, it's called as `kibibytes`.

- See also: https://en.wikipedia.org/wiki/Kibibyte
- Its related word is kilobyte, refer to: https://en.wikipedia.org/wiki/Kilobyte



|                |                    |
| -------------- | ------------------ |
| Valid formats  | 2k, 2kb, 2kB, 2KB. |
| Valid Suffixes | k, m, g, t, p, e.  |
|                |                    |

### Backstage

See also `cmdr.GetKibibytesR()`.

```go
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
			fmt.Printf("Got kilo: %v (literal: %v)\n\n", cmdr.GetKibibytesR("kb-print.kilo"), cmdr.GetStringR("kb-print.kilo"))
			return
		})

	kb.NewFlagV("1k", "size", "s").
		Description("max message size. Valid formats: 2k, 2kb, 2kB, 2KB. Suffixes: k, m, g, t, p, e.", "").
		Group("")	
	
	cmdr.NewString("1k").
		Titles("kilo", "k").
		Description("message size. Valid formats: 2k, 2kb, 2kB, 2KB. Suffixes: k, m, g, t, p, e.", "").
		Group("").AttachTo(kb)

}
```
