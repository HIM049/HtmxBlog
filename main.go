package main

import (
	"HtmxBlog/config"
	"HtmxBlog/database"
	"HtmxBlog/router"
	"HtmxBlog/template"
	"fmt"
	"net/http"
)

func main() {
	// initialize modules
	config.Init()
	database.Init()
	template.Init()

	fmt.Println("Server is running on", config.Cfg.Service.Addr())
	http.ListenAndServe(config.Cfg.Service.Addr(), router.Init())
}
