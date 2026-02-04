package main

import (
	"HtmxBlog/config"
	"HtmxBlog/database"
	"fmt"
)

func main() {
	// initialize modules
	config.Init()
	database.Init()

	fmt.Println(config.Cfg.Database)
}
