package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST /post",
	})
}
