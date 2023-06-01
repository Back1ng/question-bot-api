package questions

import (
	"math/rand"
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
