package main

import (
	"flag"
	"log"
	"os"

	"github.com/AkinoKaede/Kiririn/bot"
	"github.com/AkinoKaede/Kiririn/config"
	"github.com/AkinoKaede/Kiririn/handler"
)

func loadConfig() config.TomlConfig {
	wd, _ := os.Getwd()
	confPath := flag.String("config", wd+"/config.toml", "Config file's path.")
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
