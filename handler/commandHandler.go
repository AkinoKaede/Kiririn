package handler

import (
	"github.com/AkinoMaple/Kiririn/command"
	"github.com/AkinoMaple/Kiririn/command/censys"
)

type commandHandler struct {
	Handler *Handler
	Command *command.Command
}

func (ch *commandHandler) censysHealder() {
	if ch.Handler.Config.Censys.Enable {
		cs := censys.Censys{Command: ch.Command}
		ch.Handler.Bot.Bot.Handle("/censys", cs.CensysCommand)
	}
}

func (ch *commandHandler) CommandHandler() {
	ch.Command = &command.Command{Bot: ch.Handler.Bot, Config: ch.Handler.Config}
	ch.Command.SetModel()
	ch.censysHealder()
}
