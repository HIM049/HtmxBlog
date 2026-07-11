package model

type Tag struct {
	BaseModel
	Name  string  `gorm:"unique;not null"`
	Posts []*Post `json:"-" gorm:"many2many:post_tags"`
}
