# Examples for `cmdr`

![Go](https://github.com/hedzr/cmdr-examples/workflows/Go/badge.svg)


see also [`cmdr`](https://github.com/hedzr/cmdr), [`cmdr-addons`](https://github.com/hedzr/cmdr-addons).

## Status


### v1.11.6 and newer

golang 1.17+ required.

> Since Go Modules 1.17 can't compatible with lower versions.

### v1.9.9 and newer

golang 1.16+ required.

> **Causes**:
> 1. golang.org/x/net/http2 used errors.Is()
> 2. golang.org/x/net/http2 used os.ErrDeadlineExceeded
> 
> Since cmdr-addon v1.9.8-p3
> 
> Since cmdr v1.9.9 and later

Updates:
1. removed iris/v12 [`import "github.com/hedzr/cmdr-addons v1.9.8-p3"`]
2. seems ci not good for go1.14



## Index

TODO




## `service`

new service app with [`dex` plugin](https://github.com/hedzr/cmdr-addons/tree/master/pkg/dex):

<https://github.com/hedzr/cmdr-examples/tree/master/examples/service>




## LICENSE

Feel free with MIT.


