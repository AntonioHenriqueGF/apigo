package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET /note",
	})
}
