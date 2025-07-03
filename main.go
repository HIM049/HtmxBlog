package main

import (
	"HtmxBlog/config"
	"HtmxBlog/router"
	"fmt"
	"log"
)

func main() {
	// initialize config
	err := config.InitConfig()
	if err != nil {
		log.Fatalln("Failed to initialize config:", err)
	}

	// initialize router
	r := router.SetupRouter()
	err = r.Run(fmt.Sprintf("%s:%d", config.Cfg.NetWork.Host, config.Cfg.NetWork.Port))
	if err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}
