package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletePostHandler(ctx *gin.Context) {
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

	err = repo.Post.DeleteByID(ctx, id)

	if err != nil {
		logger.Errorf("Error deleting post: %s", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "Post deleted")
}
