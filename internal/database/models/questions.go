package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID       int64
	PresetId int64
	Preset   Preset
	Title    string   `json:"title" validate:"required"`
	Answers  []Answer `json:"answers" validate:"required"`
}

func (q *Question) CreatePoll(chatId int64) tgbotapi.SendPollConfig {

	answers := make([]string, len(q.Answers))

	for i, answer := range q.Answers {
		answers[i] = answer.Title
	}

	poll := tgbotapi.NewPoll(chatId, q.Title, answers...)

	for i, answer := range q.Answers {
		if answer.IsCorrect {
			poll.CorrectOptionID = int64(i)
		}
	}

	poll.Type = "quiz"

	return poll
}
