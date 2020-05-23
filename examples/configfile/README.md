# configfile

`cmdr` will try locating `<appname>.{yml,yaml,json,toml}` in current directory or these locations:

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

So, run and test `configfile` with the following bash commands:

```bash
cd ./examples/configfile
go run . f -i 79 -i64=131 ~~debug | grep 'app\.config-file\.'
```

and these results should be printed:

```bash
app.config-file.bool                             => true
app.config-file.int                              => 9
app.config-file.string                           => string
app.config-file.updated                          => true
```

To show the value type info in the `~~debug` output, uses `~~value-type`:

```bash
cd ./examples/configfile
go run . f -i 79 -i64=131 ~~debug ~~value-type | grep 'app\.config-file\.'
```

The result is:

```bash
app.config-file.bool                             => true (bool)
app.config-file.int                              => 9 (int)
app.config-file.string                           => string (string)
app.config-file.updated                          => true (bool)
```
