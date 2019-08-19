package handler

import (
	"github.com/AkinoKaede/Kiririn/command"
	"github.com/AkinoKaede/Kiririn/command/about"
	"github.com/AkinoKaede/Kiririn/command/censys"
)

type commandHandler struct {
	Handler *Handler
	Command *command.Command
}

func (ch *commandHandler) aboutHealder() {
	a := about.About{Command: ch.Command}
	ch.Handler.Bot.Bot.Handle("/about", a.AboutCommand)
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

	ch.aboutHealder()
	ch.censysHealder()
}
