package censys

import (
	"fmt"
	"log"

	"github.com/AkinoKaede/Kiririn/command"
	csm "github.com/AkinoKaede/Kiririn/model/censys"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Censys struct {
	Command *command.Command
}

func (c *Censys) CensysCommand(m *tb.Message) {
	if m.Payload == "" {
		_, err := c.Command.Bot.Bot.Reply(m, "Usage: /censys domain  \n", tb.ModeMarkdown)
		if err != nil {
			log.Print(err)
		}

		return
	}

	log.Printf("[%s(ID:%d)]使用Censys在聊天[%s(ID:%d Type:%s Title:%s)]查询了%s\n", m.Sender.Username, m.Sender.ID, m.Chat.Username, m.Chat.ID, m.Chat.Type, m.Chat.Title, m.Payload)

	cs := csm.Censys{Model: *c.Command.Model}
	results, err := cs.Query(m.Payload)
	if *err != nil {
		log.Print(*err)
		return
	}

	var replyErr error

	if results.Metadata.Count == 0 {
		msg := fmt.Sprintf("*查询内容: %s*  \n未能找到结果  \n", m.Payload)
		_, replyErr = c.Command.Bot.Bot.Reply(m, msg, tb.ModeMarkdown)
	} else {
		msg := fmt.Sprintf("*查询内容: %s*  \n找到 *%d* 条结果  \n  \n查询结果:  \n", m.Payload, results.Metadata.Count)
		for _, result := range results.Results {
			msg += fmt.Sprintf("IP: %s  \n协议: %s  \n位置: %s, %s, %s  \n座标: %f, %f  \n时区: %s  \n域名: %s  \n  \n", result.IP, result.Protocols, result.City, result.Province, result.Country, result.Latitude, result.Longitude, result.TimeZone, result.Domain)
		}

		_, replyErr = c.Command.Bot.Bot.Reply(m, msg, tb.ModeMarkdown)
	}

	if replyErr != nil {
		log.Print(replyErr)
	}

}
