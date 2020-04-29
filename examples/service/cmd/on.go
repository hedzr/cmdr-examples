// Copyright © 2020 Hedzr Yeh.

package cmd

import (
	"github.com/hedzr/cmdr"
	"github.com/sirupsen/logrus"
)

func modifier(daemonServerCommands *cmdr.Command) *cmdr.Command {
	if startCmd := daemonServerCommands.FindSubCommand("start"); startCmd != nil {
		startCmd.PreAction = onServerPreStart
		startCmd.PostAction = onServerPostStop
	}

	return daemonServerCommands
}

func onAppStart(cmd *cmdr.Command, args []string) (err error) {
	logrus.Debug("onAppStart")
	return
}

func onAppExit(cmd *cmdr.Command, args []string) {
	logrus.Debug("onAppExit")
}

func onServerPostStop(cmd *cmdr.Command, args []string) {
	logrus.Debug("onServerPostStop")
}

// onServerPreStart is earlier than onAppStart.
func onServerPreStart(cmd *cmdr.Command, args []string) (err error) {
	// earlierInitLogger() // deprecated by cmdr.WithLogex()
	logrus.Debug("onServerPreStart")
	return
}
