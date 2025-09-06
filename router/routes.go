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
		v1.GET("/post/:id", handler.ShowPostHandler)

		// Create Post
		v1.POST("/post", handler.CreatePostHandler)

		// Update Post
		v1.PUT("/post/:id", handler.UpdatePostHandler)

		// Delete Post
		v1.DELETE("/post/:id", handler.DeletePostHandler)

		// Show all Posts
		v1.GET("/posts", handler.ListPostsHandler)

		// Create User
		v1.POST("/user", handler.CreateUserHandler)

		// Log User
		v1.POST("/login", handler.LoginUserHandler)
	}
}
