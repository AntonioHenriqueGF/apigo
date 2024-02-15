package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowPostHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		err := errParamIsRequired("id", "urlParam")
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	post, err := repo.Post.GetByID(ctx, id)

	if err != nil {
		logger.Errorf("Error getting post: %s", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendBody(ctx, http.StatusOK, post)
}
