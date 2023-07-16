package models

import (
	"gorm.io/gorm"
)

type Preset struct {
	gorm.Model
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Questions []Question
}
