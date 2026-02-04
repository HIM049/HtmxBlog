package config

import "github.com/knadh/koanf/v2"

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
