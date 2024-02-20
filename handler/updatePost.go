package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatePostHandler(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT /post",
	})
}
