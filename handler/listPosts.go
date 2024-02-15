package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListPostsHandler(ctx *gin.Context) {

	posts, err := repo.Post.GetAll(ctx)

	if err != nil {
		logger.Errorf("Error listing posts: %s", err.Error())
	}

	logger.Debug(posts)

	ctx.JSON(http.StatusOK, posts)
}
