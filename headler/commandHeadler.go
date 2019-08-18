package headler

import (
	"github.com/AkinoMaple/Kiririn/command"
	"github.com/AkinoMaple/Kiririn/command/censys"
)

type commandHeadler struct {
	Headler *Headler
	Command *command.Command
}

func (ch *commandHeadler) censysHealder() {
	cs := censys.Censys{Command: ch.Command}
	ch.Headler.Bot.Bot.Handle("/censys", cs.CensysCommand)
}

func (ch *commandHeadler) CommandHeadler() {
	ch.Command = &command.Command{Bot: ch.Headler.Bot, Config: ch.Headler.Config}
	ch.Command.SetModel()
	ch.censysHealder()
}
