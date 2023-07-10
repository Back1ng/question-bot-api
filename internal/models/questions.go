package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Id    int64
	Title string
}
