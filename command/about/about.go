package about

import (
	"log"

	"github.com/AkinoKaede/Kiririn/command"
	tb "gopkg.in/tucnak/telebot.v2"
)

type About struct {
	Command *command.Command
}

func (about *About) AboutCommand(m *tb.Message) {
	log.Printf("[%s(ID:%d)]使用Censys在聊天[%s(ID:%d Type:%s Title:%s)]查询了关于信息\n", m.Sender.Username, m.Sender.ID, m.Chat.Username, m.Chat.ID, m.Chat.Type, m.Chat.Title)

	msg := "*Kiririn Bot Dev*  \nCreated by [秋のかえで](https://t.me/AkinoKaede)  \nSource Code: [AkinoKaede/Kiririn](https://github.com/AkinoKaede/Kiririn)  \n"
	_, sendErr := about.Command.Bot.Bot.Send(m.Chat, msg, tb.ModeMarkdown, tb.NoPreview)
	if sendErr != nil {
		log.Print(sendErr)
	}
}
