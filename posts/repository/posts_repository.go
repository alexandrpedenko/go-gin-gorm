package repository

import (
	"errors"
	"fmt"
	"gin-gorm-crud/core/app_errors"
	"gin-gorm-crud/posts/dto/request"
	"gin-gorm-crud/posts/model"

	"gorm.io/gorm"
)

type PostsRepository struct {
	Db *gorm.DB
}

func NewPostsRepository(Db *gorm.DB) *PostsRepository {
	return &PostsRepository{Db}
}

func (t *PostsRepository) FindAll() ([]model.Posts, error) {
	var posts []model.Posts
	db := t.Db.Find(&posts)

	if db.Error != nil {
		return posts, app_errors.NewUnexpectedError(
			fmt.Sprintf("Error occurred when query all posts: %v", db.Error),
		)
	}

	return posts, db.Error
}

func (t *PostsRepository) Count() (int64, error) {
	var count int64
	db := t.Db.Model(&model.Posts{}).Count(&count)

	if db.Error != nil {
		return 0, app_errors.NewUnexpectedError(
			fmt.Sprintf("Error occurred when query count for all posts: %v", db.Error),
		)
	}

	return count, nil
}

func (t *PostsRepository) FindById(id int) (model.Posts, error) {
	var post model.Posts
	db := t.Db.Where("id = ?", id).First(&post)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return post, app_errors.NewNotFoundError(
			fmt.Sprintf("Post with ID %d not found", id),
		)
	} else if db.Error != nil {
		return post, app_errors.NewUnexpectedError(
			fmt.Sprintf("Error occurred when query count for all posts: %v", db.Error),
		)
	}

	return post, db.Error
}

func (t *PostsRepository) Create(post model.Posts) error {
	db := t.Db.Create(&post)

	return db.Error
}

func (t *PostsRepository) Update(id int, post model.Posts) (model.Posts, error) {
	var updatePost = request.UpdatePostsRequest{
		Id:    post.Id,
		Title: post.Title,
	}

	var updatedPost model.Posts
	result := t.Db.Model(&post).Where("id = ?", updatePost.Id).Updates(updatedPost).Scan(&updatedPost)

	return post, result.Error
}

func (t *PostsRepository) Delete(id int) error {
	var posts model.Posts
	db := t.Db.Where("id = ?", id).Delete(&posts)

	return db.Error
}
