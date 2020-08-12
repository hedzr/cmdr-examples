// Copyright © 2020 Hedzr Yeh.

package cmd

import (
	"github.com/hedzr/cmdr"
)

func modifier(daemonServerCommands *cmdr.Command) *cmdr.Command {
	if startCmd := daemonServerCommands.FindSubCommand("start"); startCmd != nil {
		startCmd.PreAction = onServerPreStart
		startCmd.PostAction = onServerPostStop
	}

	return daemonServerCommands
}

func onAppStart(cmd *cmdr.Command, args []string) (err error) {
	cmdr.Logger.Debugf("onAppStart")
	return
}

func onAppExit(cmd *cmdr.Command, args []string) {
	cmdr.Logger.Debugf("onAppExit")
}

func onServerPostStop(cmd *cmdr.Command, args []string) {
	cmdr.Logger.Debugf("onServerPostStop")
}

// onServerPreStart is earlier than onAppStart.
func onServerPreStart(cmd *cmdr.Command, args []string) (err error) {
	// earlierInitLogger() // deprecated by cmdr.WithLogex()
	cmdr.Logger.Debugf("onServerPreStart")
	return
}
