package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	ID         int64 `json:"id"`
	QuestionId int64 `json:"question_id"`
	Question   Question
	Title      string `json:"title" validate:"required"`
	IsCorrect  bool   `json:"is_correct" validate:"required"`
}
