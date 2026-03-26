package handler

import (
	"net/http"

	"github.com/AntonioHenriqueGF/apigo/utils"
	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware checks if the user has a valid JWT token
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("auth_token")
		if err != nil {
			sendError(c, http.StatusUnauthorized, "Missing authentication token")
			c.Abort()
			return
		}

		claims, err := utils.VerifyToken(cookie)
		if err != nil {
			sendError(c, http.StatusUnauthorized, "Invalid authentication token")
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
