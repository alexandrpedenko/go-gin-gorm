package service

import (
	"gin-gorm-crud/posts/dto/request"
	"gin-gorm-crud/posts/dto/response"
)

type PostServiceInterface interface {
	FindAll() ([]response.PostResponse, error)
	Create(post request.CreatePostsRequest) error
	// FindByID(id int) (response.PostResponse, error)
	// Update(id int, post model.Posts) error
	// Delete(id int) error
}
