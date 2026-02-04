package config

import (
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var Cfg *Config

type Config struct {
	Database Database
}

// Init loads the config from the config.toml file.
// It panics when some error occurs.
func Init() {
	k := koanf.New(".")
	k.Load(file.Provider("config.toml"), toml.Parser())

	Cfg = &Config{
		Database: ReadDatabase(k),
	}
}
