# logging

`cmdr` will try locating `<appname>.{yml,yaml,json,toml}` in the current directory or these locations:

```go
		predefinedLocations: []string{
			"./ci/etc/%s/%s.yml",       // for developer
			"/etc/%s/%s.yml",           // regular location
			"/usr/local/etc/%s/%s.yml", // regular macOS HomeBrew location
			"$HOME/.config/%s/%s.yml",  // per user
			"$HOME/.%s/%s.yml",         // ext location per user
			"$THIS/%s.yml", // executable's directory
			"%s.yml",       // current directory
		},
```

Once the major config file (`<appname>.{yml,yaml,json,toml}`) found 
in a location such as `ABC`, `cmdr` would try locating the child 
directory `conf.d` under `ABC`, and loading and merging all valid
config files in this directory.



## Run and test

So, run and test `logging` with the following bash commands:

```bash
cd ./examples/logging
go run . f -i 79 -i64=131 ~~debug | grep 'app\.config-file\.'
```

and these results should be printed:

```bash
app.logging.bool                             => true
app.logging.int                              => 9
app.logging.string                           => string
app.logging.updated                          => true
```

To show the value type info in the `~~debug` output, uses `~~value-type`:

```bash
cd ./examples/logging
go run . f -i 79 -i64=131 ~~debug ~~value-type | grep 'app\.config-file\.'
```

The result is:

```bash
app.logging.bool                             => true (bool)
app.logging.int                              => 9 (int)
app.logging.string                           => string (string)
app.logging.updated                          => true (bool)
```
