package config

import (
	"fmt"
	"strconv"
)

type Service struct {
	Port        string
	AdminPasswd string
	ValidTime   int
	ThemeDir    string
}

func ReadService() Service {
	validTime, _ := strconv.Atoi(getEnv("VALID_HOUR"))
	return Service{
		Port:        getEnv("PORT"),
		AdminPasswd: getEnv("ADMIN_PASSWORD"),
		ValidTime:   validTime,
		ThemeDir:    getEnvOrDefault("THEME_DIR", "./themes"),
	}
}

func (s *Service) Addr() string {
	return fmt.Sprintf(":%s", s.Port)
}
