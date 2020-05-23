# envvars

To bind one or more envvars (`OS Environment Variables`) to a 
flag, use `.EnvKeys(k1,k2,...)`:

```go
	cmdr.NewFloat64(3.14159265358979323846264338327950288419716939937510582097494459230781640628620899).
		Titles("float64", "f64").
		Description("A float64 flag with a `PI` value", "").
		Group("2000.Float").
		EnvKeys("PI").
		AttachTo(parent)
```

Now test it with the following commands:

```bash
PI=3.1 go run ./examples/envvars/ f ~~debug|grep 'app\.flags\.float64'
PI=3.2 go run ./examples/envvars/ f ~~debug|grep 'app\.flags\.float64'
```
