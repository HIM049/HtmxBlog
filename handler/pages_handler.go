package handler

import "github.com/gin-gonic/gin"

func PageIndex(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.HTML(200, "index", nil)
}
