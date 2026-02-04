package handler

import (
	"net/http"
	"strings"

	"github.com/AntonioHenriqueGF/apigo/utils"
	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware checks if the user has a valid JWT token
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			sendError(c, http.StatusUnauthorized, "Missing authentication token")
			c.Abort()
			return
		}

		// The token should be prefixed with "Bearer "
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			sendError(c, http.StatusUnauthorized, "Invalid authentication token")
			c.Abort()
			return
		}

		tokenString = tokenParts[1]

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			sendError(c, http.StatusUnauthorized, "Invalid authentication token")
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
