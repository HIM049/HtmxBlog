package model

import "gorm.io/gorm"

type Page struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null;unique"`
	Route    string `json:"route" gorm:"not null;unique"`
	Template string `json:"template" gorm:"not null"`
	Sort     uint   `json:"sort" gorm:"index"`
}
