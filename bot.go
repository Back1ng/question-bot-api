package main

import (
	"fmt"
	"log"

	questions "gitlab.com/back1ng1/question-bot/internal/questions"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6124185313:AAH0X7cTdWfPD0ExMkXekGJPx-p1Ke2HBhQ")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			question := questions.GetRandom()

			poll := question.CreatePoll(update.Message.Chat.ID)

			bot.Send(poll)

			// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID

			// bot.Send(msg)
		}

		if update.Poll != nil {
			for id, option := range update.Poll.Options {
				if option.VoterCount > 0 {
					fmt.Printf("Selected{id: %v, value:%s}", id, option.Text)
				}
			}
		}
	}
}
