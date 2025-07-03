package router

import (
	"HtmxBlog/handler"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.HTMLRender = NewTemplateRender()

	// routers
	router.GET("/", handler.PageIndex)

	responseGroup := router.Group("/response")
	{
		responseGroup.GET("/hello", handler.ResponseHello)
	}

	return router
}

func NewTemplateRender() multitemplate.Renderer {
	r := multitemplate.New()
	r.AddFromFiles("index", "templates/index.tmpl")
	return r
}
