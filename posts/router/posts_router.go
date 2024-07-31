package router

import (
	"gin-gorm-crud/posts/controller"

	"github.com/gin-gonic/gin"
)

func PostsRoutes(router *gin.RouterGroup, postsController *controller.PostController) {
	posts := router.Group("/posts")
	{
		posts.GET("/", postsController.FindAll)
		posts.POST("/", postsController.Create)
		// posts.GET("/:id", postsController.FindByID)
		// posts.PUT("/:id", postsController.Update)
		// posts.DELETE("/:id", postsController.Delete)
	}
}
