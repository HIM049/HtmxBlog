package config

import (
	"HtmxBlog/model"
	"path/filepath"

	"github.com/knadh/koanf/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Driver string
	DSN    string
}

func ReadDatabase(k *koanf.Koanf) Database {
	driver := k.String("database.driver")
	if driver != "sqlite" {
		panic("database driver not supported")
	}

	return Database{
		Driver: driver,
		DSN:    k.String("database.dsn"),
	}
}

var DB *gorm.DB

const DB_PATH = "./app_data"

// InitDB initializes the database connection.
// It panics when some error occurs.
func InitDB() {
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
