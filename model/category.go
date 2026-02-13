package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name       string `json:"name" gorm:"not null;unique"`
	Color      string `json:"color" gorm:"not null"`
	Visibility string `json:"visibility" gorm:"default:'public'"`
}
