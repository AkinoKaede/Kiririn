package handler

import (
	"github.com/AkinoMaple/Kiririn/bot"
	"github.com/AkinoMaple/Kiririn/config"
)

type Handler struct {
	Bot    *bot.Bot
	Config *config.TomlConfig
}

func (h *Handler) commandHandle() {
	ch := commandHandler{Handler: h}
	ch.CommandHandler()
}

func (h *Handler) Handle() {
	h.commandHandle()
}
