package handler

import "github.com/gin-gonic/gin"

func ResponseHello(ctx *gin.Context) {
	ctx.String(200, "<p>Hello, World. Welcome to visit my site!</p>")
}
