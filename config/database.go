package config

import (
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Driver string
	DSN    string
}

func ReadDatabase() Database {
	return Database{
		Driver: getEnv("DB_DRIVER"),
		DSN:    getEnv("DB_DSN"),
	}
}

var DB *gorm.DB

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
}
