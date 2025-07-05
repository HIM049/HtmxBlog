package handler

import (
	"HtmxBlog/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPost(ctx *gin.Context) {
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")

	if title == "" || content == "" {
		ctx.String(http.StatusBadRequest, "Title and content cannot be empty")
		return
	}

	err := services.CreatePost(title, content)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to create post: %v", err)
		return
	}

	ctx.String(http.StatusOK, "Post created successfully")
}
