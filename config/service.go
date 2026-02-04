package config

import (
	"fmt"

	"github.com/knadh/koanf/v2"
)

type Service struct {
	Port string
}

func ReadService(k *koanf.Koanf) Service {
	return Service{
		Port: k.String("service.port"),
	}
}

func (s *Service) Addr() string {
	return fmt.Sprintf(":%s", s.Port)
}
