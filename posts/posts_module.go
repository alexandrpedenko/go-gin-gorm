package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"gin-gorm-crud/posts/controller"
	"gin-gorm-crud/posts/model"
	"gin-gorm-crud/posts/repository"
	"gin-gorm-crud/posts/router"
	"gin-gorm-crud/posts/service"
)

type PostsModuleArgs struct {
	Db          *gorm.DB
	RouterGroup *gin.RouterGroup
	Validate    *validator.Validate
}

func PostsModule(args PostsModuleArgs) {
	args.Db.Table("posts").AutoMigrate(&model.Posts{})

	// NOTE: Think if it is possible to handle the `DI` in a better way
	postsRepository := repository.NewPostsRepository(args.Db)
	postsService := service.NewPostService(postsRepository, args.Validate)
	postsController := controller.NewPostsController(postsService)

	router.PostsRoutes(args.RouterGroup, postsController)
}
