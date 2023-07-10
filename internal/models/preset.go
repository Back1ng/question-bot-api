package models

import "gorm.io/gorm"

type Preset struct {
	gorm.Model
	Id        int64
	Title     string
	Questions []Question `gorm:"many2many:presets_questions;"`
}
