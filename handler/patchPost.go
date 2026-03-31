package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/AntonioHenriqueGF/apigo/schemas"
	"github.com/gin-gonic/gin"
)

func PatchPostHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	request := PatchPostRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("Error binding request body: %s", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	post := schemas.PostPatch{
		ID:      id,
		Title:   request.Title,
		Content: request.Content,
	}

	if err := repo.Post.PatchPostById(ctx, post); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			sendError(ctx, http.StatusNotFound, "post not found")
			return
		}

		logger.Errorf("Error patching post: %s", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "Post patched")
}
