package entity

type Answer struct {
	ID         int64  `json:"id"`
	QuestionId int64  `json:"question_id"`
	Title      string `json:"title"`
	IsCorrect  bool   `json:"is_correct"`
}
