package router

import "github.com/gin-gonic/gin"

func Init() {
	// Initialize the router with default gin settings
	router := gin.Default()

	// Initializes the routes
	InitializeRoutes(router)

	// Runs the HTTP server and listens to requests on port 8080
	router.Run(":8080")
}
