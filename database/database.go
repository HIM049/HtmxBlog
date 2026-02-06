package database

import (
	"HtmxBlog/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

const DB_PATH = "./app_data/app.db"

// Init initializes the database connection.
// It panics when some error occurs.
func Init() {
	d, err := gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = d

	err = Migrate()
	if err != nil {
		panic("failed to migrate database")
	}
}

func Migrate() error {
	return db.AutoMigrate(&model.Post{})
}
