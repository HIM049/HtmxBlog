package main

import (
	"HtmxBlog/config"
	"HtmxBlog/router"
	"HtmxBlog/services"
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
	config.InitDB()
	template.Init()
	router.Init()

	// handle page change
	services.RegisterOnPageChange(func() {
		go func() {
			router.RefreshRoutes()
			template.UpdateNavigation()
		}()
	})

	fmt.Println("Server is running on", config.Cfg.Service.Addr())
	http.ListenAndServe(config.Cfg.Service.Addr(), router.HRouter)
}
