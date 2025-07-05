package handler

import (
	"HtmxBlog/services"
	"github.com/gin-gonic/gin"
)

func PageIndex(ctx *gin.Context) {
	posts, _ := services.ListPosts()
	ctx.Header("Content-Type", "text/html")
	ctx.HTML(200, "index", gin.H{
		"Posts": posts,
	})
}

func NewPageCustom(pageName string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html")
		ctx.HTML(200, pageName, nil)
	}
}
