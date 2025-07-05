package db

import (
	"HtmxBlog/config"
	"HtmxBlog/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	// Check the type of database driver
	var dialector gorm.Dialector
	switch config.Cfg.Database.Driver {
	case "sqlite":
		dialector = sqlite.Open(config.Cfg.Database.DSN)
	case "mysql":
		dialector = mysql.Open(config.Cfg.Database.DSN)
	default:
		return gorm.ErrUnsupportedDriver

	}

	// Connect to the database
	if err := connectToDatabase(dialector); err != nil {
		return err
	}

	// Migrate the database models
	if err := migrateModel(); err != nil {
		return err
	}

	return nil
}

// Migrate the database models
func migrateModel() error {
	modelsToMigrate := []interface{}{
		&models.Post{},
	}

	for _, model := range modelsToMigrate {
		err := db.AutoMigrate(model)
		if err != nil {
			return err
		}

	}
	return nil
}

// Initialize the database connection
func connectToDatabase(dialector gorm.Dialector) error {
	d, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}

	db = d
	return nil
}
