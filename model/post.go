package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string                 `json:"title"`
	Category    string                 `json:"category"`
	Tags        []string               `gorm:"serializer:json" json:"tags"`
	ContentPath string                 `gorm:"not null" json:"content_path"`
	CustomVars  map[string]interface{} `gorm:"serializer:json" json:"custom_vars"`
}
