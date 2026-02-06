package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string                 `json:"title"`
	Category    Category               `gorm:"embedded;embeddedPrefix:category_" json:"category"`
	Tags        []string               `gorm:"serializer:json" json:"tags"`
	ContentPath string                 `gorm:"not null" json:"content_path"`
	CustomVars  map[string]interface{} `gorm:"serializer:json" json:"custom_vars"`
}

type Category struct {
	Name  string `gorm:"not null" json:"name"`
	Color string `gorm:"not null" json:"color"`
}
