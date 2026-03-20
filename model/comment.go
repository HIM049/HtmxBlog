package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID    uint   `json:"post_id" gorm:"not null"`
	Parent    uint   `json:"parent" gorm:"default:0"`
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Url       string `json:"url"`
	UserAgent string `json:"user_agent"`
	IP        string `json:"ip"`
	State     string `json:"state" gorm:"default:'pending'"`
	Content   string `json:"content" gorm:"not null"`
}

const (
	StatePending  = "pending"
	StateApproved = "approved"
)
