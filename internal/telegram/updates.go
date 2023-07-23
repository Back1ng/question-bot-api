package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return bot.GetUpdatesChan(u)
}
