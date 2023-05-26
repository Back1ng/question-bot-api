package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("placeholder")

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

			poll := tgbotapi.NewPoll(update.Message.Chat.ID, "Как у тебя настроение?", "Хорошо", "Отлично", "Пошел нахуй")

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
