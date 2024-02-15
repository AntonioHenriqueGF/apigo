package handler

import (
	"net/http"
	"time"

	"github.com/AntonioHenriqueGF/apigo/schemas"
	"github.com/gin-gonic/gin"
)

func CreatePostHandler(ctx *gin.Context) {
	request := CreatePostRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	dt := time.Now()

	post := schemas.Post{
		Title:         request.Title,
		Content:       request.Content,
		Date_creation: &dt,
	}

	logger.Debug(post)
}
