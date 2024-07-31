package service

import (
	"fmt"
	"gin-gorm-crud/posts/dto/request"
	"gin-gorm-crud/posts/dto/response"
	"gin-gorm-crud/posts/model"
	"gin-gorm-crud/posts/repository"

	"github.com/go-playground/validator/v10"
)

type PostService struct {
	postRepository repository.PostsRepositoryInterface
	validate       *validator.Validate
}

func NewPostService(postRepository repository.PostsRepositoryInterface, validate *validator.Validate) *PostService {
	return &PostService{
		postRepository,
		validate,
	}
}

func (s *PostService) FindAll() ([]response.PostResponse, error) {
	posts, err := s.postRepository.FindAll()

	fmt.Println(posts)

	if err != nil {
		return nil, err
	}

	serializedPosts := make([]response.PostResponse, len(posts))
	for i, post := range posts {
		serializedPosts[i] = response.PostResponse{
			Id:    post.Id,
			Title: post.Title,
		}
	}

	return serializedPosts, err
}

// func (s *PostService) FindByID(id int) (model.Posts, error) {
// 	return nil, nil
// }

func (s *PostService) Create(post request.CreatePostsRequest) error {
	err := s.validate.Struct(post)

	fmt.Println(err)

	if err != nil {
		return err
	}

	postModel := model.Posts{
		Title: post.Title,
	}

	err = s.postRepository.Create(postModel)
	return err
}

// func (s *PostService) Update(id int, post model.Posts) (model.Posts, error) {
// 	return nil, nil
// }

// func (s *PostService) Delete(id int) error {
// 	return nil, nil
// }
