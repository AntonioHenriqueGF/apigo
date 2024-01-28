package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE /note",
	})
}
