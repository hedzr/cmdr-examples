# `service`

Its key name is `my-service`, that is the `appName` ([root.go#L32](https://github.com/hedzr/cmdr-examples/blob/master/examples/service/cmd/root.go#L32) set to cmdr.

`/my-service.yml` is the major config file, and the files under `/conf.d` would be scanned if exists.


## Sample

new service app with [`dex` plugin](https://github.com/hedzr/cmdr-addons/tree/master/pkg/dex):

<https://github.com/hedzr/cmdr-examples/tree/master/examples/service>

### NOTE

`kataras/iris/v12` can't be built on go 1.12 or lower.

