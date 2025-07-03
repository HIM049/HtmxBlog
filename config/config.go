package config

import "github.com/spf13/viper"

var Cfg Config

type Config struct {
	NetWork
}

type NetWork struct {
	Host string
	Port int
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

	Cfg = Config{
		NetWork: net,
	}

	return nil
}
