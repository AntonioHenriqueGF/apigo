package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, operation string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation success: %s", operation),
	})
}
