package router

import (
	"github.com/AntonioHenriqueGF/apigo/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// Show note
		v1.GET("/note", handler.ShowNoteHandler)

		// Create note
		v1.POST("/note", handler.CreateNoteHandler)

		// Update note
		v1.PUT("/note", handler.UpdateNoteHandler)

		// Delete note
		v1.DELETE("/note", handler.DeleteNoteHandler)

		// Show all notes
		v1.GET("/notes", handler.ListNotesHandler)
	}
}
