package model

type Category struct {
	BaseModel
	Name       string `json:"name" gorm:"not null;unique"`
	Color      string `json:"color" gorm:"not null"`
	Visibility string `json:"visibility" gorm:"default:'public'"`
}
