package entity

import (
	"errors"
	"math/rand"
)

type User struct {
	ID              int64 `json:"id"`
	ChatId          int64 `json:"chat_id"`
	PresetId        int64 `json:"preset_id"`
	Preset          Preset
	Nickname        string `json:"nickname"`
	Interval        int    `json:"interval"`
	IntervalEnabled bool   `json:"interval_enabled"`
}

func (u *User) GetQuestion() Question {
	// database.Database.DB.Preload("Preset.Questions.Answers").Find(u)

	if len(u.Preset.Questions) == 0 {
		return Question{}
	}

	return u.Preset.Questions[rand.Intn(len(u.Preset.Questions))]
}

func (u *User) UpdateInterval(i int) error {
	if i < 1 || i > 24 {
		return errors.New("given interval must be between 1-24")
	}

	u.Interval = i

	return nil
}