module github.com/hedzr/cmdr-examples

go 1.14

//replace github.com/hedzr/log => ../10.log

//replace github.com/hedzr/logex => ../15.logex

//replace github.com/hedzr/cmdr => ../50.cmdr

//replace github.com/hedzr/cmdr-addons => ../53.cmdr-addons

//replace github.com/kardianos/service => ../../kardianos/service

//replace github.com/hedzr/cmdr-examples/cmdr-gen => ./cmdr-gen

require (
	github.com/gizak/termui/v3 v3.1.0
	github.com/hedzr/cmdr v1.9.8
	github.com/hedzr/cmdr-addons v1.9.8
	github.com/hedzr/log v1.5.0
	github.com/hedzr/logex v1.5.0
	github.com/kardianos/service v1.2.1
	github.com/nsf/termbox-go v1.1.1
	github.com/sirupsen/logrus v1.8.1
	github.com/superhawk610/bar v0.0.2
	github.com/superhawk610/terminal v0.1.0 // indirect
	github.com/ttacon/chalk v0.0.0-20160626202418-22c06c80ed31
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b
	gopkg.in/AlecAivazis/survey.v1 v1.8.8
	gopkg.in/hedzr/errors.v2 v2.1.5
	gopkg.in/yaml.v2 v2.3.0
)
