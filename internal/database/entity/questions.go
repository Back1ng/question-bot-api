package entity

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Question struct {
	ID       int64 `json:"id"`
	PresetId int64 `json:"preset_id"`
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

	correctCount := 0
	for i, answer := range q.Answers {
		if answer.IsCorrect {
			poll.CorrectOptionID = int64(i)
			correctCount++
		}
	}

	if correctCount > 1 {
		poll.AllowsMultipleAnswers = true
	}

	poll.Type = "quiz"

	return poll
}
