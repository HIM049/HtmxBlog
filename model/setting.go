package model

type Setting struct {
	BaseModel
	Key   string `gorm:"uniqueIndex;not null;size:191"`
	Value string
}
