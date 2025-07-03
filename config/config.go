package config

import "github.com/spf13/viper"

var Cfg Config

type Config struct {
	NetWork
	Pages
}

type NetWork struct {
	Host string
	Port int
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

	Cfg = Config{
		NetWork: net,
		Pages:   pages,
	}

	return nil
}
