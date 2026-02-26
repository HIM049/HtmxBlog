package config

import (
	"fmt"

	"github.com/knadh/koanf/v2"
)

type Service struct {
	Port        string
	AdminPasswd string
	ValidTime   int
}

func ReadService(k *koanf.Koanf) Service {
	return Service{
		Port:        k.String("service.port"),
		AdminPasswd: k.String("service.admin_password"),
		ValidTime:   k.Int("service.vaild_hour"),
	}
}

func (s *Service) Addr() string {
	return fmt.Sprintf(":%s", s.Port)
}
