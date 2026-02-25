package model

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	Key   string `gorm:"uniqueIndex"`
	Value string
}
