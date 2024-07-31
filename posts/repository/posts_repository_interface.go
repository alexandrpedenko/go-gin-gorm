package repository

import (
	"gin-gorm-crud/posts/model"
)

type PostsRepositoryInterface interface {
	FindAll() ([]model.Posts, error)
	FindById(postsId int) (posts model.Posts, err error)
	Create(post model.Posts) error
	Update(id int, posts model.Posts) (model.Posts, error)
	Delete(postsId int) error
}
