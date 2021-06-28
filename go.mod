module github.com/hedzr/cmdr-examples

go 1.14

//replace github.com/hedzr/log => ../log

//replace github.com/hedzr/logex => ../logex

//replace github.com/hedzr/cmdr => ../cmdr

//replace github.com/hedzr/cmdr-addons => ../cmdr-addons

//replace github.com/kardianos/service => ../../kardianos/service

//replace github.com/hedzr/cmdr-examples/cmdr-gen => ./cmdr-gen

require (
	github.com/gizak/termui/v3 v3.1.0
	github.com/hedzr/cmdr v1.9.0
	github.com/hedzr/cmdr-addons v1.9.0
	github.com/hedzr/log v1.3.22
	github.com/hedzr/logex v1.3.22
	github.com/kardianos/service v1.2.0
	github.com/nsf/termbox-go v1.1.1
	github.com/sirupsen/logrus v1.8.1
	github.com/superhawk610/bar v0.0.2
	github.com/superhawk610/terminal v0.1.0 // indirect
	github.com/ttacon/chalk v0.0.0-20160626202418-22c06c80ed31
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22
	gopkg.in/AlecAivazis/survey.v1 v1.8.8
	gopkg.in/hedzr/errors.v2 v2.1.5
	gopkg.in/yaml.v2 v2.4.0
)
