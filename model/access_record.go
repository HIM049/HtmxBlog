package model

import "time"

type AccessRecord struct {
	ID uint `gorm:"primaryKey"`

	CreatedAt time.Time `gorm:"index"`

	IP        string `gorm:"size:45;index"`
	UserAgent string `gorm:"type:text"`
	Referer   string `gorm:"type:text"`

	Method string `gorm:"size:10"`
	Path   string `gorm:"size:512;index"`
	Query  string `gorm:"type:text"`
}
