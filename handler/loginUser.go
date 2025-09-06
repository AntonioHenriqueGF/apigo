package handler

import (
	"net/http"

	"github.com/AntonioHenriqueGF/apigo/utils"
	"github.com/gin-gonic/gin"
)

func LoginUserHandler(ctx *gin.Context) {
	request := LoginUserRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Errorf("failed to bind request: %v", err)
		sendError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := repo.User.Login(ctx, request.Email, request.Password)
	if err != nil {
		logger.Errorf("failed to login user: %v", err)
		sendError(ctx, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token, err := utils.GenerateToken(uint(user.ID))
	if err != nil {
		logger.Errorf("failed to generate token: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Error generating token")
	}

	sendBody(ctx, http.StatusOK, gin.H{"token": token})
}
