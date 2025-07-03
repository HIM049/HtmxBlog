package router

import (
	"HtmxBlog/config"
	"HtmxBlog/handler"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.HTMLRender = NewTemplateRender()

	// routers
	router.GET("/", handler.PageIndex)
	SetupDynmicRouter(router)

	responseGroup := router.Group("/response")
	{
		responseGroup.GET("/hello", handler.ResponseHello)
	}

	return router
}

func NewTemplateRender() multitemplate.Renderer {
	r := multitemplate.New()
	r.AddFromFiles("index", "templates/index.tmpl")

	for _, page := range config.Cfg.CustomPages {
		r.AddFromFiles(page, "templates/"+page+".tmpl")
	}

	return r
}

func SetupDynmicRouter(router *gin.Engine) {
	for _, page := range config.Cfg.CustomPages {
		router.GET("/"+page, handler.NewPageCustom(page))
	}
}
