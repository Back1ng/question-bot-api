package database_adapter

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

// что за (strategy string)???? Надо (s *DB), а уже у него вызываешь s.strategy
// или нормально создаешь метод func SetStrategyDatabase(s string) DB {return DB{strategy: s}}

// func (strategy string) SetStrategyDatabase() DB {
// 	return DB{strategy: strategy}
// }

func main() {
	//todo выбор стратегии, а потом получать. Сделать получение в зависимости от стратегий
}

// нельзя ссылаться на (_)
// func (_) Get() map[int]PollAnswer {
// 	var data map[int]PollAnswer
// 	answers := make(map[int]string)
// 	tag := "tag"
// 	question := "question"
// 	answer := "answer"

// 	for i := 0; i < 5; i++ {
// 		answers[i] = fmt.Sprint(i, answer)
// 		data[i] = PollAnswer{
// 			UserId:      uint64(i),
// 			TelegramTag: fmt.Sprint(i, tag),
// 			Question:    fmt.Sprint(i, question),
// 			Answers:     answers,
// 		}
// 	}

// 	return data
// }
