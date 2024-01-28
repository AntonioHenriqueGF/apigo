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

func CreateNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST /note",
	})
}

func UpdateNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT /note",
	})
}

func DeleteNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE /note",
	})
}

func ShowAllNotesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET /notes",
	})
}
