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
	services.UpdateConfig()
	template.Init()
	template.InitBaseApp()
	router.Init()

	// handle page change
	services.RegisterOnPageChange(func() {
		go func() {
			router.RefreshRoutes()
			template.UpdateNavigation()
		}()
	})
	// handle category change
	services.RegisterOnCategoryChange(func() {
		go func() {
			template.UpdateCategories()
		}()
	})
	services.RegisterOnPostChange(func() {
		go func() {
			template.UpdateCategories()
		}()
	})
	services.RegisterOnSettingChange(func() {
		go func() {
			services.UpdateConfig()
			template.UpdateSettings()
		}()
	})

	fmt.Println("Server is running on", config.Cfg.Service.Addr())
	http.ListenAndServe(config.Cfg.Service.Addr(), router.HRouter)
}
