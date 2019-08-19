package handler

import (
	"github.com/AkinoKaede/Kiririn/bot"
	"github.com/AkinoKaede/Kiririn/config"
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
