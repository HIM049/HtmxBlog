package model

import "gorm.io/gorm"

type Attach struct {
	gorm.Model
	Hash   string `json:"hash" gorm:"unique"`
	Uid    string `json:"uid" gorm:"unique"`
	Name   string `json:"name"`
	Mime   string `json:"mime"`
	Refers []Post `json:"refers" gorm:"many2many:post_attaches"`
}
