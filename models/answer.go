package models

import (
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	Content    string
	QuestionID uint
	Votes      int
}
