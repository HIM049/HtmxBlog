package model

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	Key   string `json:"key" gorm:"uniqueIndex;not null;size:255"`
	Value string `json:"value"`
}
