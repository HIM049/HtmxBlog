package model

import "gorm.io/gorm"

type Attach struct {
	gorm.Model
	Hash       string `json:"hash" gorm:"unique"`
	Uid        string `json:"uid" gorm:"unique"`
	Name       string `json:"name"`
	Mime       string `json:"mime"`
	Permission string `json:"permission" gorm:"default:'private'"`
	Refers     []Post `json:"refers" gorm:"many2many:post_attaches"`
}

const (
	PermissionPublic  = "public"
	PermissionLogin   = "login"
	PermissionPrivate = "private"
)
