package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	API api `toml:"api"`
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

type ParseConfig struct {
	ConfPath string
}

func (conf ParseConfig) Prase() TomlConfig {
	configText := conf.loadFile()
	var config TomlConfig
	if _, err := toml.Decode(configText, &config); err != nil {
		panic(err)
	}

	return config
}

func (conf ParseConfig) loadFile() string {
	f, err := os.Open(conf.ConfPath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	buf := make([]byte, 4096)
	var configText string

	for n, err := f.Read(buf); err == nil; n, err = f.Read(buf) {
		configText += string(buf[:n])
	}

	return configText
}
