package command

import (
	"github.com/AkinoKaede/Kiririn/bot"
	"github.com/AkinoKaede/Kiririn/config"
	"github.com/AkinoKaede/Kiririn/model"
)

type Command struct {
	Bot    *bot.Bot
	Config *config.TomlConfig
	Model  *model.Model
}

func (cmd *Command) SetModel() {
	cmd.Model = &model.Model{Config: cmd.Config}
}
