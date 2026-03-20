package config

import (
	"HtmxBlog/model"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"github.com/knadh/koanf/v2"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Driver string
	DSN    string
}

func ReadDatabase(k *koanf.Koanf) Database {
	driver := k.String("database.driver")

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
	var db *gorm.DB
	var err error
	switch Cfg.Database.Driver {
	case "sqlite":
		dbPath := filepath.Join(DB_PATH, Cfg.Database.DSN)
		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{TranslateError: true})
	case "mysql":
		db, err = gorm.Open(mysql.Open(Cfg.Database.DSN), &gorm.Config{TranslateError: true})
	case "postgres":
		db, err = gorm.Open(postgres.Open(Cfg.Database.DSN), &gorm.Config{TranslateError: true})
	default:
		panic("database driver not supported")
	}

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
	return DB.AutoMigrate(&model.Post{}, &model.Page{}, &model.Attach{}, &model.Setting{}, &model.Comment{})
}
