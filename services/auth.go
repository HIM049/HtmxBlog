package services

import (
	"HtmxBlog/config"
	"HtmxBlog/state"
	"time"
)

func IsTokenExpired() bool {
	return time.Since(state.CreateTime) > time.Duration(config.Cfg.Service.ValidTime)*time.Hour
}
