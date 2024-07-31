package controller

import (
	"gin-gorm-crud/core/app_errors"
	"gin-gorm-crud/core/dto/response"
	"gin-gorm-crud/posts/dto/request"
	"gin-gorm-crud/posts/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postsService service.PostServiceInterface
}

func NewPostsController(postsService service.PostServiceInterface) *PostController {
	return &PostController{postsService}
}

func (c *PostController) FindAll(ctx *gin.Context) {
	posts, err := c.postsService.FindAll()

	if err != nil {
		app_errors.HandleHTTPError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.ResultsResponse{
		Results: posts,
	})
}

func (c *PostController) Create(ctx *gin.Context) {
	createPostRequest := request.CreatePostsRequest{}
	err := ctx.ShouldBindJSON(&createPostRequest)

	if err != nil {
		app_errors.HandleHTTPError(ctx, app_errors.NewBadRequestError(err.Error()))
		return
	}

	err = c.postsService.Create(createPostRequest)

	if err != nil {
		app_errors.HandleHTTPError(ctx, app_errors.NewBadRequestError(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
	})
}
