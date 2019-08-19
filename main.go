package main

import (
	"flag"
	"log"

	"github.com/AkinoMaple/Kiririn/bot"
	"github.com/AkinoMaple/Kiririn/config"
	"github.com/AkinoMaple/Kiririn/handler"
)

func loadConfig() config.TomlConfig {
	confPath := flag.String("conf", "/etc/kiririn/config.toml", "Config file's path.")
	flag.Parse()

	conf := config.Config{ConfPath: *confPath}
	conf.Parse()
	return conf.Config
}

func main() {
	botConfig := loadConfig()
	var bot bot.Bot
	newErr := bot.NewBot(&botConfig)
	if newErr != nil {
		log.Panicln(newErr)
	}
	handle := handler.Handler{Bot: &bot, Config: &botConfig}
	handle.Handle()

	bot.Start()
}
