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
	github.com/hedzr/cmdr v1.7.40
	github.com/hedzr/cmdr-addons v1.7.40
	github.com/hedzr/log v0.3.8
	github.com/hedzr/logex v1.3.8
	github.com/kardianos/service v1.1.0
	github.com/nsf/termbox-go v0.0.0-20190121233118-02980233997d
	github.com/sirupsen/logrus v1.7.0
	github.com/superhawk610/bar v0.0.0-20190614064228-4fbf44d086fd
	github.com/superhawk610/terminal v0.0.0-20200123193603-cbc69427a94a // indirect
	github.com/ttacon/chalk v0.0.0-20140724125006-76b3c8b611de
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f
	gopkg.in/AlecAivazis/survey.v1 v1.8.8
	gopkg.in/hedzr/errors.v2 v2.1.3
	gopkg.in/yaml.v2 v2.4.0
)
