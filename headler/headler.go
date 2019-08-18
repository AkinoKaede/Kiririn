package headler

import (
	"github.com/AkinoMaple/Kiririn/bot"
	"github.com/AkinoMaple/Kiririn/config"
)

type Headler struct {
	Bot    *bot.Bot
	Config *config.TomlConfig
}

func (h *Headler) commandHeadle() {
	ch := commandHeadler{Headler: h}
	ch.CommandHeadler()
}

func (h *Headler) Headle() {
	h.commandHeadle()
}
