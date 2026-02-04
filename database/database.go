package database

import (
	"HtmxBlog/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database connection.
// It panics when some error occurs.
func Init() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}

func Migrate() error {
	return DB.AutoMigrate(&model.Post{})
}
