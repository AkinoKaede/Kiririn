package censys

import (
	"fmt"
	"log"

	"github.com/AkinoMaple/Kiririn/command"
	csm "github.com/AkinoMaple/Kiririn/model/censys"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Censys struct {
	Command *command.Command
}

func (c *Censys) CensysCommand(m *tb.Message) {
	if m.Payload == "" {
		_, err := c.Command.Bot.Bot.Send(m.Chat, "Usage: /censys domain\n")
		if err != nil {
			log.Println(err)
		}

		return
	}

	log.Printf("[%s]使用Censys查询了%s\n", m.Sender.Username, m.Payload)

	cs := csm.Censys{Model: *c.Command.Model}
	results, err := cs.Query(m.Payload)
	if err != nil {
		log.Println(err)
		return
	}

	var sendErr error

	if results.Metadata.Count == 0 {
		msg := fmt.Sprintf("查询内容:%s\n查询结果:无\n", m.Payload)
		_, sendErr = c.Command.Bot.Bot.Send(m.Chat, msg)
	} else {
		msg := fmt.Sprintf("查询内容:%s\n查询结果:\nIP:%s\n协议:%s\n国家:%s\n注册国家:%s\n经度:%f\n纬度:%f\n省/州:%s\n城市:%s\n时区:%s\n", m.Payload, results.Results[0].IP, results.Results[0].Protocols, results.Results[0].Country, results.Results[0].RegisteredCountry, results.Results[0].Longitude, results.Results[0].Latitude, results.Results[0].Province, results.Results[0].City, results.Results[0].TimeZone)
		_, sendErr = c.Command.Bot.Bot.Send(m.Chat, msg)
	}

	if sendErr != nil {
		log.Println(sendErr)
	}

}
