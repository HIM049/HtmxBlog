package main

import (
	"HtmxBlog/config"
	"HtmxBlog/maintain"
	"HtmxBlog/router"
	"HtmxBlog/services"
	"HtmxBlog/template"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Parse command-line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			fmt.Println("Installing...")
			if err := maintain.CheckAndInstall(); err != nil {
				fmt.Fprintf(os.Stderr, "Install failed: %v\n", err)
				os.Exit(1)
			}
			return

		case "export":
			outPath := ""
			if len(os.Args) > 2 {
				outPath = os.Args[2]
			}
			fmt.Println("Exporting database...")
			if err := maintain.ExportAll(outPath); err != nil {
				fmt.Fprintf(os.Stderr, "Export failed: %v\n", err)
				os.Exit(1)
			}
			return

		case "import":
			if len(os.Args) < 3 {
				fmt.Fprintf(os.Stderr, "Usage: %s import <file.json>\n", os.Args[0])
				os.Exit(1)
			}
			fmt.Println("Importing database...")
			if err := maintain.ImportAll(os.Args[2]); err != nil {
				fmt.Fprintf(os.Stderr, "Import failed: %v\n", err)
				os.Exit(1)
			}
			return

		default:
			fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
			fmt.Fprintf(os.Stderr, "Usage: %s [install | export [out.json] | import <file.json>]\n", os.Args[0])
			os.Exit(1)
		}
	}

	// Normal server startup
	os.MkdirAll("./app_data", 0755)
	os.MkdirAll("./app_data/posts", 0755)
	os.MkdirAll("./app_data/attaches", 0755)
	os.MkdirAll("./app_data/drafts", 0755)

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
			template.UpdateTags()
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
