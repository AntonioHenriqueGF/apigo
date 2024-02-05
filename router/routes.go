package router

import (
	"github.com/AntonioHenriqueGF/apigo/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		// Show Post
		v1.GET("/post", handler.ShowPostHandler)

		// Create Post
		v1.POST("/post", handler.CreatePostHandler)

		// Update Post
		v1.PUT("/post", handler.UpdatePostHandler)

		// Delete Post
		v1.DELETE("/post", handler.DeletePostHandler)

		// Show all Posts
		v1.GET("/posts", handler.ListPostsHandler)
	}
}
