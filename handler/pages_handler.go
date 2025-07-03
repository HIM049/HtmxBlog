package handler

import (
	"github.com/gin-gonic/gin"
)

func PageIndex(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.HTML(200, "index", nil)
}

func NewPageCustom(pageName string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html")
		ctx.HTML(200, pageName, nil)
	}
}
