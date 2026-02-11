package model

import (
	"path/filepath"

	"gorm.io/gorm"
)

const POSTS_DIR = "./app_data/posts"

type Post struct {
	gorm.Model
	Uid        string `json:"uid" gorm:"unique"`
	Visibility string `json:"visibility" gorm:"default:'private'"`
	Protect    string `json:"protected" gorm:"default:'none'"`
	State      string `json:"state" gorm:"default:'draft'"`

	Title      string                 `json:"title"`
	CategoryID uint                   `json:"category_id"`
	Category   Category               `json:"category" gorm:"foreignKey:CategoryID"`
	Tags       []string               `json:"tags" gorm:"serializer:json"`
	Attachs    []Attach               `json:"attachs" gorm:"many2many:post_attaches"`
	CustomVars map[string]interface{} `json:"custom_vars" gorm:"serializer:json"`
}

func (p *Post) ContentPath() string {
	return filepath.Join(POSTS_DIR, p.Uid)
}
