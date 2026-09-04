package model

type Page struct {
	BaseModel
	Name              string `json:"name" gorm:"not null;unique"`
	Route             string `json:"route" gorm:"not null;unique"`
	Template          string `json:"template" gorm:"not null"`
	Sort              uint   `json:"sort" gorm:"index"`
	FilterMode        string `json:"filter_mode" gorm:"default:'none'"`
	FilterCategoryIDs []uint `json:"filter_category_ids" gorm:"serializer:json"`
}
