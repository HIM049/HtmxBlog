package model

type Page struct {
	Name  string `gorm:"not null;unique"`
	Route string `gorm:"not null;unique"`
}
