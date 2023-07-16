package models

import (
	"math/rand"

	"gitlab.com/back1ng1/question-bot/internal/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int64
	ChatId   int64
	PresetId int64
	Preset   Preset
}

func (u *User) GetQuestion() Question {
	database.Database.DB.Preload("Preset.Questions.Answers").Find(u)

	return u.Preset.Questions[rand.Intn(len(u.Preset.Questions))]
}
