module github.com/hedzr/cmdr-examples

go 1.14

// replace github.com/hedzr/logex => ../logex

// replace github.com/hedzr/cmdr => ../cmdr

replace github.com/hedzr/cmdr-addons => ../cmdr-addons

// replace github.com/kardianos/service => ../../kardianos/service

require (
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/hedzr/cmdr v1.6.39
	github.com/hedzr/cmdr-addons v1.0.1
	github.com/hedzr/logex v1.1.8
	github.com/kardianos/service v1.0.0
	github.com/klauspost/compress v1.10.4 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7 // indirect
	golang.org/x/sys v0.0.0-20200519105757-fe76b779f299 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/hedzr/errors.v2 v2.0.12
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
