package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUserHandler(ctx *gin.Context) {
	request := LoginUserRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Errorf("failed to bind request: %v", err)
		sendError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := repo.User.Login(ctx, request.Email, request.Password)
	if err != nil {
		logger.Errorf("failed to login user: %v", err)
		sendError(ctx, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	sendSuccess(ctx, "logged in")
}
