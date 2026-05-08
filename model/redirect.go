package model

import (
	"gorm.io/gorm"
)

type Redirect struct {
	gorm.Model
	SourcePath string `gorm:"uniqueIndex;not null" json:"source_path"`
	TargetPath string `gorm:"not null" json:"target_path"`
	StatusCode int    `gorm:"not null;default:301" json:"status_code"`
}
