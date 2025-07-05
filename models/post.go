package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Uid     string   `json:"uid" gorm:"uniqueIndex;not null"`
	Title   string   `json:"title" gorm:"not null"`
	Tags    []string `json:"tags" gorm:"type:text[]"`
	Content string   `json:"content" gorm:"type:text;not null"`
}
