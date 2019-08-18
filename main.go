package main

import (
	"flag"
	"log"

	"github.com/AkinoMaple/Kiririn/bot"
	"github.com/AkinoMaple/Kiririn/config"
	"github.com/AkinoMaple/Kiririn/headler"
)

func loadConfig() config.TomlConfig {
	confPath := flag.String("conf", "/etc/kiririn/config.toml", "Config file's path.")
	flag.Parse()

	conf := config.Config{ConfPath: *confPath}
	conf.Prase()
	return conf.Config
}

func main() {
	botConfig := loadConfig()
	var bot bot.Bot
	newErr := bot.NewBot(&botConfig)
	if newErr != nil {
		log.Panicln(newErr)
	}
	headle := headler.Headler{Bot: &bot, Config: &botConfig}
	headle.Headle()

	bot.Start()
}
