package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Uid         string                 `json:"uid" gorm:"unique"`
	Title       string                 `json:"title"`
	Category    Category               `json:"category" gorm:"embedded;embeddedPrefix:category_"`
	Tags        []string               `json:"tags" gorm:"serializer:json"`
	ContentPath string                 `json:"content_path" gorm:"not null"`
	Attachs     []Attach               `json:"attachs" gorm:"many2many:post_attaches"`
	CustomVars  map[string]interface{} `json:"custom_vars" gorm:"serializer:json"`
}

type Category struct {
	Name  string `json:"name" gorm:"not null"`
	Color string `json:"color" gorm:"not null"`
}
