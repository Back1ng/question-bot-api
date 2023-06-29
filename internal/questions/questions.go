package questions

import (
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Question struct {
	Title   string
	Answers []Answer
}

type Answer struct {
	Title     string
	IsCorrect bool
}

var questions []Question = []Question{
	{"Привет", []Answer{{"Да", false}, {"Нет", true}}},
	{"Пока", []Answer{{"Да", true}, {"Нет", false}}},
}

func GetRandom() Question {
	return questions[rand.Intn(len(questions))]
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
