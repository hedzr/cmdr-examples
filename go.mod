module github.com/hedzr/cmdr-examples

go 1.14

// replace github.com/hedzr/logex => ../logex

replace github.com/hedzr/cmdr-addons => ../cmdr-addons

require (
	github.com/hedzr/cmdr v1.6.35
	github.com/hedzr/logex v1.1.8
	github.com/kardianos/service v1.0.0
	github.com/sirupsen/logrus v1.5.0
	gopkg.in/hedzr/errors.v2 v2.0.11
)
