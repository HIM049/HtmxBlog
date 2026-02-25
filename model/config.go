package model

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	Key   string `json:"key" gorm:"uniqueIndex"`
	Value string `json:"value"`
}
