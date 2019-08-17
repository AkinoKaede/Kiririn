package main

import (
	"flag"
	"fmt"

	"github.com/AkinoMaple/Kiririn/config"
)

func loadConfig() config.TomlConfig {
	confPath := flag.String("conf", "/etc/kiririn/config.toml", "Config file's path.")
	flag.Parse()

	conf := config.ParseConfig{ConfPath: *confPath}

	return conf.Prase()
}

func main() {
	conf := loadConfig()
	fmt.Println(conf)
}
