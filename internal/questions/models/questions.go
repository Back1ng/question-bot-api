package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	ID      int64
	Title   string   `json:"title" validate:"required"`
	Answers []Answer `json:"answers" validate:"required"`
}
