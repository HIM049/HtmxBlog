package model

import "gorm.io/gorm"

type Attach struct {
	gorm.Model
	Hash string `gorm:"unique" json:"hash"`
	Uid  string `gorm:"unique" json:"uid"`
	Name string `json:"name"`
	Mime string `json:"mime"`
}
