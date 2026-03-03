package model

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	Key   string `gorm:"uniqueIndex;not null;size:191"`
	Value string
}
