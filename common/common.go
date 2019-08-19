package common

import (
	"github.com/AkinoKaede/Kiririn/bot"
	"github.com/AkinoKaede/Kiririn/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Common struct {
	Bot    *bot.Bot
	Config *config.TomlConfig
}

func (c *Common) IsCreator(user *tb.User, chat *tb.Chat) (bool, *error) {
	chatUsers, err := c.Bot.Bot.AdminsOf(chat)
	for _, chatUser := range chatUsers {
		if *user == *chatUser.User && chatUser.Role == tb.Creator {
			return true, &err
		}
	}
	return false, &err
}

func (c *Common) IsAdmin(user *tb.User, chat *tb.Chat) (bool, *error) {
	chatUsers, err := c.Bot.Bot.AdminsOf(chat)
	for _, chatUser := range chatUsers {
		if *user == *chatUser.User && chatUser.Role == tb.Administrator {
			return true, &err
		}
	}
	return false, &err
}

func (c *Common) IsCreatorOrAdmin(user *tb.User, chat *tb.Chat) (bool, *error) {
	chatUsers, err := c.Bot.Bot.AdminsOf(chat)
	for _, chatUser := range chatUsers {
		if *user == *chatUser.User && (chatUser.Role == tb.Creator || chatUser.Role == tb.Administrator) {
			return true, &err
		}
	}
	return false, &err
}

func (c *Common) IsMaster(user *tb.User) bool {
	for _, masterID := range c.Config.Settings.Masters.MastersID {
		if user.ID == masterID {
			return true
		}
	}

	return false
}

func (c *Common) StoredMessage(messageID string, chatID int64) tb.StoredMessage {
	return tb.StoredMessage{MessageID: messageID, ChatID: chatID}
}
