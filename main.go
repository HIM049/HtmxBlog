package main

import (
	"HtmxBlog/config"
	"HtmxBlog/database"
	"HtmxBlog/router"
	"HtmxBlog/template"
	"fmt"
	"net/http"
	"os"
)

func main() {
	os.MkdirAll("./app_data", 0755)
	os.MkdirAll("./app_data/posts", 0755)
	os.MkdirAll("./app_data/attaches", 0755)

	// initialize modules
	config.Init()
	database.Init()
	template.Init()

	fmt.Println("Server is running on", config.Cfg.Service.Addr())
	http.ListenAndServe(config.Cfg.Service.Addr(), router.Init())
}
