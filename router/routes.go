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

		// Show all Posts
		v1.GET("/posts", handler.ListPostsHandler)

		// Create User
		v1.POST("/user", handler.CreateUserHandler)

		// Log User
		v1.POST("/login", handler.LoginUserHandler)
	}

	protected := router.Group("/api/protected")
	protected.Use(handler.AuthenticationMiddleware())
	{
		// Create Post
		protected.POST("/post", handler.CreatePostHandler)

		// Update Post
		protected.PUT("/post/:id", handler.UpdatePostHandler)

		// Delete Post
		protected.DELETE("/post/:id", handler.DeletePostHandler)
	}
}
