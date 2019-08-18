package bot

import (
	"time"

	"github.com/AkinoMaple/Kiririn/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Bot struct {
	Bot *tb.Bot
}

func (b *Bot) NewBot(config *config.TomlConfig) error {
	var pref tb.Settings
	if config.API.Webhook.Enable {
		poll := &tb.Webhook{
			Listen: config.API.Webhook.Listen,
			TLS: &tb.WebhookTLS{
				Key:  config.API.Webhook.WebhookTLS.Key,
				Cert: config.API.Webhook.WebhookTLS.Cert,
			},
			Endpoint: &tb.WebhookEndpoint{
				PublicURL: config.API.Webhook.WebhookEndpoint.PublicURL,
				Cert:      config.API.Webhook.WebhookEndpoint.Cert,
			},
		}

		pref = tb.Settings{
			Token:  config.API.Token,
			Poller: poll,
		}
	} else {
		poll := &tb.LongPoller{
			Timeout: time.Duration(config.API.Timeout) * time.Second,
		}

		pref = tb.Settings{
			Token:  config.API.Token,
			Poller: poll,
		}
	}

	bot, err := tb.NewBot(pref)
	if err == nil {
		b.Bot = bot
	}

	return err
}

func (b *Bot) Start() {
	b.Bot.Start()
}
