package handler

import (
	"net/http"
	"time"

	"github.com/AntonioHenriqueGF/apigo/schemas"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	dt := time.Now()

	user := schemas.User{
		Name:       request.Name,
		Email:      request.Email,
		Password:   request.Password,
		Created_at: &dt,
	}

	err := repo.User.Create(ctx, &user)

	if err != nil {
		logger.Errorf("Error creating user: %s", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "User created")
}
