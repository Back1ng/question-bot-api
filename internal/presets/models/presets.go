package models

import (
	"gitlab.com/back1ng1/question-bot/internal/questions/models"
	"gorm.io/gorm"
)

type Preset struct {
	gorm.Model
	Id        int64
	Title     string
	Questions []models.Question `gorm:"many2many:presets_questions;"`
}
