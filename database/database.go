package database

import (
	"HtmxBlog/model"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DB_PATH = "./app_data"

// Init initializes the database connection.
// It panics when some error occurs.
func Init() {
	dbPath := filepath.Join(DB_PATH, "app.db")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	err = Migrate()
	if err != nil {
		panic("failed to migrate database")
	}
}

func Migrate() error {
	return DB.AutoMigrate(&model.Post{}, &model.Page{}, &model.Attach{})
}
