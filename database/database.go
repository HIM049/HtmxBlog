package database

import (
	"HtmxBlog/model"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

const DB_PATH = "./app_data"

// Init initializes the database connection.
// It panics when some error occurs.
func Init() {
	dbPath := filepath.Join(DB_PATH, "app.db")
	d, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
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
