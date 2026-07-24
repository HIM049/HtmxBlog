package main

import (
	"HtmxBlog/config"
	"HtmxBlog/handler"
	"HtmxBlog/maintain"
	"HtmxBlog/model"
	"HtmxBlog/router"
	"HtmxBlog/services"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{ReportTimestamp: true})
	log.SetDefault(logger)

	swFunction()

	// Normal server startup
	os.MkdirAll("./app_data", 0755)
	os.MkdirAll(config.POSTS_DIR, 0755)
	os.MkdirAll(config.ATTACHES_DIR, 0755)
	os.MkdirAll(config.DRAFTS_DIR, 0755)

	// initialize modules
	config.Init()
	config.InitDB()
	config.DB.AutoMigrate(&model.Post{}, &model.Page{}, &model.Attach{}, &model.Setting{}, &model.Comment{}, &model.Redirect{}, &model.AccessRecord{})
	services.UpdateConfig()
	handler.InitI18n()
	services.Init()
	services.InitBaseApp()
	router.Init()

	handler.RefreshRoutes = router.RefreshRoutes

	log.Infof("Server is running on %s", config.Cfg.Service.Addr())
	http.ListenAndServe(config.Cfg.Service.Addr(), router.HRouter)
}

func swFunction() {
	// Parse command-line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			log.Info("Installing...")
			if err := maintain.CheckAndInstall(); err != nil {
				log.Errorf("Install failed: %v", err)
				os.Exit(1)
			}
			return

		case "export":
			outPath := ""
			if len(os.Args) > 2 {
				outPath = os.Args[2]
			}
			log.Info("Exporting database...")
			if err := maintain.ExportAll(outPath); err != nil {
				log.Errorf("Export failed: %v", err)
				os.Exit(1)
			}
			return

		case "import":
			if len(os.Args) < 3 {
				log.Errorf("Usage: %s import <file.json>", os.Args[0])
				os.Exit(1)
			}
			log.Info("Importing database...")
			if err := maintain.ImportAll(os.Args[2]); err != nil {
				log.Errorf("Import failed: %v", err)
				os.Exit(1)
			}
			return

		default:
			log.Errorf("Unknown command: %s", os.Args[1])
			log.Errorf("Usage: %s install | export [out.json] | import [file.json]", os.Args[0])
			os.Exit(1)
		}
	}
}
