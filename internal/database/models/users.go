package models

import (
	"math/rand"

	"gitlab.com/back1ng1/question-bot/internal/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int64 `json:"id"`
	ChatId   int64
	PresetId int64 `json:"preset_id"`
	Preset   Preset
	Nickname string `json:"nickname"`
}

func (u *User) GetQuestion() Question {
	database.Database.DB.Preload("Preset.Questions.Answers").Find(u)

	if len(u.Preset.Questions) == 0 {
		return Question{}
	}

	return u.Preset.Questions[rand.Intn(len(u.Preset.Questions))]
}
