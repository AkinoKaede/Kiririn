package command

import (
	"github.com/AkinoMaple/Kiririn/bot"
	"github.com/AkinoMaple/Kiririn/config"
	"github.com/AkinoMaple/Kiririn/model"
)

type Command struct {
	Bot    *bot.Bot
	Config *config.TomlConfig
	Model  *model.Model
}

func (cmd *Command) SetModel() {
	cmd.Model = &model.Model{Config: cmd.Config}
}
