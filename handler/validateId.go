package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateIdMiddleware(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		err := errParamIsRequired("id", "urlParam")
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hasPost, err := repo.Post.ExistsByID(ctx, id)

	if err != nil {
		logger.Errorf("Error obtaining post: %s", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if !hasPost {
		err := errEntrieNotFound("post")
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
}
