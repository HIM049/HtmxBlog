package model

type Page struct {
	Name  string `json:"name" gorm:"not null;unique"`
	Route string `json:"route" gorm:"not null;unique"`
}
