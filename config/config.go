package config

import "github.com/spf13/viper"

var Cfg Config

type Config struct {
	NetWork
	Pages
	Database
}

type NetWork struct {
	Host string
	Port int
}

type Database struct {
	Driver string
	DSN    string
}

type Pages struct {
	CustomPages []string
}

// InitConfig initializes the configuration from file.
func InitConfig() error {
	// read the configuration file
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// write into the global config variable
	net := NetWork{
		Host: viper.GetString("Network.Host"),
		Port: viper.GetInt("Network.Port"),
	}

	pages := Pages{
		CustomPages: viper.GetStringSlice("Pages.CustomPages"),
	}

	database := Database{
		Driver: viper.GetString("Database.Driver"),
		DSN:    viper.GetString("Database.DSN"),
	}

	Cfg = Config{
		NetWork:  net,
		Pages:    pages,
		Database: database,
	}

	return nil
}
