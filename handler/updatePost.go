package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatePostHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT /post",
	})
}
