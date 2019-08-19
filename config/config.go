package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	API    api    `toml:"api"`
	Censys censys `toml:"censys"`
}

type api struct {
	Token   string  `toml:"token"`
	URL     string  `toml:"url"`
	Timeout int     `toml:"timeout"`
	Webhook webhook `toml:"webhook"`
}

type webhook struct {
	Enable          bool            `toml:"enable"`
	Listen          string          `toml:"listen"`
	WebhookTLS      webhookTLS      `toml:"tls"`
	WebhookEndpoint webhookEndpoint `toml:"endpoint"`
}

type webhookTLS struct {
	Key  string `toml:"key"`
	Cert string `toml:"cert"`
}

type webhookEndpoint struct {
	PublicURL string `toml:"publicURL"`
	Cert      string `toml:"cert"`
}

type censys struct {
	Enable    bool   `toml:"enable"`
	ApiID     string `toml:"apiID"`
	ApiSecret string `toml:"apiSecret"`
}

type Config struct {
	ConfPath string
	Config   TomlConfig
}

func (conf *Config) check() {
	if conf.Config.API.Token == "" {
		log.Panicln("API.Token can not be null.")
	}

	if conf.Config.API.Webhook.Enable && conf.Config.API.Webhook.Listen == "" {
		log.Panicln("API.Webhook.Listen can not be null when API.Webhook.Enable is true.")
	}

	if conf.Config.Censys.Enable && (conf.Config.Censys.ApiID == "" || conf.Config.Censys.ApiSecret == "") {
		log.Panicln("Centsys' ApiID and ApiSecret must be set when the censys is enable.")
	}
}

func (conf *Config) Parse() {
	configText := conf.loadFile()
	var config TomlConfig
	if _, err := toml.Decode(configText, &config); err != nil {
		panic(err)
	}
	conf.Config = config
	conf.check()
}

func (conf *Config) loadFile() string {
	f, err := os.Open(conf.ConfPath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	buf := make([]byte, 1024)
	var configText string

	for n, err := f.Read(buf); err == nil; n, err = f.Read(buf) {
		configText += string(buf[:n])
	}

	return configText
}
