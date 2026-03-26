package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/AntonioHenriqueGF/apigo/schemas"
	"github.com/gin-gonic/gin"
)

func UpdatePostHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	request := UpdatePostRequest{}

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

	post := schemas.PostUpdate{
		ID:      id,
		Title:   request.Title,
		Content: request.Content,
	}

	err = repo.Post.UpdatePostById(ctx, post)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			sendError(ctx, http.StatusNotFound, "post not found")
			return
		}

		logger.Errorf("Error updating post: %s", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "Post updated")
}
