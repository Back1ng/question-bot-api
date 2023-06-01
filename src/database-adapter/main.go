package database_adapter

import (
	"fmt"
)

type DB struct {
	strategy string
}

type PollAnswer struct {
	UserId      uint64
	TelegramTag string
	Question    string
	Answers     map[int]string
}

type DatabaseActions interface {
	SetStrategyDatabase(strategy string) DB
	Get() map[int]PollAnswer
}

type void interface{}

func (strategy string) SetStrategyDatabase() DB {
	return DB{strategy: strategy}
}

func main() {
	//todo выбор стратегии, а потом получать. Сделать получение в зависимости от стратегий
}

func (_) Get() map[int]PollAnswer {
	var data map[int]PollAnswer
	answers := make(map[int]string)
	tag := "tag"
	question := "question"
	answer := "answer"

	for i := 0; i < 5; i++ {
		answers[i] = fmt.Sprint(i, answer)
		data[i] = PollAnswer{
			UserId:      uint64(i),
			TelegramTag: fmt.Sprint(i, tag),
			Question:    fmt.Sprint(i, question),
			Answers:     answers,
		}
	}

	return data
}
