package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	ID         int64
	QuestionId int64
	Title      string `json:"title" validate:"required"`
	IsCorrect  bool   `json:"is_correct" validate:"required"`
}
